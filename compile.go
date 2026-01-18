package main

import (
	"log"
	"os"
	"strings"
)

func CompileProject(projectName string, verbose bool) { 

	if projectName == "" { 
		dir, err := os.Getwd()
		if err != nil { 
			log.Fatalf("failed to get current directory: %s\n", err)
		}

		projectName = dir[strings.LastIndex(dir, "/")+1:]
		if verbose { 
			log.Printf("project name not passed, using this instead: %s\n", projectName)
		}
	}

	if verbose { 
		log.Printf("started compiling project [ %s ]\n", projectName)
	}


}
