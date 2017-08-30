// Package config manages the configuration files for Leman.
// Configuration files are always written in TOML, since that's really easy for us to parse
package config

import (
	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"
	"os"
	"reflect"
)

type BasicConfig struct {
	Database DatabaseConfig
	Logging  LoggingConfig
	Network  NetworkConfig
}

type DatabaseConfig struct {
	URL      string
	Password string
	User     string
	Database string
}

type LoggingConfig struct {
	LogToFile      bool
	LogToStdout    bool
	LogFile        string
	FileLogLevel   string
	StdoutLogLevel string
}

type NetworkConfig struct {
	Port int
}

func (bc *BasicConfig) Verify() error {
	var errorString string

	// Test the Database section
	// This iterates over all fields in the Database subsection
	// All these fields are strings and mandatory
	// So if any one of these is empty, the error string is appended.
	v := reflect.ValueOf(bc.Database)
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).String() == "" {
			// That weird `v.Type().Field(i).Name` gives the field name in the struct
			errorString += "Incomplete DB Config section: field " + v.Type().Field(i).Name
		}
	}

	// Test the Logging section
	if bc.Logging.LogToFile && bc.Logging.LogFile == "" {
		errorString += "Incomplete Logging Config section: File logging enabled but no file path set."
	}
	if bc.Logging.LogToFile && bc.Logging.FileLogLevel == "" {
		errorString += "Incomplete Logging Config section: File logging enabled but no log level set."
	}
	if bc.Logging.LogToStdout && bc.Logging.StdoutLogLevel == "" {
		errorString += "Incomplete Logging Config section: Stdout logging enabled but no log level set."
	}

	// Test the Network section
	if bc.Network.Port == 0 || bc.Network.Port > 65535 {
		errorString += "Invalid Network Config section: Invalid Port number."
	}

	if errorString == "" {
		return nil
	} else {
		return errors.Wrap(errors.New(errorString), "Config verification error.")
	}
}

func (bc *BasicConfig) ReadConfig(fpath string) error {
	_, err := os.Stat(fpath)
	if err != nil {
		//TODO: more informative help message
		return errors.New("Config file not found")
	}

	_, err = toml.DecodeFile(fpath, &bc)
	if err != nil {
		return errors.Wrap(err, "Config file decode error")
	}

	return bc.Verify()
}
