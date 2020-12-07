package goshortcute

import (
	"fmt"
	"math/rand"
	"strconv"
)

// RandomNumber func to get random number in x digits
func RandomNumber(digit int) string {
	digit--
	pattern := "%0" + strconv.Itoa(digit) + "d"

	low := "1" + fmt.Sprintf(pattern, 0)
	lowInt, _ := strconv.Atoi(low)

	hi := low + "0"
	hiInt, _ := strconv.Atoi(hi)

	randNumber := lowInt + rand.Intn(hiInt-lowInt)
	return strconv.Itoa(randNumber)
}

// RandomAlphaNumeric func to get random alphanumeric in x digits
func RandomAlphaNumeric(digit int) string {
	var letterRunes = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, digit)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// RandomAlphaNumericSpecialCharacter func to get random alphanumeric and special character in x digits
func RandomAlphaNumericSpecialCharacter(digit int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ~!@#$%^&*()_+=`:.,\"[]{}\\|")
	b := make([]rune, digit)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
