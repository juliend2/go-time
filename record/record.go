package record

import (
	"fmt"
	"regexp"
	"desrosiers.org/gotime/time"
)

type Record struct {
	content string
}

func NewRecord(content string) Record {
	return Record{content}
}

func (r Record) String() string {
	return r.content
}

func (r Record) getTimeRange() string {
	re := regexp.MustCompile(`^(.*)\s*:`)
	return string(re.Find([]byte(r.String())))
}

func (r Record) GetMinutes() int {
	re := regexp.MustCompile(`(\d+h\d+)\s*-\s*(\d+h\d+)`)
	matches := re.FindAllSubmatch([]byte(r.getTimeRange()), 2)
	fmt.Printf("%v\n", string(matches[0][1]))
 	from := time.NewTime(string(matches[0][1]))
 	to := time.NewTime(string(matches[0][2]))
	return to.GetMinutes() - from.GetMinutes()
}

