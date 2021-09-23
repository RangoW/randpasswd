# Generate Random Password

[![Go Report Card](https://goreportcard.com/badge/github.com/rangow/randpasswd)](https://goreportcard.com/badge/github.com/rangow/randpasswd)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/gojp/goreportcard/blob/master/LICENSE)

## Range

The default password character range is **[0-9,a-zA-Z,#!]**, which can be freely modified.

## Test

```golang
$ go test -v -run TestGenerate
=== RUN   TestGenerate
    symbol_test.go:9: <nil>       
    symbol_test.go:10: 11
    symbol_test.go:11: 6TvJSGi9jyH
--- PASS: TestGenerate (0.00s)    
PASS
ok      github.com/rangow/randpasswd    0.065
```
