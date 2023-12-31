package models

import (
	"flag"
	"os"
	"slices"
)

type Arguments struct {
	ConfigurationPath string
	Tags              []string
	Exclusions        []string
	UseWeb            bool
}

// Parses command line arguments
func GetArguments() *Arguments {
	var configPath string
	flag.StringVar(&configPath, "config", "", "Path for cmdo configuration file")

	var tags stringArray
	flag.Var(&tags, "tags", "List of tags indicating which commands include as part of the cmdo execution")

	var exclusions stringArray
	flag.Var(&exclusions, "exclusions", "List of command names indicating which commands exclude as part of the cmdo execution")

	var useWeb bool
	flag.BoolVar(&useWeb, "web", true, "Opt out of displaying output using a web dashboard")

	flag.Parse()

	if configPath == "" {
		println("Missing -config flag with path to configuration file.\n")
		println("cmdo usage:")
		flag.PrintDefaults()
		os.Exit(1)
		return nil
	}

	args := Arguments{
		ConfigurationPath: configPath,
		Tags:              tags,
		Exclusions:        exclusions,
		UseWeb:            useWeb,
	}

	return &args
}

// Filters commands using command line arguments
func (args Arguments) FilterCommands(commands []Command) []Command {
	filteredCommands := []Command{}

	for _, command := range commands {
		if len(args.Exclusions) > 0 && slices.Contains(args.Exclusions, command.Name) {
			continue
		}

		if len(args.Tags) == 0 {
			filteredCommands = append(filteredCommands, command)
			continue
		}

		for _, tag := range command.Tags {
			if slices.Contains(args.Tags, tag) {
				filteredCommands = append(filteredCommands, command)
				break
			}
		}
	}

	return filteredCommands
}

// Custom flag for parsing an array of strings
type stringArray []string

func (arr *stringArray) String() string {
	return ""
}

func (arr *stringArray) Set(value string) error {
	*arr = append(*arr, value)
	return nil
}
