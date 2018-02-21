package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/go-github/github"
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

func main() {
	var orgName string
	var orgType string

	ctx := context.Background()
	client := github.NewClient(nil)

	orgName = lookupEnv("ORG_NAME", defaultOrgName)
	orgType = lookupEnv("ORG_TYPE", defaultOrgType)
	opts := &github.RepositoryListByOrgOptions{Type: orgType}

	repos, _, err := client.Repositories.ListByOrg(ctx, orgName, opts)
	checkError(err)

	config := make(map[string]interface{})
	config["max-concurrent-indexers"] = 2
	config["dbpath"] = "data"
	configRepos := map[string]map[string]string{}
	for _, repo := range repos {
		if *repo.Fork {
			continue
		}
		container := map[string]string{"url": *repo.CloneURL}
		configRepos[*repo.Name] = container
	}
	config["repos"] = configRepos

	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", strings.Repeat(" ", 4))
	encoder.Encode(config)

	os.Exit(0)
}
