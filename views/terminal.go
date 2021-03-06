package views

import (
	"bytes"
	"strconv"
	"strings"

	termbox "github.com/nsf/termbox-go"
)

// ConvertClockToMain converts "12:34" to array of string for ease of drawing
func ConvertClockToMain(clockStr string) []string {
	result := []string{"", "", "", "", ""}
	for _, c := range clockStr {
		if string(c) == ":" {
			result = concatArrayHorizontal(result, convertIntToArray(colon))
			continue
		}
		i, err := strconv.Atoi(string(c))
		if err != nil {
			i = 0
		}
		result = concatArrayHorizontal(result, convertIntToArray(largeInt[i]))
	}
	return result
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

/*
ViewData group the necessary information needs to be on screen together so
one can just call `draw(data)` with all the necessary data
*/
type ViewData struct {
	Text   string
	Main   []string
	Footer string
}

// View encapsulates the logic to draw things on screen
type View interface {
	Draw(data ViewData)
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

// Clear clears the whole screen before drawing
func (p *Terminal) clear() {
	termbox.Clear(termbox.ColorBlack, termbox.ColorDefault)
}

func (p *Terminal) flush() {
	termbox.Flush()
}

// DrawText is used to draw a text at the bottom half of screen like hints or prompt
func (p *Terminal) drawText(text string) {
	w, h := termbox.Size()
	// calculate where to draw the text
	y := (h / 2) + (mainHeight / 2) + 1
	x := (w / 2) - (len(text) / 2)
	for _, c := range text {
		termbox.SetCell(x, y, c, termbox.ColorWhite, termbox.ColorDefault)
		x++
	}
}

// DrawMain draws the main part of screen at the center of terminal
func (p *Terminal) drawMain(main []string) {
	w, h := termbox.Size()
	// calculate where to draw the text
	y := (h / 2) - ((mainHeight - 2) / 2)
	for _, line := range main {
		x := (w / 2) - (len(line) / 2)
		for _, c := range line {
			if string(c) == "#" {
				termbox.SetCell(x, y, ' ', termbox.ColorWhite, termbox.ColorBlue)
			}
			x++
		}
		y++
	}
}

// DrawFooter draws the footer at the bottom of terminal screen
func (p *Terminal) drawFooter(text string) {
	w, h := termbox.Size()
	// calculate where to draw the text
	y := h - 1
	x := (w / 2) - (len(text) / 2)
	for _, c := range text {
		termbox.SetCell(x, y, c, termbox.ColorWhite, termbox.ColorDefault)
		x++
	}
}

// Draw draws main in the center, text right below main and footer at the bottom
// of the screen
func (p *Terminal) Draw(data ViewData) {
	p.clear()
	p.drawText(data.Text)
	p.drawMain(data.Main)
	p.drawFooter(data.Footer)
	p.flush()
}
