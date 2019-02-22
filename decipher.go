package megadecipher

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"regexp"
)

var (
	// ErrInvalidURL indicates the url input is not valid.
	ErrInvalidURL = errors.New("invalid URL on input")

	// ErrInvalidPKCS7Padding indicates PKCS7 unpad fails to bad input.
	ErrInvalidPKCS7Padding = errors.New("invalid PKCS7 padding on input")
)

func Decipher(url string) (string, error) {
	re := regexp.MustCompile("^mega://(f?)enc(2?)\\?([A-Za-z0-9-_,]+)")

	matched := re.MatchString(url)
	if !matched {
		return "", ErrInvalidURL
	}

	matches := re.FindStringSubmatch(url)
	foldertag := matches[1]
	versiontag := matches[2]
	b64str := matches[3]

	// link is for folder or not
	rootUrl := "https://mega.nz/#"
	if foldertag == "f" {
		rootUrl += "F"
	}

	var key []byte
	if versiontag == "2" {
		key = []byte{237, 31, 76, 32, 11, 53, 19, 152, 6, 178, 96, 86, 59, 61, 56, 118, 240, 17, 180, 117, 15, 58, 26, 74, 94, 253, 11, 190, 103, 85, 75, 68}
	} else {
		password := []rune{'k', '1', 'o', '6', 'A', 'l', '-', '1', 'k', 'z', '?', '!', 'z', '0', '5', 'y'}
		for len(password) < 32 {
			password = append(password, 'X')
		}
		key = []byte(string(password))
	}

	iv, _ := hex.DecodeString("79F10A01844A0B27FF5B2D4E0ED3163E")

	// decode base64 encoded url
	b64str += ("==")[((2 - len(b64str)*3) & 3):]
	ciphertext, err := base64.URLEncoding.DecodeString(b64str)
	if err != nil {
		return "", err
	}

	// decipher
	block, _ := aes.NewCipher(key)
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)

	// unpad(pksc7)
	c := ciphertext[len(ciphertext)-1]
	n := int(c)
	if n == 0 || n > len(ciphertext) {
		return "", ErrInvalidPKCS7Padding
	}
	for i := 0; i < n; i++ {
		if ciphertext[len(ciphertext)-n+i] != c {
			return "", ErrInvalidPKCS7Padding
		}
	}

	deciphertext := string(ciphertext[:len(ciphertext)-n])

	return rootUrl + deciphertext, nil
}
