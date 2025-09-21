package main

import (
	"regexp"
	// idk why, but go's regexp package doesn't support lookaheads
	"github.com/dlclark/regexp2"
)

func ValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

func StrongPassword(password string) bool {
	pattern := regexp2.MustCompile(`(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&]).{8,}`, 0)
	ans, _ := pattern.MatchString(password)
	return ans
}
