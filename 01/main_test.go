package main

import (
	"testing"
)

func TestCaptcha(t *testing.T) {
	tt := []struct {
		name    string
		digits  string
		captcha int
	}{
		{"case 1122", "1122", 3},
		{"case 1111", "1111", 4},
		{"case 1234", "1234", 0},
		{"case 91212129", "91212129", 9},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			c := Captcha(tc.digits)
			if c != tc.captcha {
				t.Fatalf("Captcha should be %d, got: %d", tc.captcha, c)
			}

		})

	}
}

func TestCaptchaRing(t *testing.T) {
	tt := []struct {
		name    string
		digits  string
		captcha int
	}{
		{"case 1122", "1122", 3},
		{"case 1111", "1111", 4},
		{"case 1234", "1234", 0},
		{"case 91212129", "91212129", 9},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			c := CaptchaRing(tc.digits)
			if c != tc.captcha {
				t.Fatalf("Captcha should be %v, got: %d", tc.captcha, c)
			}

		})

	}
}

// eof
