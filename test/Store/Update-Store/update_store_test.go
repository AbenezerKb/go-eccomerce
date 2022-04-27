package Update_Store_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	// "io"
	// "mime/multipart"
	// "os"
	// "path/filepath"

	// "gin-exercise/controller"
	"gin-exercise/entity"
	"gin-exercise/server"
	// "gin-exercise/db"

	// "gin-exercise/service"

	// "io"
	"io/ioutil"
	// "mime/multipart"
	"net/http"
	"net/http/httptest"

	// "os"
	// "path/filepath"
	// "time"

	// "github.com/gin-contrib/sessions"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)



var _ = Describe("Store Update", func() {
	
	var ID string

	Context("Accepted Tests", func() {
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
							

			body, err := ioutil.ReadAll(res.Body)
			status := res.Status
			if err != nil {
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

