package spaceapivalidator

import (
	"strings"
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
var unknownApiCompatibility = `{ "api_compatibility": [ "142" ] }`
var missingIssueReportChannel13 = `{
  "api": "0.13",
  "api_compatibility": ["14"],
  "space": "NYCResistor",
  "logo": "https://www.nycresistor.com/wp-content/uploads/2017/09/NYCR-logo-wide.png",
  "url": "https://www.nycresistor.com/",
  "location": {
    "address": "87 3rd Avenue, 4th floor, Brooklyn, NY 11217, USA",
    "lon": 40.683653,
    "lat": -73.981542,
    "timezone": "America/New_York"
  },
  "contact": {
    "email": "contact@nycresistor.com",
    "matrix": "#public:nycr.chat",
    "facebook": "https://www.facebook.com/groups/23644223800/",
    "ml": "nycresistormicrocontrollers@googlegroups.com"
  },
  "state": {
    "open": false
  },
  "links": [
   {
     "name": "eventbrite",
     "url": "http://nycresistor.eventbrite.com/"
   },
   {
     "name": "code of conduct",
     "url": "https://www.nycresistor.com/participate/"
   },
   {
     "name": "withfriends",
     "url": "https://withfriends.co/nyc_resistor/join"
   }
  ],
  "membership_plans": [
    {
      "name": "standard",
      "value": 115,
      "currency": "USD",
      "billing_interval": "monthly"
    },
    {
      "name": "teaching",
      "value": 75,
      "currency": "USD",
      "billing_interval": "monthly"   
    }
  ],
  "feeds": {
    "blog": {
      "type": "rss",
      "url": "https://www.nycresistor.com/feed/"
    },
    "wiki": {
      "type": "atom",
      "url": "https://wiki.nycresistor.com/w/api.php?hidebots=1&urlversion=1&days=30&limit=50&action=feedrecentchanges&feedformat=atom"
    },
    "calendar": {
      "type": "ical",
      "url": "https://calendar.google.com/calendar/ical/p2m2av9dhrh4n1ub7jlsc68s7o%40group.calendar.google.com/public/basic.ics"
    }
  }
}`
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

	invalidResult, _ = Validate(missingIssueReportChannel13)
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
		t.Error("Schema should have got 6 errors, got", len(invalidErrors))
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

	// Test handling of unknown schema versions
	invalidResult, err = Validate(unknownApiCompatibility)
	if err != nil {
		t.Error("validation error shouldn't show up on valid json")
	} else if invalidResult.Valid == true {
		t.Error("Expected validation to be false, got true")
	}
	invalidErrors = invalidResult.Errors
	if len(invalidErrors) != 1 {
		t.Logf("%v", invalidResult)
		t.Error("Schema should have got 1 errors, got", len(invalidErrors))
	}
	if !strings.Contains(invalidErrors[0].Description, "Endpoint declares compatibility with schema version 142, which isn't supported") {
		t.Error("Did not find expected 'unknown schema version' error:", invalidErrors[0].Description)
	}
}
