
# Project
banner is a simple raster graphics generator that generates randomly a background and two lines of text.

## Warning
[banner](https://github.com/kamchy/banner) is my first Go repository cretaed for language learning and is not ready for production.

## Example
This is default image (when no commandline options are provided)
![example](img/default.png)

## Usage
The usage options are as follows:

```bash
Usage of input:
  -alg int
    	Background painter algorithm; valid values are: 
    	4 -&gt; concentric circles offset
    	5 -&gt; random horizontal lines
    	6 -&gt; random hexagons
    	7 -&gt; random hexagons with offset
    	0 -&gt; random rectangles
    	1 -&gt; random rectangles with offset
    	2 -&gt; plain color
    	3 -&gt; concentric circles
    	 (default 5)
  -height int
    	height of the resulting image (default 600)
  -outName string
    	name of output file where banner in .png format will be saved (default &#34;out.png&#34;)
  -palette int
    	palette type; valid values are: 
    	0 -&gt; Warm
    	1 -&gt; Happy
    	
  -subtext string
    	explanatory text to display in the image below the text (default &#34;this time about really important things&#34;)
  -text string
    	text to display in the image (default &#34;My blogpost&#34;)
  -ts float
    	size of tile (default 30)
  -width int
    	width of the resulting image (default 800)

```
## Readme generator
The project also contains readme generator binary (in cmd/readmegenerator/main.go)
which takes path to image directory where it generates images, and writes this file's
contents to stdout, with image linked to this markdown file.

For details, see [the source](https://github.com/kamchy/banner/blob/main/src/readmegenerator/main.go)
### Usage

```bash
cd /cmd/readmegenerator && go build
cd ../..
./cmd/readmegenerator/readmegenerator img > README.md
```

## Images
And here are images:


### Image img/out_alg0_pal0.png
![random rectangles](img/out_alg0_pal0.png)

### Image img/out_alg0_pal1.png
![random rectangles](img/out_alg0_pal1.png)

### Image img/out_alg1_pal0.png
![random rectangles with offset](img/out_alg1_pal0.png)

### Image img/out_alg1_pal1.png
![random rectangles with offset](img/out_alg1_pal1.png)

### Image img/out_alg2_pal0.png
![plain color](img/out_alg2_pal0.png)

### Image img/out_alg2_pal1.png
![plain color](img/out_alg2_pal1.png)

### Image img/out_alg3_pal0.png
![concentric circles](img/out_alg3_pal0.png)

### Image img/out_alg3_pal1.png
![concentric circles](img/out_alg3_pal1.png)

### Image img/out_alg4_pal0.png
![concentric circles offset](img/out_alg4_pal0.png)

### Image img/out_alg4_pal1.png
![concentric circles offset](img/out_alg4_pal1.png)

### Image img/out_alg5_pal0.png
![random horizontal lines](img/out_alg5_pal0.png)

### Image img/out_alg5_pal1.png
![random horizontal lines](img/out_alg5_pal1.png)

### Image img/out_alg6_pal0.png
![random hexagons](img/out_alg6_pal0.png)

### Image img/out_alg6_pal1.png
![random hexagons](img/out_alg6_pal1.png)

### Image img/out_alg7_pal0.png
![random hexagons with offset](img/out_alg7_pal0.png)

### Image img/out_alg7_pal1.png
![random hexagons with offset](img/out_alg7_pal1.png)


