package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func (r root) LicenseScan() {
	var files []string
	dec := false
	fmt.Println("Running LICENSE.md scanner...")
	err := filepath.Walk(string(r), func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		if file == string(r)+"/LICENSE.md" {
			dec = true
			break
		}
	}
	if dec == true {
		fmt.Println("LICENSE.md found.")
	} else {
		fmt.Println("LICENSE.md not found.")
	}
}