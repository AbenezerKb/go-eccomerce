package Get_Item_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"

	// "time"

	//	"gin-exercise/db"
	"gin-exercise/entity"
	"gin-exercise/server"

	// "io/ioutil"
	"net/http"
	"net/http/httptest"

	// "time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Get Item", func() {


	Context("Accepted Criteria", func() {

		BeforeEach(func() {

			loginPayload := map[string]string{"name": "dell_computer"}
			values, _ := json.Marshal(loginPayload)

			res, err := Request("POST", "/items", bytes.NewBuffer(values))
			if err != nil {
				fmt.Println(err)
				return
			}		



			status := res.Status

			Expect(status).To(BeEquivalentTo("200 OK"))

		})

		It("get all item by name", func() {

			payload := map[string]string{"name": "dell_computer"}
			values, _ := json.Marshal(payload)

			res, err := Request("GET", "/item", bytes.NewBuffer(values))
			if err != nil {
				fmt.Println(err)
				return
			}
			

			body, err := ioutil.ReadAll(res.Body)
			status := res.Status
			
			if err != nil {
				fmt.Println(err)
				return
			}
			var savedItems []entity.Item
			json.Unmarshal(body, &savedItems)
			

			Expect(status).To(BeEquivalentTo("200 OK"))

		})

	})




	Context("Failing Senario", func() {

		
		It("searching item with wrong name or that doesn;t exist", func() {

			payload := map[string]string{"name": "computer"}
			values, _ := json.Marshal(payload)

			res, err := Request("GET", "/item", bytes.NewBuffer(values))
			if err != nil {
				fmt.Println(err)
				return
			}
			

			body, err := ioutil.ReadAll(res.Body)
			status := res.Status
		
			if err != nil {
				fmt.Println(err)
				return
			}
			var savedItems []entity.Item
			json.Unmarshal(body, &savedItems)			

			Expect(status).To(BeEquivalentTo("204 No Content"))

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
	req.Header.Add("token","eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6Im5ld3NhbXBsZUBnbWFpbC5jb20iLCJleHAiOjE2NTEwNjI4NTgsImlhdCI6MTY1MDg5MDA1OCwiaXNzIjoibmV3c2FtcGxlQGdtYWlsLmNvbSJ9.f0HslcfaX8kf4OVDAFC_3KkZNckP2JMYavXOqXl7FQI")
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}


	_, er := ioutil.ReadAll(res.Body)

	if er != nil {

		return nil, er
	}
	return res, nil
}
