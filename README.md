spaceapi-validator
==

This is a simple library for golang to validate spaceapi specs

Right now it uses the SpaceApi spec found at https://github.com/spacedirectory/schema until the spec is moved

usage
==
```go
package main

import (
	validator "github.com/gidsi/go-spaceapi-validator"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Validate with following commigithub.com/gidsi/go-spaceapi-validatort: " + validator.CommitHash)
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