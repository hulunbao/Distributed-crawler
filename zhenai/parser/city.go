package parser

import (
	"Distributed-crawler/engine"
	"regexp"
)

var cityRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)".+?alt=([^>]+)></a>`)

func ParseCity(contents []byte) engine.ParserResult {
	matches := cityRe.FindAllSubmatch(contents,-1)

	result := engine.ParserResult{}
	for _, m := range matches {
		result.Items = append(result.Items,"User " + string(m[2]))
		result.Requests = append(result.Requests,engine.Request{
			Url: string(m[1]),
			ParserFunc: engine.NilParser,
		})
	}
	return result
}
