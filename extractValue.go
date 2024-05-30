package pkg

import (
	"regexp"
	"strings"
)

func ExtractValue(html, pattern string) string {
	re := regexp.MustCompile(pattern)
	match := re.FindStringSubmatch(html)
	if len(match) >= 2 {
		return strings.TrimSpace(match[1])
	}
	return ""
}

func ExtractMultipleValues(html, pattern string) []string {
	re := regexp.MustCompile(pattern)
	match := re.FindStringSubmatch(html)
	if len(match) >= 2 {
		value := match[1]
		values := strings.Split(value, "<br>")
		for i := 0; i < len(values); i++ {
			values[i] = strings.TrimSpace(values[i])
		}
		return values
	}
	return []string{}
}
