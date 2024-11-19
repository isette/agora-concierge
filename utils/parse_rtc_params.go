package utils

import (
	"fmt"
	"strconv"
	"time"

	rtctokenbuilder2 "github.com/AgoraIO-Community/go-tokenbuilder/rtctokenbuilder"
	"github.com/gin-gonic/gin"
)

func ParseRtcParams(c *gin.Context) (channelName, tokentype, uidStr string, role rtctokenbuilder2.Role, expireTimestamp uint32, err error) {
	channelName = c.Param("channelName")
	roleStr := c.Param("role")
	tokentype = c.Param("tokentype")
	uidStr = c.Param("uid")
	expireTime := c.DefaultQuery("expiry", "3600")

	if roleStr == "publisher" {
		role = rtctokenbuilder2.RolePublisher
	} else {
		role = rtctokenbuilder2.RoleSubscriber
	}

	expireTime64, parseErr := strconv.ParseUint(expireTime, 10, 64)
	if parseErr != nil {
		err = fmt.Errorf("failed to parse expireTime: %s, causing error: %s", expireTime, parseErr)
	}

	expireTimeInSeconds := uint32(expireTime64)
	currentTimestamp := uint32(time.Now().UTC().Unix())
	expireTimestamp = currentTimestamp + expireTimeInSeconds

	return channelName, tokentype, uidStr, role, expireTimestamp, err
}
