package mega_link_decipher

import (
	"fmt"
	"regexp"
)

func DecipherUrl(url string) string {
	re := regexp.MustCompile("^mega://enc\\?([A-Za-z0-9-_,]+)")

	if matched := re.MatchString(url); matched {
		matches := re.FindStringSubmatch(url)
		fmt.Printf("%q\n", matches)
	}

	return url
}
