package utils

import "github.com/go-resty/resty/v2"

func RepoExist(repo string) bool {
	client := resty.New()

	resp, err := client.R().Get("https://api.github.com/repos/" + repo)

	if err != nil {
		return false
	}

	if resp.StatusCode() == 200 {
		return true
	} else {
		return false
	}
}
