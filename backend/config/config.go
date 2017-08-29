// Package config manages the configuration files for Leman.
// Configuration files are always written in TOML, since that's really easy for us to parse
package config

import (
	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"
	"os"
)

type BasicConfig struct {
	DB_URL      string
	DB_Password string
	DB_Database string
	DB_User     string
}

func (bc *BasicConfig) Verify() error {
	var errorString string
	if bc.DB_URL == "" {
		errorString += "Incomplete config file: no DB_URL field present\n"
	}
	if bc.DB_Database == "" {
		errorString += "Incomplete config file: no DB_URL field present\n"
	}
	if bc.DB_Password == "" {
		errorString += "Incomplete config file: no DB_URL field present\n"
	}
	if bc.DB_User == "" {
		errorString += "Incomplete config file: no DB_User field present\n"
	}
	if errorString != "" {
		return errors.New(errorString)
	} else {
		return nil
	}
}

func (bc *BasicConfig) ReadTOML(data string) error {
	_, err := toml.Decode(data, &bc)
	if err != nil {
		err = errors.Wrap(err, "Could not decode TOML")
	} else {
		err = bc.Verify()
	}
	return err
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
