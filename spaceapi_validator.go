package spaceapi_validator

import (
	"github.com/xeipuuv/gojsonschema"
	"encoding/json"
)

//go:generate go run scripts/generate.go

type spaceApiVersion struct {
	Api string
}

type ResultError struct {
	Field string `json:"field"`
	Context string `json:"context"`
	Description string `json:"description"`
}

type ValidationResult struct {
	Valid bool `json:"valid"`
	Errors []ResultError `json:"errors"`
}

func Validate(document string) (ValidationResult, error) {
	documentLoader := gojsonschema.NewStringLoader(document)

	suppliedVersion := spaceApiVersion{}
	json.Unmarshal([]byte(document), &suppliedVersion)

	schemaString, ok := SpaceApiSchemas[suppliedVersion.Api]
	if !ok {
		schemaString = SpaceApiSchemas["13"]
	}
	var schema = gojsonschema.NewStringLoader(schemaString)

	result, err := gojsonschema.Validate(schema, documentLoader)

	var myErrors = []ResultError{}
	for _, bar := range result.Errors() {
		myErrors = append(myErrors, ResultError{
			bar.Field(),
			bar.Context().String(),
			bar.Description(),
		})
	}

	myResult := ValidationResult{
		result.Valid(),
		myErrors,
	}

	return myResult, err
}