package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"os"
	"strings"
	"time"
)

type AomineScraper struct {
	col *colly.Collector

	IDs   []string
	Works []ScrapedWork
}

type WorkStats struct {
	Published string `json:"published,omitempty"`
	Words     int    `json:"words,omitempty"`
	Chapters  string `json:"chapters,omitempty"`
	Comments  int    `json:"comments,omitempty"`
	Kudos     int    `json:"kudos,omitempty"`
	Bookmarks int    `json:"bookmarks,omitempty"`
	Hits      int    `json:"hits,omitempty"`
}

type ScrapedWork struct {
	ID            string    `json:"id" json:"id,omitempty"`
	Title         string    `json:"title,omitempty"`
	Author        string    `json:"author,omitempty"`
	Ratings       []string  `json:"ratings,omitempty"`
	Warnings      []string  `json:"warnings,omitempty"`
	Categories    []string  `json:"categories,omitempty"`
	Fandoms       []string  `json:"fandoms,omitempty"`
	Characters    []string  `json:"characters,omitempty"`
	Relationships []string  `json:"relationships,omitempty"`
	FreeformTags  []string  `json:"freeform_tags,omitempty"`
	Language      string    `json:"language,omitempty"`
	Summary       string    `json:"summary,omitempty"`
	EndNotes      string    `json:"end_notes,omitempty"`
	StartNotes    string    `json:"start_notes,omitempty"`
	Chapters      string    `json:"chapters,omitempty"`
	Stats         WorkStats `json:"stats"`
}

func NewScraper(filepath string) (s *AomineScraper, err error) {
	f, err := os.OpenFile(filepath, os.O_RDONLY, 0755)

	if err != nil {
		return
	}

	defer f.Close()

	s = &AomineScraper{
		col: BaseColly(),
	}
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		s.IDs = append(s.IDs, scanner.Text())
	}

	return
}

func bindTags(dom *goquery.Selection, tagName string, arr *[]string) {
	dom.Find(fmt.Sprintf(".%s.tags > ul > li > a", tagName)).Each(func(i int, selection *goquery.Selection) {
		*arr = append(*arr, selection.Text())
	})
}

func (s *AomineScraper) Start() (err error) {
	var work ScrapedWork

	s.col.OnHTML("#main", func(element *colly.HTMLElement) {
		bindTags(element.DOM, "character", &work.Characters)
		bindTags(element.DOM, "freeform", &work.FreeformTags)
		bindTags(element.DOM, "relationship", &work.Relationships)
		bindTags(element.DOM, "fandom", &work.Fandoms)
		bindTags(element.DOM, "warning", &work.Warnings)
		bindTags(element.DOM, "rating", &work.Ratings)
		bindTags(element.DOM, "category", &work.Categories)

		work.Title = strings.TrimSpace(element.DOM.Find("h2.title.heading").First().Text())
		work.Author = strings.TrimSpace(element.DOM.Find("a[rel=\"author\"]").First().Text())

		chap, _ := element.DOM.Find("#chapters").First().Html()
		work.Chapters = chap

		work.Summary = element.DOM.Find(".summary.module > .userstuff").First().Text()
		work.StartNotes = element.DOM.Find(".notes.module > .jump").First().Text()
		work.EndNotes = element.DOM.Find("#work_endnotes > .userstuff").First().Text()

		s.Works = append(s.Works, work)
	})

	for k, v := range s.IDs {
		work = ScrapedWork{
			ID: v,
		}
		err = s.col.Visit(fmt.Sprintf("https://archiveofourown.org/works/%s", v))
		if err != nil {
			return
		}

		fmt.Printf("%d/%d | Scraped %s - %s\n", k+1, len(s.IDs), v, work.Title)

		time.Sleep(time.Second * 3)
	}

	fmt.Println("Completed!")

	return
}

func (s *AomineScraper) SaveJSON(filepath string) (err error) {
	f, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, 0755)

	if err != nil {
		return
	}

	err = json.NewEncoder(f).Encode(s.Works)

	if err != nil {
		return
	}

	return f.Close()
}
