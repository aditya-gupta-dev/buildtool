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

	fname := "main.c"
	err = createFile(projectName, fname)
	if err != nil { 
		log.Fatalf("failed to create project file [%s]: %s\n", fname, err)
	}

	fname = "include"
	err = createDir(projectName, fname)	
	if err != nil {
		log.Fatalf("failed to create project file [%s]: %s\n", fname, err)
	}

	fname = "lib"
	err = createDir(projectName, fname)
	if err != nil { 
		log.Fatalf("failed to create project file [%s]: %s\n", fname, err)
	}
}

func createFile(projectName, filename string) error {
	chdir, err := os.Getwd()
	if err != nil { 
		return err 
	}

	file, err := os.Create(filepath.Join(chdir, projectName, filename))
	if err != nil { 
		return err 
	}
	file.Close() 

	return nil 
}

func createDir(projectName, dirname string) error {
	err := os.MkdirAll(filepath.Join(projectName, dirname), 0755)
	return err
}

