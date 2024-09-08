package internal

import (
	"log"

	"github.com/jexroid/neojex/utils"

	"github.com/charmbracelet/huh/spinner"
)

func GitClone(repo string) {
	err := spinner.New().
		Title("Cloning the repository").
		Action(func() {
			err := utils.ExecuteCmd("git", []string{"clone", repo})
			if err != nil {
				log.Fatal(err)
			}

		}).
		Run()
	if err != nil {
		log.Fatal(err)
	}
}
