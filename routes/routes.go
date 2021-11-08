package router

import (
	card "Card-Request-Manager/controller/card"
	env "Card-Request-Manager/env"
	"net/http"

	"github.com/gin-gonic/gin"
)

func StartGin() {
	router := gin.Default()
	// router.Use(cors.Default())
	router.Use(CORSMiddleware())
	router.Use(gin.Recovery())
	router.GET("qr/:qrCode", card.GetCardByQrcode)
	router.GET("/:cardName", card.GetCardByUsername)
	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus((http.StatusNotFound))
	})
	port := env.GoDotEnvVariable("PORT")
	router.Run(port)
}

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}