package main

import (
	"github.com/gocolly/colly"
)

type Schedule struct {
	Date   string
	Sport  *sport
	Events []*Event
}

type Event struct {
	Time        string
	Description string
	Location    string
}
type sport struct {
	Name string
	Link string
}

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("en.wikipedia.org"),
	)
}
