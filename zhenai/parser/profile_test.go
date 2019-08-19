package parser

import (
	"io/ioutil"
	"testing"

	"github.com/hulunbao/Distributed-crawler/model"
)

func TestParseProfile(t *testing.T) {
	//contents, err := fetcher.Fetch("http://album.zhenai.com/u/1534498366")
	contents, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}
	result := ParseProfile(contents, "张一")
	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 element; but was %v", result.Items)
	}
	profile := result.Items[0].(model.Profile)
	expected := model.Profile{
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
	}
	if profile != expected {
		t.Errorf("expected %v, but was %v", expected, profile)
	}

}
