package isogram

import (
    "strings"
    "unicode"
)

func IsIsogram(word string) bool {
	replacer := strings.NewReplacer("-", "", " ", "")
    cleanStr := replacer.Replace(word)
    wordMap := make(map[rune]int)

    for _, v := range cleanStr {       
        wordMap[unicode.ToLower(v)]++
    }

    for _, v := range wordMap {
        if v > 1 {
            return false
        }
    }
    return true
}
