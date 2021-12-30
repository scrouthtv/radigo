package main

import (
	"os"
)

func main() {
	p := Player{S: Stations, G: NewGui(), R: NewBeep(), Active: -1}

	p.G.Open()
	p.R.Open()

	p.G.Loop(&p)
}

type Player struct {
	S []Station
	Active int
	G *Gui
	R Renderer
}

func (p *Player) Exit(code int) {
	p.R.Stop()
	p.G.Close()
	os.Exit(code)
}

type Renderer interface {
	Open()
	Play(*Station) error
	Stop()
	Close()
}

