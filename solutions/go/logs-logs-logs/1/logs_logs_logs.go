package logs

import "unicode/utf8"

// Application identifies the application emitting the given log.
func Application(log string) string {
	application := map[rune]string{
        '❗': "recommendation", 
        '🔍': "search", 
        '☀': "weather",}
    for _,v := range log {
        if char, ok := application[v]; ok {
            return char
        }
    }
    return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
    word := ""
	for _, v := range log {
        if v == oldRune {
            word += string(newRune)
   			continue
        }
        word+= string(v)
    }
    return word
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	numberOfRunes := utf8.RuneCountInString(log)
    return numberOfRunes <= limit
}
