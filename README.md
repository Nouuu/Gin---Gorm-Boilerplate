# Gorm Gin boilerplate

## Description

This is a boilerplate for a golang project using gin and gorm.

It includes a basic book crud with some logging and error handling.

## Prerequisites

- [Golang](https://golang.org/doc/install)
- Currently, using version 1.20

> Get the compile daemon for hot reloading

```shell
go get github.com/githubnemo/CompileDaemon
go install github.com/githubnemo/CompileDaemon
```

### Dependencies

| Dependency                 | Version | 	Description                                 | 
|----------------------------|---------|----------------------------------------------|
| github.com/caarlos0/env/v7 | v7.0.0  | Go library for parsing environment variables | 
| github.com/gin-gonic/gin   | v1.9.0	 | Web framework for Go                         | 
| github.com/joho/godotenv	  | v1.5.1  | Go library for loading environment variables | 
| gorm.io/driver/postgres	   | v1.4.8  | Postgres driver for GORM, a Go ORM           |
| gorm.io/gorm               | v1.24.5 | 	Go ORM with support for multiple databases  | 

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

## Installation

> Clone the repository

```shell
git clone https://github.com/Nouuu/Gin-Gorm-Boilerplate.git
```

> Install the dependencies

```shell
cd gin-gorm-boilerplate
go mod download
```

> Copy the .env.example file to .env

```shell
cp .env.example .env
```

> Edit the .env file to match your requirements

| Variable Name | Value Type | Description                                                         |
|---------------|------------|---------------------------------------------------------------------|
| PORT          | int        | The port number on which the server listens                         |
| DB_HOST       | string     | The hostname of the database server                                 |
| DB_PORT       | int        | The port number of the database server                              |
| DB_NAME       | string     | The name of the database to connect to                              |
| DB_USERNAME   | string     | The username to use when connecting to the database                 |
| DB_PASSWORD   | string     | The password to use when connecting to the database                 |
| DB_SYNC       | bool       | Whether to synchronize the database schema with the models          |
| RELEASE_MODE  | bool       | Whether the server runs in release mode                             |
| LOG_LEVEL     | string     | The minimum level of logging messages to output (info, warn, debug) |
| LOG_FILE      | string     | The path to the log file to use                                     |

## Usage

> Run the server

```shell
go run main.go
```

### Initializers

This package is called in the main.go file and initializes the following :

- Variables from the .env file
- The logger
- The database connection
- The gin router

The function [`Init()`](initializers/main.go) is called in the main.go file and initializes all of theses.

#### LoadEnv

This function loads the variables from the .env file into the environment variables.

Then, it uses the [caarlos0/env](github.com/caarlos0/env/v7) to load the environment variables into
the `EnvironmentConf` struct.

[initializers/env.go](initializers/env.go)
```go
package initializers

type environmentConf struct {
 Port        string `env:"PORT" envDefault:"8080"`
 DbHost      string `env:"DB_HOST" envDefault:"localhost"`
 DbPort      string `env:"DB_PORT" envDefault:"5432"`
 DbName      string `env:"DB_NAME" envDefault:"postgres"`
 DbUsername  string `env:"DB_USERNAME" envDefault:"postgres"`
 DbPassword  string `env:"DB_PASSWORD" envDefault:"postgres"`
 DbSync      bool   `env:"DB_SYNC" envDefault:"true"`
 ReleaseMode bool   `env:"RELEASE_MODE" envDefault:"false"`
 LogLevel    string `env:"LOG_LEVEL" envDefault:"debug"`
 LogFile     string `env:"LOG_FILE" envDefault:"gin.log"`
}

var envConf *environmentConf

func loadEnvVariables() error {
 var envCfg = &environmentConf{}

 err := godotenv.Load(".env")
 if err != nil && !os.IsNotExist(err) {
  return err
 }

 if err := env.Parse(envCfg); err != nil {
  return err
 }

 envConf = envCfg
 return nil
}
```


#### InitLogger

This function initializes several loggers :
- **WarningLogger**
- **InfoLogger**
- **DebugLogger**
- **ErrorLogger**

Each logger is initialized with a different log level and an output file.

[initializers/logger.go](initializers/logger.go)
```go
package initializers

func initLogger() error {
 f, err := os.OpenFile(envConf.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
 if err != nil {
  return err
 }
 loggerOutput := io.MultiWriter(f, os.Stdout)
 infoEnabled := envConf.LogLevel == "info" || envConf.LogLevel == "warn" || envConf.LogLevel == "debug"
 warningEnabled := envConf.LogLevel == "warn" || envConf.LogLevel == "debug"
 debugEnabled := envConf.LogLevel == "debug"

 logs.InitLoggers(loggerOutput, infoEnabled, warningEnabled, debugEnabled)
 return nil
}
```
[logs/logs.go](logs/logs.go)
```go
package logs

var (
	warningLogger *log.Logger
	infoLogger    *log.Logger
	debugLogger   *log.Logger
	errorLogger   *log.Logger
)

type logWriter struct {
	writer io.Writer
	enable bool
	level  string
}

func (writer logWriter) Write(bytes []byte) (int, error) {
	if !writer.enable {
		return 0, nil
	}
	message := []byte(fmt.Sprintf("[%s][%s] %s", time.Now().Format(time.RFC3339), writer.level, bytes))
	return writer.writer.Write(message)
}

func InitLoggers(output io.Writer, info, warn, debug bool) {
	warningLogger = log.New(logWriter{writer: output, enable: warn, level: "WARN"}, "", 0)
	infoLogger = log.New(logWriter{writer: output, enable: info, level: "INFO"}, "", 0)
	debugLogger = log.New(logWriter{writer: output, enable: debug, level: "DEBUG"}, "", log.Lshortfile)
	errorLogger = log.New(logWriter{writer: output, enable: true, level: "ERROR"}, "", log.Lshortfile)
}

func InfoPrint(v ...any) {
	infoLogger.Print(v...)
}

func InfoPrintf(format string, v ...any) {
	infoLogger.Printf(format, v...)
}

func InfoPrintln(v ...any) {
	infoLogger.Println(v...)
}

func WarnPrint(v ...any) {
	warningLogger.Print(v...)
}
...
```

The function `Write` customizes the output of the loggers.

#### InitDatabase

This function initializes the database connection and the gorm instance.

if the `DB_SYNC` variable is set to true, the database schema will be synchronized with the models.

[initializers/database.go](initializers/database.go)
```go
package initializers

func connectToDatabase() error {
 uri := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
  envConf.DbHost,
  envConf.DbUsername,
  envConf.DbPassword,
  envConf.DbName,
  envConf.DbPort,
 )

 db, err := gorm.Open(postgres.Open(uri), &gorm.Config{})

 if err != nil {
  return err
 }
 repositories.InitDB(db)

 if envConf.DbSync {
  return autoMigrate(db)
 }

 return nil
}

func autoMigrate(db *gorm.DB) error {
 err := db.AutoMigrate(&entities.BookEntity{})
 if err != nil {
  return err
 }
 return nil
}
```

#### InitGinEngine

When init the gin engine, we set various options :
- Set the release mode
- Set the logger
- Set the recovery middleware
- Set the routes


## License

[![mit](https://img.shields.io/github/license/nouuu/Gin-Gorm-Boilerplate?style=for-the-badge)](LICENSE.md)

## Authors and Contributors

<a href="https://github.com/nouuu/Gin-Gorm-Boilerplate/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=nouuu/Gin-Gorm-Boilerplate"  alt="contributors"/>
</a>