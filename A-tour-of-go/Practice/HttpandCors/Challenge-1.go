// modern web speaks json api
/*
starts a  normal https  server no neeed for https and tls for this one just use http.ListenAndserve(":8080",mux)
it should have post/user end point
this end point shoudl read json bodfy from request that looks like tthis
*/
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// user struct to decode json into (Exported type with capital U)
type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// response struct for the json response
type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// HandleUser handles post /user endpoint
func handleUser(w http.ResponseWriter, r *http.Request) {
	// check if tis post request
	if r.Method != http.MethodPost {
		http.Error(w, "invalid method", http.StatusMethodNotAllowed)
		return
	}

	var user User

	// Fix: Decoder -> Decode
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "invalid Json", http.StatusBadRequest)
		return
	}

	// creating response
	response := Response{
		Status:  "success",
		Message: fmt.Sprintf("User %s is created!", user.Name),
	}

	// Fix: set response header
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)
}

func main() {
	// Register the handler
	// Fix: Missing comma in HandleFunc
	http.HandleFunc("/user", handleUser)
	addr := ":8080"
	fmt.Printf("server listening on http://localhost%s\n", addr)
	
	// Fix: ListenAndServe takes two arguments (addr, handler)
	log.Fatal(http.ListenAndServe(addr, nil))
}