package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	rtctokenbuilder2 "github.com/AgoraIO-Community/go-tokenbuilder/rtctokenbuilder"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/isette/agoraio-service/handlers"
)

func init() {
	handlers.LoadEnv()
	log.Println("Environment variables loaded successfully.")
}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	channelName := request.PathParameters["channelName"]
	roleStr := request.PathParameters["role"]
	tokenType := request.PathParameters["tokentype"]
	uidStr := request.PathParameters["uid"]

	if channelName == "" || roleStr == "" || tokenType == "" || uidStr == "" {
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       `{"error": "Missing required path parameters"}`,
		}, nil
	}

	var role rtctokenbuilder2.Role
	if roleStr == "publisher" {
		role = rtctokenbuilder2.RolePublisher
	} else {
		role = rtctokenbuilder2.RoleSubscriber
	}

	expireTimestamp := uint32(3600)
	if expireStr, exists := request.QueryStringParameters["expireTimestamp"]; exists {
		expire, err := strconv.Atoi(expireStr)
		if err == nil {
			expireTimestamp = uint32(expire)
		}
	}

	rtcToken, err := handlers.GenerateRtcToken(channelName, uidStr, tokenType, role, expireTimestamp)
	if err != nil {
		log.Println("Error generating RTC token:", err)
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       fmt.Sprintf(`{"error": "Failed to generate RTC token: %s"}`, err.Error()),
		}, nil
	}

	response := map[string]string{
		"rtcToken": rtcToken,
	}
	responseBody, _ := json.Marshal(response)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(responseBody),
	}, nil
}

func main() {
	lambda.Start(Handler)
}
