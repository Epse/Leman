package config

import (
	"fmt"
	"testing"
)

func TestVerify(t *testing.T) {
	var bc BasicConfig
	if bc.Verify() == nil {
		t.Fail()
	}

	bc.Database.URL = "localhost"
	bc.Database.Password = "test"
	bc.Database.User = "user"
	bc.Database.Database = "db"
	bc.Logging.LogToFile = false
	bc.Logging.LogToStdout = true
	bc.Logging.StdoutLogLevel = "INFO"
	bc.Network.Port = 65536

	if bc.Verify() == nil {
		t.Fail()
	}

	bc.Network.Port = 0
	if bc.Verify() == nil {
		t.Fail()
	}

	bc.Network.Port = 8080
	err := bc.Verify()
	if err != nil {
		t.Fail()
		fmt.Println(err)
	}
}

func TestReadConfig(t *testing.T) {
	var bc BasicConfig
	err := bc.ReadConfig("./testfiles/doesnotexist.toml")
	if err == nil {
		fmt.Println(err)
		t.Fail()
	}
	err = bc.ReadConfig("./testfiles/empty.toml")
	if err == nil {
		fmt.Println(err)
		t.Fail()
	}
	err = bc.ReadConfig("./testfiles/complete.toml")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
}
