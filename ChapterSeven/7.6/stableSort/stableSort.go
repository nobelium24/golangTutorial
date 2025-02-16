package stablesort

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
	City string
}

var people = []Person{
	{"Alice", 25, "New York"},
	{"Bob", 30, "Los Angeles"},
	{"Charlie", 20, "Chicago"},
	{"David", 30, "New York"},
	{"Eve", 25, "Los Angeles"},
	{"Frank", 20, "Chicago"},
	{"Grace", 30, "New York"},
	{"Heidi", 25, "Los Angeles"},
	{"Ivan", 20, "Chicago"},
}

type ByAge []Person

func (a ByAge) Len() int {
	return len(a)
}

func (a ByAge) Swap(i int, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByAge) Less(i int, j int) bool {
	return a[i].Age < a[j].Age
}

type ByName []Person

func test() {
	sort.Stable(ByAge(people))
	fmt.Print(people)
}
