package main

import (
	"net/http"

	sort "7.6/sort"
)

func main() {
	http.HandleFunc("/display", sort.Display)
	http.HandleFunc("/sort", sort.Sort)
	http.ListenAndServe(":8880", nil)
}
