package main

import (
	"fmt"
	"log"
	"os"

	decipher "github.com/yistLin/megadecipher"
)

func main() {
	url := os.Args[1]
	fmt.Println("Ciphered link =>", url)

	decipheredUrl, err := decipher.Decipher(url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Original link =>", decipheredUrl)
}
