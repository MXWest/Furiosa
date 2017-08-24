package main

import (
	"fmt"
	"net/http"
	"os"
	"io/ioutil"
	"strings"
	"github.com/andygrunwald/go-jira"
)

const GITHUBQL_URL = "https://api.github.com/graphql"



func graphql_query(query string) string {
	return "{\"query\": \"" + query + "\"}"
}

func main() {
	/*
	We want dag to (I think) take an argument:
	- org/Repositories
	- desired output format and location. Let's start with GFM and a file
	- Uses GitHub GraphQL https://developer.github.com/v4/
	 */

	jiraClient, err := jira.NewClient(nil, "https://urbnit.atlassian.net/")
	if err != nil {
		panic(err)
	}
	username := os.Getenv("FURIOSA_JIRA_USERNAME")
	password := os.Getenv("FURIOSA_JIRA_PASSWORD")
	jiraClient.Authentication.SetBasicAuth(username, password)

	issue, _, err := jiraClient.Issue.Get("A15-10390", nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s: %+v\n", issue.Key, issue.Fields.Summary)

	fmt.Println("Dag reports about repositories")
	auth_token := os.Getenv("FURIOSA_GITHUB_TOKEN")
	req, err := http.NewRequest("POST", GITHUBQL_URL,
		strings.NewReader(graphql_query("{viewer{login}}")))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Authorization", "bearer "+auth_token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println("Status: ", resp.Status)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Body:", string(body))
}
