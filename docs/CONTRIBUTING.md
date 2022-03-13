# Contributing Guide

- [Development Setup](#development-setup)
    - [Install Dependencies](#install-dependencies)
    - [Helpful Commands](#helpful-commands)
    - [Resolve Protobuf Imports](#resolve-protobuf-imports)
    - [Running It Locally](#running-it-locally)
- [Package Structure](#package-structure)
- [Making a Change](#making-a-change)

## ⚙️ Development Setup

This guide will walk you through, step by step, how to get a local environment setup for development.

There will be a few precursory dependencies to install, a few commands to get familiar with, and some 
package structure to walk through in order for you to confidently develop on the project.

### 📦 Install Dependencies

Before we get going, we'll need a few software dependencies. All of these are to ease build/test time and to put the burden of **magic** on the tools, instead of you. There might be more than a few here but we'd like to think its because we're being explicit about exactly what's here.

Each will come with an explanation on its' purpose and installation preference, although you may choose to install any way you feel comfortable.

_note: some of the dependencies are installed via `brew`. if you are unfamiliar or would like to check it out, read more at [brew.sh](https://brew.sh/)_

Dependencies:

- `go`:
    - description: `go` is a the language the codebase is written in.
    - installation: 
        ```sh
        $ brew install go
        ```
    - reference: https://go.dev/doc/install
- `protoc`:
    - description: `protoc` is a compiler that will convert `.proto` files, protocol buffer service and message definitions, into language-specific code.
    - installation: 
        ```sh
        $ brew install protobuf
        ```
    - reference: https://grpc.io/docs/protoc-installation/
 - `buf`:
    - description: `buf` is a wrapper around `protoc` that aims to lint, detect breaking changes and reduce `protoc` complexity for developers.
    - installation: 
        ```sh
        $ brew tap bufbuild/buf
        $ brew install buf
        ```
    - reference: https://grpc.io/docs/protoc-installation/
- `protoc-gen-go`:
    - description: `protoc-gen-go` is a plugin that will handle converting the `.proto` file to `.go` file conversion inside the `protoc` toolchain.
    - installation: 
        ```sh
        $ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
        ```
    - reference: https://developers.google.com/protocol-buffers/docs/reference/go-generated
- `protoc-gen-go-grpc`:
    - description: `protoc-gen-go-grpc` is a plugin that will handle converting the services defined in the `.proto` files to a gRPC service interface we can implement against.
    - installation: 
        ```sh
        $ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
        ```
    - reference: https://developers.google.com/protocol-buffers/docs/reference/go-generated
- `golangci-lint`:
    - description: `golangci-lint` is a go linter, aggregation of linters really, that will check for common syntactical misteps throughout the codebase. It does the nitpicking so we can focus on the actual changes during PRs.
    - installation: 
        ```sh
        $ go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.43.0
        ```
    - reference: https://grpc.io/docs/protoc-installation/

### 💻 Helpful Commands

There is a `Makefile` included in the root of the repository with a number of targets, or commands, to help you
do anything from getting a local docker environment spun up to running a thorough test suite.

There is also a helper built into the `Makefile` to tell you exactly what commands are available, see it by running `make help`.

```sh
$ make
 make help                 -> help: display make targets
 make up                   -> runtime: start local environment
 make status               -> runtime: check local environment status
 make down                 -> runtime: stop local environment
 make restart              -> runtime: restart environment
 make build-clean          -> build: clean build workspace
 make build-binary         -> build: build binary file
 make build-proto          -> build: generate proto files
 make test-clean           -> test: clean test workspace
 make test-lint            -> test: check for lint failures
 make test-unit            -> test: execute unit test suite
```

As you can see the commands are put into categories such as `runtime` (without a target prefix), for local runtime start/stop operations, `build` for compiling new proto files or a distribution binary, and `test` for linting and test suites.

### ✨ Resolving Protobuf Imports

If you're using VSCode, there's a little secret to being able resolve project and plugin protobuf imports.

We recommend adding a workspace, non-global, `settings.json` file with the following contents:
```json
{
    "protoc": {
        "options": [
        "--proto_path=${CHECKOUT_LOCATION}/internal/proto",
        ]
    }
}
```

_note: we recommend the use of [global gitignores](https://mike.place/2020/global-gitignore/) rather than include directories and files unbeknownst to the project itself._

### 🏃‍♀️ Running It Locally

If you'd like to poke about the API directly using something like [Insomnia](https://insomnia.rest/) then being able to spin up the service(s) locally is a great way to test things out.

To spin up a complete environment, you can run `make up`. After that you can check that things are healthy and running as expected using `make status`.

Example:
```sh
$ make up
2022-03-11 13:50:44 -0800 [cardmod] cleaning build workspace
2022-03-11 13:50:44 -0800 [cardmod] building binary
2022-03-11 13:50:44 -0800 [cardmod] starting local environment
<... docker vomit ...>

[+] Running 2/2
 ⠿ Network cardmod_default  Created
 ⠿ Container cardmod_api    Started
$ make status
2022-01-15 23:26:12 -0800 [cardmod] checking environment status
NAME                COMMAND             SERVICE             STATUS              PORTS
cardmod_api         "./cardmodd"        api                 running             0.0.0.0:9000->9000/tcp
cardmod_api  | {"level":"INFO","timestamp":"2022-03-11T21:50:46Z","caller":"cardmodd/main.go:54","message":"starting gRPC server","service":{"name":"cardmodd","version":"v0.1.0-alpha"}}
```

## 📁 Package Structure

To see the must up-to-date structure, you can run the command below. 

This is the current structure of the packages with a few comments on the intented purpose for each sub-structure.

```awk
$ tree -aC -I '.git' -I '.vscode' --dirsfirst -d | less -FRX
.
├── .github
│   └── workflows
├── cmd
│   ├── cardmodd        # daemon program for API server(s)
├── docs                # contributing documentation
├── internal
│   ├── api             # concrete gRPC API implementations
│   ├── config          # application configuration
│   ├── daos            # DAO interface definitions
│   ├── logger          # application logger
│   ├── models          # model interface definitions
│   ├── proto
│   │   └── iamnande
│   │       └── cardmod # <resource>/<version>/*.proto
│   ├── repositories    # concrete DAO and model implementations
│   └── server
└── pkg
    └── api             # generated gRPC interfaces
```

## 🖌️ Making a Change

TODO:

- forking
- branch naming
- a good commit structure and message
- testing
- pr description and labels
- rebase vs squash vs merge (who does the merge)