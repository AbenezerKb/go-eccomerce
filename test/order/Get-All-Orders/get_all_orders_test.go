package Get_All_Orders_test

import (
	"bytes"
	"encoding/json"
	"fmt"

	// "gin-exercise/db"
	// "gin-exercise/entity"
	"gin-exercise/entity"
	"gin-exercise/server"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Store Functionalities", func() {
	var _ = Describe("Get Store", func() {

		BeforeEach(func() {


			//Add order
			order := entity.Order{OrderOwner: "Abenezer", Items_quantity: 20, Total_price: 1300}
			values, _ := json.Marshal(order)

			HT := httptest.NewServer(server.TestingServer())
			url := HT.URL

			client := &http.Client{}

			req, eror := http.NewRequest("POST", url+"/orders", bytes.NewBuffer(values))

			if eror != nil {
				fmt.Println(eror)
				return
			}

			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Add("Content-Type", "application/json")
			req.Header.Add("token","eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6Im5ld3NhbXBsZUBnbWFpbC5jb20iLCJleHAiOjE2NTEwNjQ3NDUsImlhdCI6MTY1MDg5MTk0NSwiaXNzIjoibmV3c2FtcGxlQGdtYWlsLmNvbSJ9.PaQ9GBixzNTAYVmQw-k0BVzp1_HX3yhGb0Rf1rj53t4")
			res, err := client.Do(req)
			if err != nil {
				fmt.Println(err)
				return
			}

			status := res.Status
			if err != nil {
				return
			}

			Expect(status).To(BeEquivalentTo("200 OK"))
			







			//Add another order
			order1 := entity.Order{OrderOwner: "Abenezer", Items_quantity: 20, Total_price: 1300}
			values1, _ := json.Marshal(order1)

			HT1 := httptest.NewServer(server.TestingServer())
			url1 := HT1.URL

			client1 := &http.Client{}

			req1, err := http.NewRequest("POST", url1+"/orders", bytes.NewBuffer(values1))

			Expect(err).To(BeNil())

			req1.Header.Set("X-Custom-Header", "myvalue")
			req1.Header.Add("Content-Type", "application/json")
			req.Header.Add("token","eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6Im5ld3NhbXBsZUBnbWFpbC5jb20iLCJleHAiOjE2NTEwNjQ3NDUsImlhdCI6MTY1MDg5MTk0NSwiaXNzIjoibmV3c2FtcGxlQGdtYWlsLmNvbSJ9.PaQ9GBixzNTAYVmQw-k0BVzp1_HX3yhGb0Rf1rj53t4")
			res1, err := client1.Do(req1)
			Expect(err).To(BeNil())
			status1 := res1.Status
		

			Expect(status1).To(BeEquivalentTo("200 OK"))
			




			//Add another order
			order2 := entity.Order{OrderOwner: "Abenezer", Items_quantity: 20, Total_price: 1300}
			values2, _ := json.Marshal(order2)

			HT2 := httptest.NewServer(server.TestingServer())
			url2 := HT2.URL

			client2 := &http.Client{}

			req2, eror2 := http.NewRequest("POST", url2+"/orders", bytes.NewBuffer(values2))

			if eror2 != nil {
				fmt.Println(eror2)
				return
			}

			req2.Header.Set("X-Custom-Header", "myvalue")
			req2.Header.Add("Content-Type", "application/json")
			req.Header.Add("token","eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6Im5ld3NhbXBsZUBnbWFpbC5jb20iLCJleHAiOjE2NTEwNjQ3NDUsImlhdCI6MTY1MDg5MTk0NSwiaXNzIjoibmV3c2FtcGxlQGdtYWlsLmNvbSJ9.PaQ9GBixzNTAYVmQw-k0BVzp1_HX3yhGb0Rf1rj53t4")
			res2, err2 := client2.Do(req2)
			if err2 != nil {
				fmt.Println(err2)
				return
			}

			status2 := res2.Status
			if err != nil {
				return
			}

			Expect(status2).To(BeEquivalentTo("200 OK"))
			
		})

		It("Accepted Criteria", func() {

			payload := map[string]int{"page": 1, "size": 3}
			values, _ := json.Marshal(payload)

			HT := httptest.NewServer(server.TestingServer())
			url := HT.URL

			client := &http.Client{}

			req, eror := http.NewRequest("GET", url+"/orders", bytes.NewBuffer(values))

			if eror != nil {
				fmt.Println(eror)
				return
			}

			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Add("Content-Type", "application/json")
			req.Header.Add("token","eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6Im5ld3NhbXBsZUBnbWFpbC5jb20iLCJleHAiOjE2NTEwNjQ3NDUsImlhdCI6MTY1MDg5MTk0NSwiaXNzIjoibmV3c2FtcGxlQGdtYWlsLmNvbSJ9.PaQ9GBixzNTAYVmQw-k0BVzp1_HX3yhGb0Rf1rj53t4")
			res, err := client.Do(req)
			fmt.Println("the request token: ",req.Header.Get("token"))
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

			fmt.Println("the searrch result", string(body))

			Expect(status).To(BeEquivalentTo("200 OK"))

			
		})
	})
})