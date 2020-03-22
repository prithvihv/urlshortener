package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"math/rand"
	"time"
)

func hash(url string) string {
	h := sha256.New()
	h.Write([]byte(url))
	byteString := h.Sum(nil)
	stringEncoded := base64.StdEncoding.EncodeToString(byteString[:])
	return stringEncoded
}

func randomStringOfSize(l int, s string) string {
	var buffer bytes.Buffer
	for len(s) != 7 {
		buffer.WriteString(pickRandomChar(s))
	}
	return buffer.String()
}

func pickRandomChar(s string) string {
	charIndex := random(0, len(s))
	return string(s[charIndex])
}

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
