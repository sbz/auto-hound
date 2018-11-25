package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"os"
	"strings"
)

const (
	defaultOrgName = "Nitro"
	defaultOrgType = "public"
)

func checkError(err interface{}) {
	if err != nil {
		fmt.Printf("Error %v exiting...\n", err)
		os.Exit(1)
	}
}

func PrettyPrint(v interface{}) {
	b, _ := json.MarshalIndent(v, "", "  ")
	println(string(b))
}

func isEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

func lookupEnv(key string, defaultValue string) string {
	var ret string

	ret = os.Getenv(key)
	if isEmpty(ret) {
		return defaultValue
	}

	return ret
}

func topicMatches(repoTopics []string, wantedTopics []string) bool {
	for _, topic := range repoTopics {
		for _, wanted := range wantedTopics {
			if topic == wanted {
				return true
			}
		}
	}

	return false
}

func main() {
	var orgName string
	var orgType string
	var client *github.Client

	ctx := context.Background()
	if !isEmpty(orgToken) {
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: orgToken},
		)
		tc := oauth2.NewClient(ctx, ts)

		client = github.NewClient(tc)
	} else {
		client = github.NewClient(nil)
	}

	orgName = lookupEnv("ORG_NAME", defaultOrgName)
	orgType = lookupEnv("ORG_TYPE", defaultOrgType)

	opts := &github.RepositoryListByOrgOptions{Type: orgType, ListOptions: github.ListOptions{PerPage: 50}}
	var allRepos []*github.Repository

	for {
		repos, response, err := client.Repositories.ListByOrg(ctx, orgName, opts)
		checkError(err)

		allRepos = append(allRepos, repos...)
		if response.NextPage == 0 {
			break
		}
		opts.Page = response.NextPage
	}

	config := make(map[string]interface{})
	config["max-concurrent-indexers"] = 2
	config["dbpath"] = "data"
	configRepos := map[string]map[string]string{}
	for _, repo := range allRepos {
		if *repo.Fork || *repo.Archived {
			continue
		}
		println(*repo.Name)
		container := map[string]string{"url": *repo.CloneURL}
		configRepos[*repo.Name] = container
	}
	//os.Exit(1)
	config["repos"] = configRepos

	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", strings.Repeat(" ", 4))
	encoder.Encode(config)

	os.Exit(0)
}
