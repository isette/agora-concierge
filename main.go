package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/gin-gonic/gin"
	"github.com/isette/agora.io-service/handlers"
	"github.com/isette/agora.io-service/runners"
)

func init() {
	handlers.LoadEnv()
}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Received body: ", request.Body)

	return events.APIGatewayProxyResponse{Body: request.Body, StatusCode: 200}, nil
}

func main() {
	api := gin.Default()

	Port := handlers.GetEnvWithKey("PORT")
	fmt.Println("Running in http://localhost:" + Port)

	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	api.GET("rtc/:channelName/:role/:tokentype/:uid/", runners.GetRtcToken)

	api.Run(":" + Port)
	// lambda.Start(Handler)
}
