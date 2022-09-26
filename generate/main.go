package main

import (
	"crypto/rand"
	"log"
)

func main() {
	log.Println(generateOTP())
}

func generateOTP() (string, error) {
	otp := make([]byte, 6)
	if _, err := rand.Read(otp); err != nil {
		return "", err
	}

	for i := 0; i < 6; i++ {
		otp[i] = 48 + (otp[i] % 10)
	}

	return string(otp), nil
}
