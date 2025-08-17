package luhn

import (
    "strings"
    "regexp"
    "strconv"
    )

func Valid(id string) bool {
    cleanId := strings.ReplaceAll(id, " ", "")
    matched, _ := regexp.MatchString(`^\d+$`, cleanId)
	str := []rune(cleanId)
    if len(str) <= 1 || !matched {
        return false
    }
    sum,num := 0, 0
    for i, _ := range str {
        conv, _ := strconv.Atoi(string(str[i]))
        if (len(str)-i)%2 == 0 {
            num = conv * 2
            if num > 9 {
                num = num - 9
            }
        } else {
            num = conv
        }
        sum += num
    } 
    return sum % 10 == 0
}
