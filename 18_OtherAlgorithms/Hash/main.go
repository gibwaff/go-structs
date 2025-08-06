package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

// Задание 1
func md5Hash(s string) string {
	hash := md5.Sum([]byte(s))
	return hex.EncodeToString(hash[:])
}

func same5lastChars() (string, string) {
	seen := make(map[string]string)
	for i := 0; ; i++ {
		s := fmt.Sprintf("str%d", i)
		h := md5Hash(s)
		last5 := h[len(h)-5:]
		if val, ok := seen[last5]; ok && val != s {
			return val, s
		}
		seen[last5] = s
	}
}

// Задание 2
func md5Key(secret string) []byte {
	sum := md5.Sum([]byte(secret))
	return sum[:]
}

func pkcs7Pad(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padtext...)
}

func encrypt(strToEncrypt, secret string) (string, error) {
	key := md5Key(secret)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	plaintext := pkcs7Pad([]byte(strToEncrypt), block.BlockSize())
	iv := key[:block.BlockSize()] // Для примера используем часть ключа как IV
	mode := cipher.NewCBCEncrypter(block, iv)
	ciphertext := make([]byte, len(plaintext))
	mode.CryptBlocks(ciphertext, plaintext)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// Задание 3
func pkcs7Unpad(data []byte) []byte {
	length := len(data)
	unpadding := int(data[length-1])
	return data[:(length - unpadding)]
}

func decrypt(strToDecrypt, secret string) (string, error) {
	key := md5Key(secret)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	iv := key[:block.BlockSize()]
	data, err := base64.StdEncoding.DecodeString(strToDecrypt)
	if err != nil {
		return "", err
	}
	mode := cipher.NewCBCDecrypter(block, iv)
	plaintext := make([]byte, len(data))
	mode.CryptBlocks(plaintext, data)
	plaintext = pkcs7Unpad(plaintext)
	return string(plaintext), nil
}

func main() {
	// Задание 1
	s1, s2 := same5lastChars()
	fmt.Println("String1:", s1)
	fmt.Println("String2:", s2)
	fmt.Println("MD5-1:", md5Hash(s1))
	fmt.Println("MD5-2:", md5Hash(s2))

	// Задание 2
	enc, _ := encrypt("mypassword", "mysecret")
	fmt.Println("Encrypted:", enc)

	// Задание 3
	dec, _ := decrypt(enc, "mysecret")
	fmt.Println("Decrypted:", dec)
}
