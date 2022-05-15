package main

import (
	"github.com/alecthomas/kong"
)

type AomineCLIArgs struct {
	Collect struct {
		Url      string `arg:"" help:"The AO3 Search URL"`
		Filepath string `arg:"" optional:"" help:"CSV output file" default:"output.csv"`
		Offset   int    `help:"Offset the search - Defaults to 0"`
		Limit    int    `help:"Limit the amount of work IDs collected - Defaults to 100" default:"100"`
	} `cmd:"" help:"Collects the IDs of works based of a search url, and outputs it to a CSV file"`
	Scrape struct {
		IDFilepath     string `arg:"" help:"CSV input filepath containing the IDs" default:"output.csv"`
		OutputFilepath string `arg:"" optional:"" help:"JSON output filepath" default:"scrape_output.json"`
	} `cmd:"" help:"Scrape works based of a CSV file containing the IDs of the work, and outputs it to a JSON file"`
}

func main() {
	var args AomineCLIArgs

	ctx := kong.Parse(&args)

	switch ctx.Command() {
	case "collect <url>", "collect <url> <filepath>":
		collector := NewCollector(args.Collect.Url, args.Collect.Offset, args.Collect.Limit)
		collector.Start()

		err := collector.SaveCSV(args.Collect.Filepath)

		if err != nil {
			panic(err)
		}
	case "scrape <id-filepath>", "scrape":
		scraper, err := NewScraper(args.Scrape.IDFilepath)

		if err != nil {
			panic(err)
		}

		if err = scraper.Start(); err != nil {
			panic(err)
		}

		if err = scraper.SaveJSON(args.Scrape.OutputFilepath); err != nil {
			panic(err)
		}
	default:
		panic(ctx.Command())
	}
}
