package models

import (
	"os"
	"os/exec"

	"github.com/fatih/color"
	"github.com/ryanachten/cmdo/events"
)

type Command struct {
	Name             string   `json:"name"`
	Executable       string   `json:"executable"`
	Arguments        []string `json:"arguments"`
	WorkingDirectory string   `json:"workingDirectory"`
	Tags             []string `json:"tags"`
}

func (c Command) Create(commandColour color.Attribute, broadcastChannel events.CommandOutputChannel, useWeb bool) exec.Cmd {
	cmd := exec.Command(c.Executable, c.Arguments...)

	cmd.Stdout = &commandWriter{
		color:            color.New(commandColour),
		isError:          false,
		commandName:      c.Name,
		broadcastChannel: broadcastChannel,
		useWeb:           useWeb,
	}
	cmd.Stderr = &commandWriter{
		color:            color.New(color.FgRed),
		isError:          true,
		commandName:      c.Name,
		broadcastChannel: broadcastChannel,
		useWeb:           useWeb,
	}

	cmd.Dir = c.WorkingDirectory

	return *cmd
}

// Formats standard output for each command
type commandWriter struct {
	color            *color.Color
	commandName      string
	isError          bool
	broadcastChannel events.CommandOutputChannel
	useWeb           bool
}

// Logs to both writer and broadcast channel
func (cw *commandWriter) Write(p []byte) (n int, err error) {
	message := string(p)
	writer := os.Stdout
	if cw.isError {
		writer = os.Stderr
	}

	cw.color.Fprintf(writer, "[%s]: %s", cw.commandName, message)
	if cw.useWeb {
		if cw.isError {
			cw.broadcastChannel <- events.CommandOutputError(cw.commandName, message)
		} else {
			cw.broadcastChannel <- events.CommandOutputInformation(cw.commandName, message)
		}
	}

	return len(p), nil
}
