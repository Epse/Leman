package main

import (
	"fmt"
	"net/http"
	"strings"
)

// This function handles all requests to /items/operation
// operation is one of the following: list, instock, rented
// list gives all items, instock all those not rented and rented all those rented out
func itemListHandler(w http.ResponseWriter, r *http.Request) {
	operation := string(r.URL.Path[len("/items/")])
	operation = strings.TrimSuffix(operation, "/")

	if strings.Contains(operation, "/") {
		w.WriteHeader(http.StatusBadRequest)
		//TODO: more informative error message
		fmt.Fprintf(w, "400 BAD REQUEST.")
	}

	if operation == "list" {
		//TODO: return all items in neat JSON
		w.WriteHeader(http.StatusNotImplemented)
		fmt.Fprintf(w, "501 not implemented.")
	} else if operation == "instock" {
		//TODO: return all items in stock in neat JSON
		w.WriteHeader(http.StatusNotImplemented)
		fmt.Fprintf(w, "501 not implemented.")
	} else if operation == "rented" {
		//TODO: return all rented items as neat JSON
		w.WriteHeader(http.StatusNotImplemented)
		fmt.Fprintf(w, "501 not implemented.")
	}
}

func itemViewHandler(w http.ResponseWriter, r *http.Request) {
	//TODO
	w.WriteHeader(http.StatusNotImplemented)
	fmt.Fprintf(w, "501 not implemented.")
}

func itemUpdateHandler(w http.ResponseWriter, r *http.Request) {
	//TODO
	w.WriteHeader(http.StatusNotImplemented)
	fmt.Fprintf(w, "501 not implemented.")
}

func itemNewHandler(w http.ResponseWriter, r *http.Request) {
	//TODO
	w.WriteHeader(http.StatusNotImplemented)
	fmt.Fprintf(w, "501 not implemented.")
}

func main() {
	http.HandleFunc("/items/", itemListHandler)
	http.HandleFunc("/item/view/", itemViewHandler)
	http.HandleFunc("/item/update/", itemUpdateHandler)
	http.HandleFunc("/item/new/", itemNewHandler)
	http.ListenAndServe(":8080", nil)
}
