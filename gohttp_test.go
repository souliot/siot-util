package feutil

import (
	"net/url"

	"github.com/souliot/feutil"
)

func gohttp_test() string {
	uri := "http://openapi.ecois.info/v2/poi/detail?id=2"
	var param = url.Values{}
	// param.Add("regcode", "yuandong_meter")
	// param.Add("data", string(json_nodeRealTimeDatas))
	// param.Add("type", "3")
	// mime := "application/json; charset=utf-8"
	mime := "application/x-www-form-urlencoded"
	// beego.Info(string(json_nodeRealTimeDatas))
	header := map[string][]string{
		"X-Access-Token": {"81EGSR7ptRYyEsoCIZw5yC8hyRKR68oW"},
		"Content-Type":   {mime},
	}
	str := feutil.HttpDo("GET", uri, param.Encode(), header)
	return str
}
