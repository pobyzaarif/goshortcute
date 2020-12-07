package goshortcute

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTranslator(t *testing.T) {
	// Lists of normal case scenarios.
	listTest := map[string]string{
		// Case|Transalator_function : Input|Output
		"Test1|MD5":           "ThisIsTest1|1fb81916b94ae73ddd71ac6fcf5a6e01",
		"Test2|Base64Encode":  "ThisIsTest2|VGhpc0lzVGVzdDI=",
		"Test3|Base64Decode":  "VGhpc0lzVGVzdDM=|ThisIsTest3",
		"Test4|SHA256":        "ThisIsTest4|f9964fc0c93157234071446069c72b0d571918f6d737f30054adc7ba516db380",
		"Test5|AESCBCEncrypt": "ThisIsTest5#testkeytestkeyyy|Y+sHtwHdsy9CgOzEo7pgHQ==",
		"Test6|AESCBCDecrypt": "+LRW8mPVLSgPEBMi47gZRQ==#testkeytestkeyyy|ThisIsTest6",
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
		case "AESCBCEncrypt":
			t := strings.Split(inout[0], "#")
			out, _ = AESCBCEncrypt([]byte(t[0]), []byte(t[1]))
		case "AESCBCDecrypt":
			t := strings.Split(inout[0], "#")
			out, _ = AESCBCDecrypt([]byte(t[0]), []byte(t[1]))
		}

		assert.Equal(t, inout[1], out)
	}
}

func TestTranslatorPKCS7(t *testing.T) {
	listTest := []string{
		"Test1",
		"Test1234567890123456",
		"Test1234567890123456789012345678901234567",
	}

	for _, v := range listTest {
		pad := PKCS7Padding([]byte(v), 16)
		unpad := PKCS7UnPadding(pad)
		assert.Equal(t, v, string(unpad))
	}
}
