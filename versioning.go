package spaceapivalidator

type VersionSchema int

const (
	// versioning scheme introduced in v0.12: ".api" key
	V12 VersionSchema = iota
	// versioning scheme introduced in v14: ".api_compatibility" list
	V14
)

var SpaceApiVersioning = map[string]VersionSchema {
	"12": V12,
	"13": V12,
	"14": V14,
	"15": V14,
	"16": V14,
}
