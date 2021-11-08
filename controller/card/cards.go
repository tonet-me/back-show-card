package card

import (
	env "Card-Request-Manager/env"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

var netClient = &http.Client{}
var cardRequestServer = env.GoDotEnvVariable("CARD_REQUEST_URI")

func GetCardByUsername(c *gin.Context) {
	var userName string = c.Param(("cardName"))

	tonetCardReq, err := http.NewRequest("GET", cardRequestServer+"un/" + userName, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": "false", "message": err.Error()})
		return
	}

	getUserAgent(c,tonetCardReq)

	// tonetCardReq.Header.Add("user-agent-orig","masoood")
	resp, err := netClient.Do(tonetCardReq)
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


func GetCardByQrcode(c *gin.Context) {
	var userName string = c.Param(("qrCode"))

	tonetCardReq, err := http.NewRequest("GET", cardRequestServer+"qr/" + userName, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": "false", "message": err.Error()})
		return
	}

	getUserAgent(c,tonetCardReq)

	// tonetCardReq.Header.Add("user-agent-orig","masoood")
	resp, err := netClient.Do(tonetCardReq)
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

func getUserAgent(c *gin.Context, req *http.Request) {
	uaStringFromReq := c.Request.Header.Get("User-Agent")
	req.Header.Add("user-agent",string(uaStringFromReq))
	c.Next()
}