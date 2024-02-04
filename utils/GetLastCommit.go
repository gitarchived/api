package utils

import "github.com/go-resty/resty/v2"

type CommitApiResponse struct {
	Sha string `json:"sha"`
}

func GetLastCommit(repo string) (string, error) {
	client := resty.New()

	resp, err := client.R().
		SetHeader("Accept", "application/vnd.github.v3+json").
		SetResult([]CommitApiResponse{}).
		Get("https://api.github.com/repos/" + repo + "/commits")

	if err != nil {
		return "", err
	}

	result := resp.Result().(*[]CommitApiResponse)
	commit := (*result)[0].Sha

	if resp.StatusCode() == 200 {
		return commit, nil
	} else {
		return "", nil
	}
}
