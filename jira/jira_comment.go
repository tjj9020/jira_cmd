package jira

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	http "net/http"

	config "github.com/tjj9020/jira_cmd/config"
)

type JiraComment struct {
	Body   string
	Ticket string
	Config config.JiraConfig
}

func NewJiraComment(body string, ticket string, config config.JiraConfig) JiraComment {
	return JiraComment{
		Body:   body,
		Ticket: ticket,
		Config: config,
	}
}

func (j JiraComment) AddComment() (string, error) {
	dataMap := createBodyMap(j.Body)
	dataReader, err := convertToReader(dataMap)
	if err != nil {
		return "", err
	}

	client := http.Client{}
	url := fmt.Sprintf("%s/rest/api/2/issue/%s/comment", j.Config.JiraURL, j.Ticket)
	req, err := http.NewRequest("POST", url, dataReader)
	req.SetBasicAuth(j.Config.JiraUserName, j.Config.JiraUserToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	myResponse := responseTranslator(resp.StatusCode)
	return myResponse, nil
}

func responseTranslator(resp int) string {
	switch resp {
	case 401:
		myResponse := fmt.Sprintf("There was a problem, authentication failed, recieved: %v", resp)
		return myResponse
	case 404:
		myResponse := fmt.Sprintf("There was a problem, check the ticket number, recieved: %v", resp)
		return myResponse
	}
	return "Success!"
}

func createBodyMap(body string) map[string]string {
	data := make(map[string]string)
	data["body"] = body
	return data
}

func convertToReader(bodyMap map[string]string) (io.Reader, error) {
	dataBytes, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, err
	}
	dataReader := bytes.NewReader(dataBytes)
	return dataReader, nil
}
