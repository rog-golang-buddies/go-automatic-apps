

# Example App 1

Example app to use GAA

## Setup

```shell
go mod init exampleapp1
go get github.com/rog-golang-buddies/go-automatic-apps
go get -d entgo.io/ent/cmd/ent
```

## Add model

```shell
go run entgo.io/ent/cmd/ent init Todo
```

Then, generate its code

go generate ./ent

## Run

```shell
go run main.go
```

