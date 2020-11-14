// Package spaceapivalidator provides a validator for the SpaceAPI schemas
// check spaceapi.io for more information
package spaceapivalidator

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/xeipuuv/gojsonschema"
	"strings"
)

//go:generate go run scripts/generate.go

type spaceAPIVersion struct {
	API              interface{} `json:"api"`
	APICompatibility []string    `json:"api_compatibility"`
}

// ResultError tells you whats wrong with specific attributes of your SpaceApi file
type ResultError struct {
	Field       string `json:"field"`
	Context     string `json:"context"`
	Description string `json:"description"`
}

// VersionValidationResult validation result per version
type VersionValidationResult struct {
	Version	string			`json:"version"`
	Valid	bool			`json:"valid"`
	Errors	[]ResultError	`json:"errors"`
}

// ValidationResult tells you if the provided string is a valid SpaceApi schema
// and if not tells you what needs to be fixed
type ValidationResult struct {
	Valid	bool						`json:"valid"`
	Schemas	[]VersionValidationResult	`json:"version_validation"`
	Errors	[]ResultError				`json:"errors"`
}

// Validate a string to match jsonschema of SpaceApi
func Validate(document string) (ValidationResult, error) {
	myResult := ValidationResult{}
	if document == "" {
		return myResult, errors.New("document is empty")
	}
	documentLoader := gojsonschema.NewStringLoader(document)

	suppliedVersion := spaceAPIVersion{}
	err := json.Unmarshal([]byte(document), &suppliedVersion)
	if err != nil {
		return myResult, err
	}

	versionList := getVersionList(suppliedVersion)

	for _, version := range versionList {
		schemaString, _ := SpaceAPISchemas[version]
		var schema = gojsonschema.NewStringLoader(schemaString)
		result, err := gojsonschema.Validate(schema, documentLoader)
		if err != nil {
			return myResult, err
		}

		var myErrors []ResultError
		for _, error := range result.Errors() {
			myErrors = append(myErrors, ResultError{
				error.Field(),
				error.Context().String(),
				error.Description(),
			})
		}

		myResult.Schemas = append(myResult.Schemas, VersionValidationResult{
			version,
			result.Valid(),
			myErrors,
		})

		myResult.Errors = append(myResult.Errors, myErrors...)

		if result.Valid() {
			myResult.Valid = true
		}
	}

	return myResult, err
}

func getVersionList(suppliedVersion spaceAPIVersion) []string {
	versionList := suppliedVersion.APICompatibility
	oldVersion := strings.Replace(fmt.Sprintf("%v", suppliedVersion.API), "0.", "", 1)
	if oldVersion != "<nil>" {
		versionList = append(versionList, oldVersion)
	}

	if len(versionList) == 0 {
		versionList = []string{"14"}
	}
	return versionList
}
