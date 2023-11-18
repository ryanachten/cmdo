package main

import (
	"log"

	"github.com/ryanachten/cmdo/models"
	"github.com/ryanachten/cmdo/services"
)

func main() {
	arguments := models.GetArguments()

	config, err := models.ParseConfigurationFile(arguments.ConfigurationPath)
	if err != nil {
		log.Fatalln(err)
	}

	commands := arguments.FilterCommands(config.Commands)
	if len(commands) == 0 {
		log.Fatalln("No commands selected using the provided arguments")
	}

	var broadcastChannel = make(models.BroadcastChannel)
	var commandRequestChannel = make(models.CommandRequestChannel)

	if arguments.UseWeb {
		webServer := services.WebServer{
			BroadcastChannel:      broadcastChannel,
			CommandRequestChannel: commandRequestChannel,
		}
		go webServer.Start()
	}

	commander := services.Commander{
		Commands:              commands,
		BroadcastChannel:      broadcastChannel,
		CommandRequestChannel: commandRequestChannel,
		UseWeb:                arguments.UseWeb,
	}
	commander.Start()
}
