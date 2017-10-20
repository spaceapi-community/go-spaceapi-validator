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
	fmt
)

func main() {
	fmt.Println("Validate with following commit: " + validator.CommitHash)
	validator.Validate(/* spaceapi json you want to validate */)
}
```

Todo
==
* Update schemas.go file when something is pushed into schema respository