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
}

func (c Commander) Start() {
	var waitGroup sync.WaitGroup

	for i, command := range c.Commands {
		cmd := command.Create(terminalColours[i%len(terminalColours)], c.BroadcastChannel, c.UseWeb)

		go func(cmd *exec.Cmd) {
			defer waitGroup.Done()

			err := cmd.Start()
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
		}(&cmd)

		// Watch for updates in the command request channel
		go func(currentCommand models.Command, sysCommand *exec.Cmd) {
			for {
				req := <-c.CommandRequestChannel
				if currentCommand.Name == req.CommandName {
					killProcess(cmd.Process)
				}
			}
		}(command, &cmd)

		waitGroup.Add(1)
	}

	// Wait for all processes to finish or a signal to stop.
	go func() {
		waitGroup.Wait()
		close(c.BroadcastChannel)
		close(c.CommandRequestChannel)
	}()

	// Handle signals to stop all processes (e.g., Ctrl+C).
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)
	<-signalChannel
	close(signalChannel)
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
