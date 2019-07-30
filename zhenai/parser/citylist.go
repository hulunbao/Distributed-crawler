package parser

import (
	"Distributed-crawler/engine"
	"regexp"
)

var cityListRe = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)

// ParseCityList 爬取城市列表
func ParseCityList(contents []byte) engine.ParserResult {
	matches := cityListRe.FindAllSubmatch(contents, -1)

	result := engine.ParserResult{}
	limit := 2
	for _, m := range matches {
		result.Items = append(result.Items, "City "+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity,
		})
		limit--
		if limit == 0 {
			break
		}
	}
	return result
}
