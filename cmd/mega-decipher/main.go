package main

import (
	"os"
	"fmt"
	"log"

	decipher "github.com/yistLin/mega_link_decipher"
)

func main() {
	url := os.Args[1]
	fmt.Println("Ciphered link =>", url)

	decipheredUrl, err := decipher.DecipherUrl(url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Original link =>", decipheredUrl)
}
