package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/new", db.new)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	bytes, err := json.Marshal(db)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, string(bytes))
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if price, ok := db[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
		return
	}

	w.WriteHeader(http.StatusNotFound)
}

func (db database) new(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	priceStr := req.URL.Query().Get("price")

	if item == "" || priceStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "item and price query is required")
		return
	}

	if _, exists := db[item]; exists {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "already exists: %q\n", item)
		return
	}

	price, err := strconv.ParseFloat(priceStr, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "price query must be float value: %q\n", priceStr)
		return
	}

	db[item] = dollars(price)
	w.WriteHeader(http.StatusCreated)
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	priceStr := req.URL.Query().Get("price")

	if item == "" || priceStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "item and price query is required")
		return
	}

	if _, exists := db[item]; !exists {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "not exists: %q\n", item)
		return
	}

	price, err := strconv.ParseFloat(priceStr, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "price query must be float value: %q\n", priceStr)
		return
	}

	db[item] = dollars(price)

	w.WriteHeader(http.StatusOK)
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")

	if item == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "item query is required")
		return
	}

	if _, exists := db[item]; !exists {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "not exists: %q\n", item)
		return
	}

	delete(db, item)

	w.WriteHeader(http.StatusOK)
}
