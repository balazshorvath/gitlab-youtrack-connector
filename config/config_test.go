package config

import (
	"io/ioutil"
	"testing"

	"gopkg.in/yaml.v2"
)

func TestConfigParsing(t *testing.T) {
	file, err := ioutil.ReadFile("config.yaml")
	assertErr(t, err, "could not open config file")
	conf := &Service{}
	err = yaml.Unmarshal(file, conf)
	assertErr(t, err, "could not unmarshal config")
	assertStr(t, conf.GitlabToken, "token", "GitlabToken")
	assert(t, conf.YouTrack != nil, "YouTrack property should not be empty")
	assertStr(t, conf.YouTrack.Url, "yturl", "YouTrack.Url")
	assertStr(t, conf.YouTrack.ProjectId, "projectId", "YouTrack.ProjectId")
	assertStr(t, conf.YouTrack.Token, "yttoken", "YouTrack.Token")
}

func assert(t *testing.T, condition bool, message string) {
	if condition {
		return
	}
	t.Fatalf("assertion failed: %s", message)
}

func assertStr(t *testing.T, result, expected, name string) {
	if result == expected {
		return
	}
	t.Fatalf("assertion failed (%s): %s is expected to be equal to %s", name, result, expected)
}

func assertErr(t *testing.T, err error, message string) {
	if err == nil {
		return
	}
	t.Fatalf("error occured: %s, error: %v", message, err)
}
