package utils

import (
	"regexp"
	"strings"
)

type Query struct {
	Owner string
	Name  string
}

func ExamineQuery(query string) Query {
	// Pattern: [flag]:[value] name
	var finalQuery Query

	flagsPattern := regexp.MustCompile(`(\w+):(\w+)`)
	flags := flagsPattern.FindAllStringSubmatch(query, -1)
	pureQuery := flagsPattern.ReplaceAllString(query, "")

	for _, flag := range flags {
		switch flag[1] {
		case "owner":
			finalQuery.Owner = flag[2]
		}
	}

	if finalQuery.Owner == "" {
		finalQuery.Owner = pureQuery
	}

	finalQuery.Name = strings.TrimSpace(pureQuery)

	return finalQuery
}
