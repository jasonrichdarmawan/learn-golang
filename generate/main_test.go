package main

import (
	"crypto/rand"
	"testing"
)

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

func TestGenerateOTP(t *testing.T) {
	iteration := 365
	m := make(map[string]int)

	for i := 0; i < iteration; i++ {
		otp, err := generateOTP()
		if err != nil {
			t.Fatal(err)
		}
		m[otp] += 1
		if m[otp] > 4 {
			t.Fatalf("otp %s is already existed for %d times", otp, m[otp])
		}
	}

	if len(m) < iteration*95/100 {
		t.Fatalf("with %dth iteration, expected len m: %d actual len m: %d",
			iteration, iteration*95/100, len(m))
	}
}
