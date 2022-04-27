package Get_Orders_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	// "gin-exercise/db"
	"gin-exercise/entity"
	"gin-exercise/server"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Order Functionalities", func() {
	var _ = Describe("Get Order", func() {
		var ID string
		BeforeEach(func() {

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
			res.Body.Close()

			order1 := entity.Order{OrderOwner: "Abenezer", Items_quantity: 20, Total_price: 1300}
			values1, _ := json.Marshal(order1)

			HT1 := httptest.NewServer(server.TestingServer())
			url1 := HT1.URL

			client1 := &http.Client{}

			req1, eror1 := http.NewRequest("POST", url1+"/orders", bytes.NewBuffer(values1))

			if eror1 != nil {
				fmt.Println(eror1)
				return
			}

			req1.Header.Set("X-Custom-Header", "myvalue")
			req1.Header.Add("Content-Type", "application/json")

			res1, err1 := client1.Do(req1)
			if err1 != nil {
				fmt.Println(err1)
				return
			}

			status1 := res1.Status
			if err != nil {
				return
			}

			Expect(status1).To(BeEquivalentTo("200 OK"))
			res.Body.Close()

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
			res2.Body.Close()
		})

		It("Accepted Criteria", func() {

			payload := map[string]string{"name": "shoa"}
			values, _ := json.Marshal(payload)

			HT := httptest.NewServer(server.TestingServer())
			url := HT.URL

			client := &http.Client{}
ID = "06fe2b90-4e8d-4ab4-98a5-60ac85d97e4f"
			req, eror := http.NewRequest("GET", url+"/orders/"+ID, bytes.NewBuffer(values))

			if eror != nil {
				fmt.Println(eror)
				return
			}

			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Add("Content-Type", "application/json")

			res, err := client.Do(req)

			//	res, eror := Request("GET", "/store", bytes.NewBuffer(values))
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

			fmt.Println("the searrch result in response", string(body))

			Expect(status).To(BeEquivalentTo("200 OK"))

		})
		// It("Failing Senario, trying to search with store that doesn't exist", func() {

		// 	payload := map[string]string{"name": "bilos"}
		// 	values, _ := json.Marshal(payload)

		// 	HT := httptest.NewServer(server.TestingServer())
		// 	url := HT.URL

		// 	client := &http.Client{}

		// 	req, eror := http.NewRequest("GET", url+"/store", bytes.NewBuffer(values))

		// 	if eror != nil {
		// 		fmt.Println(eror)
		// 		return
		// 	}

		// 	req.Header.Set("X-Custom-Header", "myvalue")
		// 	req.Header.Add("Content-Type", "application/json")

		// 	res, err := client.Do(req)
		// 	if err != nil {
		// 		fmt.Println(err)
		// 		return
		// 	}
		// 	defer res.Body.Close()

		// 	body, err := ioutil.ReadAll(res.Body)
		// 	status := res.Status
		// 	if err != nil {
		// 		fmt.Println(err)
		// 		return
		// 	}

		// 	Expect(body).To(BeEmpty())

		// 	Expect(status).To(BeEquivalentTo("200 OK"))

		// })
		// AfterEach(func() {
		// 	db.StoreDelete("shoa", "name")
		// })
	})
})

// func Request(method string, api string, buffer *bytes.Buffer) (*http.Response, error) {
// 	HT := httptest.NewServer(server.TestingServer())
// 	url := HT.URL

// 	client := &http.Client{}
// 	req, eror := http.NewRequest(method, url+api, buffer)

// 	if eror != nil {
// 		fmt.Println(eror)
// 		return nil, eror
// 	}

// 	req.Header.Set("X-Custom-Header", "myvalue")
// 	req.Header.Add("Content-Type", "application/json")

// 	res, err := client.Do(req)
// 	if err != nil {
// 		fmt.Println(err)
// 		return nil, err
// 	}
// 	defer res.Body.Close()

// 	_, er := ioutil.ReadAll(res.Body)

// 	if er != nil {

// 		return nil, er
// 	}
// 	return res, nil
// }
