package main

import (
	"fmt"
	"github.com/Epse/Leman/backend/config"
	"github.com/pkg/errors"
	"net/http"
)

var conf config.BasicConfig

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

func unresolvedRouteHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintf(w, "BAD REQUEST: route does not exist")
}

func main() {
	// Load the config file
	err := conf.ReadConfig("./config.toml")
	if err != nil {
		panic(errors.Wrap(err, "Configuration error"))
	}

	http.HandleFunc("/items/list/", itemListHandler)
	http.HandleFunc("/items/instock/", itemsInStockHandler)
	http.HandleFunc("/items/rented/", itemsRentedHandler)
	http.HandleFunc("/item/view/", itemViewHandler)
	http.HandleFunc("/item/update/", itemUpdateHandler)
	http.HandleFunc("/item/new/", itemNewHandler)
	http.HandleFunc("*", unresolvedRouteHandler)
	http.ListenAndServe(":8080", nil)
}
