package handlers

import (
	"fmt"
	"log"
	"strconv"

	rtctokenbuilder2 "github.com/AgoraIO-Community/go-tokenbuilder/rtctokenbuilder"
	rtmtokenbuilder2 "github.com/AgoraIO-Community/go-tokenbuilder/rtmtokenbuilder"
)

func GenerateRtcToken(channelName, uidStr, tokenType string, role rtctokenbuilder2.Role, expireTimestamp uint32) (rtcToken string, err error) {

	if tokenType == "userAccount" {
		log.Printf("Building Token with userAccount: %s\n", uidStr)
		rtcToken, err = rtctokenbuilder2.BuildTokenWithAccount(
			GetEnvWithKey("AGORA_APP_ID"),
			GetEnvWithKey("AGORA_APP_CERTIFICATE"),
			channelName,
			uidStr,
			role,
			expireTimestamp,
		)
		return rtcToken, err

	} else if tokenType == "uid" {
		uid64, parseErr := strconv.ParseUint(uidStr, 10, 64)
		if parseErr != nil {
			err = fmt.Errorf("failed to parse uidStr: %s, to uint causing error: %s", uidStr, parseErr)
			return "", err
		}

		uid := uint32(uid64)
		log.Printf("Building Token with uid: %d\n", uid)
		rtcToken, err = rtctokenbuilder2.BuildTokenWithUid(
			GetEnvWithKey("AGORA_APP_ID"),
			GetEnvWithKey("AGORA_APP_CERTIFICATE"),
			channelName,
			uid,
			role,
			expireTimestamp,
		)

		return rtcToken, err

	} else {
		err = fmt.Errorf("failed to generate RTC token for Unknown Tokentype: %s", tokenType)
		log.Println(err)
		return "", err
	}
}

func GenerateRteToken(channelName, uidStr, tokenType string, role rtctokenbuilder2.Role, expireTimestamp uint32) (rtcToken string, err error) {

	if tokenType == "userAccount" {
		log.Printf("Building Token with userAccount: %s\n", uidStr)
		rtcToken, err = rtctokenbuilder2.BuildTokenWithAccount(
			GetEnvWithKey("AGORA_APP_ID"),
			GetEnvWithKey("AGORA_APP_CERTIFICATE"),
			channelName,
			uidStr,
			role,
			expireTimestamp,
		)
		return rtcToken, err

	} else if tokenType == "uid" {
		uid64, parseErr := strconv.ParseUint(uidStr, 10, 64)
		if parseErr != nil {
			err = fmt.Errorf("failed to parse uidStr: %s, to uint causing error: %s", uidStr, parseErr)
			return "", err
		}

		uid := uint32(uid64)
		log.Printf("Building Token with uid: %d\n", uid)
		rtcToken, err = rtctokenbuilder2.BuildTokenWithUid(
			GetEnvWithKey("AGORA_APP_ID"),
			GetEnvWithKey("AGORA_APP_CERTIFICATE"),
			channelName,
			uid,
			role,
			expireTimestamp,
		)

		return rtcToken, err

	} else {
		err = fmt.Errorf("failed to generate RTC token for Unknown Tokentype: %s", tokenType)
		log.Println(err)
		return "", err
	}
}

func GenerateRtmToken(channelName, uidStr, tokenType string, role rtctokenbuilder2.Role, expireTimestamp uint32) (rtcToken string, err error) {
	rtmToken, tokenErr := rtmtokenbuilder2.BuildToken(
		GetEnvWithKey("AGORA_APP_ID"),
		GetEnvWithKey("AGORA_APP_CERTIFICATE"),
		string(uidStr),
		expireTimestamp,
		"",
	)

	log.Printf("uidStr: %s, expireTimestamp: %d", string(uidStr), expireTimestamp)
	log.Printf("appID: %s, appCertificate: %s", GetEnvWithKey("AGORA_APP_ID"), GetEnvWithKey("AGORA_APP_CERTIFICATE"))
	log.Printf("Generated RTM Token: %s", rtmToken)

	return rtmToken, tokenErr

}
