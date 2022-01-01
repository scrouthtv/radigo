package main

import (
	"github.com/nsf/termbox-go"
	"github.com/mattn/go-runewidth"
)

type Gui struct {
	pos int
	status string
}

func NewGui() *Gui {
	return &Gui{pos: 0}
}

func (g *Gui) Open() {
	termbox.Init()
}

func (g *Gui) Redraw(p *Player) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	for i, s := range p.S {
		col := termbox.ColorDefault
		if i == p.Active && p.R.Playing() {
			col |= termbox.AttrReverse
		}

		text := s.Name
		if i == g.pos {
			text = " -> " + text
		} else {
			text = "    " + text
		}

		tprint(2, i+1, col, col, text)
	}

	_, h := termbox.Size()
	tprint(1, h-2, termbox.AttrBold, termbox.ColorDefault, g.status)

	termbox.Flush()
}

func (g *Gui) Loop(p *Player) {
	for {
		g.Redraw(p)

		ev := termbox.PollEvent()

		switch ev.Type {
		case termbox.EventKey:
			if ev.Ch == 'q' {
				p.Exit(0)
			} else if ev.Ch == 'j' || ev.Key == termbox.KeyArrowDown {
				if g.pos < len(p.S) - 1 {
					g.pos++
				}
			} else if ev.Ch == 'k' || ev.Key == termbox.KeyArrowUp {
				if g.pos > 0 {
					g.pos--
				}
			} else if ev.Key == termbox.KeySpace || ev.Key == termbox.KeyEnter {
				if g.pos == p.Active && p.R.Playing() {
					p.R.Stop()
				} else {
					err := p.Play(g.pos)
					if err != nil {
						g.status = err.Error()
					} else {
						g.status = "ok."
					}
				}
			}
		}
	}
}

func (g *Gui) Close() {
	termbox.Close()
}

func tprint(x, y int, fg, bg termbox.Attribute, s string) {
	for _, c := range s {
		termbox.SetCell(x, y, c, fg, bg)
		x += runewidth.RuneWidth(c)
	}
}
