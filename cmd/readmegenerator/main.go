// blabla
package main

import (
	"bytes"
	"html/template"

	ba "github.com/kamchy/banner"
)

type Repo struct {
	Name  string
	Url   string
	Usage string
}

func generateReadme(r Repo, t *template.Template) string {
	var b bytes.Buffer
	t.Execute(&b, r)
	return b.String()
}

func main() {
	t, err := getTemplate()
	if err != nil {
		panic(err)
	}
	s := generateReadme(
		Repo{"banner", "https://github.com/kamchy/banner", generateHelpMessage()}, t)
	println(s)

}

func getTemplate() (*template.Template, error) {
	var tmpl, err = template.New("test").Parse(`
# Project
{{.Name}} is a simple raster graphics generator that generates randomly a background and two lines of text.

## Warning
[{{.Name}}]({{.Url}}) is my first Go repository cretaed for language learning and is not ready for production.

## Usage
The usage options are as follows:
` + "```" +
		`bash
{{.Usage}}
` + "```" +
		`
## Images
And here are images:
`)

	return tmpl, err
}

func generateHelpMessage() string {
	ifs, _ := ba.InputFlagSet()
	var b bytes.Buffer
	ifs.SetOutput(&b)
	ifs.Usage()
	return b.String()
}
