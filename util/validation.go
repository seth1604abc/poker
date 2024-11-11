package util

import (
	"fmt"
	"regexp"
	"unicode"
)

func IsContain(str string, upper bool, lower bool, number bool) bool {
	var (
		hasUpper bool
		hasLower bool
		hasNumber bool
	)

	for _, char := range str {
		switch {
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsDigit(char):
			hasNumber = true
		}
	}

	return (!upper || hasUpper) && (!lower || hasLower) && (!number || hasNumber)
}

func IsValidEmail(email string) bool{
	re := regexp.MustCompile(`^[a-zA-Z0-9.]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	if !re.MatchString(email) {
		return false
	}

	return true
}

func IsValidAccount(account string) bool {
	re := regexp.MustCompile(`[a-zA-Z0-9]{6,20}$`)

	if !re.MatchString(account) {
		return false
	}

	return true
}

func IsValidPassword(password string) bool {
	re, err := regexp.Compile(`^[a-zA-Z\d]{10,20}$`)

	if err != nil {
		fmt.Println(err)
		return false
	}

	if !re.MatchString(password) {
		// fmt.Println("not fit")
		return false
	}

	if contain := IsContain(password, true, true, true); !contain {
		return false
	}

	return true
}