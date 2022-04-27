package Create_Item_test

import (
	"bytes"
	"encoding/json"
	"fmt"

	// "fmt"
	"io/ioutil"

	// "time"

	//	"gin-exercise/db"
	// "gin-exercise/entity"
	"gin-exercise/entity"
	"gin-exercise/server"

	// "io/ioutil"
	"net/http"
	"net/http/httptest"

	// "time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Create Item", func() {

	
	
	It("Accepted Criteria", func() {
	
			loginPayload := map[string]string{"name": "dell_computer","image":"item_pic.jpg" }
			values, _ := json.Marshal(loginPayload)

			HT := httptest.NewServer(server.TestingServer())
			url := HT.URL
		
			client := &http.Client{}
			
			req, err := http.NewRequest("POST", url+"/items", bytes.NewBuffer(values))						
			Expect(err).To(BeNil())
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Add("Content-Type", "application/json")
			req.Header.Add("token","eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6Im5ld3NhbXBsZUBnbWFpbC5jb20iLCJleHAiOjE2NTEwNjI4NTgsImlhdCI6MTY1MDg5MDA1OCwiaXNzIjoibmV3c2FtcGxlQGdtYWlsLmNvbSJ9.f0HslcfaX8kf4OVDAFC_3KkZNckP2JMYavXOqXl7FQI")
			res, err := client.Do(req)

			Expect(err).To(BeNil())
			
			Imgbody, err := ioutil.ReadAll(res.Body)
			Expect(err).To(BeNil())
			
			var theitem entity.Item
			json.Unmarshal(Imgbody,&theitem)		
			status := res.Status
			fmt.Println("the status code: ", status)
			Expect(status).To(BeEquivalentTo("200 OK"))
			res.Body.Close()
	
	})

	It("Failing Senario, creating item with empty name", func() {
	
		loginPayload := map[string]string{"name": ""}
		values, _ := json.Marshal(loginPayload)


		HT := httptest.NewServer(server.TestingServer())
		url := HT.URL
	
		client := &http.Client{}
		
		req, err := http.NewRequest("POST", url+"/items", bytes.NewBuffer(values))
	
		Expect(err).To(BeNil())
	
		req.Header.Set("X-Custom-Header", "myvalue")
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("token","eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6Im5ld3NhbXBsZUBnbWFpbC5jb20iLCJleHAiOjE2NTEwNTg3MjYsImlhdCI6MTY1MDg4NTkyNiwiaXNzIjoibmV3c2FtcGxlQGdtYWlsLmNvbSJ9.-aXk_ZVB131w6r8QlWCAvGxPzccqSuk9jT3XKij4gho")
		res, err := client.Do(req)
		Expect(err).To(BeNil())	
	
		_, err = ioutil.ReadAll(res.Body)
	
		Expect(err).To(BeNil())

		status := res.Status
	
		Expect(status).To(BeEquivalentTo("400 Bad Request"))
		res.Body.Close()

})



})



func Request(method string, api string, buffer *bytes.Buffer) (*http.Response, error) {
	HT := httptest.NewServer(server.TestingServer())
	url := HT.URL

	client := &http.Client{}
	
	req, err := http.NewRequest("POST", url+api, buffer)

	Expect(err).To(BeNil())

	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	
	Expect(err).To(BeNil())

	_, err = ioutil.ReadAll(res.Body)

	Expect(err).To(BeNil())
	return res, nil
}
