![main pipeline](https://github.com/gregod-com/grgd/actions/workflows/go.yml/badge.svg) [![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/gregod-com/grgd.svg)](https://github.com/gregod-com/grgd) [![GoReportCard](https://goreportcard.com/badge/github.com/gregod-com/grgd)](https://goreportcard.com/report/github.com/gregod-com/grgd) [![GitHub license](https://img.shields.io/github/license/gregod-com/grgd)](https://github.com/gregod-com/grgd/blob/master/LICENSE)

# grgd

> go package for grgd

grgd is a cli framework extending "github.com/urfave/cli" with user, project and service definitions. Create your own powerful cli with persistence.

## Getting started

Test the embeded example cli:

```shell
go run example/main.go
```

Here you should say what actually happens when you execute the code above.

```bash

                                            _          version   0.0.1                         |
   ___  __  __   __ _   _ __ ___    _ __   | |   ___   profile   gregor                        |
  / _ \ \ \/ /  / _` | | '_ ` _ \  | '_ \  | |  / _ \  project   examplestack                      |
 |  __/  >  <  | (_| | | | | | | | | |_) | | | |  __/  grgdDir   /Users/gregor/.grgd           |
  \___| /_/\_\  \__,_| |_| |_| |_| | .__/  |_|  \___|  hackDir   /Users/gregor/.grgd/hack      |
                                   |_|                 |

example [global options] command [command options] [arguments...]
COMMANDS:
  /// group-1 \\\
      my-commands-group-1  ->  do some stuff

  /// settings \\\
      update               ->  Check and load updates
      profile              ->  Configuration for profiles
      project              ->  Configuration for projects
      service              ->  Configuration for services

  Flags:
      --profile value, -p value  (default: "gregor") [$USER]
      --log-level value          (default: "info")
      --version, -v              print the version (default: false)

```

### Initial Configuration

grgd uses an extendable persistence backend to configure your profiles, projects and services. Per default a sqlite DB is used and safed at ~/.grgd/

## Developing

To start your own implementation define which dependecies you want to rewrite and which you want to reuse:

```go
func main() {
	log := logger.ProvideLogrusLogger()
	dependecies := map[string]interface{}{
		"ILogger":     logger.ProvideLogrusLogger,
		"IConfig":     config.ProvideConfig,
		"IHelper":     helper.ProvideHelper,
		"INetworker":  helper.ProvideNetworker,
		"IDAL":        gormdal.ProvideDAL,
		"IProfile":    profile.ProvideProfile,
		"IUIPlugin":   view.ProvideFallbackUI,
		"my-commands": ProvideCommands,
	}
	core, err := core.RegisterDependecies(dependecies)
	if err != nil {
		log.Fatalf("Error with register dependencies: %s", err.Error())
	}

	grgd.NewApp(core, "example", "0.0.1", nil)
}
```

Feel free to add any function pointers to the list of dependencies. In this example the key `my-commands` is added with the value of pointer to function that returns customized commands.

```go
		"my-commands": ProvideCommands,
```

### Building

Run go build:

```shell
go build -o myapp *.go
```

<!-- ### Deploying / Publishing

In case there's some step you have to take that publishes this project to a
server, this is the right time to state it.

```shell
packagemanager deploy awesome-project -s server.com -u username -p password
```

And again you'd need to tell what the previous code actually does. -->

## Features

<!-- - What's the main functionality
- You can also do another thing
- If you get really randy, you can even do this -->

## Configuration

<!-- Here you should write what are all of the configurations a user can enter when
using the project. -->

## Contributing

"If you'd like to contribute, please fork the repository and use a feature
branch. Pull requests are warmly welcome."

## Links

<!-- - Project homepage: https://your.github.com/awesome-project/ -->

- Repository: https://github.com/gregod-com/grgd
- Issue tracker: https://github.com/gregod-com/grgd/issues
- Related projects:
  - urfave cli: https://github.com/urfave/cli

## Licensing

"The code in this project is licensed under MIT license."
