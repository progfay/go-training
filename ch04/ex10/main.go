package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/progfay/go-training/ch04/ex10/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	var (
		oneMonthBefore = time.Now().AddDate(0, -1, 0)
		oneYearBefore  = time.Now().AddDate(-1, 0, 0)

		lessThanOneMonthIssues = make([]*github.Issue, 0)
		lessThanOneYearIssues  = make([]*github.Issue, 0)
		moreThanOneYearIssues  = make([]*github.Issue, 0)
	)

	for _, item := range result.Items {
		switch {
		case item.CreatedAt.After(oneMonthBefore):
			lessThanOneMonthIssues = append(lessThanOneMonthIssues, item)

		case item.CreatedAt.After(oneYearBefore):
			lessThanOneYearIssues = append(lessThanOneYearIssues, item)

		default:
			moreThanOneYearIssues = append(moreThanOneYearIssues, item)
		}
	}

	fmt.Println("Issues created at less than a month:")
	for _, issue := range lessThanOneMonthIssues {
		fmt.Printf("  #%-5d %9.9s %.55s\n",
			issue.Number, issue.User.Login, issue.Title)
	}

	fmt.Println("\nIssues created at less than a year:")
	for _, issue := range lessThanOneYearIssues {
		fmt.Printf("  #%-5d %9.9s %.55s\n",
			issue.Number, issue.User.Login, issue.Title)
	}

	fmt.Println("\nIssue created over a year ago:")
	for _, issue := range moreThanOneYearIssues {
		fmt.Printf("  #%-5d %9.9s %.55s\n",
			issue.Number, issue.User.Login, issue.Title)
	}
}
