package goshortcute

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
)

// StringtoMD5Hash : encode text string to md5 string.
// Input	: "ThisIsTest1"
// Output	: "1fb81916b94ae73ddd71ac6fcf5a6e01"
func StringtoMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	encode := hex.EncodeToString(hasher.Sum(nil))
	return encode
}

// StringtoBase64Encode : encode text string to base64 string.
// Input	: "ThisIsTest2"
// Output	: "VGhpc0lzVGVzdDI="
func StringtoBase64Encode(text string) string {
	encode := base64.StdEncoding.EncodeToString([]byte(text))
	return string(encode)
}

// StringtoBase64Decode : decode base64 string to text string.
// Input	: "VGhpc0lzVGVzdDM="
// Output	: "ThisIsTest3"
func StringtoBase64Decode(text string) string {
	decode, _ := base64.StdEncoding.DecodeString(string(text))
	return string(decode)
}

// StringtoSHA256Hash : hash text string to SHA256 string.
// Input 	: "hello"
// Output : "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824"
func StringtoSHA256Hash(text string) string {
	textByte := []byte(text)
	hash := sha256.Sum256(textByte)
	return fmt.Sprintf("%x", hash[:])
}

// PKCS7Padding : ...
func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	if blockSize == 0 {
		blockSize = aes.BlockSize
	}
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// PKCS7UnPadding : ...
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// AESCBCEncrypt : AES Encryption function
// Input 	: []byte("1234567890")
// Key   	: []byte("testtesttesttesttesttesttesttest")
// Output : "XyFYQ7yRapfsYEwYpqEuZA=="
// Error	: nil
func AESCBCEncrypt(plainText, key []byte) (encText string, err error) {
	encText = string(plainText)
	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	blockSize := block.BlockSize()
	plainText = PKCS7Padding(plainText, blockSize)
	ciphertext := make([]byte, blockSize+len(plainText))

	mode := cipher.NewCBCEncrypter(block, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	mode.CryptBlocks(ciphertext[blockSize:], plainText)

	encText = base64.StdEncoding.EncodeToString(ciphertext[blockSize:])

	return
}

// AESCBCDecrypt : AES Decryption function
// Input 	: []byte("XyFYQ7yRapfsYEwYpqEuZA==")
// Key   	: []byte("testtesttesttesttesttesttesttest")
// Output : "1234567890"
// Error	: nil
func AESCBCDecrypt(encText, key []byte) (plainText string, err error) {
	plainText = string(encText)
	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	encText, err = base64.StdEncoding.DecodeString(plainText)
	if err != nil {
		return
	}

	blockSize := block.BlockSize()

	if len(encText) < blockSize {
		err = errors.New("ciphertext too short")
		return
	}

	if len(encText)%blockSize != 0 {
		err = errors.New("ciphertext is not a multiple of the block size")
		return
	}

	mode := cipher.NewCBCDecrypter(block, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	decrypted := make([]byte, len(encText))
	mode.CryptBlocks(decrypted, encText)
	encText = PKCS7UnPadding(decrypted)

	plainText = string(encText)

	return
}
