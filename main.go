package main

import (
	"github.com/hulunbao/Distributed-crawler/engine"
	"github.com/hulunbao/Distributed-crawler/persist"
	"github.com/hulunbao/Distributed-crawler/scheduler"
	"github.com/hulunbao/Distributed-crawler/zhenai/parser"
)

func main() {
	itemChan, err := persist.ItemSaver("dating_profile")
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    itemChan,
	}
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
