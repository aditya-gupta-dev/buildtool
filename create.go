package main

import (
	"log"
	"os"
	"strings"
	"path/filepath"
)

func CreateProject(projectName string) {
	if projectName == "" {
		dir, err := os.Getwd()
		if err != nil { 
			log.Fatalf("failed to get pwd: %s\n", err)
		}

		projectName = dir[strings.LastIndex(dir, "/")+1:]
	}

	err := os.MkdirAll(projectName, 0755)
	if err != nil {
		log.Fatalf("failed to create directory: %s\n", err)
	}


}

func createSubFiles(projectName string) error {
	chdir, err := os.Getwd()
	if err != nil { 
		return err 
	}

	file, err := os.Create(filepath.Join(chdir, projectName, "main.c"))
	if err != nil { 
		return err 
	}
	file.Close() 

	return nil 
}

