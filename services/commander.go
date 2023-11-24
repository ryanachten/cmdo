package services

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/fatih/color"
	"github.com/ryanachten/cmdo/events"
	"github.com/ryanachten/cmdo/models"
)

var terminalColours = []color.Attribute{color.FgCyan, color.FgMagenta, color.FgGreen, color.FgBlue}

type Commander struct {
	Commands       []models.Command
	EventBus       events.EventBus
	UseWeb         bool
	commandColours map[string]color.Attribute // command name / colour lookup
	userCommands   map[string]models.Command  // command name / command input lookup
	systemCommands map[string]*exec.Cmd       // command name / exec command lookup
}

func (c *Commander) Start() {
	c.commandColours = make(map[string]color.Attribute)
	c.userCommands = make(map[string]models.Command)
	c.systemCommands = make(map[string]*exec.Cmd)

	for i, command := range c.Commands {
		c.commandColours[command.Name] = terminalColours[i%len(terminalColours)]
		c.userCommands[command.Name] = command

		go c.startCommand(command)
	}

	go c.handleCommandRequests()

	// Handle signals to stop all processes (e.g., Ctrl+C).
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)
	<-signalChannel
	c.EventBus.Close()
}

func (c *Commander) startCommand(command models.Command) {
	cmd := command.Create(c.commandColours[command.Name], c.EventBus.CommandOutput, c.UseWeb)
	err := cmd.Start()

	c.systemCommands[command.Name] = &cmd

	if err != nil {
		fmt.Printf("Error starting command: %v\n", err)
		c.EventBus.CommandState <- events.CommandStateFail(command.Name, err.Error())
		return
	}

	c.EventBus.CommandState <- events.CommandStateStart(command.Name)

	// Wait for the command to finish.
	err = cmd.Wait()
	if err != nil {
		fmt.Printf("Command %s exited with error: %v\n", command.Name, err)
	} else {
		fmt.Printf("Command %s exited successfully\n", command.Name)
	}
}

// Watch for updates in the command request channel
func (c *Commander) handleCommandRequests() {
	for {
		req := <-c.EventBus.CommandRequest
		systemCommand := c.systemCommands[req.CommandName]

		// If the user is requesting a command to stop, we kill the process
		if req.RequestedState == events.CommandRequestStop {
			err := killProcess(systemCommand.Process)
			if err != nil {
				log.Printf("Error killing process: %v", err)
				c.EventBus.CommandState <- events.CommandStateFail(req.CommandName, err.Error())
				continue
			}
			c.EventBus.CommandState <- events.CommandStateStop(req.CommandName)
		} else {
			// Otherwise, if the user is requesting a command to start, we start it
			userCommand := c.userCommands[req.CommandName]
			go c.startCommand(userCommand)
		}
	}
}

// Kills process and child process tree
func killProcess(process *os.Process) error {
	log.Printf("Killing process and child processes: %v", process.Pid)

	if runtime.GOOS == "windows" {
		err := exec.Command("taskkill", "/F", "/T", "/PID", fmt.Sprint(process.Pid)).Run()
		return err
	}

	err := process.Signal(syscall.SIGKILL) // TODO: test on MacOS and Linux
	return err
}
