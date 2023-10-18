package main

import (
	"commando/models"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"sync"
	"syscall"

	"github.com/fatih/color"
)

var terminalColours = []color.Attribute{color.FgCyan, color.FgMagenta, color.FgGreen, color.FgBlue}

func main() {
	arguments, err := models.GetArguments()
	if err != nil {
		log.Fatalln(err)
	}
	config, err := models.ParseConfigurationFile(arguments.ConfigurationPath)
	if err != nil {
		log.Fatalln(err)
	}

	commands := arguments.FilterCommands(config.Commands)
	if len(commands) == 0 {
		log.Fatalln("No commands selected using the provided arguments")
	}

	var waitGroup sync.WaitGroup
	stopChannel := make(chan struct{})

	for i, command := range commands {
		cmd := command.Create(terminalColours[i%len(terminalColours)])

		go func(cmd *exec.Cmd) {
			defer waitGroup.Done()
			defer close(stopChannel)

			err := cmd.Start()
			if err != nil {
				fmt.Printf("Error starting command: %v\n", err)
				return
			}

			// Wait for the command to finish.
			err = cmd.Wait()
			if err != nil {
				fmt.Printf("Command exited with error: %v\n", err)
			} else {
				fmt.Printf("Command exited successfully\n")
			}

			// Signal other Goroutines to stop.
			select {
			case <-stopChannel:
				fmt.Printf("Command %v exiting\n", command.Name)
				cmd.Process.Kill()
				return
			default:
				// Continue execution to close other processes.
			}
		}(&cmd)

		waitGroup.Add(1)
	}

	// Wait for all processes to finish or a signal to stop.
	go func() {
		waitGroup.Wait()
		close(stopChannel)
	}()

	// Handle signals to stop all processes (e.g., Ctrl+C).
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)
	<-signalChannel
	close(signalChannel)
	close(stopChannel)
}
