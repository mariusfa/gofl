# App log package

This is a package for general logging.

## Usage

Here is an example of how to use this package.
Start with creating a package in `internal/logging` and add a file `logging.go` with the following content:
```go
package logging

import applog "github.com/mariusfa/gofl/v2/logging/app-log"

var AppLogger *applog.AppLogger

func SetupAppLogger(appName string) {
	AppLogger = applog.NewAppLogger(appName)
}
```

In the main.go file, you can use the logger like this:
```go
package main

import (
	"<my-module>/internal/logging"

	"github.com/mariusfa/gofl/v2/config"
)

func main() {
	logging.SetupAppLogger("todo")
	appLogger := logging.AppLogger

    // ... rest of the code
}
