package Update_Store_test

import (
	"bytes"
	"encoding/json"
	"fmt"	
	"gin-exercise/entity"
	"gin-exercise/server"	
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)



var _ = Describe("Store Status Update", func() {
	
	var ID string

	Context("Accepted Tests", func() {
		BeforeEach(func() {		
			
			loginPayload := map[string]string{"id": 6}
			values, _ := json.Marshal(loginPayload)
			HT := httptest.NewServer(server.TestingServer())
			url := HT.URL
		
			client := &http.Client{}
			
			req, eror := http.NewRequest("GET", url+"/stores/6/status",  bytes.NewBuffer(values))
		
			if eror != nil {
				rest_error.Unable_to_read(eror.Error())

				return 
			}
		
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Add("Content-Type", "application/json")
		
			res, err1 := client.Do(req)
			if err1 != nil {
				rest_error.Unable_to_save(err1.Error())
				return 
			}
							

			body, err2 := ioutil.ReadAll(res.Body)
			status := res.Status
			if err2 != nil {
				rest_error.Unable_to_save(err2.Error())
				return
			}
			
			var Savedstore entity.Store
			json.Unmarshal(body, &Savedstore)
			ID = Savedstore.ID
			
			Expect(status).To(BeEquivalentTo("200 OK"))
			res.Body.Close()
		})

		It("Update registered Store", func() {

			loginPayload := map[string]string{"name": "lomiad", "location": "Ayat,Addis Ababa" }
			values, _ := json.Marshal(loginPayload)
			

			HT := httptest.NewServer(server.TestingServer())
			url := HT.URL
		
			client := &http.Client{}
			
			req, eror := http.NewRequest("PATCH", url+"/stores/"+ID,  bytes.NewBuffer(values))
		
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
							

			body, err := ioutil.ReadAll(res.Body)
			status := res.Status
			if err != nil {

				return
			}

			var savedstore entity.Store
			json.Unmarshal(body, &savedstore)		
			Expect(status).To(BeEquivalentTo("200 OK"))
			res.Body.Close()
		})

		

	})

	
})

