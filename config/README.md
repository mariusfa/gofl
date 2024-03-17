# Config package
This is a package for default config setup

This package for local development reads `.env` file. For production it reads environment variables. If missing ENV variables it will return error.

All variables in config struct will be required by default.

## Usage
Here is an example of how to use this package.
```go
logger := logging.NewLoggerWithInfoFunction()

type Config struct {
	Port string
}

var config Config

err := GetConfig(fake, ".env", &config)
if err != nil {
	t.Fatal(err)
}

```

An example of the logger impl you can find in the logging package app-log.


For local dev use `.env file. Example of file:
```bash
PORT=8080
```