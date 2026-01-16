package main

type Args struct {
	Create      bool
	ProjectName string `arg:"positional"`
}
