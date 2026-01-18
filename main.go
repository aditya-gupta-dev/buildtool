package main

import (
	"github.com/alexflint/go-arg"
)

func main() {
	var args Args
	arg.MustParse(&args)

	if args.Create { 
		CreateProject(args.ProjectName, args.Verbose)
	}

	if args.Compile { 
		CompileProject(args.ProjectName, args.Verbose)	
	}
}
