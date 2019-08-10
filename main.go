package main

import (
	"fmt"

	"github.com/tjj9020/jira_cmd/config"
)

func main() {
	config, err := config.GetConfig()
	if err != nil {
		fmt.Println("Could not get config")
	}
	fmt.Printf("JiraUrl: %v, JiraUser: %v, JiraToken: %v", config.JiraURL, config.JiraUserName, config.JiraUserToken)
	if err != nil {
		fmt.Println(err)
	}
}
