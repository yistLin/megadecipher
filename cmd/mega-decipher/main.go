package main

import (
	"os"
	"fmt"

	decipher "github.com/yistLin/mega_link_decipher"
)

func main() {
	url := os.Args[1]
	fmt.Println("Ciphered link =>", url)

	decipheredUrl := decipher.DecipherUrl(url)
	fmt.Println("Original link =>", decipheredUrl)
}
