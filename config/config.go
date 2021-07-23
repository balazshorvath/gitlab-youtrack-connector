package config

import (
	"io/ioutil"

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

func New(path string) *Service {
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
