package sutil

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

func HttpDo(method string, uri string, param string, header http.Header) string {
	client := &http.Client{}
	var data io.Reader
	if header.Get("Content-Type") == "application/json" {
		json := []byte(param)
		data = bytes.NewBuffer(json)
	} else {
		data = strings.NewReader(param)
	}

	req, err := http.NewRequest(method, uri, data)
	if err != nil {
		// handle error
		// beego.Error(err)
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
		// beego.Error(err)
		return ""
	}

	return string(body)
}
