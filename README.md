# Go Automatic Apps (GAA)
[![pre-commit](https://img.shields.io/badge/pre--commit-enabled-brightgreen?logo=pre-commit&logoColor=white)](https://github.com/pre-commit/pre-commit)
[![pre-commit.ci status](https://results.pre-commit.ci/badge/github/rog-golang-buddies/golang-template-repository/main.svg)](https://results.pre-commit.ci/latest/github/rog-golang-buddies/golang-template-repository/main)


Generate CRUD applications automatically with the least effort and as beautiful as possible

## How Go Automatic Apps is used? (To be done)

1. Create a new Go project
2. Install Go Automatic Apps dependency
3. Define your database file (schema name and connection parameters)
4. Define your table models (using ent.io)
5. Instantiate Go Automatic Apps server
6. Done

# How it works? (To be done)

## GAA CLI

Command line to execute actions:

- Create a new model
- Create migrations
- Run migrations

## GAA Server

The server that handles the CRUD UI (React) and the API.

The server does it automatically by reading the models definitions.

1. Autogenerated UI for Create/Read/Update/Delete (CRUD) rows from a model
2. Autogenerated API to respond to the UI actions


# Contribute

## Setup

* Go 1.18

## Folders

* `architecture`: Documents and diagrams
* `cmd`: Code for the **gaa** command
* `pkg`: Exportable libraries
* `internal`: Utility functions
* `server`: Server side code

