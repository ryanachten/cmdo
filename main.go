package main

import (
	"log"

	"github.com/ryanachten/cmdo/models"
	"github.com/ryanachten/cmdo/services"
)

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

	var broadcastChannel = make(chan models.BroadcastMessage)

	if arguments.UseWeb {
		webServer := services.WebServer{
			BroadcastChannel: broadcastChannel,
		}
		go webServer.Start()
	}

	commander := services.Commander{
		Commands:         commands,
		BroadcastChannel: broadcastChannel,
		UseWeb:           arguments.UseWeb,
	}
	commander.Start()
}
