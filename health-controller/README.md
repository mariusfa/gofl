# Health Controller

This is package for creating a health controller for stdlib http server.
The package is based of the `health` package.

## Usage
```go
router := http.NewServeMux()
hc := health.NewController()
hc.RegisterRoutes(router)

err := http.ListenAndServe(":8080", router)
if err != nil {
    log.Fatal(err)
}
```

