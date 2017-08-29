// Package config manages the configuration files for Leman.
// Configuration files are always written in TOML, since that's really easy for us to parse
package config

import (
	"errors"
	"github.com/BurntSushi/toml"
	"os"
)

type BasicConfig struct {
	DB_URL      string
	DB_Password string
	DB_Database string
}

func (bc BasicConfig) ReadTOML(data string) error {
	_, err := toml.Decode(data, &bc)
	return errors.Wrap(err, "Couldn't decode TOML config file.")
}

func (bc BasicConfig) ReadConfig(fpath string) error {
	_, err := os.Stat(fpath)
	if err != nil {
		//TODO: more informative help message
		panic("Missing config file!")
	}

	_, err = toml.DecodeFile(fpath, &bc)
	if err != nil {
		panic("Can't decode config file!")
	}

	// Verifying the contents of the config file
	if bc.DB_URL == "" {
		return errors.New("Inclomplete config file: no DB_URL field present")
	}
	if bc.DB_Password == "" {
		return errors.New("Inclomplete config file: no DB_Password field present")
	}
	if bc.DB_Database == "" {
		return errors.New("Inclomplete config file: no DB_Database field present")
	}
	return nil
}
