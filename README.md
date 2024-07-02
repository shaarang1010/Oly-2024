## 2024 Olympics Schedule scraper

### Why?

The Paris Olympics webiste has times listed in Paris time.

But I want to watch Track and Field, Swimming, Cycling and other sports in AEST time.

Plus, would like to know when are the main events rather than getting generic schedule and would like to get an overall view rather than filtering per sport.

In addition, great way to test [Colly](https://github.com/gocolly/colly) library.

### How?

Scrape Paris Schedule website and store details in a `csv` or `json` file.
Serve file contents via Chi-router package.
