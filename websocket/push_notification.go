package websocket

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/nrechn/musubi/database"
	"github.com/nrechn/musubi/pushbullet"
	"github.com/nrechn/musubi/utils"
)

// PushNoti handles notification pushing via websocket.
//
// If target device is offline, PushNoti will send notificatrion via Pushbullet
// as long as Pushbullet's token is set.
func PushNoti(c *gin.Context, hub *Hub) error {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Println(err)
		return errors.New("internal error.")
	}
	msg := utils.ReadJson(string(body))
	b, err := json.Marshal(msg.Data)
	if err != nil {
		log.Println(err)
		return errors.New("internal error.")
	}
	var send bool
	for i := 0; i < len(msg.Destination); i++ {
		isOl := isOnline(hub, msg.Destination[i])
		isToke := database.IsToken(msg.Destination[i])
		switch {
		case isToke && isOl:
			for client := range hub.clients {
				if client.token == msg.Destination[i] {
					client.send <- b
				}
			}
			send = true
		case isToke && !isOl:
			defaultNotiRec := utils.GetStringSlice("defaultNotificationReceiver")
			for i := 0; i < len(defaultNotiRec); i++ {
				destName := database.GetName(msg.Destination[i])
				if destName != "" && (defaultNotiRec[i] == msg.Destination[i] || defaultNotiRec[i] == destName) {
					if err := pushbullet.PushNote(msg.Data["Title"], msg.Data["Text"]); err != nil {
						log.Println("write close:", err)
						return errors.New("internal error.")
					}
					send = true
				}
			}
			if !send {
				return errors.New("error! destination is offline.")
			}
		case !isToke && (!isOl || isOl):
			return errors.New("error! wrong destination token.")
		default:
			break
		}
	}
	if !send {
		hub.broadcast <- b
	}
	return nil
}

// isOnline checks if given destination is online.
func isOnline(hub *Hub, toke string) bool {
	for client := range hub.clients {
		if client.token == toke {
			return true
		}
	}
	return false
}
