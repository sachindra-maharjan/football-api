package util

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestCsvUtil_CreateDirectoryStructure(t *testing.T) {
	path, err := ioutil.TempDir("", "league")
	if err != nil {
		t.Fatalf("TempDir returned error %v", err)
	}
	path += "/league.csv"
	t.Log("filepath : " + path)
	dir := filepath.Dir(path)
	file, err := GetFile(path)
	defer os.RemoveAll(dir)

	if err != nil {
		t.Fatalf("CreateNewFile returned error %v", err)
	}

	defer file.Close()

	if err != nil {
		t.Fatalf("Read file returned error %v", err)
	}

}

func TestCsvUtil_WriteData(t *testing.T) {
	records := [][]string{
		{"first_name", "last_name", "username"},
		{"Rob", "Pike", "rob"},
		{"Ken", "Thompson", "ken"},
		{"Robert", "Griesemer", "gri"},
	}

	Write("~/test.csv", records)

}
