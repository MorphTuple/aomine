package main

import (
	"github.com/alecthomas/kong"
)

type AomineCLIArgs struct {
	Collect struct {
		Url      string `arg:"" help:"The AO3 Search URL"`
		Filepath string `arg:"" optional:"" help:"CSV Output File" default:"output.csv"`
		Offset   int    `help:"Offset the search - Defaults to 0"`
		Limit    int    `help:"Limit the amount of work IDs collected - Defaults to 100" default:"100"`
	} `cmd:"" help:"Collects the IDs of works based of a search url"`
}

func main() {
	var args AomineCLIArgs

	ctx := kong.Parse(&args)

	switch ctx.Command() {
	case "collect <url>":
		collector := NewCollector(args.Collect.Url, args.Collect.Offset, args.Collect.Limit)
		collector.Start()

		err := collector.SaveCSV(args.Collect.Filepath)

		if err != nil {
			panic(err)
		}
	default:
		panic(ctx.Command())
	}
}
