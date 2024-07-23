## Acknowledgements

This project includes code from the [Google's UUID package](https://pkg.go.dev/github.com/google/uuid). The original code can be found [here](https://github.com/google/uuid/tree/v1.6.0).

The original contributors are:

Paul Borman <borman@google.com>
bmatsuo
shawnps
theory
jboverfelt
dsymonds
cd1
wallclockbuilder
dansouza

# uuid

The uuid package generates and inspects UUIDs based on
[RFC 4122](https://datatracker.ietf.org/doc/html/rfc4122)
and DCE 1.1: Authentication and Security Services.

This package is based on the github.com/pborman/uuid package (previously named
code.google.com/p/go-uuid). It differs from these earlier packages in that
a UUID is a 16 byte array rather than a byte slice. One loss due to this
change is the ability to represent an invalid UUID (vs a NIL UUID).

###### Install

```sh
go get github.com/google/uuid
```

###### Documentation

[![Go Reference](https://pkg.go.dev/badge/github.com/google/uuid.svg)](https://pkg.go.dev/github.com/google/uuid)

Full `go doc` style documentation for the package can be viewed online without
installing this package by using the GoDoc site here:
https://pkg.go.dev/github.com/google/uuid
