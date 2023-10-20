![commando](./docs/commando.gif)

# Commando

Runs multiple commands in parallel

## Usage

Supply a configuration file containing different commands you want to execute in parallel

### Arguments

```bash
go run .\main.go --config .\example-config.json --tags backend --exclusions EnsembleApi --web=false
```

| command        | required | type         | default | description                                                                     |
| -------------- | -------- | ------------ | ------- | ------------------------------------------------------------------------------- |
| `--config`     | true     | string       |         | points to a configuration file using the schema below                           |
| `--tags`       | false    | string array | []      | when defined, only commands _with_ supplied `tags` in configuration will be run |
| `--exclusions` | false    | string array | []      | when defined, only commands _without_ supplied `name` will be run               |
| `--web`        | false    | bool         | true    | opts out of web view and only outputs using stdout and stderror                 |

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
