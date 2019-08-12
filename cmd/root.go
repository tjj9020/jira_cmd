package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var jiraTopic string       //comment
var jiraAction string      //add
var jiraTicket string      //1234
var jiraCommentText string //holds comment text

func init() {

}

var rootCmd = &cobra.Command{
	Use:   "jiracmd",
	Short: "Tool for interacting with Jira on the command-line",
	Long:  `Tool for interacting with Jira on the command-line`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
