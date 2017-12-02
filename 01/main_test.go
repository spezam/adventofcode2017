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
				t.Errorf("Captcha should be %d, got: %d", tc.captcha, c)
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
				t.Errorf("Captcha should be %v, got: %d", tc.captcha, c)
			}

		})

	}
}

func TestHalfCaptcha(t *testing.T) {
	tt := []struct {
		name    string
		digits  string
		captcha int
	}{
		{"case 1212", "1212", 6},
		{"case 1221", "1221", 0},
		{"case 123425", "123425", 4},
		{"case 123123", "123123", 12},
		{"case 12131415", "12131415", 4},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			c := HalfCaptcha(tc.digits)
			if c != tc.captcha {
				t.Errorf("Captcha should be %v, got: %d", tc.captcha, c)
			}

		})

	}
}

// eof
