![commando](./docs/commando.gif)

# Commando

Runs multiple commands in parallel

## Usage

Supply a configuration file containing different commands you want to execute in parallel

```bash
dotnet run -- --config .\example-config.json
```

The configuration file must conform to the following API

```jsonc
{
  "commands": [
    {
      "name": "EnsembleFrontend", // name of the command
      "fileName": "yarn.cmd", // executable to perform the command
      "arguments": "run dev", // arguments to be supplied as part of the command
      "workingDirectory": "C:\\dev\\ensemble\\client" // working directory to perform the command in
    },
    {
      "name": "EnsembleApi",
      "fileName": "go.exe",
      "arguments": "run .",
      "workingDirectory": "C:\\dev\\ensemble\\api"
    }
  ]
}
```
