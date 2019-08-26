package persist

import (
	"context"
	"encoding/json"
	"testing"

	"gopkg.in/olivere/elastic.v5"

	"github.com/hulunbao/Distributed-crawler/engine"
	"github.com/hulunbao/Distributed-crawler/model"
)

func TestSave(t *testing.T) {
	expected := engine.Item{
		Url:  "https://album.zhenai.com/u/1534498366",
		Type: "zhenai",
		Id:   "1534498366",
		Payload: model.Profile{
			Name:       "张一",
			Age:        28,
			Height:     172,
			Weight:     61,
			Income:     "8千-1.2万",
			Gender:     "",
			Xinzuo:     "魔羯座(12.22-01.19)",
			Marriage:   "未婚",
			Education:  "大专",
			Occupation: "计算机/互联网",
			Hokou:      "阿坝汶川",
			House:      "已购房",
			Car:        "未买车",
		},
	}

	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	const index = "dating_test"
	err = save(client, expected, index)

	if err != nil {
		panic(err)
	}

	client, err = elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	resp, err := client.Get().Index(index).Type(expected.Type).Id(expected.Id).Do(context.Background())
	if err != nil {
		panic(err)
	}

	t.Logf("%s", resp.Source)

	var actual engine.Item
	json.Unmarshal(*resp.Source, &actual)

	actualProfile, _ := model.FromJsonObj(actual.Payload)
	actual.Payload = actualProfile

	if actual != expected {
		t.Errorf("got %v;expected %v", actual, expected)
	}
}
