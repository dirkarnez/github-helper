package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/google/go-github/v41/github"
)

var (
	repoName string
)

/*
	@echo off

	set USER_NAME=dirkarnez
	set PASSWORD={PAT}

	github-helper.exe --repo 1 2 3
	pause
*/
func main() {
	flag.StringVar(&repoName, "repo", "", "repo name(s)")
	flag.Parse()

	tail := flag.Args()
	listOfRepos := tail
	listOfRepos = append(listOfRepos, "")  // Making space for the new element
	copy(listOfRepos[1:], listOfRepos[0:]) // Shifting elements
	listOfRepos[0] = repoName              // Copying/inserting the value

	var userName = os.Getenv("USER_NAME")
	var password = os.Getenv("PASSWORD")

	tp := github.BasicAuthTransport{
		Username: strings.TrimSpace(userName),
		Password: strings.TrimSpace(password),
	}

	client := github.NewClient(tp.Client())

	for _, s := range listOfRepos {
		fmt.Printf("creating: \"%s\"\n", s)

		var private = true
		var autoInit = false

		repo := github.Repository{
			Name:     &s,
			Private:  &private,
			AutoInit: &autoInit,
		}

		_, _, err := client.Repositories.Create(context.Background(), "", &repo)

		if err != nil {
			fmt.Printf("Error when creating \"%s\": %v\n", s, err)
			return
		} else {
			fmt.Printf("Succesfully created\"%s\"\n", s)
		}
	}
}
