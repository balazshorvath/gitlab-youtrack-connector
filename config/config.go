package config

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Service struct {
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
		YouTrack: &YouTrack{
			Token:     getEnvOrPanic("YT_TOKEN"),
			Url:       getEnvOrPanic("YT_URL"),
			ProjectId: getEnvOrPanic("YT_PROJECT_ID"),
		},
	}
}

func NewFile(path string) *Service {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	conf := &Service{}
	err = yaml.Unmarshal(content, conf)
	if err != nil {
		panic(err)
	}
	return conf
}

func getEnvOrPanic(name string) string {
	value := os.Getenv(name)
	if value == "" {
		log.Fatalf("failed to get env var %s", name)
	}
	return value
}
