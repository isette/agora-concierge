package main

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/isette/appointments-concierge/handlers"
	"github.com/isette/appointments-concierge/pkg"
	"github.com/isette/appointments-concierge/runners"
)

func init() {
	handlers.LoadEnv("../.env")
	pkg.InitGORM()
	// pkg.Migrate()

	log.Println("Environment variables loaded successfully.")
}

func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Content-Length", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.OPTIONS("/*path", func(c *gin.Context) {
		c.JSON(200, struct{}{})
	})

	// router.GET("/socket.io/", handlers.HandleConnection)

	group := router.Group("/v1")
	group.GET("rtc/:channelName/:role/:tokenType/:uid", runners.GetRtcToken)
	group.GET("rte/:channelName/:role/:tokenType/:uid", runners.GetBothTokens)
	group.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	router.Run(":" + handlers.GetEnvWithKey("PORT"))
}
