package main

import (
	"fmt"
	"github.com/Epse/Leman/backend/config"
	"github.com/op/go-logging"
	"github.com/pkg/errors"
	"net/http"
	"os"
)

var conf config.BasicConfig
var log = logging.MustGetLogger("main")

func itemListHandler(w http.ResponseWriter, r *http.Request) {
	//TODO
	w.WriteHeader(http.StatusNotImplemented)
	fmt.Fprintf(w, "501 NOT IMPLEMENTED")
	log.Error("itemListHandler not implemented but requested")
}

func itemsInStockHandler(w http.ResponseWriter, r *http.Request) {
	//TODO
	w.WriteHeader(http.StatusNotImplemented)
	fmt.Fprintf(w, "501 NOT IMPLEMENTED")
	log.Error("itemsInStockHandler not implemented but requested")
}

func itemsRentedHandler(w http.ResponseWriter, r *http.Request) {
	//TODO
	w.WriteHeader(http.StatusNotImplemented)
	fmt.Fprintf(w, "501 NOT IMPLEMENTED")
	log.Error("itemsRentedHandler not implemented but requested")
}

func itemViewHandler(w http.ResponseWriter, r *http.Request) {
	//TODO
	w.WriteHeader(http.StatusNotImplemented)
	fmt.Fprintf(w, "501 not implemented.")
	log.Error("itemViewHandler not implemented but requested")
}

func itemUpdateHandler(w http.ResponseWriter, r *http.Request) {
	//TODO
	w.WriteHeader(http.StatusNotImplemented)
	fmt.Fprintf(w, "501 not implemented.")
	log.Error("itemUpdateHandler not implemented but requested")
}

func itemNewHandler(w http.ResponseWriter, r *http.Request) {
	//TODO
	w.WriteHeader(http.StatusNotImplemented)
	fmt.Fprintf(w, "501 not implemented.")
	log.Error("itemNewHandler not implemented but requested")
}

func unresolvedRouteHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintf(w, "BAD REQUEST: route does not exist")
	log.Error("Unresolved route requested, route: " + r.URL.String())
}

func main() {
	// Set up logging
	var logFormat = logging.MustStringFormatter(
		`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
	)
	logBackend := logging.NewLogBackend(os.Stdout, "", 0)
	backendFormatter := logging.NewBackendFormatter(logBackend, logFormat)
	logging.SetBackend(backendFormatter)
	log.Info("Logging initialized")

	// Load the config file
	err := conf.ReadConfig("./config.toml")
	if err != nil {
		panic(errors.Wrap(err, "Configuration error"))
	}
	log.Info("Config file read")

	http.HandleFunc("/items/list/", itemListHandler)
	http.HandleFunc("/items/instock/", itemsInStockHandler)
	http.HandleFunc("/items/rented/", itemsRentedHandler)
	http.HandleFunc("/item/view/", itemViewHandler)
	http.HandleFunc("/item/update/", itemUpdateHandler)
	http.HandleFunc("/item/new/", itemNewHandler)
	http.HandleFunc("*", unresolvedRouteHandler)
	log.Info("Preparing to serve")
	http.ListenAndServe(":8080", nil)
}
