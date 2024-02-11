package utils

import (
	"errors"

	"github.com/go-resty/resty/v2"
)

// This feature will be subject to change as more hosts are supported
// The check happens through a request to the repo url
// This change comes true at the support of new hosts

type RepoResponse struct {
	StargazersCount int `json:"stargazers_count"`
}

func IsEligible(owner string, name string) (bool, error) {
	client := resty.New()

	resp, err := client.R().
		SetResult(&RepoResponse{}).
		Get("https://api.github.com/repos/" + owner + "/" + name)

	if err != nil {
		return false, errors.New("Error while fetching repository")
	}

	if resp.StatusCode() == 200 {
		repoResponse := resp.Result().(*RepoResponse)

		if repoResponse.StargazersCount > 200 {
			return true, nil
		} else {
			return false, errors.New("Repository has less than 100 stars")
		}
	} else {
		return false, errors.New("Repository not found")
	}
}
