package toposort

import (
	"fmt"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)
	onPath := make(map[string]bool)

	visitAll = func(items []string) {
		for _, item := range items {
			if onPath[item] {
				fmt.Printf("Cycle detected. Course %s detected twice", item)
				return
			}
			if !seen[item] {
				seen[item] = true
				onPath[item] = true
				visitAll(m[item])
				onPath[item] = false
				order = append(order, item)
			}
		}
	}

	allCourses := make(map[string]bool)
	for courses, dependencies := range m {
		allCourses[courses] = true
		for _, dep := range dependencies {
			allCourses[dep] = true
		}
	}

	var keys []string
	for key := range allCourses {
		keys = append(keys, key)
	}
	visitAll(keys)
	return order
}
