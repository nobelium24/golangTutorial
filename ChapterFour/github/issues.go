package github

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	var lessThanAMonth []Issue
	var lessThanAYear []Issue
	var moreThanAYear []Issue
	setTime := time.Now()

	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d Issues found:\n", result.TotalCount)

	for _, item := range result.Items {
		if item.CreatedAt.Before(setTime.AddDate(0, -1, 0)) {
			lessThanAMonth = append(lessThanAMonth, *item)
		} else if item.CreatedAt.Before(setTime.AddDate(-1, 0, 0)) {
			lessThanAYear = append(lessThanAYear, *item)
		} else {
			moreThanAYear = append(moreThanAYear, *item)
		}
	}

	fmt.Println("\nIssues less than a month old:")
	for _, issue := range lessThanAMonth {
		fmt.Printf("#%-5d %9.9s %.55s\n", issue.Number, issue.User.Login, issue.Title)
	}

	fmt.Println("\nIssues less than a year old:")
	for _, issue := range lessThanAYear {
		fmt.Printf("#%-5d %9.9s %.55s\n", issue.Number, issue.User.Login, issue.Title)
	}

	fmt.Println("\nIssues more than a year old:")
	for _, issue := range moreThanAYear {
		fmt.Printf("#%-5d %9.9s %.55s\n", issue.Number, issue.User.Login, issue.Title)
	}
}
