package main

import (
	"fmt"
	"net/http"
	"os"
	"io/ioutil"
	"strings"
	"github.com/andygrunwald/go-jira"
	"encoding/json"
)

const GITHUBQL_URL = "https://api.github.com/graphql"

type SampleRepoQuery struct {
	Data struct {
		Viewer struct {
			Login string `json:"login"`
			StarredRepositories struct {
				TotalCount int `json:"totalCount"`
			} `json:"starredRepositories"`
			Repositories struct {
				Edges []struct {
					Node struct {
						Name string `json:"name"`
						Stargazers struct {
							TotalCount int `json:"totalCount"`
						} `json:"stargazers"`
						Forks struct {
							TotalCount int `json:"totalCount"`
						} `json:"forks"`
						Watchers struct {
							TotalCount int `json:"totalCount"`
						} `json:"watchers"`
						Issues struct {
							TotalCount int `json:"totalCount"`
						} `json:"issues"`
					} `json:"node"`
				} `json:"edges"`
			} `json:"repositories"`
		} `json:"viewer"`
	} `json:"data"`
}

const A_QUERY = `
{
	viewer {
		login
		starredRepositories {
			totalCount
		}
		repositories(first: 3) {
			edges {
				node {
					name
					stargazers {
						totalCount
					}
					forks {
						totalCount
					}
					watchers {
						totalCount
					}
					issues(states:[OPEN]) {
						totalCount
					}
				}
			}
		}
	}
}`

func graphql_query(query string) string {
	r := strings.NewReplacer("\n", " ", "\t", " ")
	my_q := "{\"query\": \"" + r.Replace(query) + "\"}"
	return my_q
}

func main() {
	/*
	We want dag to (I think) take an argument:
	- org/Repositories
	- desired output format and location. Let's start with GFM and a file
	- Uses GitHub GraphQL https://developer.github.com/v4/
	 */
	username := os.Getenv("FURIOSA_JIRA_USERNAME")
	password := os.Getenv("FURIOSA_JIRA_PASSWORD")
	jiraUrl := os.Getenv("FURIOSA_JIRA_URL")
	jiraClient, err := jira.NewClient(nil, jiraUrl)
	if err != nil {
		panic(err)
	}
	jiraClient.Authentication.SetBasicAuth(username, password)

	issue, _, err := jiraClient.Issue.Get("A15-10390", nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s: %+v\n", issue.Key, issue.Fields.Summary)

	fmt.Println("Dag reports about repositories")
	auth_token := os.Getenv("FURIOSA_GITHUB_TOKEN")
	req, err := http.NewRequest("POST", GITHUBQL_URL,
		strings.NewReader(graphql_query(A_QUERY)))
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

	var s SampleRepoQuery
	json.Unmarshal(body, &s)
	fmt.Println("Body:", string(body))
	fmt.Printf("%v\n", s)
	fmt.Printf("login: %v\n", s.Data.Viewer.Login)
	fmt.Printf("Number of starred repositories: %v\n", s.Data.Viewer.StarredRepositories.TotalCount)
}
