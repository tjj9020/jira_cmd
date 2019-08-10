package config

import (
	"os"
	"testing"
)

func TestGetConfig(t *testing.T) {
	myConfig, err := GetConfig()
	if err != nil {
		t.Errorf("Could not load configuration file: %v", err)
	}
	if myConfig.JiraURL != "https://someurl.atlassian.net" {
		t.Errorf("Expected jira_url to be \"https://someurl.atlassian.net\", got %v", myConfig.JiraURL)
	}
}

func TestGetJiraUserToken(t *testing.T) {
	os.Setenv("JIRA_USER_TOKEN", "1")
	token, err := getJiraUserToken()
	if err != nil {
		t.Errorf("Could not get token was: nil")
	}
	if token != "1" {
		t.Errorf("Expected token to be \"1\", got: %v", token)
	}

}
