# KStub

The Kubernetes manifest stub generator for the rest of us. Because copy and pasting from the Documentation, as well trying to remember the manifest specs are a waste of time and error-prone. This tool does not presume to be a Kubernetes CLI, but rather a simple fast generator for starting points of kubernetes manifests that you can further customize to your needs.

<p align="center">
    <img src="https://cdn.rawgit.com/nicbet/kstub/master/demo.svg">
</p>

## Installation

Run `go get github.com/nicbet/kstub`

## Usage

```sh
KStub is a very fast generator for Kubernetes manifests

Usage:
  kstub [flags] command

Available Commands:
  help        Help about any command
  version     Print the version number of KStub

  deployment  Generate a deployment manifest
  service     Generate a service manifest
  ingress     Generate a ingress manifest
  
Flags:
      --config string      /path/to/config.yml
  -h, --help               help for kstub
      --log-level string   Output level of logs (TRACE, DEBUG, INFO, WARN, ERROR, FATAL) (default "INFO")
  -v, --version            Display the current version of this CLI

Use "kstub [command] --help" for more information about a command.
```

Most flags can be used across commands. For instance to generate a stack of manifests for an app called "echo", one could run:

```sh
kstub --name echo --port 80 deployment
kstub --name echo --port 80 service
kstub --name echo --port 80 ingress
```