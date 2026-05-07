package main

// Station represents a lo-fi radio station.
type Station struct {
	Name string
	URL  string
}

var stations = []Station{
	{"Lofi 24/7", "https://usa9.fastcast4u.com/proxy/jamz?mp=/1"},
	{"Nightride FM - Chillsynth", "https://stream.nightride.fm/chillsynth.mp3"},
	{"Chillsky - Chillhop", "http://chill.radioca.st/stream"},
}
