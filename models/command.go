package models

import (
	"os"
	"os/exec"
)

type Command struct {
	Name             string   `json:"name"`
	Executable       string   `json:"executable"`
	Arguments        []string `json:"arguments"`
	WorkingDirectory string   `json:"workingDirectory"`
}

func (c Command) Create() exec.Cmd {
	cmd := exec.Command(c.Executable, c.Arguments...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = c.WorkingDirectory
	return *cmd
}
