package config

import (
	"log"
	"os"
)

type Service struct {
	Port        string    `yaml:"port"`
	GitlabToken string    `yaml:"gitlab-token"`
	YouTrack    *YouTrack `yaml:"you-track"`
}

type YouTrack struct {
	Token     string `yaml:"token"`
	Url       string `yaml:"url"`
	ProjectId string `yaml:"project-id"`
}

func NewEnv() *Service {
	return &Service{
		GitlabToken: getEnvOrPanic("GITLAB_TOKEN"),
		Port:        getEnvOrPanic("PORT"),
		YouTrack: &YouTrack{
			Token:     getEnvOrPanic("YT_TOKEN"),
			Url:       getEnvOrPanic("YT_URL"),
			ProjectId: getEnvOrPanic("YT_PROJECT_ID"),
		},
	}
}

func getEnvOrPanic(name string) string {
	value := os.Getenv(name)
	if value == "" {
		log.Fatalf("failed to get env var %s", name)
	}
	return value
}
