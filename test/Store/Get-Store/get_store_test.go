package Get_Store_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gin-exercise/db"
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


			
			loginPayload := map[string]string{"name": "shoa", "location": "Bole,Addis Ababa", "image": "pit"}
			values, _ := json.Marshal(loginPayload)
			
			HT := httptest.NewServer(server.TestingServer())
			url := HT.URL
		
			client := &http.Client{}
			
			req, eror := http.NewRequest("POST", url+"/stores",  bytes.NewBuffer(values))
		
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
			})
		
		It("Accepted Criteria", func() {
			

				payload := map[string]string{"name": "shoa"}
				values, _ := json.Marshal(payload)


				HT := httptest.NewServer(server.TestingServer())
				url := HT.URL
			
				client := &http.Client{}
				
				req, eror := http.NewRequest("GET", url+"/store",  bytes.NewBuffer(values))
			
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
			
					fmt.Println("the searrch result",string(body))
			
					Expect(status).To(BeEquivalentTo("200 OK"))

			
			})
			It("Failing Senario, trying to search with store that doesn't exist", func() {
				
	
					payload := map[string]string{"name": "bilos"}
					values, _ := json.Marshal(payload)
	
					HT := httptest.NewServer(server.TestingServer())
				url := HT.URL
			
				client := &http.Client{}
				
				req, eror := http.NewRequest("GET", url+"/store",  bytes.NewBuffer(values))
			
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
					defer res.Body.Close()
	
					body, err := ioutil.ReadAll(res.Body)
					status := res.Status
					if err != nil {
						fmt.Println(err)
						return
					}
	
				
						Expect(body).To(BeEmpty())
				
						Expect(status).To(BeEquivalentTo("200 OK"))				
				
				})
			AfterEach(func(){
				db.StoreDelete("shoa", "name")
			})		
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
