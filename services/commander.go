package services

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"sync"
	"syscall"

	"github.com/fatih/color"
	"github.com/ryanachten/cmdo/models"
)

var terminalColours = []color.Attribute{color.FgCyan, color.FgMagenta, color.FgGreen, color.FgBlue}

type Commander struct {
	Commands              []models.Command
	BroadcastChannel      models.BroadcastChannel
	CommandRequestChannel models.CommandRequestChannel
	UseWeb                bool
	wg                    sync.WaitGroup
	commandColours        map[string]color.Attribute // command name / colour lookup
	userCommands          map[string]models.Command  // command name / command input lookup
	systemCommands        map[string]*exec.Cmd       // command name / exec command lookup
}

func (c *Commander) Start() {
	c.commandColours = make(map[string]color.Attribute)
	c.userCommands = make(map[string]models.Command)
	c.systemCommands = make(map[string]*exec.Cmd)

	for i, command := range c.Commands {
		c.commandColours[command.Name] = terminalColours[i%len(terminalColours)]
		c.userCommands[command.Name] = command

		go c.startCommand(command)

		c.wg.Add(1)
	}

	// Watch for updates in the command request channel
	go func() {
		for {
			req := <-c.CommandRequestChannel
			systemCommand := c.systemCommands[req.CommandName]
			if req.RequestedState == models.CommandRequestStop {
				err := killProcess(systemCommand.Process)
				if err != nil {
					log.Printf("Error killing process: %v", err)
				}
			} else {
				userCommand := c.userCommands[req.CommandName]
				go c.startCommand(userCommand)
			}
		}
	}()

	// Wait for all processes to finish or a signal to stop.
	go func() {
		c.wg.Wait()
		close(c.BroadcastChannel)
		close(c.CommandRequestChannel)
	}()

	// TODO: is this signalChannel still needed?
	// Handle signals to stop all processes (e.g., Ctrl+C).
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)
	<-signalChannel
	close(signalChannel)
}

func (c *Commander) startCommand(command models.Command) {
	// TODO: we need to reconsider how this wait group is used in the new control flow
	// defer c.wg.Done()

	cmd := command.Create(c.commandColours[command.Name], c.BroadcastChannel, c.UseWeb)
	err := cmd.Start()

	c.systemCommands[command.Name] = &cmd

	if err != nil {
		fmt.Printf("Error starting command: %v\n", err)
		return
	}

	// Wait for the command to finish.
	err = cmd.Wait()
	if err != nil {
		fmt.Printf("Command %s exited with error: %v\n", command.Name, err)
	} else {
		fmt.Printf("Command %s exited successfully\n", command.Name)
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
