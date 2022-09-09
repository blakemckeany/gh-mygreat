package main

import (
	"fmt"
	"github.com/cli/go-gh"
	"log"
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
	// Ask the user for the owner of the repository
	// and the name of the repository.
	fmt.Print("Enter the owner of the repository (FROM):")
	var ownerFrom string
	fmt.Scanln(&ownerFrom)
	fmt.Print("Enter the name of the repository (FROM):")
	var repoFrom string
	fmt.Scanln(&repoFrom)

	// Ask the user for the owner of the repository
	// and the name of the repository. fmt.Print("Enter the owner of the repository (TO):")
	fmt.Print("Enter the owner of the repository (TO):")
	var ownerTo string
	fmt.Scanln(&ownerTo)
	fmt.Print("Enter the name of the repository (TO):")
	var repoTo string
	fmt.Scanln(&repoTo)

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

			fmt.Println(labels)
			labels = labels[:len(labels)-1]
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
		fmt.Println("Issue created: " + issue.Title)
	}
	// client, err := gh.RESTClient(nil) // if err != nil {
	// 	log.Fatal(err)
	// }
	// for _, issue := range issues {
	// 	fmt.Println(issue.Title)
	// 	_, err := client.Issues.Create(owner, repo, &gh.IssueRequest{Title: issue.Title, Body: issue.Body})
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }
	// client, err := gh.RESTClient(nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for _, issue := range issues {
	// 	// Create io.Reader for the issue body.
	// 	var buf bytes.Buffer
	// 	err := json.NewEncoder(&buf).Encode(issue)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	client.Post("/repos/"+owner+"/"+repo+"/issues", &buf, nil)
	// }
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
