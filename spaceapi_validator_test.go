package spaceapivalidator

import (
	"testing"
)

var wrongVersion = `{ "api": "0.5" }`
var invalid13 = `{ "api": "0.13" }`
var valid13 = `{ "api": "0.13", "open": true, "space": "example", "url": "https://example.com", "logo": "https://example.com/logo.png", "location": { "lon": 42, "lat": 23 }, "state": { "open": true }, "contact": {}, "issue_report_channels": [ "email" ] }`
var valid14 = `{ "api_compatibility": [ "14" ], "space": "example", "url": "https://example.com", "logo": "https://example.com/logo.png", "location": { "lon": 42, "lat": 23 }, "state": { "open": true }, "contact": {} }`
var invalid14 = `{ "api_compatibility": [ "14" ], "space_invalid": "example", "url": "https://example.com", "logo": "https://example.com/logo.png", "location": { "lon": 42, "lat": 23 }, "state": { "open": true }, "contact": {} }`
var valid15 = `{ "api_compatibility": [ "15" ], "space": "example", "url": "https://example.com", "logo": "https://example.com/logo.png", "location": { "lon": 42, "lat": 23 }, "state": { "open": true }, "contact": {} }`
var invalid15 = `{ "api_compatibility": [ "15" ], "space_invalid": "example", "url": "https://example.com", "logo": "https://example.com/logo.png", "location": { "lon": 42, "lat": 23 }, "state": { "open": true }, "contact": {} }`
var wrongVersionNumeric = `{ "api": 0.13, "open": true, "space": "example", "url": "https://example.com", "logo": "https://example.com/logo.png", "location": { "lon": 42, "lat": 23 }, "state": { "open": true }, "contact": {}, "issue_report_channels": [ "email" ] }`
var noVersion = `{ "data": "asd" }`

func TestValidate(t *testing.T) {
	invalidResult, _ := Validate(invalid13)
	if invalidResult.Valid == true {
		t.Error("Expected validation to be false, got", invalidResult.Valid)
	}

	invalidErrors := invalidResult.Errors
	if len(invalidErrors) != 7 {
		t.Error("Schema should have got 7 errors, got", len(invalidErrors))
	}

	validResult, _ := Validate(valid13)
	if validResult.Valid == false {
		t.Error("Expected validation to be true, got", validResult.Valid)
	}
	validErrors := validResult.Errors
	if len(validErrors) != 0 {
		t.Error("Schema should have got no errors, got", len(validErrors))
	}

	validResult, _ = Validate(valid14)
	if validResult.Valid == false {
		t.Error("Expected validation to be true, got", validResult.Valid)
	}
	validErrors = validResult.Errors
	if len(validErrors) != 0 {
		t.Error("Schema should have got no errors, got", len(validErrors))
	}

	invalidResult, _ = Validate(invalid14)
	if invalidResult.Valid == true {
		t.Error("Expected validation to be false, got", invalidResult.Valid)
	}
	invalidErrors = invalidResult.Errors
	if len(invalidErrors) != 1 {
		t.Error("Schema should have got 1 errors, got", len(invalidErrors))
	}

	validResult, _ = Validate(valid15)
	if validResult.Valid == false {
		t.Error("Expected validation to be true, got", validResult.Valid)
	}
	validErrors = validResult.Errors
	if len(validErrors) != 0 {
		t.Error("Schema should have got no errors, got", len(validErrors))
	}

	invalidResult, _ = Validate(invalid15)
	if invalidResult.Valid == true {
		t.Error("Expected validation to be false, got", invalidResult.Valid)
	}
	invalidErrors = invalidResult.Errors
	if len(invalidErrors) != 1 {
		t.Error("Schema should have got 1 errors, got", len(invalidErrors))
	}

	invalidResult, _ = Validate(noVersion)
	if invalidResult.Valid == true {
		t.Error("Expected validation to be false, got", invalidResult.Valid)
	}
	invalidErrors = invalidResult.Errors
	if len(invalidErrors) != 6 {
		t.Logf("%v", invalidResult)
		t.Error("Schema should have got 1 errors, got", len(invalidErrors))
	}

	validResult, err := Validate("")
	if err == nil {
		t.Error("Should provide an error on faulty json")
	}

	invalidResult, _ = Validate(wrongVersion)
	if invalidResult.Valid == true {
		t.Error("Expected validation to be false, got", invalidResult.Valid)
	}

	invalidResult, err = Validate(wrongVersionNumeric)
	if err != nil {
		t.Error("validation error shouldn't show up on valid json")
	} else if invalidResult.Valid == true {
		t.Error("Expected validation to be false, got", invalidResult.Valid)
	}
}
