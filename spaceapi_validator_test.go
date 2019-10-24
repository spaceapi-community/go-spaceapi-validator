package spaceapivalidator

import (
	"testing"
)

var invalid13 = `{ "api": "0.13" }`
var valid13 = `{ "api": "0.13", "open": true, "space": "example", "url": "https://example.com", "logo": "https://example.com/logo.png", "location": { "lon": 42, "lat": 23 }, "state": { "open": true }, "contact": {}, "issue_report_channels": [ "email" ] }`

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
}
