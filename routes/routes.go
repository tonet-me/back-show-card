package router

import (
	card "Card-Request-Manager/controller/card"
	env "Card-Request-Manager/env"
	"net/http"

	"github.com/gin-gonic/gin"
)

func StartGin() {
	router := gin.Default()
	router.Use(gin.Recovery())
	router.GET("/:userName", card.GetCard)
	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus((http.StatusNotFound))
	})
	port := env.GoDotEnvVariable("PORT")
	router.Run(port)
}
