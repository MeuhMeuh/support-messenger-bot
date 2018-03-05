package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"fmt"
)

func isSignatureMatching(appSecret string, bytes []byte, expectedSignature string) bool {
	mac := hmac.New(sha1.New, []byte(appSecret))
	mac.Write(bytes)

	return fmt.Sprintf("%x", mac.Sum(nil)) == expectedSignature
}
