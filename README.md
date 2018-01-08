go-spaceapi-validator
==
[![Go Report Card](https://goreportcard.com/badge/github.com/gidsi/go-spaceapi-validator)](https://goreportcard.com/report/github.com/gidsi/go-spaceapi-validator)

This is a simple library for golang to validate spaceapi specs

usage
==
```go
package main

import (
	validator "github.com/spaceapi-community/go-spaceapi-validator"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Validate with following commit: " + validator.CommitHash)
	result, err := validator.Validate(`{}`/* spaceapi json you want to validate */)
	
	if err != nil {
		panic(err)
	}
	
	fmt.Println("Provided data is valid: " + strconv.FormatBool(result.Valid))
	fmt.Println(result.Errors)
}
```

Todo
==
* Update schemas.go file when something is pushed into schema respository
