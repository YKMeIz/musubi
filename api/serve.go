package api

import (
	"github.com/gin-gonic/gin"
	"github.com/nrechn/musubi/pushbullet"
	"github.com/nrechn/musubi/utils"
	"github.com/nrechn/musubi/websocket"
)

// Execute starts the main program.
//
// It provides routing and handling requests.
func Execute() {
	url := utils.GetString("domainName") + ":" + utils.GetString("portNumber")
	r := gin.Default()

	if utils.Websocket() {
		hub := websocket.Init()
		r.GET("/ws", func(c *gin.Context) {
			websocket.Execute(hub, c.Writer, c.Request)
		})
		r.POST("/nc", func(c *gin.Context) {
			notiViaWS(c, hub)
		})
	} else {
		r.POST("/nc", notiViaPb)
	}

	if utils.Secure() {
		certFile := utils.GetString("certChain")
		keyFile := utils.GetString("certKey")
		r.RunTLS(url, certFile, keyFile)
	} else {
		r.Run(url)
	}
}

// Send notification via websocket.
func notiViaWS(c *gin.Context, hub *websocket.Hub) {
	err := websocket.PushNoti(c, hub)
	checkErr(err, c)
}

// Send notification via Pushbullet.
func notiViaPb(c *gin.Context) {
	// ToDo: more push action type need to be available.
	err := pushbullet.Pushbullet(c, "note")
	checkErr(err, c)
}
