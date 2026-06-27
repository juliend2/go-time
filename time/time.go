package time

import (
	"regexp"
	"strconv"
)

type Time struct {
	content string
}

func NewTime(content string) Time {
	return Time{content}
}

func (t Time) GetMinutes() int {
	re := regexp.MustCompile(`(\d+)h(\d+)`)
	matches := re.FindAllSubmatch([]byte(t.content), 2)
	h, _ := strconv.Atoi(string(matches[0][1]))
	m, _ := strconv.Atoi(string(matches[0][2]))
	return (h * 60) + m
}


