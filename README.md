![cmdo banner](./docs/cmdo_promo.png)

## cmdo

`cmdo` ("commando") is a command-line tool which runs multiple commands in parallel.

A common problem when working with microservices is the need to run multiple applications in unison when developing locally; often using different languages and frameworks.

Existing solutions like Docker Compose exist, however, come with the overhead of creating and using Docker containers. Additionally, local development features such as hot-reloading can prove hard to implement when relying on containerisation.

![command view](./docs/cmdo_commands.jpg)

### Features

cmdo provides a simple configuration-driven approach to running services in parallel without the need for containerisation. The key features currently provided are:

- Run multiple commands in different working directories in parallel
- Stdout/stderr output view
- Web output view (see below)
- Easily include and exclude different commands at runtime using tags

#### Web view

The cmdo has an optional web view for displaying command stdout and stderr output. This web view provides the following additional capabilities:

- Command grid view
- Inline unified command view
- Global log searching and command-specific log searching

![terminal view](./docs/cmdo_terminal.jpg)

## Usage

Supply a configuration file containing different commands you want to execute in parallel

### Installation

#### Via Go (recommended)

The simplest way to use and stay up to date with the latest version of cmdo is to retrieve the latest version via a go installation.

This automatically adds it to your path for use:

```bash
go install github.com/ryanachten/cmdo@main # install latest from main branch
cmdo --config .\example-config.json # use cmdo!
```

#### Release builds

If you prefer using prebuilt executables, the latest release can be downloaded from [GitHub](https://github.com/ryanachten/cmdo/releases).

You'll need to add this to your path yourself and download the latest release to receive updates.

![unified view](./docs/cmdo_unified.jpg)

### Arguments

```bash
cmdo --config .\example-config.json --tags backend --exclusions EnsembleApi --web=false
```

| command        | required | type     | default | description                                                                     |
| -------------- | -------- | -------- | ------- | ------------------------------------------------------------------------------- |
| `--config`     | true     | string   |         | points to a configuration file using the schema below                           |
| `--tags`       | false    | string[] | []      | when defined, only commands _with_ supplied `tags` in configuration will be run |
| `--exclusions` | false    | string[] | []      | when defined, only commands _without_ supplied `name` will be run               |
| `--web`        | false    | bool     | true    | opts out of web view and only outputs using stdout and stderror                 |

### Configuration

The configuration file must conform to the following API:

```typescript
  commands: {
    name: string, // label to be presented in the UI
    executable: string, // command to to be executed
    arguments: string[], // arguments to be passed to the command
    workingDirectory: string, // working directory where the command should be executed
    tags?: string[] // (optional) tags used to include exclude commands
  }[],
```

#### Examples

Examples of different configuration can be found in the [examples](./examples/) directory.

One example using the [ensemble](https://github.com/ryanachten/ensemble) project is as follows:

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

## Development

### Architecture

![cmdo model](./docs/cmdo_model.png)

cmdo is comprised of two main parts; a command runner (`commander` in the diagram above), which executes a terminal process for each command the configuration.

The command output is piped is broadcasted via the `webserver` using websockets and piped to stdout/stderr. Commands can be individually stopped or started via the web view and the REST API, the resulting state update is then broadcasted via websockets. The command state and log history is also stored and served via a REST API so it can be persisted on page refresh.

A web app listening to websocket and REST API endpoints is served by cmdo. This intentionally uses a no-build process to keep the dev experience simple. Preact, JS modules, JSDoc and CSS imports are used in favour of alternatives requiring a frontend build step. The web app provides different view and searching capabilities not found in the CLI tool alone.

### Prerequisites

- Go installed locally (>v1.21)

### Running locally

```bash
# from project root directory
go run . --config .\examples\ensemble.cmdo.json
```
