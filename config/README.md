# Config package
This is a package for default config setup

## Usage
The `GetConfig` function takes 3 arguments:
- logger - a logger that implements the `Logger` interface
- envFile - a string that represents the path to the `.env` file
- config - a pointer to a struct that represents the configuration

The `GetConfig` function returns an error if the configuration is not valid.

The `config` struct requires all the member variables to be strings. All the member variables are required by default. If you want to make a member variable optional, you can add the `required:"false"` tag to the member variable.

Here is an example of how to use this package.
```go
logger := logging.NewLoggerWithInfoFunction()

type Config struct {
	Port string // This is required by default
	OptionalSetting string `required:"false"`
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

## TODO
See if there is a way to skip sending in logger. Maybe just remove the print all together when not finding the .env file.