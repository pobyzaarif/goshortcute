package goshortcute

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
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

// StringtoSHA256Hash : hash text string to SHA256 string.
// input : hello | output : 2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824
func StringtoSHA256Hash(text string) string {
	textByte := []byte(text)
	hash := sha256.Sum256(textByte)
	return fmt.Sprintf("%x", hash[:])
}
