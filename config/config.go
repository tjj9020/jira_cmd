package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type JiraConfig struct {
	JiraURL       string `mapstructure:"jira_url"`
	JiraUserName  string `mapstructure:"jira_user_name"`
	JiraUserToken string
}

func GetConfig() (JiraConfig, error) {
	var config JiraConfig

	configName := "jiracmd"
	viper.SetConfigName(configName)
	viper.AddConfigPath(fmt.Sprintf("$HOME/.jiracmd/%s", configName))
	viper.AddConfigPath(".")
	viper.AddConfigPath("./test_data/")

	err := viper.ReadInConfig()
	if err != nil {
		return JiraConfig{}, err
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		return JiraConfig{}, err
	}
	token, err := getJiraUserToken()
	if err != nil {
		return JiraConfig{}, errors.New("Could not get a jira user token")
	}
	config.JiraUserToken = token

	return config, nil
}

func getJiraUserToken() (string, error) {
	if os.Getenv("JIRA_USER_TOKEN") == "" {
		return "", errors.New("No token defined in config or in Environment \"JIRA_USER_TOKEN\"")
	}
	envToken := os.Getenv("JIRA_USER_TOKEN")
	return envToken, nil
}
