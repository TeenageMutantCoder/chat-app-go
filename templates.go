package main

import (
	"embed"
	"html/template"
)

//go:generate npm run build

const (
	StartingTemplateDirectory = "templates/"
	TemplateFileType          = ".go.html" // Make sure this is in sync with .prettierrc, package.json, and tailwind.config.js
)

//go:embed static
var staticFiles embed.FS

var chatTemplate = template.Must(ParseFiles("base", "chat"))

func ParseFiles(filenames ...string) (*template.Template, error) {
	for index := range filenames {
		filenames[index] = StartingTemplateDirectory + filenames[index] + TemplateFileType
	}
	return template.ParseFiles(filenames...)
}
