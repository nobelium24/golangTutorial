package sort

import (
	"encoding/json"
	"html/template"
	"io"
	"net/http"
	"sort"
)

type Person struct {
	Name string
	Age  int
	City string
}

type Tracker struct {
	columns []int
}

type Table struct {
	table   []Person
	tracker Tracker
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

const (
	NameColumn = iota
	AgeColumn
	CityColumn
)

func (t Table) Len() int {
	return len(t.table)
}

func (t Table) Less(i int, j int) bool {
	// return t.table[i].Age < t.table[j].Age
	for _, col := range t.tracker.columns {
		if col == NameColumn {
			if t.table[i].Name != t.table[j].Name {
				return t.table[i].Name < t.table[j].Name
			}
		} else if col == AgeColumn {
			if t.table[i].Age != t.table[j].Age {
				return t.table[i].Age < t.table[j].Age
			}
		} else if col == CityColumn {
			if t.table[i].City != t.table[j].City {
				return t.table[i].City < t.table[j].City
			}
		}
	}
	return false
}

func (t Table) Swap(i int, j int) {
	t.table[i], t.table[j] = t.table[j], t.table[i]
}

func (t *Table) Update(column int) {
	for i, col := range t.tracker.columns {
		if col == column {
			t.tracker.columns = append(t.tracker.columns[:i], t.tracker.columns[i+1:]...)
			break
		}
	}
	t.tracker.columns = append([]int{column}, t.tracker.columns...)
}

type ReqBody struct {
	Column string `json:"column"`
}

func Display(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("./template.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = template.Execute(w, people)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Sort(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var reqBody ReqBody
	err = json.Unmarshal(body, &reqBody)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	table := Table{table: people}
	sort.Sort(table)

	err = json.NewEncoder(w).Encode(table.table)
	if err != nil {
		http.Error(w, "Unable to encode response", http.StatusInternalServerError)
	}

}
