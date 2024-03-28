# gofl
Go framework library

## Packages
- [config](./config/README.md)
- [health](./health/README.md)
- [health-controller](./health-controller/README.md)
- logging
    - [app-log](./logging/app-log/README.md)
    - [transaction-log](./logging/transaction-log/README.md)
- [circuit-breaker](./circuit-breaker/README.md)
- [database](./database/README.md)


## Install 
```bash
go get github.com/mariusfa/gofl/v2
```

## Local development
Add this to `go.mod`
```bash
// For local dev
replace github.com/mariusfa/gofl => ../../gofl
```

