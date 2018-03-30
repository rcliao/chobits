package views

import (
	"bytes"
	"fmt"
	"strings"

	termbox "github.com/nsf/termbox-go"
)

var mainHeight = 5

var largeInt = map[int]string{
	1: `
   ##
   ##
   ##
   ##
   ##
`,
	2: `
#####
   ##
#####
##
#####
`,
	3: `
#####
   ##
#####
   ##
#####
`,
	4: `
## ##
## ##
#####
   ##
   ##
`,
	5: `
#####
##
#####
   ##
#####
`,
	6: `
#####
##
#####
## ##
#####
`,
	7: `
#####
   ##
   ##
   ##
   ##
`,
	8: `
#####
## ##
#####
## ##
#####
`,
	9: `
#####
## ##
#####
   ##
   ##
`,
}

func convertIntToArray(intStr string) []string {
	result := []string{}
	for _, s := range strings.Split(intStr, "\n") {
		if len(s) == 0 {
			continue
		}
		result = append(result, s)
	}
	return result
}

func concatArrayHorizontal(intArray1 []string, intArray2 []string) []string {
	result := []string{}
	for i, a1 := range intArray1 {
		var buffer bytes.Buffer
		buffer.WriteString(a1)
		buffer.WriteString(" ")
		buffer.WriteString(intArray2[i])
		result = append(result, buffer.String())
	}
	return result
}

// View encapsulates the logic to draw things on screen
type View interface {
	DrawText(text string)
	DrawMain(main []string)
	DrawFooter(text string)
}

/*
Terminal is specification for drawing in the terminal screen like below:

```
+-----------------+
|                 |
|  #  #     #  #  |
|  #  #  #  #  #  |
|  #  #     #  #  |
|  #  #  #  #  #  |
|  #  #     #  #  |
|                 |
|  Hello world!   |
|                 |
|  Footer         |
+-----------------+
```

Where clock at the center is the "Main" and text below is the "Text"
*/
type Terminal struct {
}

// NewTerminalView creates a new terminal view to draw
func NewTerminalView() (*Terminal, error) {
	result := &Terminal{}
	err := result.init()
	if err != nil {
		return result, err
	}
	return result, nil
}

// Init the Terminal view
func (p *Terminal) init() error {
	return termbox.Init()
}

// Close the terminal view
func (p *Terminal) Close() {
	termbox.Close()
}

// DrawText is used to draw a text at the bottom half of screen like hints or prompt
func (p *Terminal) DrawText(text string) {
	termbox.Clear(termbox.ColorBlack, termbox.ColorDefault)
	w, h := termbox.Size()
	// calculate where to draw the text
	y := (h / 2) + (mainHeight / 2) + 1
	x := (w / 2) - (len(text) / 2)
	fmt.Println(w, h, x, y, text)
	for _, c := range text {
		termbox.SetCell(x, y, c, termbox.ColorWhite, termbox.ColorDefault)
		x++
	}
	termbox.Flush()
}

// DrawMain draws the main part of screen at the center of terminal
func (p *Terminal) DrawMain(main []string) {

}

// DrawFooter draws the footer at the bottom of terminal screen
func (p *Terminal) DrawFooter(text string) {

}
