package repository

import (
	"os"
	"path/filepath"
)

func createDir(address string) error {
	err := os.MkdirAll(address, 0777)
	if err != nil && !os.IsExist(err) {
		return err
	}
	return nil
}

func createFile(address string, name string) error {
	err := createDir(address)
	if err != nil && !os.IsExist(err) {
		return err
	}
	add := filepath.Join(address, name)
	_, err = os.Create(add)
	return err
}

func isFileExists(address string) (bool, error) {
	info, err := os.Stat(address)
	if os.IsNotExist(err) {
		return false, nil
	}
	return !info.IsDir(), nil
}

func isDirExists(address string) (bool, error) {
	info, err := os.Stat(address)
	if os.IsNotExist(err) {
		return false, nil
	}
	return info.IsDir(), nil
}

func listDirFiles(address string) ([]string, error) {
	var contents []string
	entries, err := os.ReadDir(address)
	if err != nil {
		return nil, err
	}

	for _, e := range entries {
		if !e.IsDir() {
			contents = append(contents, e.Name())
		}
	}
	return contents, nil
}
