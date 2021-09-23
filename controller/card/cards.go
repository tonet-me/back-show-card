package card

import (
	env "Card-Request-Manager/env"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

var netClient = &http.Client{}
var cardRequestServer = env.GoDotEnvVariable("CARD_REQUEST_URI")

func GetCard(c *gin.Context) {
	var userName string = c.Param(("userName"))
	resp, err := netClient.Get(cardRequestServer + userName)
	defer netClient.CloseIdleConnections()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": "false", "message": err.Error()})
		return
	}
	//Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": "false", "message": err.Error()})
		return
	}
	jsonData := []byte(string(body))
	c.Data(http.StatusOK, "application/json", jsonData)
}
