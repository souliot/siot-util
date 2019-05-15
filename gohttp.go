package sutil

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func HttpDo(method string, uri string, header http.Header, data string, params string) string {
	client := &http.Client{}

	req_data := bytes.NewBuffer([]byte(data))

	req, err := http.NewRequest(method, uri, req_data)

	req.URL.RawQuery = params

	if err != nil {
		return ""
	}

	req.Header = header

	resp, err := client.Do(req)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}

	return string(body)
}
