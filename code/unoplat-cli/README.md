# unoplat-cli

`unoplat-cli` is a command-line interface (CLI) application that allows you to install unoplat's components, manage and uninstall them in a simple way.

## Prerequisites

Ensure that you have Go installed on your system and the Go bin directory is in your PATH. Please refer  for the same

Please note that for `unoplat-cli` to work following needs to be installed on your system:-

1. **Go**[Golang](https://go.dev/doc/install)

2. **Step**[Step](https://smallstep.com/docs/step-cli/installation/)

3. **Kubectl**[Kubectl](https://kubernetes.io/docs/reference/kubectl/overview/)

4. **Helm**[helm](https://helm.sh/docs/intro/install/)


## Make File Commands (Local Setup)

1. `make`: Builds the project (equivalent to make all).

2. `make install`: Installs the project.

3. `make test`: Runs tests.

4. `make clean`: Cleans up generated files and binaries.


## Installation

To install `unoplat-cli`, you can use the `go install` command. 

```
go install github.com/unoplat/unoplat-cli/code/unoplat-cli
```

## Usage

Once unoplat-cli is installed, you can run it from the command line by typing:

```
unoplat-cli [command] [options]
```

- Install

```
unoplat-cli install
```
