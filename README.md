# Gorm Gin boilerplate

## Prerequisites

> Get the compile daemon for hot reloading
```shell
go get github.com/githubnemo/CompileDaemon
go install github.com/githubnemo/CompileDaemon
```

## Hot reloading

> Run the following command to start the server

**Windows**
```shell
CompileDaemon -color=true -build="go build -o main.exe" -command=".\main.exe"
```

**Linux**
```shell
CompileDaemon -color=true -build="go build -o main" -command="./main"
```