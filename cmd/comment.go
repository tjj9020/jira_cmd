package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tjj9020/jira_cmd/config"
	"github.com/tjj9020/jira_cmd/jira"
)

//var jiraTopic string       //comment
//var jiraAction string      //add
var jiraTicket string      //1234
var jiraCommentText string //holds comment text

func init() {
	rootCmd.AddCommand(commentAddCmd)
	commentAddCmd.Flags().StringVarP(&jiraTicket, "ticket-identifier", "t", "", "Ticket identifier")
	commentAddCmd.MarkFlagRequired("ticket")
	commentAddCmd.Flags().StringVarP(&jiraCommentText, "comment", "c", "", "Comment surrounded by quotes")
	commentAddCmd.MarkFlagRequired("comment")
}

var commentAddCmd = &cobra.Command{
	Use:   "comment add",
	Short: "jira action for comments",
	Long:  `jira action for comments`,
	Run: func(cmd *cobra.Command, args []string) {
		// comment := "test comment"
		// ticket := "PBT-15163"
		config, err := config.GetConfig()
		jiraComment := jira.NewJiraComment(jiraCommentText, jiraTicket, config)
		response, err := jiraComment.AddComment()
		if err != nil {
			fmt.Printf("There was an error, %v", err)
		}
		fmt.Println(response)
	},
}
