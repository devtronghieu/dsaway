package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	s := "A man, a plan, a canal: Panama"
	assert.Equal(t, true, isPalindrome(s))

	s = "race a car"
	assert.Equal(t, false, isPalindrome(s))

	s = " "
	assert.Equal(t, true, isPalindrome(s))
}

func isUppercase(char byte) bool {
	return char >= 65 && char <= 90
}

func isLowercase(char byte) bool {
	return char >= 97 && char <= 122
}

func isNumeric(char byte) bool {
	return char >= 48 && char <= 57
}

func isAlphanumeric(char byte) bool {
	if isUppercase(char) || isLowercase(char) || isNumeric(char) {
		return true
	}
	return false
}

func toLowercase(char byte) byte {
	if isUppercase(char) {
		return char + 32
	}
	return char
}

func isPalindrome(s string) bool {
	leftPointer := 0
	rightPointer := len(s) - 1

	for rightPointer > leftPointer {
		if !isAlphanumeric(s[leftPointer]) {
			leftPointer++
			continue
		}

		if !isAlphanumeric(s[rightPointer]) {
			rightPointer--
			continue
		}

		if toLowercase(s[leftPointer]) != toLowercase(s[rightPointer]) {
			return false
		}
		leftPointer++
		rightPointer--
	}

	return true
}
