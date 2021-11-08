package router

import (
	card "Card-Request-Manager/controller/card"
	env "Card-Request-Manager/env"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func StartGin() {
	router := gin.Default()
	 router.Use(cors.Default())
	router.Use(gin.Recovery())
	router.GET("qr/:qrCode", card.GetCardByQrcode)
	router.GET("/:cardName", card.GetCardByUsername)
	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus((http.StatusNotFound))
	})
	port := env.GoDotEnvVariable("PORT")
	router.Run(port)
}

