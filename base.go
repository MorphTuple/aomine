package main

import (
	"github.com/gocolly/colly"
	"time"
)

const ao3domain = "archiveofourown.org"

func BaseColly() *colly.Collector {
	col := colly.NewCollector(colly.AllowedDomains(ao3domain))
	col.AllowURLRevisit = true

	col.OnRequest(func(request *colly.Request) {
		request.Headers.Set("cookie", "view_adult=true;")
	})

	col.Limit(&colly.LimitRule{
		Delay: time.Second * 5,
	})

	return col
}
