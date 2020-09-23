package goshortcute

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTranslator(t *testing.T) {

	// Lists of normal case scenarios.
	listTest := map[string]string{
		// Case|Transalator_function : Input|Output
		"Tes1|MD5":          "ThisIsTest1|1fb81916b94ae73ddd71ac6fcf5a6e01",
		"Tes2|Base64Encode": "ThisIsTest2|VGhpc0lzVGVzdDI=",
		"Tes3|Base64Decode": "VGhpc0lzVGVzdDM=|ThisIsTest3",
		"Tes4|SHA256":       "ThisIsTest4|f9964fc0c93157234071446069c72b0d571918f6d737f30054adc7ba516db380",
	}

	var out string
	for k, v := range listTest {
		testcase := strings.Split(k, "|")
		inout := strings.Split(v, "|")

		switch testcase[1] {
		case "MD5":
			out = StringtoMD5Hash(inout[0])
		case "SHA256":
			out = StringtoSHA256Hash(inout[0])
		case "Base64Encode":
			out = StringtoBase64Encode(inout[0])
		case "Base64Decode":
			out = StringtoBase64Decode(inout[0])
		}

		assert.Equal(t, inout[1], out, fmt.Sprintf("%s matched to translate %s", testcase[0], testcase[1]))
	}
}
