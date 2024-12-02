package runners

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/isette/agoraio-service/handlers"
	"github.com/isette/agoraio-service/utils"
)

func GetRtcToken(c *gin.Context) {
	log.Printf("rtc token\n")

	fmt.Println(c)

	channelName, tokenType, uidStr, role, expireTimestamp, err := utils.ParseRtcParams(c)

	fmt.Println(channelName, uidStr, tokenType, role, expireTimestamp)

	if err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(400, gin.H{
			"message": "Error Generating RTC token: " + err.Error(),
			"status":  400,
		})
		return
	}

	fmt.Println(channelName, uidStr, tokenType, role, expireTimestamp)

	rtcToken, tokenErr := handlers.GenerateRtcToken(channelName, uidStr, tokenType, role, expireTimestamp)

	if tokenErr != nil {
		log.Println(tokenErr)
		c.Error(tokenErr)
		errMsg := "Error Generating RTC token - " + tokenErr.Error()
		c.AbortWithStatusJSON(400, gin.H{
			"status": 400,
			"error":  errMsg,
		})
	} else {
		log.Println("RTC Token generated")
		c.JSON(200, gin.H{
			"rtcToken": rtcToken,
		})
	}
}

func GetBothTokens(c *gin.Context) {
	log.Printf("dual token\n")
	channelName, tokenType, uidStr, role, expireTimestamp, err := utils.ParseRtcParams(c)

	fmt.Println("tokenType", tokenType)

	if err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(400, gin.H{
			"message": "Error Generating RTE token: " + err.Error(),
			"status":  400,
		})
		return
	}

	rtcToken, rtcTokenErr := handlers.GenerateRteToken(channelName, uidStr, tokenType, role, expireTimestamp)

	rtmToken, rtmTokenErr := handlers.GenerateRtmToken(channelName, uidStr, tokenType, role, expireTimestamp)

	if rtcTokenErr != nil {
		log.Println(rtcTokenErr) // token failed to generate
		c.Error(rtcTokenErr)
		errMsg := "Error Generating RTC token - " + rtcTokenErr.Error()
		c.AbortWithStatusJSON(400, gin.H{
			"status": 400,
			"error":  errMsg,
		})
	} else if rtmTokenErr != nil {
		log.Println(rtmTokenErr) // token failed to generate
		c.Error(rtmTokenErr)
		errMsg := "Error Generating RTC token - " + rtmTokenErr.Error()
		c.AbortWithStatusJSON(400, gin.H{
			"status": 400,
			"error":  errMsg,
		})
	} else {
		log.Println("RTC Token generated")
		c.JSON(200, gin.H{
			"rtcToken": rtcToken,
			"rtmToken": rtmToken,
		})
	}

}
