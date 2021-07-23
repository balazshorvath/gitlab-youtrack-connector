package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/webhooks/v6/gitlab"
	"gitlab-youtrack-connector/youtrack"
)

func handle(hook *gitlab.Webhook, youtrack *youtrack.YouTrack, ytProjectId string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		payload, err := hook.Parse(r, gitlab.IssuesEvents)
		if err != nil {
			ok(w)
			return
		}
		switch payload.(type) {
		case gitlab.IssueEventPayload:
			issue := payload.(gitlab.IssueEventPayload)
			if issue.ObjectAttributes.Action != "open" {
				ok(w)
				return
			}
			err = youtrack.CreateIssue(
				fmt.Sprintf("%d", issue.ObjectAttributes.IID),
				issue.ObjectAttributes.Title,
				issue.ObjectAttributes.Description,
				issue.ObjectAttributes.URL,
				ytProjectId,
			)
			if err != nil {
				fail(w, err)
				return
			}
		}
		ok(w)
	}
}

func fail(w http.ResponseWriter, err error) {
	log.Printf("[Server] Request resulted in error: %v", err)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	_, _ = w.Write([]byte(`{"status":"ERROR"}`))
}
func ok(w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{"status":"OK"}`))
}
