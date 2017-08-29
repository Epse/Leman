package config

import (
	"fmt"
	"testing"
)

func TestReadTOMLEmptyString(t *testing.T) {
	var bc BasicConfig
	err := bc.ReadTOML("")
	if err == nil {
		t.Fail()
	}
}

func TestReadTOMLBasicConfigReuse(t *testing.T) {
	var bc BasicConfig
	bc.DB_URL = "localhost"
	err := bc.ReadTOML("DB_URL=\"test\"\nDB_Password=\"pswd\"\nDB_Database=\"testdb\"\nDB_User=\"usr\"")
	if err != nil {
		t.Fail()
	}
	if bc.DB_URL != "test" {
		t.Fail()
	}
	if bc.DB_Password != "pswd" {
		t.Fail()
	}
	if bc.DB_Database != "testdb" {
		t.Fail()
	}
	if bc.DB_User != "usr" {
		t.Fail()
	}
}

func TestVerify(t *testing.T) {
	var bc BasicConfig
	err := bc.Verify()
	if err == nil {
		fmt.Println(err)
		t.Fail()
	}

	bc.DB_URL = "test"
	err = bc.Verify()
	if err == nil {
		fmt.Println(err)
		t.Fail()
	}

	bc.DB_Database = "test2"
	err = bc.Verify()
	if err == nil {
		fmt.Println(err)
		t.Fail()
	}

	bc.DB_User = "usr"
	err = bc.Verify()
	if err == nil {
		fmt.Println(err)
		t.Fail()
	}

	bc.DB_Password = "pass"
	err = bc.Verify()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	bc.DB_URL = ""
	err = bc.Verify()
	if err == nil {
		fmt.Println(err)
		t.Fail()
	}

	bc.DB_Database = ""
	err = bc.Verify()
	if err == nil {
		fmt.Println(err)
		t.Fail()
	}

	bc.DB_User = ""
	err = bc.Verify()
	if err == nil {
		fmt.Println(err)
		t.Fail()
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
