package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

type Naira float32
type DataBase map[string]Naira

type ReqBody struct {
	Item  string `json:"item"`
	Price string `json:"price"`
}

type UpdateReqBody struct {
	OldItem string `json:"old_item"`
	NewItem string `json:"new_item"`
	Price   string `json:"price"`
}

func (db DataBase) list(w http.ResponseWriter, r *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %g\n", item, price)
	}
}

func (db DataBase) listOne(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Item %s does not exist", item)
	}
	fmt.Fprintf(w, "%s: %g\n", item, price)
}

func (db DataBase) create(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "An error occurred: %s", err)
	}
	defer r.Body.Close()

	var reqBody ReqBody
	err = json.Unmarshal(body, &reqBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "An error occurred: %s", err)
	}
	price, err := strconv.Atoi(reqBody.Price)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "An error occurred: %s", err)
	}
	db[reqBody.Item] = Naira(price)
	fmt.Fprintf(w, "Item %s added to db", reqBody.Item)
}

func (db DataBase) update(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "An error occurred: %s", err)
	}
	defer r.Body.Close()

	var reqBody UpdateReqBody
	err = json.Unmarshal(body, &reqBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "An error occurred: %s", err)
	}

	_, ok := db[reqBody.OldItem]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "%s does not exist in database", reqBody.OldItem)
	}

	delete(db, reqBody.OldItem)

	price, err := strconv.Atoi(reqBody.Price)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "An error occurred: %s", err)
	}

	db[reqBody.NewItem] = Naira(price)

	fmt.Fprintf(w, "Item %s updated in db", reqBody.OldItem)
}

func (db DataBase) delete(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")

	_, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "%s does not exist in database", item)
	}

	delete(db, item)
	fmt.Fprintf(w, "Item %s deleted in db", item)
}

func main() {
	db := DataBase{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/list/:item", db.listOne)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete/:item", db.delete)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
