package main

import (
	"commando/models"
	"commando/services"
	"log"
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

	webServer := services.WebServer{
		BroadcastChannel: broadcastChannel,
	}
	go webServer.Start()

	commander := services.Commander{
		Commands: commands,
	}
	commander.Start(broadcastChannel)
}
