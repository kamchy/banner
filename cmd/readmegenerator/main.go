package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"text/template"

	ba "github.com/kamchy/banner"
)

type Repo struct {
	Name  string
	Url   string
	Usage string
}

type Image struct {
	Alt string
	Url string
}
type TemplateData struct {
	Repo
	Images []Image
}

func generateReadme(wr io.Writer, r TemplateData, t *template.Template) {
	t.Execute(wr, r)
}

func makeInput(algIdx int, palidx int, fname string) ba.Input {
	t1, t2 := "Go programming", "for fun and profit"
	w, h := 400, 200
	tileSize := 30.0
	return ba.Input{
		W:        &w,
		H:        &h,
		Texts:    []*string{&t1, &t2},
		AlgIdx:   &algIdx,
		TileSize: &tileSize,
		Pt:       &palidx,
		OutName:  &fname}
}

func generateImages(cwd string, dirName string) ([]Image, error) {
	algsCount := len(ba.PainterAlgs)
	infosCount := len(ba.PaletteInfos)
	images := make([]Image, algsCount*infosCount, algsCount*infosCount)
	var abs string
	for algidx, alg := range ba.PainterAlgs {
		for palidx, _ := range ba.PaletteInfos {
			fName := fmt.Sprintf("out_alg%d_pal%d.png", algidx, palidx)
			if filepath.IsAbs(dirName) {
				abs = path.Join(dirName, fName)
			} else {
				abs = path.Join(cwd, dirName, fName)
			}
			relToCurrent, err := filepath.Rel(cwd, abs)
			if err != nil {
				panic(err)
			}
			images[algidx*infosCount+palidx] = Image{alg.Desc(), relToCurrent}
			ba.GenerateBanner(makeInput(algidx, palidx, abs))
		}
	}
	return images, nil
}

func getTemplate() (*template.Template, error) {
	var tmpl, err = template.New("test").Parse(`
# Project
{{.Name}} is a simple raster graphics generator that generates randomly a background and two lines of text.

## Warning
[{{.Name}}]({{.Url}}) is my first Go repository cretaed for language learning and is not ready for production.

## Example
This is default image (when no commandline options are provided)
![example](img/default.png)

## Usage
The usage options are as follows:

` + "```" +
		`bash
{{.Usage}}
` + "```" +
		`
## Readme generator
The project also contains readme generator binary (in cmd/readmegenerator/main.go)
which takes path to image directory where it generates images, and writes this file's
contents to stdout, with image linked to this markdown file.

For details, see [the source](https://github.com/kamchy/banner/blob/main/src/readmegenerator/main.go)
### Usage

` + "```" +
		`bash
cd /cmd/readmegenerator && go build
cd ../..
./cmd/readmegenerator/readmegenerator img > README.md
` + "```" +
		`

## Images
And here are images:

{{ range .Images }}
### Image {{.Url}}
![{{.Alt}}]({{.Url}})
{{ end }}
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

func main() {
	t, err := getTemplate()
	if err != nil {
		panic(err)
	}
	repo := Repo{
		"banner",
		"https://github.com/kamchy/banner",
		generateHelpMessage()}

	dirName := flag.String("o", "/tmp", "output directory for generated images")
	readmeName := flag.String("r", "/tmp/README.md", "output file name")
	flag.Parse()

	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	images, err := generateImages(cwd, *dirName)
	if err != nil {
		panic(err)
	}

	td := TemplateData{repo, images}
	tf, err := os.Create(*readmeName)
	defer func() { fmt.Println("Closing ", *readmeName); tf.Close() }()

	if err != nil {
		panic(err)
	}
	generateReadme(tf, td, t)

}
