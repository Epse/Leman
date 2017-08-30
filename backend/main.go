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
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "404 NOT FOUND:  route does not exist")
	log.Error("Unresolved route requested, route: " + r.URL.String())
}

func loggingSetup() error {
	var logFormat = logging.MustStringFormatter(
		`%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x} %{message}`,
	)
	var logConsoleFormat = logging.MustStringFormatter(
		`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
	)

	var loggingLeveledFileBackend logging.LeveledBackend
	var loggingLeveledStdoutBackend logging.LeveledBackend

	if conf.Logging.LogToStdout {
		logStdoutBackend := logging.NewLogBackend(os.Stdout, "", 0)
		backendStdoutFormatter := logging.NewBackendFormatter(logStdoutBackend, logConsoleFormat)
		backendStdoutFormatterLeveled := logging.AddModuleLevel(backendStdoutFormatter)
		switch conf.Logging.StdoutLogLevel {
		case "DEBUG":
			backendStdoutFormatterLeveled.SetLevel(logging.DEBUG, "")
		case "INFO":
			backendStdoutFormatterLeveled.SetLevel(logging.INFO, "")
		case "NOTICE":
			backendStdoutFormatterLeveled.SetLevel(logging.NOTICE, "")
		case "WARNING":
			backendStdoutFormatterLeveled.SetLevel(logging.WARNING, "")
		case "ERROR":
			backendStdoutFormatterLeveled.SetLevel(logging.ERROR, "")
		case "CRITICAL":
			backendStdoutFormatterLeveled.SetLevel(logging.CRITICAL, "")
		default:
			return errors.New("invalid stdout log level")
		}
		loggingLeveledStdoutBackend = backendStdoutFormatterLeveled
	}
	if conf.Logging.LogToFile {
		f, err := os.Create(conf.Logging.LogFile)
		if err != nil {
			return errors.Wrap(err, "Could not create log file")
		}
		logFileBackend := logging.NewLogBackend(f, "", 0)
		backendFileFormatter := logging.NewBackendFormatter(logFileBackend, logFormat)
		backendFileFormatterLeveled := logging.AddModuleLevel(backendFileFormatter)
		switch conf.Logging.FileLogLevel {
		case "DEBUG":
			backendFileFormatterLeveled.SetLevel(logging.DEBUG, "")
		case "INFO":
			backendFileFormatterLeveled.SetLevel(logging.INFO, "")
		case "NOTICE":
			backendFileFormatterLeveled.SetLevel(logging.NOTICE, "")
		case "WARNING":
			backendFileFormatterLeveled.SetLevel(logging.WARNING, "")
		case "ERROR":
			backendFileFormatterLeveled.SetLevel(logging.ERROR, "")
		case "CRITICAL":
			backendFileFormatterLeveled.SetLevel(logging.CRITICAL, "")
		default:
			return errors.New("invalid file log level")
		}
		loggingLeveledFileBackend = backendFileFormatterLeveled
	}

	if conf.Logging.LogToFile && conf.Logging.LogToStdout {
		logging.SetBackend(loggingLeveledFileBackend, loggingLeveledStdoutBackend)
	} else if conf.Logging.LogToFile {
		logging.SetBackend(loggingLeveledFileBackend)
	} else if conf.Logging.LogToStdout {
		logging.SetBackend(loggingLeveledStdoutBackend)
	}
	log.Info("Logging initialized")

	return nil
}

func main() {
	// Load the config file
	err := conf.ReadConfig("./config.toml")
	if err != nil {
		panic(errors.Wrap(err, "Configuration error"))
	}

	// Set up logging
	err = loggingSetup()
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/items/list/", itemListHandler)
	http.HandleFunc("/items/instock/", itemsInStockHandler)
	http.HandleFunc("/items/rented/", itemsRentedHandler)
	http.HandleFunc("/item/view/", itemViewHandler)
	http.HandleFunc("/item/update/", itemUpdateHandler)
	http.HandleFunc("/item/new/", itemNewHandler)
	http.HandleFunc("/", unresolvedRouteHandler)
	log.Info("Preparing to serve")
	http.ListenAndServe(":8080", nil)
}
