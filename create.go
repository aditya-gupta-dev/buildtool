package main

import (
	"log"
	"os"
	"fmt"
	"strings"
)

func CreateProject(projectName string) {
	if projectName == "" {
		dir, err := os.Getwd()
		if err != nil { 
			log.Fatalf("failed to get pwd: %s\n", err)
		}

		projectName = dir[strings.LastIndex(dir, "/")+1:]
	}

	fmt.Println(projectName)
}
