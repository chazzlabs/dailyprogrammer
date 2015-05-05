package main

import (
    "fmt"
    "bytes"
    "os"
)

var dictionary map[rune][]string

func createDictionary() {
    dictionary = make(map[rune][]string)

    dictionary['0'] = []string{"", ""}
    dictionary['1'] = []string{"eleventy", "-one"}
    dictionary['2'] = []string{"twenty", "-two"}
    dictionary['3'] = []string{"thirty", "-three"}
    dictionary['4'] = []string{"fourty", "-four"}
    dictionary['5'] = []string{"fifty", "-five"}
    dictionary['6'] = []string{"sixty", "-six"}
    dictionary['7'] = []string{"seventy", "-seven"}
    dictionary['8'] = []string{"eighty", "-eight"}
    dictionary['9'] = []string{"ninety", "-nine"}
    dictionary['A'] = []string{"atta", "-a"}
    dictionary['B'] = []string{"bibbity", "-bee"}
    dictionary['C'] = []string{"city", "-cee"}
    dictionary['D'] = []string{"dickety", "-dee"}
    dictionary['E'] = []string{"ebbity", "-ee"}
    dictionary['F'] = []string{"fleventy", "-eff"}
}

func pronounce(hexValue string) string {
    var result bytes.Buffer

    hexValue = hexValue[2:]

    for index, letter := range hexValue {
        if index == 2 {
            result.WriteString(" bitey ")
        }

        result.WriteString(dictionary[letter][index % 2])
    }

    return result.String()
}

func main() {
    createDictionary()

    inputString := os.Args[1]
    result := pronounce(inputString)

    fmt.Println(result)
}
