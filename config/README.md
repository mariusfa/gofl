# Config package
This is a package for default config setup

This package for local development reads `.env` file. For production it reads environment variables. If missing ENV variables it will return error.

## Usage
Here is an example of how to use this package.
```go
logger := logging.NewLoggerWithInfoFunction()
config, err := config.GetConfig(logger)
if err != nil {
	panic(err)
}
```

An example of the logger impl you can find in the logging package app-log.
