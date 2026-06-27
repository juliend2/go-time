package main

import (
	"fmt"
	"regexp"
	"desrosiers.org/gotime/record"
)

const input string = `
4h34 - 4h44 : project 1
4h44 - 5h01 : project 1
`


func parse(s string) {
	var records []record.Record
	re := regexp.MustCompile(`(.+)\n`)
	found := re.FindAllSubmatch([]byte(input), -1)
	for _, line := range found {
		records = append(records, record.NewRecord(string(line[1])))
	}
	for _, rec := range records {
		fmt.Println("rec", rec.GetMinutes())
	}
}

func main() {
	parse(input)
}


