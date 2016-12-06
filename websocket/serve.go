package websocket

import (
	"encoding/json"
	"log"
	"net/http"

	"fmt"
	"github.com/gorilla/websocket"
	"github.com/nrechn/musubi/utils"
	"io/ioutil"
	"strings"
	"time"
)

func Init() *Hub {
	hub := newHub()
	go hub.run()
	return hub
}

// Execute handles websocket requests from the peer.
func Execute(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	go tokenRegister(hub, w, r, conn)
}

// tokenRegister checks connected client(device)'s identification.
//
// It waits 30 second for client declaring identification.
// tokenRegister will reject connection if:
// - client keeps silence
// - client sends repeated name/token
// - client sends wrong name/token
//
// Identification declaration format:
// { "Name": "DEVICE NAME", "Token": "DEVICE TOKEN"}
func tokenRegister(hub *Hub, w http.ResponseWriter, r *http.Request, conn *websocket.Conn) {
	token := make(chan AuthMessage)
	defer close(token)
	go func() {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
		}
		token <- readToken(string(message))
	}()

	select {
	case res := <-token:
		fmt.Println("[Websocket] Received AuthMessage. Token: " + res.Token)
		if ifRepeat(hub, res) {
			if utils.Authenticate(res.Name, res.Token) {
				go response(conn, `{"Status": "ok!"}`)
				client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256), name: res.Name, token: res.Token}
				client.hub.register <- client
				go client.writePump()
				go client.readPump(hub)
			} else {
				go response(conn, `{"Status": "error! wrong user name or token."}`)
			}
		} else {
			go response(conn, `{"Status": "error! your device is already online."}`)
		}
	case <-time.After(time.Second * 30):
		fmt.Println("error: websocket connection request timeout.")
		go response(conn, `{"Status": "error! id authentication is required."}`)
	}
}

// readToken reads message and return in type "AuthMessage".
func readToken(message string) AuthMessage {
	var m AuthMessage
	dec := json.NewDecoder(strings.NewReader(message))
	if err := dec.Decode(&m); err != nil {
		log.Println(err)
	}
	return m
}

// ifRepeat checks if name/token is repeated.
// It compares with online devices.
func ifRepeat(hub *Hub, m AuthMessage) bool {
	for client := range hub.clients {
		if m.Name == client.name {
			return false
		}
		if m.Token == client.token {
			return false
		}
	}
	return true
}

// response makes a response to device(client) trying to make a connection.
func response(conn *websocket.Conn, reason string) {
	r := strings.NewReader(reason)
	b, err := ioutil.ReadAll(r)
	if err != nil {
		log.Println(err)
	}
	if err := conn.WriteMessage(websocket.TextMessage, b); err != nil {
		log.Println(err)
	}
	if reason != `{"Status": "ok!"}` {
		conn.Close()
	}
}
