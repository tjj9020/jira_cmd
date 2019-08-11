package jira

import (
	"testing"

	config "github.com/tjj9020/jira_cmd/config"
)

func TestCreateBodyMap(t *testing.T) {
	body := "My Body"
	result := createBodyMap(body)
	if result["body"] != "My Body" {
		t.Errorf("Map body should have been \"My Body\", got %v", result)
	}
}

func TestSaveComment(t *testing.T) {
	body := "test"
	ticket := "1234"
	myConfig := config.JiraConfig{
		JiraURL:       "https://url.com",
		JiraUserName:  "user",
		JiraUserToken: "token",
	}

	jiraComment := SaveComment(body, ticket, myConfig)
	if myConfig.JiraURL != "https://url.com" {
		t.Errorf("JiraURL should have been \"https://url.com\", got: %v", jiraComment.Config.JiraURL)
	}

	if jiraComment.Body != "test" {
		t.Errorf("Body should have been \"test\", got: %v", jiraComment.Body)
	}

	if jiraComment.Ticket != "1234" {
		t.Errorf("Ticket should have been \"1234\", got: %v", jiraComment.Ticket)
	}
}
