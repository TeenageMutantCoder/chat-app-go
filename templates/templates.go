package templates

import (
	"embed"
	"html/template"
)

//go:generate npm run tw:build

const (
	StartingTemplateDirectory = "templates/"
	TemplateFileType          = ".go.html" // Make sure this is in sync with .prettierrc, package.json, and tailwind.config.js
)

//go:embed static
var StaticFiles embed.FS

var (
	ChatTemplate          = template.Must(ParseFiles("base", "chat"))
	NameSelectionTemplate = template.Must(ParseFiles("base", "name_selection"))
)

func ParseFiles(filenames ...string) (*template.Template, error) {
	for index := range filenames {
		filenames[index] = StartingTemplateDirectory + filenames[index] + TemplateFileType
	}
	return template.ParseFiles(filenames...)
}
