package main

import (
    "fmt"
    "os"
    "unicode"
    "strings"
)

const Consonants = "BCDGFHJKLMNPQRSTVWXYZbcdfghjklmnpqrstvwxyz"

func encode(decodedString string) string {
    var encodedString []rune

    for _, letter := range decodedString {
        encodedString = append(encodedString, letter)

        if strings.ContainsRune(Consonants, letter) {
            encodedString = append(encodedString, 'o')
            encodedString = append(encodedString, unicode.ToLower(letter))
        }
    }

    return string(encodedString)
}

func main() {
    inputString := os.Args[1]
    encodedString := encode(inputString)
    fmt.Println(encodedString)
}
