package main

import (
	"Distributed-crawler/engine"
	"Distributed-crawler/zhenai/parser"
)

func main() {
	engine.SimpleEngine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
