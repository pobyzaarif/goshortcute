package goshortcute

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
)

// StringtoMD5Hash : encode text string to md5 string.
// Input "ThisIsTest1"
// Output "1fb81916b94ae73ddd71ac6fcf5a6e01"
func StringtoMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	encode := hex.EncodeToString(hasher.Sum(nil))
	return encode
}

// StringtoBase64Encode : encode text string to base64 string.
// Input "ThisIsTest2"
// Output "VGhpc0lzVGVzdDI="
func StringtoBase64Encode(text string) string {
	encode := base64.StdEncoding.EncodeToString([]byte(text))
	return string(encode)
}

// StringtoBase64Decode : decode base64 string to text string.
// Input "VGhpc0lzVGVzdDM="
// Output "ThisIsTest3"
func StringtoBase64Decode(text string) string {
	decode, _ := base64.StdEncoding.DecodeString(string(text))
	return string(decode)
}
