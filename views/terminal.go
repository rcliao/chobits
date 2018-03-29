package views

import termbox "github.com/nsf/termbox-go"

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

// Init the Terminal view
func (p *Terminal) Init() error {
	return termbox.Init()
}

// Close the terminal view
func (p *Terminal) Close() {
	termbox.Close()
}

// DrawText is used to draw a text at the bottom half of screen like hints or prompt
func (p *Terminal) DrawText(text string) {

}

// DrawMain draws the main part of screen at the center of terminal
func (p *Terminal) DrawMain(main []string) {

}

// DrawFooter draws the footer at the bottom of terminal screen
func (p *Terminal) DrawFooter(text string) {

}
