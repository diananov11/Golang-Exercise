package raindrops

import (
    "strconv"
    "strings"
    )

func Convert(number int) string {
	var word strings.Builder

    if number % 3 == 0 {
        word.WriteString("Pling")
    }

    if number % 5 == 0 {
        word.WriteString("Plang")
    }

    if number % 7 == 0 {
        word.WriteString("Plong")
    }

    if word.Len() == 0 {
        word.WriteString(strconv.Itoa(number))
    }

    return word.String()
}
