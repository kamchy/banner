Generates a banner of given sizes with main and support text as .png. 

Usage:
```bash

Usage of palette:
  -height int
    	height of the resulting image (default 600)
  -subtext string
    	explanatory text to display in the image below the text (default "this time about really important things")
  -text string
    	text to display in the image (default "My blogpost")
  -width int
    	width of the resulting image (default 800)
```

Result is always saved to out.png in current working directory.

## Example:

### Default 800x600 image:
`./palette`:

![default](/img/default.png)

### Banner 400x300 with text and subtext
`./palette -width 400 -height 300 -text "Go programming" -subtext "A way to become enlightened"`

![small quote](/img/small.png)
