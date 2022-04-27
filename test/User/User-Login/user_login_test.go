package User_Login_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	// "gin-exercise/entity"
	"gin-exercise/server"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	// "time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("User Registration", func() {

	HT := httptest.NewServer(server.TestingServer())
	url := HT.URL

	Context("Accepted Criteria", func() {
		// BeforeEach(
		// 	func() {
		// 		payload := map[string]string{"firstname": "Abenezer", "secondname": "kebede", "lastname": "Angamo", "email": "newsample@gmail.com", "profile": `C:\Users\Administrator\Documents\E-commerce\abenezer-dev-prep\temp\3.jpg`, "password": "wordpass", "phonenumber": "251911111100"}
		// 		values, _ := json.Marshal(payload)

		// 		res, eror := Request("POST", "/users", bytes.NewBuffer(values))
		// 		if eror != nil {
		// 			return
		// 		}
		// 		body, err := ioutil.ReadAll(res.Body)
		// 		status := res.Status
		// 		if err != nil {
		// 			return
		// 		}
		// 		var registerd_user entity.User
				
		// 		json.Unmarshal(body, &registerd_user)
		// 		fmt.Println(string(body))
				
		// 		fmt.Println("the body", string(body))
		// 		fmt.Println("the registerd user", registerd_user)
		// 		time.Sleep(time.Second*10)
		// 			Expect(status).To(BeEquivalentTo("200 OK"))
				
		// 	})
		It("when new user is login with correct email and password given that the user is already registerd", func() {

			loginPayload := map[string]string{"email": "newsample@gmail.com", "password": "wordpass"}
			values, _ := json.Marshal(loginPayload)

			client := &http.Client{}
			req, err := http.NewRequest("POST", url+"/login", bytes.NewBuffer(values))

			if err != nil {
				fmt.Println(err)
				return
			}

			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Add("Content-Type", "application/json")

			response, err := client.Do(req)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer response.Body.Close()

		fmt.Println("the response: ",response)
				Expect(response.Status).To(BeEquivalentTo("200 OK"))
		
		})
	})

	Context(" Rejected ", func() {

		Context("when new user is loging in with incorrect email and correct password", func() {

			loginPayload := map[string]string{"email": "wsample@gmail.com", "password": "wordpass"}
			values, _ := json.Marshal(loginPayload)

			client := &http.Client{}
			req, err := http.NewRequest("POST", url+"/login", bytes.NewBuffer(values))

			if err != nil {
				fmt.Println(err)
				return
			}

			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Add("Content-Type", "application/json")

			response, err := client.Do(req)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer response.Body.Close()

			It("has status of code", func() {
				Expect(response.Status).To(BeEquivalentTo("401 Unauthorized"))
			})
		})

	})

	Context("when new user is loging in with correct email and incorrect password", func() {

		loginPayload := map[string]string{"email": "newsample@gmail.com", "password": "rdpass"}
		values, _ := json.Marshal(loginPayload)

		client := &http.Client{}
		req, err := http.NewRequest("POST", url+"/login", bytes.NewBuffer(values))

		if err != nil {
			fmt.Println(err)
			return
		}

		req.Header.Set("X-Custom-Header", "myvalue")
		req.Header.Add("Content-Type", "application/json")

		response, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer response.Body.Close()

		It("has status of code", func() {
			Expect(response.Status).To(BeEquivalentTo("401 Unauthorized"))
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
