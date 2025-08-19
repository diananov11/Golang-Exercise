package isogram

import (
    "strings"
    "unicode"
)

func IsIsogram(word string) bool {
    lowerWord := strings.ToLower(word)
	for _, char := range lowerWord {
        if !unicode.IsLetter(char) {
            continue
        }

        if strings.Count(lowerWord, string(char)) > 1 {
            return false
        }
    }
    return true
}
