package pushbullet

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nrechn/musubi/utils"
)

// pushNote sends a POST request to Pushbullet in order to make a
// Pushbullet push notification.
func PushNote(title, body string) error {
	token := utils.GetString("pushbullet.token")
	jsonStr := `{"body":"` + body + `","title":"` + title + `","type":"note"}`
	url := "https://api.pushbullet.com/v2/pushes"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(jsonStr)))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Access-Token", token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return errors.New(string(body))
	}
	return nil
}

// Pushbullet processes http request body and implement a Pushbullet
// push action.
//
// ToDo: more push action type need to be available.
func Pushbullet(c *gin.Context, typ string) error {
	switch typ {
	case "note":
		msg, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			return errors.New("internal error.")
		}

		m := utils.ReadJson(string(msg)).Data
		if err := PushNote(m["Title"], m["Text"]); err != nil {
			return errors.New("internal error.")
		}
		return nil
	default:
		return errors.New("message type not found.")
	}

}
