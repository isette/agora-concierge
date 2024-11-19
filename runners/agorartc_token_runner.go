package runners

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/isette/agora.io-service/handlers"
	"github.com/isette/agora.io-service/utils"
)

func GetRtcToken(c *gin.Context) {
	log.Printf("rtc token\n")
	channelName, tokentype, uidStr, role, expireTimestamp, err := utils.ParseRtcParams(c)

	if err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(400, gin.H{
			"message": "Error Generating RTC token: " + err.Error(),
			"status":  400,
		})
		return
	}

	rtcToken, tokenErr := handlers.GenerateRtcToken(channelName, uidStr, tokentype, role, expireTimestamp)

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
