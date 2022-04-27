package Get_All_Item_test

import (
	"bytes"
	"encoding/json"	
	"io/ioutil"

	// "time"	
	"gin-exercise/entity"
	"gin-exercise/server"

	// "io/ioutil"
	"net/http"
	"net/http/httptest"

	// "time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Get List of Items in Store", func() {

	Context("Accepted Criteria", func() {

		BeforeEach(func() {

			//adding item

			loginPayload := map[string]string{"name": "dell_computer"}
			values, _ := json.Marshal(loginPayload)

			res, err := Request("POST", "/items", bytes.NewBuffer(values))

			Expect(err).To(BeNil())
			//defer res.Body.Close()

			status := res.Status

			Expect(status).To(BeEquivalentTo("200 OK"))

			//adding another item

			loginPayload1 := map[string]string{"name": "hp_computer"}
			values1, _ := json.Marshal(loginPayload1)

			res1, err := Request("POST", "/items", bytes.NewBuffer(values1))
			Expect(err).To(BeNil())

			// defer res.Body.Close()

			status1 := res1.Status

			Expect(status1).To(BeEquivalentTo("200 OK"))

			//adding another item

			loginPayload2 := map[string]string{"name": "hp_computer"}
			values2, _ := json.Marshal(loginPayload2)

			res2, err := Request("POST", "/items", bytes.NewBuffer(values2))
			Expect(err).To(BeNil())

			// defer res.Body.Close()

			status2 := res2.Status

			Expect(status2).To(BeEquivalentTo("200 OK"))

			//adding another item

			loginPayload3 := map[string]string{"name": "apple_computer"}
			values3, _ := json.Marshal(loginPayload3)

			res3, err := Request("POST", "/items", bytes.NewBuffer(values3))
			Expect(err).To(BeNil())
			// defer res.Body.Close()

			status3 := res3.Status

			Expect(status3).To(BeEquivalentTo("200 OK"))

			//addin another item

			loginPayload4 := map[string]string{"name": "toshiba_computer"}
			values4, _ := json.Marshal(loginPayload4)

			res4, err := Request("POST", "/items", bytes.NewBuffer(values4))

			Expect(err).To(BeNil())

			status4 := res4.Status

			Expect(status4).To(BeEquivalentTo("200 OK"))
			// count, _ := db.Count(entity.Item{})
			// fmt.Println("the item count: ", count)

		})

		It("get all items with specifc store id", func() {

			payload := map[string]int{"page": 1, "size": 6}
			values, _ := json.Marshal(payload)

			HT := httptest.NewServer(server.TestingServer())
	url := HT.URL

	client := &http.Client{}

	req, err := http.NewRequest("GET", url+"/items", bytes.NewBuffer(values))
	Expect(err).To(BeNil())

	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)		

	Expect(err).To(BeNil())
		//	defer res.Body.Close()

			body, err := ioutil.ReadAll(res.Body)
			Expect(err).To(BeNil())
			status := res.Status			
			Expect(err).To(BeNil())
			var savedItems []entity.Item
			json.Unmarshal(body, &savedItems)			

			Expect(status).To(BeEquivalentTo("200 OK"))
			res.Body.Close()

		})

	})

})

func Request(method string, api string, buffer *bytes.Buffer) (*http.Response, error) {
	HT := httptest.NewServer(server.TestingServer())
	url := HT.URL

	client := &http.Client{}

	req, err := http.NewRequest(method, url+api, buffer)
	Expect(err).To(BeNil())

	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token","eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6Im5ld3NhbXBsZUBnbWFpbC5jb20iLCJleHAiOjE2NTEwNjI4NTgsImlhdCI6MTY1MDg5MDA1OCwiaXNzIjoibmV3c2FtcGxlQGdtYWlsLmNvbSJ9.f0HslcfaX8kf4OVDAFC_3KkZNckP2JMYavXOqXl7FQI")
	res, err := client.Do(req)
	Expect(err).To(BeNil())
//	defer res.Body.Close()

	_, err = ioutil.ReadAll(res.Body)

	Expect(err).To(BeNil())
	return res, nil
}
