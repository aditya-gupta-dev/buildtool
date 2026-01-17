package main

type Args struct {
	Create      bool `help:"create a project"`
	ProjectName string `arg:"positional" help:"specify the name of your project"`
}
