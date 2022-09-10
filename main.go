package main

import (
	"fmt"
	"github.com/cli/go-gh"
	"github.com/fatih/color"
	"log"
	"strings"
)

// Label struct
type Label struct {
	Id          int
	Name        string
	Description string
	Color       string
}

// Assignee struct
type Assignee struct {
	Id    int
	Login string
	Name  string
}

// Issue struct
type Issue struct {
	Title     string
	Body      string
	Labels    []Label
	Assignees []Assignee
}

func main() {

	fmt.Print("Enter the owner/repository (FROM): ")
	var ownerRepoFrom string
	fmt.Scanln(&ownerRepoFrom)

	fmt.Print("Enter the owner/repository (TO): ")
	var ownerRepoTo string
	fmt.Scanln(&ownerRepoTo)

	ownerFrom := ownerRepoFrom[0:strings.Index(ownerRepoFrom, "/")]
	repoFrom := ownerRepoFrom[strings.Index(ownerRepoFrom, "/")+1 : len(ownerRepoFrom)]

	ownerTo := ownerRepoTo[0:strings.Index(ownerRepoTo, "/")]
	repoTo := ownerRepoTo[strings.Index(ownerRepoTo, "/")+1 : len(ownerRepoTo)]

	issues := GetIssues(ownerFrom, repoFrom)
	CreateIssues(ownerTo, repoTo, issues)
}

// Create the issues in the repository
func CreateIssues(owner string, repo string, issues []Issue) {
	for _, issue := range issues {
		labels := ""
		for _, label := range issue.Labels {
			labels += fmt.Sprintf("%s,", label.Name)
			if label.Name == "" {
				continue
			}
			labelArgs := []string{"label", "create", label.Name, "--description", label.Description, "--color", label.Color, "--force", "-R", owner + "/" + repo}
			_, _, err := gh.Exec(labelArgs...)
			if err != nil {
				log.Fatal(err)
			}

		}
		if len(labels) > 0 {

			labels = labels[:len(labels)-1]
			fmt.Println("Issue labels: " + labels)
		}

		assignees := ""
		for _, assignee := range issue.Assignees {
			assignees += fmt.Sprintf("%s,", assignee.Login)
		}

		if len(assignees) > 0 {
			assignees = assignees[:len(assignees)-1]
		}

		fmt.Println("Creating issue: " + issue.Title)
		issueArgs := []string{"issue", "create", "-R", owner + "/" + repo, "-t", issue.Title, "-b", issue.Body, "-l", labels, "-a", assignees}
		_, _, err := gh.Exec(issueArgs...)
		if err != nil {
			log.Fatal(err)
		}
		color.Green("Issue created: " + issue.Title)
	}
}

// Get a list of issues from a repository
func GetIssues(owner string, repo string) []Issue {
	client, err := gh.RESTClient(nil)
	if err != nil {
		log.Fatal(err)
	}
	var issues []Issue
	err = client.Get("repos/"+owner+"/"+repo+"/issues", &issues)
	if err != nil {
		log.Fatal(err)
	}
	return issues
}
