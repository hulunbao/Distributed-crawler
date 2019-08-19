package engine

import (
	"log"

	"github.com/hulunbao/Distributed-crawler/fetcher"
)

type SimpleEngine struct{}

// Run 系统引擎部分
func (e SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		ParseResult, err := worker(r)
		if err != nil {
			continue
		}
		requests = append(requests, ParseResult.Requests...)

		for _, item := range ParseResult.Items {
			log.Printf("Got item %v", item)
		}
	}
}

func worker(r Request) (ParseResult, error) {
	log.Printf("Fetching %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error fetching url %s: %v", r.Url, err)
		return ParseResult{}, err
	}
	return r.ParserFunc(body), nil
}
