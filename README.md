c (v0.0.4)

A build tool for c (like cargo for rust).

(for linux mint!)

## Status

### Working commands:

version - prints this tools version.

init - scaffolds new c project in current directory.

run - builds and runs the project (you can pass cli args to 'run' after one space!)

## Dependencies

- os: linux mint

- shell: bash

- compiler: gcc

- golang 1.25.1 or above

## Building the project

- Clone the repo.

- Inside project folder that you cloned, run:

  `go build`

- The binary will be created in project root.

- After that, just place the output binary where you like and add it to the PATH environment variable.

## Usage

- Print Version:

  `c version`

- New project in current directory:

  `c init`

- Build and run the project:

  `c run`

  To pass cli args:

  `c run hello world`

## License

MIT
