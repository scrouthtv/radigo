package main

import (
	"net/http"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/mp3"
)

const (
	sr = 44100
)

type PlayerBeep struct {
	stream beep.Streamer
	play bool
}

func NewBeep() *PlayerBeep {
	return &PlayerBeep{stream: nil, play: false}
}

func (b *PlayerBeep) Play(s *Station) error {
	web, err := http.Get(s.Url)
	if err != nil {
		return err
	}

	streamer, format, err := mp3.Decode(web.Body)
	if err != nil {
		return err
	}

	resampled := beep.Resample(4, format.SampleRate, sr, streamer)

	b.stream = resampled
	b.play = true

	return nil
}

func (b *PlayerBeep) Open() {
	sr := beep.SampleRate(sr)
	speaker.Init(sr, sr.N(time.Second/10))
	speaker.Play(b)
}

func (b *PlayerBeep) Stop() {
	b.play = false
}

func (b *PlayerBeep) Close() {
	// speaker.Close() // not needed?
}

func (b *PlayerBeep) Playing() bool {
	return b.play
}

func (b *PlayerBeep) Err() error {
	return nil
}

func (b *PlayerBeep) Stream(samples [][2]float64) (n int, ok bool) {
	filled := 0

	for filled < len(samples) {
		// There are no streamers in the queue, so we stream silence.
		if !b.play || b.stream == nil {
			for i := range samples[filled:] {
				samples[i][0] = 0
				samples[i][1] = 0
			}
			break
		}

		n, ok := b.stream.Stream(samples[filled:])
		if !ok {
			b.stream = nil
		}

		filled += n
	}

	return len(samples), true
}

