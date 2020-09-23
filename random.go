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
