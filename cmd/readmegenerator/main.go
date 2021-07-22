package main

import (
	"bytes"
	"flag"
	"fmt"
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

func makeDefaultInput(fileName string) ba.Input {
	var id = ba.InpData{
		W:   ba.DEF_WIDTH,
		H:   ba.DEF_HEIGHT,
		T:   ba.DEF_TITLE,
		St:  ba.DEF_SUB,
		Alg: ba.DEF_ALG,
		Ts:  ba.DEF_TILE,
		P:   ba.DEF_PAL,
		O:   fileName,
	}
	return new(ba.Input).From(id)
}
func absolutePath(cwd string, dirName string, fName string) string {
	var abs string
	if filepath.IsAbs(dirName) {
		abs = path.Join(dirName, fName)
	} else {
		abs = path.Join(cwd, dirName, fName)
	}
	return abs
}

func generateImages(imgDirName string) ([]Image, error) {
	algsCount := len(ba.PainterAlgs)
	infosCount := len(ba.PaletteInfos)
	images := make([]Image, algsCount*infosCount, algsCount*infosCount)
	for algidx, alg := range ba.PainterAlgs {
		for palidx, _ := range ba.PaletteInfos {
			fName := fmt.Sprintf("out_alg%d_pal%d.png", algidx, palidx)
			images[algidx*infosCount+palidx] = Image{alg.Desc(), filepath.Join("img", fName)}
			GenerateBanner(makeInput(algidx, palidx, filepath.Join(imgDirName, fName)))
		}
	}
	GenerateBanner(makeDefaultInput(filepath.Join(imgDirName, "default.png")))
	return images, nil
}

func GenerateBanner(i ba.Input) {
	fmt.Printf("Generating  %+v\n", new(ba.InpData).From(i))
	ba.GenerateBanner(i)
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
cd /cmd/readmegenerator && go build .
cd ../..
./cmd/readmegenerator/readmegenerator ../..
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
func createImgPath(targetDir string) string {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	imgPath := filepath.Join(cwd, targetDir, "img")
	return imgPath
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

	targetDir := flag.String("o", ".", "directory where img subdirectory will be created and filled with images and README.md file generated")
	flag.Parse()

	imageDescriptions, err := generateImages(createImgPath(*targetDir))
	if err != nil {
		panic(err)
	}

	td := TemplateData{repo, imageDescriptions}

	rfile := filepath.Join(*targetDir, "README.md")
	tf, err := os.Create(rfile)
	fmt.Println("Creating ", rfile)
	defer func() { fmt.Println("Closing ", rfile); tf.Close() }()

	if err != nil {
		panic(err)
	}
	t.Execute(tf, td)

}
