![commando](./docs/commando.gif)

# Commando

Runs multiple commands in parallel

## Usage

Supply a configuration file containing different commands you want to execute in parallel

```bash
go run .\main.go --config .\example-config.json
```

The configuration file must conform to the following API

```jsonc
{
  "commands": [
    {
      "name": "EnsembleFrontend",
      "executable": "yarn",
      "arguments": ["run", "dev"],
      "workingDirectory": "C:\\dev\\ensemble\\client"
    },
    {
      "name": "EnsembleApi",
      "executable": "go",
      "arguments": ["run", "."],
      "workingDirectory": "C:\\dev\\ensemble\\api"
    }
  ]
}
```
