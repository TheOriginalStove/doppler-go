# Doppler-Go

An unofficial Go wrapper for the (https://docs.doppler.com/reference#download)[Doppler API]


## Installation

Using Go Modules in your project (There will be a `go.mod` file in its root if it is).

Otherwise to begin using Go Modules use the following command:

```
go mod init
```

Then reference doppler-go in a Go program with import 

```go
import (
    "github.com/theoriginalstove/doppler-go"
)
```

Then run any of the normal `go` commands to use the Go toolchain in order to resolve and fetch the doppler-go module automatically.

An alternative is to explicitly call `go get` the package into a project.

```sh
go get -u github.com/theoriginalstove/doppler-go"
```


