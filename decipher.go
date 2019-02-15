package mega_link_decipher

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"encoding/hex"
	"encoding/base64"
)

func DecipherUrl(url string) string {
	re := regexp.MustCompile("^mega://enc2\\?([A-Za-z0-9-_,]+)")

	if matched := re.MatchString(url); matched {
		matches := re.FindStringSubmatch(url)
		fmt.Printf("%q\n", matches)

		str := matches[1]

		password := []byte{237, 31, 76, 32, 11, 53, 19, 152, 6, 178, 96, 86, 59, 61, 56, 118, 240, 17, 180, 117, 15, 58, 26, 74, 94, 253, 11, 190, 103, 85, 75, 68}
		decoded, _ := hex.DecodeString("79F10A01844A0B27FF5B2D4E0ED3163E")

		fmt.Printf("[ ")
		for _, s := range decoded {
			fmt.Printf("%d ", s)
		}
		fmt.Printf("]\n")

		// Recover Base64 from encoded string
		b64str := str
		b64str += ("==")[((2 - len(b64str) * 3) & 3):]
		fmt.Printf("%s\n", b64str)

		b64str = strings.Replace(b64str, "-", "+", -1)
		b64str = strings.Replace(b64str, "_", "/", -1)
		b64str = strings.Replace(b64str, ",", "", -1)
		fmt.Printf("%s\n", b64str)

		// Decoding step
		decoded, err := base64.StdEncoding.DecodeString(b64str)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d\n", len(decoded))

		warr := byteArrayToWordArray(decoded)
		fmt.Printf("[ ")
		for _, s := range warr {
			fmt.Printf("%d ", s)
		}
		fmt.Printf("]\n")

		key := byteArrayToWordArray(password)
		fmt.Printf("[ ")
		for _, s := range key {
			fmt.Printf("%d ", s)
		}
		fmt.Printf("]\n")
	}

	return url
}

func byteArrayToWordArray(arr []byte) []int32 {
	warr := make([]int32, (len(arr) + 1) / 4)
	for i := 0; i < len(arr); i++ {
		warr[(i / 4) | 0] |= int32(arr[i]) << uint32(24 - 8 * (i % 4))
	}
	return warr
}
