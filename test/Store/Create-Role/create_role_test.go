package Create_Role_test

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

var _ = Describe("Clerk Registration", func() {

	Context("Accepted Criteria", func() {

		Context("when user registers new clerks", func() {

			loginPayload := map[string]string{"name": "lomiad", "location": "Ayat,Addis Ababa", "image": "pit"}
			values, _ := json.Marshal(loginPayload)

			response, err := Request("POST", "/stores", bytes.NewBuffer(values))
			if err != nil {
				fmt.Println(err)
				return
			}
			defer response.Body.Close()
			



			It("has status of code", func() {
				Expect(response.Status).To(BeEquivalentTo("200 OK"))
			})
		})
	})

	Context("Rejected Criteria", func() {

		Context("when user registers new store empty name", func() {

			loginPayload := map[string]string{"name": "", "location": "Ayat,Addis Ababa", "Image": "piture.jpg"}
			values, _ := json.Marshal(loginPayload)

			response, err := Request("POST", "/stores", bytes.NewBuffer(values))
			if err != nil {
				fmt.Println(err)
				return
			}
			defer response.Body.Close()

			It("has status of code", func() {
				Expect(response.Status).To(BeEquivalentTo("400 Bad Request"))
			})
		})

		Context("when user registers new store with empty location", func() {

			loginPayload := map[string]string{"name": "safeway", "location": "", "Image": "picture.jpg"}
			values, _ := json.Marshal(loginPayload)

			response, err := Request("POST", "/stores", bytes.NewBuffer(values))
			if err != nil {
				fmt.Println(err)
				return
			}
			defer response.Body.Close()

			It("has status of code", func() {
				Expect(response.Status).To(BeEquivalentTo("400 Bad Request"))
			})
		})

		Context("when user registers new store with empty image", func() {

			It("has status of code", func() {
			loginPayload := map[string]string{"name": "safeway", "location": "Ayat, Addis Ababa", "Image": ""}
			values, _ := json.Marshal(loginPayload)

			response, err := Request("POST", "/stores", bytes.NewBuffer(values))
			Expect(err).To(BeNil())
			defer response.Body.Close()
		
			Expect(response.Status).To(BeEquivalentTo("200 OK"))
			})
		})
	})

})

func Request(method string, api string, buffer *bytes.Buffer) (*http.Response, error) {
	HT := httptest.NewServer(server.TestingServer())
	// http.Serve
	url := HT.URL

	client := &http.Client{}
	
	req, eror := http.NewRequest("POST", url+api, buffer)
	Expect(eror).To(BeNil())
	
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token","eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6ImJldGlAZW1haWwuY29tIiwiZXhwIjoxNjUwNDUxNTQzLCJpYXQiOjE2NTAyNzg3NDMsImlzcyI6ImJldGlAZW1haWwuY29tIn0.vBnwkUiX0ZBz8aiGiSxdDYBtYxrFpkVkmh_RsjVU1Cw")
	res, err := client.Do(req)
	Expect(err).To(BeNil())
	// defer res.Body.Close()

	_, er := ioutil.ReadAll(res.Body)

	Expect(er).To(BeNil())
	return res, nil
}
