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

func TestResponseTranslator(t *testing.T) {
	responseCode := 401
	response := responseTranslator(responseCode)
	if response != "There was a problem, authentication failed, recieved: 401" {
		t.Errorf("Wrong reponse received, got: %s", response)
	}

	responseCode = 404
	response = responseTranslator(responseCode)
	if response != "There was a problem, check the ticket number, recieved: 404" {
		t.Errorf("Wrong reponse received, got: %s", response)
	}

	responseCode = 201
	response = responseTranslator(responseCode)
	if response != "Success!" {
		t.Errorf("Wrong reponse received, got: %s", response)
	}
}

func TestNewJiraComment(t *testing.T) {
	body := "test"
	ticket := "1234"
	myConfig := config.JiraConfig{
		JiraURL:       "https://url.com",
		JiraUserName:  "user",
		JiraUserToken: "token",
	}

	jiraComment := NewJiraComment(body, ticket, myConfig)
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
