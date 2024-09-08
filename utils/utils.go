package utils

import (
	"bytes"
	"os/exec"
	"path"
)

func ExecuteCmd(name string, args []string) error {
	command := exec.Command(name, args...)
	command.Dir = path.Join()
	var out bytes.Buffer
	command.Stdout = &out
	if err := command.Run(); err != nil {
		return err
	}
	return nil
}
