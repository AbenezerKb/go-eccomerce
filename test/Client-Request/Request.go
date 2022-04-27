package client_request

import (
	"bytes"
	"fmt"
	"gin-exercise/server"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

func Request(method string, api string, buffer *bytes.Buffer) (*http.Response, error) {
	HT := httptest.NewServer(server.TestingServer())
	url := HT.URL

	client := &http.Client{}
	req, eror := http.NewRequest("POST", url+api, buffer)

	if eror != nil {
		fmt.Println(eror)
		return nil, eror
	}

	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	_, er := ioutil.ReadAll(res.Body)
	// _ := res.Status
	if er != nil {
		//fmt.Println(err)
		return nil, er
	}
	return res, nil
}
