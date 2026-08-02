package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`

	/*
	Book struct bundles related data together (ID, Title, Author, Price). This is encapsulation — grouping state into one unit.
	*/
}
/*
inventory and nextID are lowercase, meaning they are unexported (private to this package). 
No other package can directly access or corrupt your data store. That's Go's version of private in C++.

*/

var inventory = make(map[string]Book)
var nextID = 1

func main() {
	http.HandleFunc("/books/", handleBooks)
	fmt.Println("server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	id := strings.TrimPrefix(r.URL.Path, "/books/")

	switch r.Method {

	case http.MethodPost:
		var newBook Book
		if err := json.NewDecoder(r.Body).Decode(&newBook); err != nil {
			http.Error(w, "Invalid JSON body", http.StatusBadRequest)
			return
		}
		newBook.ID = strconv.Itoa(nextID)
		nextID++
		inventory[newBook.ID] = newBook
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newBook)

	case http.MethodGet:
		if id == "" {
			var allBooks []Book
			for _, book := range inventory {
				allBooks = append(allBooks, book)
			}
			json.NewEncoder(w).Encode(allBooks)
		} else {
			book, exists := inventory[id]
			if !exists {
				http.Error(w, "book not found", http.StatusNotFound)
				return
			}
			json.NewEncoder(w).Encode(book)
		}

	case http.MethodPut:
		if id == "" {
			http.Error(w, "ID is required for PUT", http.StatusBadRequest)
			return
		}
		_, exists := inventory[id]
		if !exists {
			http.Error(w, "book not found", http.StatusNotFound)
			return
		}
		var updateBook Book
		if err := json.NewDecoder(r.Body).Decode(&updateBook); err != nil {
			http.Error(w, "invalid JSON body", http.StatusBadRequest)
			return
		}
		updateBook.ID = id
		inventory[id] = updateBook
		json.NewEncoder(w).Encode(updateBook)

	case http.MethodPatch:
		if id == "" {
			http.Error(w, "ID is required for PATCH", http.StatusBadRequest)
			return
		}
		existingBook, exists := inventory[id]
		if !exists {
			http.Error(w, "book not found", http.StatusNotFound)
			return
		}
		var partialUpdate Book
		if err := json.NewDecoder(r.Body).Decode(&partialUpdate); err != nil {
			http.Error(w, "invalid JSON Body", http.StatusBadRequest)
			return
		}
		if partialUpdate.Title != "" {
			existingBook.Title = partialUpdate.Title
		}
		if partialUpdate.Author != "" {
			existingBook.Author = partialUpdate.Author
		}
		if partialUpdate.Price != 0 {
			existingBook.Price = partialUpdate.Price
		}
		inventory[id] = existingBook
		json.NewEncoder(w).Encode(existingBook)

	case http.MethodDelete:
		if id == "" {
			http.Error(w, "ID is required for DELETE", http.StatusBadRequest)
			return
		}
		_, exists := inventory[id]
		if !exists {
			http.Error(w, "book not found", http.StatusNotFound)
			return
		}
		delete(inventory, id)
		w.WriteHeader(http.StatusNoContent)

	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}
