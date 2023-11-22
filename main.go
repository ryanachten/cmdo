package main

import (
	"log"

	"github.com/ryanachten/cmdo/events"
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

	eventBus := events.CreateEventBus()

	if arguments.UseWeb {
		webServer := services.WebServer{
			EventBus: eventBus,
		}
		go webServer.Start()
	}

	commander := services.Commander{
		Commands: commands,
		EventBus: eventBus,
		UseWeb:   arguments.UseWeb,
	}
	commander.Start()
}
