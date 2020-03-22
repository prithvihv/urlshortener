package main

import (
	"crypto/sha256"
	"encoding/base64"
	"math/rand"
	"strings"
)

func hash(url string) string {
	h := sha256.New()
	h.Write([]byte(url))
	byteString := h.Sum(nil)
	stringEncoded := base64.StdEncoding.EncodeToString(byteString[:])
	return stringEncoded
}

func randomStringOfSize(l int, s string) string {
	var str strings.Builder
	for c := 0; c != 7; c++ {
		char := pickRandomChar(s) // post issue here
		println(char)
		str.WriteString(char)
	}
	return str.String()
}

func pickRandomChar(s string) string {
	charIndex := random(0, len(s))
	return string(s[charIndex])
}

var r = rand.New(rand.NewSource(99))

func random(min, max int) int {
	return r.Intn(max-min) + min
}
