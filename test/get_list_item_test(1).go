package main_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	// "io/ioutil"
	// "time"

	"gin-exercise/db"
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

	HT := httptest.NewServer(server.TestingServer())
	url := HT.URL
	
	Context("Accepted Criteria", func() {
	
		BeforeEach(func() {
			
	
				// Expect(db.Count(entity.Item{})).To(BeZero())
	
			
		})
		It("Registering Item", func() {			
			
		var item entity.Item
		item.Name="pc-bag"
		item.StoreID="2"
		
			values, _ := json.Marshal(item)

			client := &http.Client{}
			req, err := http.NewRequest("POST", url+"/items", bytes.NewBuffer(values))

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

			defer res.Body.Close()
		
			status := res.Status
		
			Expect(status).To(BeEquivalentTo("200 OK"))

		})

	AfterEach(func() {
			
	
				Expect(db.Count(entity.User{})).To(BeZero())
	
			
		})
	

	})
	

})
