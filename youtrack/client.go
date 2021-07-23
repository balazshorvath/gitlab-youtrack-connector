package youtrack

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type YouTrack struct {
	url   string
	token string
}

func New(url, token string) *YouTrack {
	return &YouTrack{
		url:   url,
		token: token,
	}
}

func (y *YouTrack) CreateIssue(id, summary, description, project string) error {
	req, err := json.Marshal(
		&CreateIssueRequest{
			Summary:     fmt.Sprintf("#%s - %s", id, summary),
			Description: description,
			Project: struct {
				Id string `json:"id"`
			}{
				Id: project,
			},
		},
	)
	if err != nil {
		return err
	}
	r, err := http.NewRequest("POST", fmt.Sprintf("%s/api/issues", y.url), bytes.NewBuffer(req))
	if err != nil {
		return err
	}
	y.headers(r)
	client := http.Client{}
	response, err := client.Do(r)
	if err != nil {
		return err
	}
	respBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	log.Printf("[YouTrack] Got response (%d): %s\n", response.StatusCode, respBytes)
	if response.StatusCode != http.StatusOK {
		return errors.New("failed to create issue")
	}
	return nil
}

func (y *YouTrack) headers(r *http.Request) {
	r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", y.token))
	r.Header.Add("Accept", "application/json")
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Cache-Control", "no-cache")
}

type CreateIssueRequest struct {
	Summary     string `json:"summary"`
	Description string `json:"description"`
	Project     struct {
		Id string `json:"id"`
	} `json:"project"`
}
