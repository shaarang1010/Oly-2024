// package main

// import (
// 	"github.com/gocolly/colly"
// )

// type Schedule struct {
// 	Date   string
// 	Sport  *sport
// 	Events []*Event
// }

// type Event struct {
// 	Time        string
// 	Description string
// 	Location    string
// }
// type sport struct {
// 	Name string
// 	Link string
// }

//	func main() {
//		c := colly.NewCollector(
//			colly.AllowedDomains("en.wikipedia.org"),
//		)
//	}
package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly/v2"
)

type athlete struct {
	Name string
}

type sport struct {
	Name     string `selector:"div.zephr-feature_article-content-body > h3"`
	Athletes []*athlete
}

func main() {

	s := &sport{}
	athletes := make([]*athlete, 0)
	s.Athletes = athletes
	// Create a new Colly collector
	c := colly.NewCollector()

	// Define the URL you want to scrape
	url := "https://www.sportingnews.com/au/olympics/news/australian-olympics-squad-athletes-2024-paris/67ba52fe0d6d0d815ef8e8ec"

	// Set up callbacks to handle scraping events
	c.OnHTML(".zephr-feature_article-content-body", func(e *colly.HTMLElement) {
		// Extract data from HTML elements
		quote := e.ChildText("h3")
		fmt.Println(quote)
		// author := e.ChildText("small.author")
		// tags := e.ChildText("div.tags")

		// // Clean up the extracted data
		// quote = strings.TrimSpace(quote)
		// author = strings.TrimSpace(author)
		// tags = strings.TrimSpace(tags)

		// // Print the scraped data
		// fmt.Printf("Quote: %s\nAuthor: %s\nTags: %s\n\n", quote, author, tags)

		// e.Unmarshal(s)
		// s.Name = e.Request.URL.Query().Get("sport")
		// fmt.Println(s.Name)
	})

	// Visit the URL and start scraping
	err := c.Visit(url)
	if err != nil {
		log.Fatal(err)
	}
}
