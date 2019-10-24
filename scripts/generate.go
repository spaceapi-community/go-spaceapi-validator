package main

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

var repoRef = "spaceapi/schema"

type githubFile struct {
	Url      string
	Name     string
	Content  string
	Encoding string
}

type githubCommitObject struct {
	Sha string
}

type githubCommit struct {
	Object githubCommitObject
}

func main() {
	writeSchemaGoFile(
		getCommitHash(),
		getSchemaFiles(),
	)
}

func getCommitHash() string {
	headCommitResponse, err := http.Get("https://api.github.com/repos/" + repoRef + "/git/refs/heads/master")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer headCommitResponse.Body.Close()
	githubCommit := githubCommit{}
	json.NewDecoder(headCommitResponse.Body).Decode(&githubCommit)
	return githubCommit.Object.Sha
}

func getSchemaFiles() []githubFile {
	response, err := http.Get("https://api.github.com/repos/" + repoRef + "/contents")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer response.Body.Close()

	var githubFiles = []githubFile{}
	json.NewDecoder(response.Body).Decode(&githubFiles)

	for index, file := range githubFiles {
		if strings.HasSuffix(file.Name, ".json") {
			response, err := http.Get(file.Url)
			if err == nil {
				json.NewDecoder(response.Body).Decode(&file)
				githubFiles[index] = file
			}
			response.Body.Close()
		}
	}

	return githubFiles
}

func writeSchemaGoFile(commitHash string, files []githubFile) {
	out, _ := os.Create("schemas.go")
	out.Write([]byte("package spaceapivalidator\n\n// CommitHash contains the hash of the commit the Validate function validates against\nvar CommitHash = \"" + commitHash + "\"\n\n// SpaceAPISchemas load from the repository as a map\nvar SpaceAPISchemas = map[string]string{\n"))
	for _, f := range files {
		if strings.HasSuffix(f.Name, ".json") {
			fileContent, err := b64.StdEncoding.DecodeString(f.Content)

			if err == nil {
				out.Write([]byte("\t\"" + strings.TrimSuffix(strings.TrimSuffix(f.Name, ".json"), "-draft") + "\": `"))
				out.Write(fileContent)
				out.Write([]byte("`,\n"))
			}
		}
	}
	out.Write([]byte("}\n"))
}
