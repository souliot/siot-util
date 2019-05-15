package sutil

import (
	"net/url"
)

func gohttp_test() string {
	uri := "http://example.com/path"

	// GET
	var param = url.Values{}
	param.Add("page", "1")
	param.Add("pageSize", "2")
	data := param.Encode()

	// POST
	// m := map[string]interface{}{
	// 	"page":     1,
	// 	"pageSize": 2,
	// }
	// mjson, _ := json.Marshal(m)
	// data := string(mjson)

	mime := "application/json"

	header := map[string][]string{
		"X-Access-Token": {"UlRZ2SSToJ0fcaoeaJSA8g=="},
		"Content-Type":   {mime},
	}
	str := HttpDo("GET", uri, data, header)
	return str
}
