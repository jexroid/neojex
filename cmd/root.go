package cmd

import (
	"log"

	"github.com/jexroid/neojex/pkg"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/huh"
)

var (
	repositories, _ = pkg.GetRepos()

	reposURL string

	myCustomKeyMap = &huh.KeyMap{
		FilePicker: huh.FilePickerKeyMap{
			Up: key.NewBinding(
				key.WithKeys("k", "up"),        // actual keybindings
				key.WithHelp("↑/k", "move up"), // corresponding help text
			),
			Down: key.NewBinding(
				key.WithKeys("j", "down"),        // actual keybindings
				key.WithHelp("↓/j", "move down"), // corresponding help text
			),
		},
		Quit: key.NewBinding(
			key.WithKeys("q", "ctrl+c"), // actual keybindings
			key.WithHelp("q / Ctrl+c", "quit"),
		),
	}
)

func optionFunction() []huh.Option[string] {
	options := make([]huh.Option[string], len(repositories))
	for i, repo := range repositories {
		options[i] = huh.NewOption(repo.Name, repo.Url)
	}
	return options
}

func RepoForm() string {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Select your repository.").
				OptionsFunc(optionFunction, &reposURL).
				Value(&reposURL),
		).WithKeyMap(myCustomKeyMap).WithShowHelp(true),
	)

	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}

	return reposURL
}
