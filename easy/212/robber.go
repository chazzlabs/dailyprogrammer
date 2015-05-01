package main

import (
    "fmt"
    "unicode"
    "os"
    "strings"
)

const Consonants = "bcdfghjklmnpqrstvwxyz"

func encode(decodedString string) string {
    var encodedString []rune

    for _, letter := range decodedString {
        letter = unicode.ToLower(letter)

        encodedString = append(encodedString, letter)

        if strings.ContainsRune(Consonants, letter) {
            encodedString = append(encodedString, 'o')
            encodedString = append(encodedString, letter)
        }
    }

    return string(encodedString)
}

func main() {
    inputString := os.Args[1]
    encodedString := encode(inputString)
    fmt.Println(encodedString)
}
