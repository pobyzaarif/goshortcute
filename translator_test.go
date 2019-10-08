package shortcute

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
	}

	var out string
	for k, v := range listTest {
		testcase := strings.Split(k, "|")
		inout := strings.Split(v, "|")

		switch testcase[1] {
		case "Base64Encode":
			out = StringtoBase64Encode(inout[0])
		case "Base64Decode":
			out = StringtoBase64Decode(inout[0])
		default:
			out = StringtoMD5Hash(inout[0])
		}

		assert.Equal(t, inout[1], out, fmt.Sprintf("%s matched to translate %s", testcase[0], testcase[1]))
	}
}
