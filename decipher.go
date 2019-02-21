package mega_link_decipher

import (
	"fmt"
	"errors"
	"regexp"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"encoding/base64"
)

var (
	// ErrInvalidURL indicates the url input is not valid.
	ErrInvalidURL = errors.New("invalid URL on input")

	// ErrInvalidPKCS7Padding indicates PKCS7 unpad fails to bad input.
	ErrInvalidPKCS7Padding = errors.New("invalid PKCS7 padding on input")
)

func DecipherUrl(url string) (string, error) {
	re := regexp.MustCompile("^mega://enc2\\?([A-Za-z0-9-_,]+)")

	matched := re.MatchString(url)
	if !matched {
		return "", ErrInvalidURL
	}

	matches := re.FindStringSubmatch(url)
	fmt.Printf("%q\n", matches)

	str := matches[1]

	password := []byte{237, 31, 76, 32, 11, 53, 19, 152, 6, 178, 96, 86, 59, 61, 56, 118, 240, 17, 180, 117, 15, 58, 26, 74, 94, 253, 11, 190, 103, 85, 75, 68}
	iv, _ := hex.DecodeString("79F10A01844A0B27FF5B2D4E0ED3163E")

	fmt.Printf("iv(decoded) => [ ")
	for _, s := range iv {
		fmt.Printf("%d ", s)
	}
	fmt.Printf("]\n")

	// Recover Base64 from encoded string
	b64str := str
	b64str += ("==")[((2 - len(b64str) * 3) & 3):]
	fmt.Printf("b64str(padded) => %s\n", b64str)

	// b64str = strings.Replace(b64str, "-", "+", -1)
	// b64str = strings.Replace(b64str, "_", "/", -1)
	// b64str = strings.Replace(b64str, ",", "", -1)
	// fmt.Printf("b64str(replaced) => %s\n", b64str)
	// fmt.Printf("length of b64str => %d\n", len(b64str))

	// Decoding step
	ciphertext, err := base64.URLEncoding.DecodeString(b64str)
	if err != nil {
		return "", err
	}

	// warr := byteArrayToWordArray(ciphertext)
	// fmt.Printf("ciphertext(word array) => [ ")
	// for _, s := range warr {
	// 	fmt.Printf("%d ", s)
	// }
	// fmt.Printf("]\n")

	// key := byteArrayToWordArray(password)
	// fmt.Printf("key(word array) => [ ")
	// for _, s := range key {
	// 	fmt.Printf("%d ", s)
	// }
	// fmt.Printf("]\n")

	// try cryptor/aes
	block, err := aes.NewCipher(password)
	if err != nil {
		return "", err
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)
	fmt.Printf("deciphered text => %s\n", ciphertext)
	fmt.Printf("deciphertext => [ ")
	for _, s := range ciphertext {
		fmt.Printf("%d ", s)
	}
	fmt.Printf("]\n")

	// unpad(pksc7)
	c := ciphertext[len(ciphertext) - 1]
	n := int(c)
	if n == 0 || n > len(ciphertext) {
		return "", ErrInvalidPKCS7Padding
	}
	for i := 0; i < n; i++ {
		if ciphertext[len(ciphertext)-n+i] != c {
			return "", ErrInvalidPKCS7Padding
		}
	}

	deciphered := string(ciphertext[:len(ciphertext)-n])

	return "https://mega.nz/#" + deciphered, nil
}

// func byteArrayToWordArray(arr []byte) []int32 {
// 	warr := make([]int32, (len(arr) + 1) / 4)
// 	for i := 0; i < len(arr); i++ {
// 		warr[(i / 4) | 0] |= int32(arr[i]) << uint32(24 - 8 * (i % 4))
// 	}
// 	return warr
// }
