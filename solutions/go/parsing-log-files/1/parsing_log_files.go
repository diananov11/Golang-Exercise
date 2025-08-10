package parsinglogfiles
import 
("regexp"
 "fmt")

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
    return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[~*=-]*>`)
    return re.Split(text,-1)
}

func CountQuotedPasswords(lines []string) int {
	sum := 0
    re  := regexp.MustCompile(`"[^"]*(?i:password)[^"]*"`)
    for _, v := range lines {
        if re.MatchString(v) {
            sum++
        }
    }
    return sum
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line[\d*]*`)
    return re.ReplaceAllString(text,"")
}

func TagWithUserName(lines []string) []string {
    re := regexp.MustCompile(`User\s+([a-zA-Z]+\d*)`)
    
    for i, line := range lines {
        if matches := re.FindStringSubmatch(line); matches != nil {
            lines[i] = fmt.Sprintf("[USR] %s %s", matches[1], line)
        }
    }
    
    return lines
}

// func TagWithUserName(lines []string) []string {
//     re := regexp.MustCompile(`User\s+([a-zA-Z]+\d*)`)
//     result := make([]string, 0, len(lines))
    
//     for _, line := range lines {
//         if matches := re.FindStringSubmatch(line); matches != nil {
//             taggedLine := fmt.Sprintf("[USR] %s %s", matches[1], line)
//             result = append(result, taggedLine)
//         } else {
//             result = append(result, line)
//         }
//     }
    
//     return result
// }
