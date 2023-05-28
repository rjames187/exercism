package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	pattern := `^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`
	re := regexp.MustCompile(pattern)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	pattern := "<([~*=-]+)?>"
	re := regexp.MustCompile(pattern)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	pattern := `"(?i).*password.*"`
	re := regexp.MustCompile(pattern)
	count := 0
	for _, line := range lines {
		if re.MatchString(line) {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	pattern := "end-of-line[0-9]+"
	re := regexp.MustCompile(pattern)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	res := []string{}

	for _, line := range lines {
		re := regexp.MustCompile("User ")
		if !re.MatchString(line) {
			res = append(res, line)
			continue
		}
		re = regexp.MustCompile("[ ]+")
		split := re.Split(line, -1)
		for i, chunk := range split {
			if chunk == "User" {
				user := split[i+1]
				newStr := fmt.Sprintf("[USR] %s %s", user, line)
				res = append(res, newStr)
				break
			}
		}
	}

	return res
}
