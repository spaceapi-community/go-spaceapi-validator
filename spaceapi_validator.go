package spaceapiValidator

import (
	"encoding/json"
	"github.com/xeipuuv/gojsonschema"
)

//go:generate go run scripts/generate.go

type spaceApiVersion struct {
	Api string
}

// Result error tells you whats wrong with specific attributes of your SpaceApi file
type ResultError struct {
	Field       string `json:"field"`
	Context     string `json:"context"`
	Description string `json:"description"`
}

// Validation result tells you if the provided string is a valid SpaceApi schema
// and if not tells you what needs to be fixed
type ValidationResult struct {
	Valid  bool          `json:"valid"`
	Errors []ResultError `json:"errors"`
}

// Validate a string to match jsonschema of SpaceApi
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
