package goshortcute

import (
	"fmt"
	"math/rand"
	"strconv"
)

// RandomNumber func to get random number in x digits
func RandomNumber(digit int) string {
	if digit <= 0 {
		return ""
	}

	// special-case 1-digit numbers: 1..9
	if digit == 1 {
		return strconv.Itoa(rand.Intn(9) + 1)
	}

	digit--
	pattern := "%0" + strconv.Itoa(digit) + "d"

	low := "1" + fmt.Sprintf(pattern, 0)
	lowInt, err := strconv.Atoi(low)
	if err != nil {
		return ""
	}

	hi := low + "0"
	hiInt, err := strconv.Atoi(hi)
	if err != nil {
		return ""
	}

	rangeVal := hiInt - lowInt
	if rangeVal <= 0 {
		return strconv.Itoa(lowInt)
	}

	randNumber := lowInt + rand.Intn(rangeVal)
	return strconv.Itoa(randNumber)
}

// RandomAlphaNumeric func to get random alphanumeric in x digits
func RandomAlphaNumeric(digit int) string {
	var letterRunes = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	if digit <= 0 {
		return ""
	}

	b := make([]rune, digit)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// RandomAlphaNumericSpecialCharacter func to get random alphanumeric and special character in x digits
func RandomAlphaNumericSpecialCharacter(digit int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ~!@#$%^&*()_+=`:.,\"[]{}\\|")
	if digit <= 0 {
		return ""
	}

	b := make([]rune, digit)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
