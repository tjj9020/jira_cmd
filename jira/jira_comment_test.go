package jira

import "testing"

func TestCreateBodyMap(t *testing.T) {
	body := "My Body"
	result := createBodyMap(body)
	if result["body"] != "My Body" {
		t.Errorf("Map body should have been \"My Body\", got %v", result)
	}

}
