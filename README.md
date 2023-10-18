![commando](./docs/commando.gif)

# Commando

Runs multiple commands in parallel

## Usage

Supply a configuration file containing different commands you want to execute in parallel

### Arguments

```bash
go run .\main.go --config .\example-config.json --tags backend --exclusions EnsembleApi
```

- `--config` (required) - points to a configuration file using the schema described below
- `--tags` (optional) - when defined, only commands _with_ supplied `tags` in their configuration will be run
- `--exclusions` (optional) - when defined, only commands _without_ a supplied `name` in their configuration will be run

### Configuration

The configuration file must conform to the following API

```json
{
  "commands": [
    {
      "name": "EnsembleFrontend",
      "executable": "yarn",
      "arguments": ["run", "dev"],
      "workingDirectory": "C:\\dev\\ensemble\\client",
      "tags": ["frontend"]
    },
    {
      "name": "EnsembleApi",
      "executable": "go",
      "arguments": ["run", "."],
      "workingDirectory": "C:\\dev\\ensemble\\api",
      "tags": ["backend"]
    }
  ]
}
```
