package main

import (
	"flag"
	"fmt"
	"strings"
	"unicode"
)

const consonants = "BCDGFHJKLMNPQRSTVWXZbcdfghjklmnpqrstvwxz"

func encode(inputString string) string {
	var encodedString []rune

	for _, letter := range inputString {
		encodedString = append(encodedString, letter)

		if strings.ContainsRune(consonants, letter) {
			encodedString = append(encodedString, 'o')
			encodedString = append(encodedString, unicode.ToLower(letter))
		}
	}

	return string(encodedString)
}

func decode(inputString string) string {
	var decodedString []rune
	inputRunes := []rune(inputString)

	for index := 0; index < len(inputRunes); index++ {
		letter := inputRunes[index]

		decodedString = append(decodedString, letter)

		if strings.ContainsRune(consonants, letter) {
			index += 2
		}
	}

	return string(decodedString)
}

func main() {
	var decodeFlag = flag.Bool("d", false, "Usage: '-d <message>'")
	var encodeFlag = flag.Bool("e", false, "Usage: '-e <message>'")
	flag.Parse()

	if *decodeFlag {
		result := decode(flag.Args()[0])
		fmt.Println(result)
		return
	} else if *encodeFlag {
		result := encode(flag.Args()[0])
		fmt.Println(result)
		return
	}

	flag.PrintDefaults()
}
