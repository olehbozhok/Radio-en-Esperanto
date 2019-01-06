package tgbothandlers

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

const (
	subscribeCmd    = "subs"
	unsubscribeCmd  = "unSubs"
	toSubskribePage = "toSubsPage"
)

func subscribeOn(channelID [md5.Size]byte) string {
	return subscribeCmd + ":" + hex.EncodeToString(channelID[:])
}

func unsubscribeOn(channelID [md5.Size]byte) string {
	return unsubscribeCmd + ":" + hex.EncodeToString(channelID[:])
}

func md5StrToByte(data string) [md5.Size]byte {
	var md5Id [md5.Size]byte
	dataByte, _ := hex.DecodeString(data)

	for i, b := range dataByte {
		if i > (md5.Size - 1) {
			break
		}
		md5Id[i] = b
	}
	return md5Id
}

func moveToPageSubsribe(page int) string {
	return toSubskribePage + ":" + fmt.Sprintf("%d", page)
}

func parseInlineData(data string) (cmd, payload string, err error) {
	if data == "_" || data == "" {
		// nothing to do
		return
	}
	spl := strings.Split(data, ":")
	if len(spl) != 2 {
		err = fmt.Errorf("error parseInlineData data: %s", data)
		return
	}
	cmd = spl[0]
	payload = spl[1]
	return
}
