package parser

import (
	"regexp"
	"strconv"

	"github.com/hulunbao/Distributed-crawler/engine"

	"github.com/hulunbao/Distributed-crawler/model"
)

var ageRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>([0-9]+)[^<]+</div>`)
var heightRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>([^<]+)cm</div>`)
var weightRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>([^<]+)kg</div>`)
var incomeRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>月收入:([^<]+)</div>`)
var genderRe = regexp.MustCompile(`<td><span class="label">性别：</span><span field>([^<]+)</span></td>`)
var xinzuoRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>(白羊座\(03.21-04.19\)|金牛座\(04.20-05.20\)|双子座\(05.21-06.21\)|巨蟹座\(06.22-07.22\)|狮子座\(07.23-08.22\)|处女座\(08.23-09.22\)|天秤座\(09.23-10.23\)|天蝎座\(10.24-11.22\)|射手座\(11.22-12.21\)|魔羯座\(12.22-01.19\)|水瓶座\(01.20-02.18\)|双鱼座\(02.19-03.20\))</div>`)
var marriageRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>(未婚|离异|丧偶)</div>`)
var educationRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>(高中及以下|中专|大专|大学本科|硕士|博士)</div>`)
var occupationRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>(外贸专员|计算机/互联网|公务员)</div>`)
var hokouRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>工作地:([^<]+)</div>`)
var houseRe = regexp.MustCompile(`<div class="m-btn pink" data-v-bff6f798>(已购房|未购房)</div>`)
var carRe = regexp.MustCompile(`<div class="m-btn pink" data-v-bff6f798>(已买车|未买车)</div>`)

// ParseProfile 爬取个人信息
func ParseProfile(contents []byte, name string) engine.ParseResult {
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

	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	return result

}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	}
	return ""
}
