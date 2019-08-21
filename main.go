package main

import (
	"github.com/hulunbao/Distributed-crawler/engine"
	"github.com/hulunbao/Distributed-crawler/persist"
	"github.com/hulunbao/Distributed-crawler/scheduler"
	"github.com/hulunbao/Distributed-crawler/zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    persist.ItemSaver(),
	}
	// e.Run(engine.Request{
	// 	Url:        "http://www.zhenai.com/zhenghun",
	// 	ParserFunc: parser.ParseCityList,
	// })
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun/shanghai",
		ParserFunc: parser.ParseCity,
	})
}
