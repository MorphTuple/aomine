package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"net/url"
	"os"
	"strconv"
	"strings"
)

type AomineCollector struct {
	Url    string
	Offset int
	Limit  int
	Page   int
	Passed int
	Peaked bool

	col       *colly.Collector
	Collected []string
}

func NewCollector(searchUrl string, offset int, limit int) *AomineCollector {
	parsedUrl, err := url.Parse(searchUrl)

	if err != nil {
		return nil
	}

	q := parsedUrl.Query()
	pageQuery := q.Get("page")

	page := 1

	if pageInt, err := strconv.Atoi(pageQuery); err == nil {
		page = pageInt
	}

	return &AomineCollector{
		Url:    searchUrl,
		Offset: offset,
		Limit:  limit,
		Page:   page,
		col:    BaseColly(),
	}
}

func (c *AomineCollector) getUrl() string {
	parsedUrl, _ := url.Parse(c.Url)
	q := parsedUrl.Query()

	q.Set("page", fmt.Sprintf("%d", c.Page))

	parsedUrl.RawQuery = q.Encode()

	return parsedUrl.String()
}

func (c *AomineCollector) Start() {
	c.col.OnHTML("#main", func(el *colly.HTMLElement) {
		var returnedCount int

		el.ForEachWithBreak("[role=\"article\"]", func(i int, articleEl *colly.HTMLElement) bool {
			if len(c.Collected) >= c.Limit {
				c.Peaked = true
				return false
			}

			returnedCount++

			if c.Passed < c.Offset {
				fmt.Println("Skipping...")
				c.Passed++
				return true
			}

			s, _ := articleEl.DOM.Attr("id")
			c.Collected = append(c.Collected, s[5:])
			fmt.Printf("#%d %s\n", len(c.Collected), s[5:])

			return true
		})

		if returnedCount == 0 {
			fmt.Println("Reached end of page!")
			c.Peaked = true
		}
	})

	fmt.Printf("Fetching Page %d\n", c.Page)
	c.col.Visit(c.getUrl())

	for !c.Peaked && len(c.Collected) < c.Limit {
		c.Page++
		fmt.Printf("Fetching Page %d\n", c.Page)
		c.col.Visit(c.getUrl())
	}
}

func (c *AomineCollector) SaveCSV(filepath string) (err error) {
	f, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, 0755)

	if err != nil {
		return
	}

	_, err = f.WriteString(strings.Join(c.Collected, "\n"))

	if err != nil {
		return
	}

	return f.Close()
}
