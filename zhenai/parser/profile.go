package parser

import (
	"Distributed-crawler/engine"
	"Distributed-crawler/model"
	"fmt"
	"regexp"
	"strconv"
)

var ageRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798="">([\d]+岁)</div>`)
var heightRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798="">([^<]+)cm</div>`)
var weightRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798="">([^<]+)kg</div>`)
var incomeRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798="">月收入:([^<]+)</div>`)
var genderRe = regexp.MustCompile(`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)
var xinzuoRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798="">(.+座[^<]+)</div>`)
var marriageRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798="">未婚</div>`)
var educationRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798="">大专</div>`)
var occupationRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798="">外贸专员</div>`)
var hokouRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798="">工作地:([^<]+)</div>`)
var houseRe = regexp.MustCompile(`<div class="m-btn pink" data-v-bff6f798="">(.+房)</div>`)
var carRe = regexp.MustCompile(`<div class="m-btn pink" data-v-bff6f798="">(.+车)</div>`)

// ParseProfile 爬取个人信息
func ParseProfile(contents []byte, name string) engine.ParserResult {
	profile := model.Profile{}

	profile.Name = name
	age, err := strconv.Atoi(extractString(contents, ageRe))
	if err == nil {
		profile.Age = age
	}

	profile.Marriage = extractString(contents, marriageRe)

	height, err := strconv.Atoi(extractString(contents, heightRe))
	if err == nil {
		profile.Height = height
	}

	weight, err := strconv.Atoi(extractString(contents, weightRe))
	if err == nil {
		profile.Weight = weight
	}

	profile.Income = extractString(contents, incomeRe)
	profile.Gender = extractString(contents, genderRe)
	profile.Xinzuo = extractString(contents, xinzuoRe)
	profile.Marriage = extractString(contents, marriageRe)
	profile.Education = extractString(contents, educationRe)
	profile.Occupation = extractString(contents, occupationRe)
	profile.Hokou = extractString(contents, hokouRe)
	profile.House = extractString(contents, houseRe)
	profile.Car = extractString(contents, carRe)

	result := engine.ParserResult{
		Items: []interface{}{profile},
	}
	return result

}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	fmt.Println(match)
	fmt.Println(len(match))
	if len(match) >= 2 {
		return string(match[1])
	}
	return ""
}
