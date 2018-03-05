package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"fmt"
	"log"
)

func isSignatureMatching(appSecret string, bytes []byte, expectedSignature string) bool {
	log.Println(appSecret, expectedSignature)
	mac := hmac.New(sha1.New, []byte(appSecret))
	mac.Write(bytes)

	log.Println(fmt.Sprintf("Calculated secret : %x", mac.Sum(nil)))

	return fmt.Sprintf("%x", mac.Sum(nil)) == expectedSignature
}
