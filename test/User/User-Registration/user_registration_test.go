package main_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	// "gin-exercise/db"
	"gin-exercise/entity"
	"gin-exercise/server"
	"net/http"
	"net/http/httptest"

	"io/ioutil"

	// "github.com/casbin/casbin/v2"
	// gormadapter "github.com/casbin/gorm-adapter/v3"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("User Registration", func() {

	Context("Accepted Criteria", func() {
		// var Registerd_User_ID string
		fmt.Println("starting testing")
		It("when new user registers with unique email and phone number", func() {

			payload := map[string]string{"firstname": "Abenezer", "secondname": "kebede", "lastname": "Angamo", "email": "newsample@gmail.com", "profile": `C:\Users\Administrator\Documents\E-commerce\abenezer-dev-prep\temp\3.jpg`, "password": "wordpass", "phonenumber": "251922331100"}
			values, _ := json.Marshal(payload)

			//res,eror := Request("POST", "/users", bytes.NewBuffer(values))

			HT := httptest.NewServer(server.TestingServer())
			url := HT.URL

			client := &http.Client{}
			req, eror := http.NewRequest("POST", url+"/users", bytes.NewBuffer(values))

			if eror != nil {
				fmt.Println(eror)
				return //nil, eror
			}

			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Add("Content-Type", "application/json")

			res, eror := client.Do(req)

			if eror!=nil{
				fmt.Println("error occurred in response",eror)
			}
			body, err := ioutil.ReadAll(res.Body)
			status := res.Status
			if err != nil {
				return
			}
			var registerd_user entity.User
			// It("It has response of data", func() {
				json.Unmarshal(body, &registerd_user)
				fmt.Println(string(body))
			// })
			fmt.Println("the body", string(body))
			fmt.Println("the registerd user", registerd_user)
			// It("has status of code", func() {
				Expect(status).To(BeEquivalentTo("200 OK"))
			// })
		})
	})

	// Context(" Rejected ", func() {
	// 	It("when a user registers with existing email", func() {
	// 		payload := map[string]string{"firstname": "test_person", "secondname": "kebede", "email": "newemail@email.com", "profile": "picture_", "password": "wordpass", "phonenumber": "0923001100"}
	// 		values, _ := json.Marshal(payload)

	// 		res, err := Request("POST", "/users", bytes.NewBuffer(values))
	// 		if err != nil {
	// 			fmt.Println(err)
	// 			return
	// 		}
	// 		defer res.Body.Close()

	// 		body, err := ioutil.ReadAll(res.Body)
	// 		status := res.Status
	// 		if err != nil {
	// 			fmt.Println(err)
	// 			return
	// 		}

	// 		It("It has response of data", func() {
	// 			fmt.Println(string(body))
	// 		})
	// 		It("has status of code", func() {
	// 			Expect(status).To(BeEquivalentTo("400 Bad Request"))
	// 		})

	// 	})

	// 	Context("when a user registers with existing phone number", func() {
	// 		payload := map[string]string{"firstname": "test_person", "secondname": "kebede", "email": "thenewemail@email.com", "profile": "picture_", "password": "wordpass", "phonenumber": "0923002200"}
	// 		values, _ := json.Marshal(payload)

	// 		res, err := Request("POST", "/users", bytes.NewBuffer(values))
	// 		// res, err := client.Do(req)
	// 		if err != nil {
	// 			fmt.Println(err)
	// 			return
	// 		}
	// 		defer res.Body.Close()

	// 		body, err := ioutil.ReadAll(res.Body)
	// 		status := res.Status
	// 		if err != nil {
	// 			fmt.Println(err)
	// 			return
	// 		}

	// 		It("It has response of data", func() {
	// 			fmt.Println(string(body))
	// 		})
	// 		It("has status of code", func() {
	// 			Expect(status).To(BeEquivalentTo("400 Bad Request"))
	// 		})

	// 	})

	// 	Context("when a user registers with empty first name", func() {
	// 		payload := map[string]string{"firstname": "", "secondname": "kebede", "email": "thenewemail@email.com", "profile": "picture_", "password": "wordpass", "phonenumber": "0923002200"}
	// 		values, _ := json.Marshal(payload)

	// 		res, err := Request("POST", "/users", bytes.NewBuffer(values))
	// 		if err != nil {
	// 			fmt.Println(err)
	// 			return
	// 		}
	// 		defer res.Body.Close()

	// 		body, err := ioutil.ReadAll(res.Body)
	// 		status := res.Status
	// 		if err != nil {
	// 			fmt.Println(err)
	// 			return
	// 		}

	// 		It("It has response of data", func() {
	// 			fmt.Println(string(body))
	// 		})
	// 		It("has status of code", func() {
	// 			Expect(status).To(BeEquivalentTo("400 Bad Request"))
	// 		})

	// 	})

	// 	Context("when a user registers with empty second name", func() {
	// 		payload := map[string]string{"firstname": "Abenezer", "secondname": "", "lastname": "Angamo", "email": "thenewemail@email.com", "profile": "picture_", "password": "wordpass", "phonenumber": "0923002200"}
	// 		values, _ := json.Marshal(payload)

	// 		res, err := Request("POST", "/users", bytes.NewBuffer(values))
	// 		if err != nil {
	// 			fmt.Println(err)
	// 			return
	// 		}
	// 		defer res.Body.Close()

	// 		body, err := ioutil.ReadAll(res.Body)
	// 		status := res.Status
	// 		if err != nil {
	// 			fmt.Println(err)
	// 			return
	// 		}

	// 		It("It has response of data", func() {
	// 			fmt.Println(string(body))
	// 		})
	// 		It("has status of code", func() {
	// 			Expect(status).To(BeEquivalentTo("400 Bad Request"))
	// 		})

	// 	})

	// 	Context("when a user registers with empty last name", func() {
	// 		payload := map[string]string{"firstname": "Abenezer", "secondname": "Kebede", "lastname": "", "email": "thenewemail@email.com", "profile": "picture_", "password": "wordpass", "phonenumber": "0923002200"}
	// 		values, _ := json.Marshal(payload)

	// 		res, err := Request("POST", "/users", bytes.NewBuffer(values))
	// 		if err != nil {
	// 			fmt.Println(err)
	// 			return
	// 		}
	// 		defer res.Body.Close()

	// 		body, err := ioutil.ReadAll(res.Body)
	// 		status := res.Status
	// 		if err != nil {
	// 			fmt.Println(err)
	// 			return
	// 		}

	// 		It("It has response of data", func() {
	// 			fmt.Println(string(body))
	// 		})
	// 		It("has status of code", func() {
	// 			Expect(status).To(BeEquivalentTo("400 Bad Request"))
	// 		})

	// 	})

	// 	Context("when a user registers with empty email", func() {
	// 		payload := map[string]string{"firstname": "Abenezer", "secondname": "Kebede", "lastname": "Angamo", "email": "", "profile": "picture_", "password": "wordpass", "phonenumber": "0923002200"}
	// 		values, _ := json.Marshal(payload)

	// 		res, err := Request("POST", "/users", bytes.NewBuffer(values))
	// 		if err != nil {
	// 			fmt.Println(err)
	// 			return
	// 		}
	// 		defer res.Body.Close()

	// 		body, err := ioutil.ReadAll(res.Body)
	// 		status := res.Status
	// 		if err != nil {
	// 			fmt.Println(err)
	// 			return
	// 		}

	// 		It("It has response of data", func() {
	// 			fmt.Println(string(body))
	// 		})
	// 		It("has status of code", func() {
	// 			Expect(status).To(BeEquivalentTo("400 Bad Request"))
	// 		})

	// 	})

	// 	Context("when a user registers with empty profile", func() {
	// 		payload := map[string]string{"firstname": "Abenezer", "secondname": "Kebede", "lastname": "Angamo", "email": "thenewemail@email.com", "profile": "", "password": "wordpass", "phonenumber": "0923002200"}
	// 		values, _ := json.Marshal(payload)

	// 		res, err := Request("POST", "/users", bytes.NewBuffer(values))
	// 		if err != nil {
	// 			fmt.Println(err)
	// 			return
	// 		}
	// 		defer res.Body.Close()

	// 		body, err := ioutil.ReadAll(res.Body)
	// 		status := res.Status
	// 		if err != nil {
	// 			fmt.Println(err)
	// 			return
	// 		}

	// 		It("It has response of data", func() {
	// 			fmt.Println(string(body))
	// 		})
	// 		It("has status of code", func() {
	// 			Expect(status).To(BeEquivalentTo("400 Bad Request"))
	// 		})

	// 	})

	// 	Context("when a user registers with empty password", func() {
	// 		payload := map[string]string{"firstname": "Abenezer", "secondname": "Kebede", "lastname": "Angamo", "email": "thenewemail@email.com", "profile": "picture_", "password": "", "phonenumber": "0923002200"}
	// 		values, _ := json.Marshal(payload)

	// 		res, err := Request("POST", "/users", bytes.NewBuffer(values))
	// 		if err != nil {
	// 			fmt.Println(err)
	// 			return
	// 		}
	// 		defer res.Body.Close()

	// 		body, err := ioutil.ReadAll(res.Body)
	// 		status := res.Status
	// 		if err != nil {
	// 			fmt.Println(err)
	// 			return
	// 		}

	// 		It("It has response of data", func() {
	// 			fmt.Println(string(body))
	// 		})
	// 		It("has status of code", func() {
	// 			Expect(status).To(BeEquivalentTo("400 Bad Request"))
	// 		})

	// 	})

	// 	Context("when a user registers with empty phone number", func() {
	// 		payload := map[string]string{"firstname": "Abenezer", "secondname": "Kebede", "lastname": "Angamo", "email": "thenewemail@email.com", "profile": "picture_", "password": "wordpass", "phonenumber": ""}
	// 		values, _ := json.Marshal(payload)

	// 		res, err := Request("POST", "/users", bytes.NewBuffer(values))
	// 		if err != nil {
	// 			fmt.Println(err)
	// 			return
	// 		}
	// 		defer res.Body.Close()

	// 		body, err := ioutil.ReadAll(res.Body)
	// 		status := res.Status
	// 		if err != nil {
	// 			fmt.Println(err)
	// 			return
	// 		}

	// 		It("It has response of data", func() {
	// 			fmt.Println(string(body))
	// 		})
	// 		It("has status of code", func() {
	// 			Expect(status).To(BeEquivalentTo("400 Bad Request"))
	// 		})

	// 	})

	// })

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
