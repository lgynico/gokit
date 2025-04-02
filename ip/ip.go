package ip

import (
	"strings"

	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
)

var searcher *xdb.Searcher

func init() {
	vi, err := xdb.LoadVectorIndexFromFile("./ip2region.xdb")
	if err != nil {
		panic(err)
	}

	searcher, err = xdb.NewWithVectorIndex("./ip2region.xdb", vi)
	if err != nil {
		panic(err)
	}
}

func GetRegion(ip string) (*Region, error) {
	value, err := searcher.SearchByStr(ip)
	if err != nil {
		return nil, err
	}

	var (
		country  string
		region   string
		province string
		city     string
		isp      string
		ss       = strings.Split(value, "|")
	)

	if ss[0] != "0" {
		country = ss[0]
	}
	if ss[1] != "0" {
		region = ss[1]
	}
	if ss[2] != "0" {
		province = ss[2]
	}
	if ss[3] != "0" {
		city = ss[3]
	}
	if ss[4] != "0" {
		isp = ss[4]
	}

	reg := &Region{
		Country:  country,
		Region:   region,
		Province: province,
		City:     city,
		ISP:      isp,
	}

	return reg, nil
}
