package feutil

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/astaxie/beego"
)

func HttpDo(method string, uri string, param string, header http.Header) string {
	client := &http.Client{}

	req, err := http.NewRequest(method, uri, strings.NewReader(param))
	if err != nil {
		// handle error
		beego.Error(err)
		return ""
	}

	// req.Header.Set("X-Access-Token", "81EGSR7ptRYyEsoCIZw5yC8hyRKR68oW")
	// req.Header.Set("Content-Type", mime)
	// req.Header.Set("Cookie", cookie)
	req.Header = header

	resp, err := client.Do(req)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		beego.Error(err)
		return ""
	}

	return string(body)
}
