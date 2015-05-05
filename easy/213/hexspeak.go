package main

import (
    "fmt"
    "bytes"
    "os"
)

type Order struct {
    High, Low string
}

var dictionary map[rune]Order

func createDictionary() {
    dictionary = make(map[rune]Order)

    dictionary['0'] = Order{"", ""}
    dictionary['1'] = Order{"eleventy", "-one"}
    dictionary['2'] = Order{"twenty", "-two"}
    dictionary['3'] = Order{"thirty", "-three"}
    dictionary['4'] = Order{"fourty", "-four"}
    dictionary['5'] = Order{"fifty", "-five"}
    dictionary['6'] = Order{"sixty", "-six"}
    dictionary['7'] = Order{"seventy", "-seven"}
    dictionary['8'] = Order{"eighty", "-eight"}
    dictionary['9'] = Order{"ninety", "-nine"}
    dictionary['A'] = Order{"atta", "-a"}
    dictionary['B'] = Order{"bibbity", "-bee"}
    dictionary['C'] = Order{"city", "-cee"}
    dictionary['D'] = Order{"dickety", "-dee"}
    dictionary['E'] = Order{"ebbity", "-ee"}
    dictionary['F'] = Order{"fleventy", "-eff"}
}

func pronounce(hexValue string) string {
    var result bytes.Buffer

    hexValue = hexValue[2:]

    for index, letter := range hexValue {
        switch index {
        case 0:
            result.WriteString(dictionary[letter].High)
        case 1:
            result.WriteString(dictionary[letter].Low)
        case 2:
            result.WriteString(" bitey " + dictionary[letter].High)
        case 3:
            result.WriteString(dictionary[letter].Low)
        }
    }

    return result.String()
}

func main() {
    createDictionary()

    inputString := os.Args[1]
    result := pronounce(inputString)

    fmt.Println(result)
}
