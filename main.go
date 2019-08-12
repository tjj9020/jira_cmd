package main

import "github.com/tjj9020/jira_cmd/cmd"

func main() {
	// config, err := config.GetConfig()
	// if err != nil {
	// 	fmt.Println("Could not get config")
	// }

	// comment := "test comment"
	// ticket := "PBT-15163"
	// jiraComment := jira.NewJiraComment(comment, ticket, config)
	// response, err := jiraComment.AddComment()
	// if err != nil {
	// 	fmt.Printf("There was an error, %v", err)
	// }
	// fmt.Println(response)
	cmd.Execute()
}
