package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	srv "github.com/balazshorvath/go-srv"
	"github.com/go-playground/webhooks/v6/gitlab"
	"gitlab-youtrack-connector/config"
	"gitlab-youtrack-connector/youtrack"
)

type server struct {
	srv.BasicHttpServer
	Config *config.Service
}

func (s *server) Init() {
	service, err := gitlab.New()
	if err != nil {
		panic(err)
	}
	mux := http.NewServeMux()
	yt := youtrack.New(s.Config.YouTrack.Url, s.Config.YouTrack.Token)
	// Routing
	mux.HandleFunc("/gitlab/issue", handle(service, yt, s.Config.YouTrack.ProjectId))
	s.Srv = &http.Server{
		Addr:    fmt.Sprintf(":%s", s.Config.Port),
		Handler: mux,
	}
}

func New(ctx context.Context, group *sync.WaitGroup) srv.Server {
	return &server{
		BasicHttpServer: srv.BasicHttpServer{
			BasicServer: srv.BasicServer{
				Ctx:   ctx,
				Group: group,
			},
		},
		Config: config.NewEnv(),
	}
}
