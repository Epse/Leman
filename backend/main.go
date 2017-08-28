package main

import (
	"fmt"
	"net/http"
	"strings"
)

func itemListHandler(w http.ResponseWriter, r *http.Request) {
	//TODO
	w.WriteHeader(http.StatusNotImplemented)
	fmt.Fprintf(w, "501 NOT IMPLEMENTED")
}

func itemsInStockHandler(w http.ResponseWriter, r *http.Request) {
	//TODO
	w.WriteHeader(http.StatusNotImplemented)
	fmt.Fprintf(w, "501 NOT IMPLEMENTED")
}

func itemsRentedHandler(w http.ResponseWriter, r *http.Request) {
	//TODO
	w.WriteHeader(http.StatusNotImplemented)
	fmt.Fprintf(w, "501 NOT IMPLEMENTED")
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
	http.HandleFunc("/items/list/", itemListHandler)
	http.HandleFunc("/items/instock/", itemsInStockHandler)
	http.HandleFunc("/items/rented/", itemsRentedHandler)
	http.HandleFunc("/item/view/", itemViewHandler)
	http.HandleFunc("/item/update/", itemUpdateHandler)
	http.HandleFunc("/item/new/", itemNewHandler)
	http.ListenAndServe(":8080", nil)
}
