package main

import (
	"github.com/hulunbao/Distributed-crawler/engine"
	"github.com/hulunbao/Distributed-crawler/zhenai/parser"
)

func main() {
	engine.SimpleEngine{}.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
