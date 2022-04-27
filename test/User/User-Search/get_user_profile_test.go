package User_Search_test

import (
	"bytes"
	"encoding/json"
	"fmt"

	// "gin-exercise/controller"
	// "gin-exercise/db"
	"gin-exercise/entity"
	"gin-exercise/server"

	// "gin-exercise/service"
	// "io"
	"io/ioutil"
	// "mime/multipart"
	"net/http"
	"net/http/httptest"

	// "os"
	// "path/filepath"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)


var _ = Describe("User Registration", func() {
	HT := httptest.NewServer(server.TestingServer())
	url := HT.URL	

	Context("Accepted Tests", func() {
		BeforeEach(func() {

			loginPayload := map[string]string{"firstname": "john", "secondname": "smith", "lastname": "steven", "email": "newtesting@gmail.com", "password": "wordpass", "phonenumber": "251911111122", "imageurl": `C:\Users\Administrator\Documents\E-commerce\abenezer-dev-prep\temp\3.jpg`}
			values, _ := json.Marshal(loginPayload)

			res,eror := Request("POST", "/users", bytes.NewBuffer(values))			
			if eror != nil {
				fmt.Println(eror)
				return
			}
			defer res.Body.Close()

			body, err := ioutil.ReadAll(res.Body)
			status := res.Status
			if err != nil {
				return
			}
			var saveduser entity.User
			json.Unmarshal(body, &saveduser)

			Expect(status).To(BeEquivalentTo("200 OK"))

		})

		It("Get registered user profile", func() {
			
			loginPayload := map[string]string{"email": "newtesting@gmail.com"}
			values, _ := json.Marshal(loginPayload)

			client := &http.Client{}
			req, err := http.NewRequest("GET", url+"/user", bytes.NewBuffer(values))

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

			body, err := ioutil.ReadAll(res.Body)
			status := res.Status
			if err != nil {
				return
			}
			var saveduser entity.User
			json.Unmarshal(body, &saveduser)		

			Expect(status).To(BeEquivalentTo("200 OK"))

		})

	})
	

	Context("Failing Tests", func() {
		
		It("Get registered user profile with wrong email", func() {

			time.Sleep(2 * time.Second)

			loginPayload := map[string]string{"email": "testing@gmail.com"}
			values, _ := json.Marshal(loginPayload)

			client := &http.Client{}
			req, err := http.NewRequest("GET", url+"/user", bytes.NewBuffer(values))

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

			body, err := ioutil.ReadAll(res.Body)
			status := res.Status
			if err != nil {			
				return
			}
			var saveduser entity.User
			json.Unmarshal(body, &saveduser)

		
			fmt.Println(saveduser)
		
			Expect(status).To(BeEquivalentTo("400 Bad Request"))
		

		})

	})

})





func Request(method string, api string, buffer *bytes.Buffer) (*http.Response, error) {
	HT := httptest.NewServer(server.TestingServer())
	url := HT.URL

	client := &http.Client{}
	req, eror := http.NewRequest("POST", url+api, buffer)

	if eror != nil {
		fmt.Println(eror)
		return nil, eror
	}

	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	_, er := ioutil.ReadAll(res.Body)
	// _ := res.Status
	if er != nil {
		//fmt.Println(err)
		return nil, er
	}
	return res, nil
}
