package models

import (
	"io"
	"os"
	"os/exec"

	"github.com/fatih/color"
)

type Command struct {
	Name             string   `json:"name"`
	Executable       string   `json:"executable"`
	Arguments        []string `json:"arguments"`
	WorkingDirectory string   `json:"workingDirectory"`
	Tags             []string `json:"tags"`
}

func (c Command) Create(commandColour color.Attribute, broadcastChannel BroadcastChannel, useWeb bool) exec.Cmd {
	cmd := exec.Command(c.Executable, c.Arguments...)

	cmd.Stdout = &commandWriter{
		color:            color.New(commandColour),
		writer:           os.Stdout,
		commandName:      c.Name,
		broadcastChannel: broadcastChannel,
		useWeb:           useWeb,
	}
	cmd.Stderr = &commandWriter{
		color:       color.New(color.FgRed),
		writer:      os.Stderr,
		commandName: c.Name,
		useWeb:      useWeb,
	}

	cmd.Dir = c.WorkingDirectory

	return *cmd
}

// Formats standard output for each command
type commandWriter struct {
	color            *color.Color
	commandName      string
	writer           io.Writer
	broadcastChannel BroadcastChannel
	useWeb           bool
}

// Logs to both writer and broadcast channel
func (cw *commandWriter) Write(p []byte) (n int, err error) {
	message := string(p)

	cw.color.Fprintf(cw.writer, "[%s]: %s", cw.commandName, message)
	if cw.useWeb {
		cw.broadcastChannel <- BroadcastMessage{
			CommandName: cw.commandName,
			MessageBody: message,
		}
	}

	return len(p), nil
}
