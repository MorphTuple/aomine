# aomine

archiveofourown.org scraper in Go

## How to use

First, get the search URL of the works that you want to scrape, for example :

```
https://archiveofourown.org/works/search?work_search%5Bquery%5D=&work_search%5Btitle%5D=&work_search%5Bcreators%5D=&work_search%5Brevised_at%5D=&work_search%5Bcomplete%5D=&work_search%5Bcrossover%5D=&work_search%5Bsingle_chapter%5D=0&work_search%5Bword_count%5D=&work_search%5Blanguage_id%5D=&work_search%5Bfandom_names%5D=&work_search%5Brating_ids%5D=&work_search%5Bcharacter_names%5D=Mitsurugi+Reiji+%7C+Miles+Edgeworth&work_search%5Brelationship_names%5D=&work_search%5Bfreeform_names%5D=&work_search%5Bhits%5D=&work_search%5Bkudos_count%5D=&work_search%5Bcomments_count%5D=&work_search%5Bbookmarks_count%5D=&work_search%5Bsort_column%5D=_score&work_search%5Bsort_direction%5D=desc&commit=Search
```

After that, get the IDs of the work by doing so

```
aomine collect "https://archiveofourown.org/works/search?work_search%5Bquery%5D=&work_search%5Btitle%5D=&work_search%5Bcreators%5D=&work_search%5Brevised_at%5D=&work_search%5Bcomplete%5D=&work_search%5Bcrossover%5D=&work_search%5Bsingle_chapter%5D=0&work_search%5Bword_count%5D=&work_search%5Blanguage_id%5D=&work_search%5Bfandom_names%5D=&work_search%5Brating_ids%5D=&work_search%5Bcharacter_names%5D=Mitsurugi+Reiji+%7C+Miles+Edgeworth&work_search%5Brelationship_names%5D=&work_search%5Bfreeform_names%5D=&work_search%5Bhits%5D=&work_search%5Bkudos_count%5D=&work_search%5Bcomments_count%5D=&work_search%5Bbookmarks_count%5D=&work_search%5Bsort_column%5D=_score&work_search%5Bsort_direction%5D=desc&commit=Search" edgeworth.csv
```

A file named `edgeworth.csv` would be saved in the directory, it contains the work IDs that is collected.

And then you can scrape the works by so

```
aomine scrape edgeworth.csv edgeworth-works.json
```

## Help

```shell
Usage: aomine.exe <command>

Flags:
  -h, --help    Show context-sensitive help.

Commands:
  collect <url> [<filepath>]
    Collects the IDs of works based of a search url, and outputs it to a CSV
    file

  scrape [<id-filepath> [<output-filepath>]]
    Scrape works based of a CSV file containing the IDs of the work, and outputs
    it to a JSON file

Run "aomine.exe <command> --help" for more information on a command.
```

## License

GPL-3.0