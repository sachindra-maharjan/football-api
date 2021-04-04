package csvutil

import (
	"encoding/csv"
	"log"
	"os"
	"path/filepath"
)

//Write creates if not exists or appends to an existing file
//writes data in the file
func Write(file string, data [][]string) error {
	csvFile, err := GetFile(file)
	if err != nil {
		return err
	}
	defer csvFile.Close()

	w := csv.NewWriter(csvFile)

	for _, str := range data {
		if err := w.Write(str); err != nil {
			log.Fatal("writing to file returned error", err)
			return err
		}
	}
	w.Flush()

	return nil
}

//GetFile Creates new file in the file path
func GetFile(path string) (*os.File, error) {
	dir, _ := filepath.Split(path)

	err := os.MkdirAll(dir, 0775)
	if err != nil {
		return nil, err
	}

	createOrOpenFile := func(path string) (*os.File, error) {
		file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return nil, err
		}
		return file, nil
	}

	file, err := createOrOpenFile(path)

	if err != nil {
		return nil, err
	}

	return file, nil
}

//FileExists Checks if file exists in specified path
func FileExists(file string) bool {
	info, err := os.Stat(file)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
