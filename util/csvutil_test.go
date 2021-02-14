package csvutil

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
