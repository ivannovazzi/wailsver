package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Release struct {
	TagName   string `json:"tag_name"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	Url       string `json:"html_url"`
	Current   bool
}

func getReleases(config *Config) ([]Release, error) {
	token := config.GitHubToken
	owner := config.Owner
	repo := config.Repo
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases", owner, repo)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch releases: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var releases []Release
	err = json.Unmarshal(body, &releases)
	if err != nil {
		return nil, err
	}

	return releases, nil
}
