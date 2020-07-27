# Go Playground

This repo is setup to learn Go.

## Installation

Everything is expected to be run within a docker container, using the provided Make commands.
Install [Docker-Compose](https://docs.docker.com/install/) and run `docker-compose up` from the root of the project.

This will spin up a single container setting your working directory to the app location.

## Running

To start the webserver, simply run `make run` from the container.

# Packages

* github.com/canthefason/go-watcher
  * This package is used to automatically reload the app on any changes to the Go files.
* github.com/gorilla/mux
  * gorilla/mux is a common routing library, that allows easier access to route components
* [github.com/golang-migrate/migrate](https://github.com/golang-migrate/migrate)
  * Manages database migrations
