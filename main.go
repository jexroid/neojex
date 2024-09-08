package main

import (
	"neojex/cmd"
	"neojex/internal"
	"os"

	"github.com/fatih/color"
)

var version string = "0.1.0"

func main() {
	_, available := os.LookupEnv("GH_TOKEN")
	if !available {
		color.New(color.BgRed, color.Bold).Println(" WARNING: GH_TOKEN is not configured in environment variables ")
		os.Exit(1)
	}
	color.New(color.BgGreen, color.Bold).Printf("\n %s \n", version)

	repo := cmd.RepoForm()
	internal.GitClone(repo)
}
