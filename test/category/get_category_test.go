package main_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"

	// "gin-exercise/db"
	"gin-exercise/entity"
	"gin-exercise/server"

	// "io/ioutil"
	"net/http"
	"net/http/httptest"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("User Registration", func() {

	time.Sleep(time.Second * 10)

	Context("Accepted Criteria", func() {

		var _ = Describe("User Registration", func() {
			HT := httptest.NewServer(server.TestingServer())
			url := HT.URL
			// var ID string

			BeforeEach(func() {

				// loginPayload := map[string]string{"firstname": "john", "secondname": "smith", "lastname": "steven", "email": "newemailtesting@gmail.com", "password": "wordpass", "phonenumber": "0911111122", "imageurl": "picture_"}
				loginPayload := map[string]string{"name": "Food and Beverage"}
				values, _ := json.Marshal(loginPayload)

				client := &http.Client{}
				req, err := http.NewRequest("POST", url+"/categories", bytes.NewBuffer(values))

				if err != nil {
					fmt.Println("client detail error: ", err)
					return
				}

				req.Header.Set("X-Custom-Header", "myvalue")
				req.Header.Add("Content-Type", "application/json")

				res, err := client.Do(req)

				if err != nil {
					fmt.Println(err)
					return
				}

				


				status := res.Status
			 
				Expect(status).To(BeEquivalentTo("200 OK"))

				res.Body.Close()

			})

			Context("get all categories", func() {

				payload := map[string]int{"page": 1, "size": 3}
				values, _ := json.Marshal(payload)

				client := &http.Client{}
				req, err := http.NewRequest("GET", url+"/categories", bytes.NewBuffer(values))

				if err != nil {
					fmt.Println(err)
					return
				}
				req.Header.Add("token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRoZWVtYWlsQGdtYWlsLmNvbSIsImV4cCI6MTY0OTc4NDY2NCwiaWF0IjoxNjQ5NjExODY0LCJpc3MiOiJ0aGVlbWFpbEBnbWFpbC5jb20ifQ.dP4ArdetAa4PuToQrDqXFe_v_0KAfVEQHLV54muCDuA")
				req.Header.Set("X-Custom-Header", "myvalue")
				req.Header.Add("Content-Type", "application/json")

				res, err := client.Do(req)
				if err != nil {
					fmt.Println(err)
					return
				}
				defer res.Body.Close()

				body, err := ioutil.ReadAll(res.Body)
				status := res.Status
				if err != nil {
					fmt.Println(err)
					return
				}
				var savedCategories []entity.Category
				json.Unmarshal(body, &savedCategories)
				fmt.Println(savedCategories)
			
				// It("It has response of data", func() {
				// 	fmt.Println(string(body))
				// })
				It("has status of code", func() {
					Expect(status).To(BeEquivalentTo("200 OK"))
				})

			})
			AfterEach(func() {

				//	It("deletes registerd user", func() {

			//	Expect(db.Delete("Food and Beverage", "name")).To(BeTrue())

				//	})

			})

		})
	})

})
