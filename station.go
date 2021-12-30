package main

type Station struct {
	Name string
	Url string
}

var (
	Stations = []Station{
		Station{Name: "Dasding", Url: "https://liveradio.swr.de/sw282p3/dasding/"},
		Station{Name: "SWR3", Url: "https://liveradio.swr.de/sw282p3/swr3/"},
	}
)
