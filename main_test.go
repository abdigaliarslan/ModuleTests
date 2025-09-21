package main

import (
	"testing"
)

func TestValidEmail(t *testing.T) {
	tests := []struct {
		email string
		valid bool
	}{
		{"user@gmail.com", true},
		{"", false},
		{"user@.com", false},
		{"user@domain", false},
		{"abc@", false},
		{"wtf", false},
	}

	for _, test := range tests {
		result := ValidEmail(test.email)
		if result != test.valid {
			t.Errorf("ValidEmail(%q) = %v; want %v", test.email, result, test.valid)
		}
	}
}

func TestStrongPassword(t *testing.T) {
	tests := []struct {
		password string
		valid    bool
	}{
		{"A12345", false},
		{"Aaaa321", false},
		{"A@123!BC", true},
		{"ABC123", false},
		{"lolkek123!A", true},
		{"", false},
		{"ganjklDFAN432@!#vzsd", true},
	}

	for _, test := range tests {
		result := StrongPassword(test.password)
		if result != test.valid {
			t.Errorf("StrongPassword(%q) = %v; want %v", test.password, result, test.valid)
		}
	}
}
