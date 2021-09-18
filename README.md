# Generate Random Password

[![Go Report Card](https://goreportcard.com/report/github.com/rangow/randpasswd)](https://goreportcard.com/report/github.com/rangow/randpasswd)

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