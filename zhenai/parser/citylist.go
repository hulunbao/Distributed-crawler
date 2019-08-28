package parser

import (
	"regexp"

	"github.com/hulunbao/Distributed-crawler/engine"
)

var cityListRe = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)

//var genderRe = regexp.MustCompile(``)

// ParseCityList 爬取城市列表
func ParseCityList(contents []byte) engine.ParseResult {
	matches := cityListRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity,
		})
	}
	return result
}
