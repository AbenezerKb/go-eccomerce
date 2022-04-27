package Get_All_Store_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gin-exercise/server"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Store Functionalities", func() {

	Context("Accepted Criteria", func() {

		It("when user is requesting for list of stores", func() {

			payload := map[string]int{"page": 1, "size": 3}
			values, _ := json.Marshal(payload)

			res, err := Request("GET", "/stores", bytes.NewBuffer(values))
			if err != nil {
				fmt.Println(err)
				return
			}
	//		defer res.Body.Close()

			body, err := ioutil.ReadAll(res.Body)
			status := res.Status
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println("the reponse from get all: "+string(body))

			Expect(status).To(BeEquivalentTo("200 OK"))

		})
	})

	Context("Failed Scenario", func() {
		

		It("when user is requesting for list of stores with grater than the existing pages and size", func() {

			payload := map[string]int{"page": 3, "size": 3}
			values, _ := json.Marshal(payload)

			res, err := Request("GET", "/stores", bytes.NewBuffer(values))
			if err != nil {
				fmt.Println(err)
				return
			}
			
		//	defer res.Body.Close()

			body, err := ioutil.ReadAll(res.Body)
			status := res.Status
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println("the response1: ",string(body))

			Expect(status).To(BeEquivalentTo("200 OK"))

		})
	})

})



func Request(method string, api string, buffer *bytes.Buffer) (*http.Response, error) {
	HT := httptest.NewServer(server.TestingServer())
	url := HT.URL

	client := &http.Client{}
	req, eror := http.NewRequest(method, url+api, buffer)

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
//	defer res.Body.Close()

	_, er := ioutil.ReadAll(res.Body)

	if er != nil {

		return nil, er
	}
	return res, nil
}
