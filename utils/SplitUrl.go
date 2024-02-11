package utils

import "strings"

func SplitUrl(url string) (string, string, string) {
	split := strings.Split(url, "/")

	domain := split[2]
	owner := split[3]
	name := split[4]

	return domain, owner, name
}
