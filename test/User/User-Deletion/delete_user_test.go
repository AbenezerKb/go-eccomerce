package User_deletion_test

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

var _ = Describe("User Registration", func() {

	var _ = Describe("User Registration", func() {
		HT := httptest.NewServer(server.TestingServer())
		url := HT.URL
		var ID string

		Context("Accepted Criteria", func() {

			Context("when new user is registering with unique email and phone number", func() {

				payload := map[string]string{"firstname": "Abenezer", "secondname": "kebede","lastname":"Angamo", "email": "newsample@gmail.com", "profile": "picture_", "password": "wordpass", "phonenumber": "251911111100"}
				values, _ := json.Marshal(payload)

				client := &http.Client{}
				req, err := http.NewRequest("POST", url+"/users", bytes.NewBuffer(values))

				if err != nil {
					fmt.Println(err)
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
				var Saveduser entity.User
				json.Unmarshal(body, &Saveduser)
				ID = Saveduser.Id

				if err != nil {
					//fmt.Println(err)
					return
				}

				It("It has response of data", func() {
					fmt.Println(string(body))
				})
				It("has status of code", func() {
					Expect(status).To(BeEquivalentTo("200 OK"))
				})
			})

			It("Delete registered user", func() {

				loginPayload := map[string]string{"id": ID}
				values, _ := json.Marshal(loginPayload)
fmt.Println("the ID: ",ID)
				client := &http.Client{}
				req, err := http.NewRequest("DELETE", url+"/user/"+ID, bytes.NewBuffer(values))

				if err != nil {
					fmt.Println("client detail error: ", err)
					return
				}
				req.Header.Add("token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Im5ld2VtYWlsQGdtYWlsLmNvbSIsImV4cCI6MTY0OTU4NjQ1OSwiaWF0IjoxNjQ5NDEzNjU5LCJpc3MiOiJuZXdlbWFpbEBnbWFpbC5jb20ifQ.e3VjjA4w-rApEuC9N6BbmjRiRDRP9tb6g71y07HQhAc")
				req.Header.Set("X-Custom-Header", "myvalue")
				req.Header.Add("Content-Type", "application/json")

				res, err := client.Do(req)
				if err != nil {
					fmt.Println(err)
					return
				}
				defer res.Body.Close()

				//body, err := ioutil.ReadAll(res.Body)
				status := res.Status
				if err != nil {

					return
				}

				Expect(status).To(BeEquivalentTo("200 OK"))
				//	})
			})

		})
	})
})

// 	Context(" Rejected ", func() {
// 		Context("when a user registers with existing email", func() {
// 			payload := map[string]string{"firstname": "test_person", "secondname": "kebede", "email": "newemail@email.com", "profile": "picture_", "password": "wordpass", "phonenumber": "0923001100"}
// 			values, _ := json.Marshal(payload)

// 			client := &http.Client{}
// 			req, err := http.NewRequest("POST", UserRegistrationurl, bytes.NewBuffer(values))

// 			if err != nil {
// 				fmt.Println(err)
// 				return
// 			}

// 			req.Header.Set("X-Custom-Header", "myvalue")
// 			req.Header.Add("Content-Type", "application/json")

// 			res, err := client.Do(req)
// 			if err != nil {
// 				fmt.Println(err)
// 				return
// 			}
// 			defer res.Body.Close()

// 			body, err := ioutil.ReadAll(res.Body)
// 			status := res.Status
// 			if err != nil {
// 				fmt.Println(err)
// 				return
// 			}

// 			It("It has response of data", func() {
// 				fmt.Println(string(body))
// 			})
// 			It("has status of code", func() {
// 				Expect(status).To(BeEquivalentTo("400 Bad Request"))
// 			})

// 		})

// 		Context("when a user registers with existing phone number", func() {
// 			payload := map[string]string{"firstname": "test_person", "secondname": "kebede", "email": "thenewemail@email.com", "profile": "picture_", "password": "wordpass", "phonenumber": "0923002200"}
// 			values, _ := json.Marshal(payload)

// 			client := &http.Client{}
// 			req, err := http.NewRequest("POST", UserRegistrationurl, bytes.NewBuffer(values))

// 			if err != nil {
// 				fmt.Println(err)
// 				return
// 			}

// 			req.Header.Set("X-Custom-Header", "myvalue")
// 			req.Header.Add("Content-Type", "application/json")

// 			res, err := client.Do(req)
// 			if err != nil {
// 				fmt.Println(err)
// 				return
// 			}
// 			defer res.Body.Close()

// 			body, err := ioutil.ReadAll(res.Body)
// 			status := res.Status
// 			if err != nil {
// 				fmt.Println(err)
// 				return
// 			}

// 			It("It has response of data", func() {
// 				fmt.Println(string(body))
// 			})
// 			It("has status of code", func() {
// 				Expect(status).To(BeEquivalentTo("400 Bad Request"))
// 			})

// 		})

// 		Context("when a user registers with empty first name", func() {
// 			payload := map[string]string{"firstname": "", "secondname": "kebede", "email": "thenewemail@email.com", "profile": "picture_", "password": "wordpass", "phonenumber": "0923002200"}
// 			values, _ := json.Marshal(payload)

// 			client := &http.Client{}
// 			req, err := http.NewRequest("POST", UserRegistrationurl, bytes.NewBuffer(values))

// 			if err != nil {
// 				fmt.Println(err)
// 				return
// 			}

// 			req.Header.Set("X-Custom-Header", "myvalue")
// 			req.Header.Add("Content-Type", "application/json")

// 			res, err := client.Do(req)
// 			if err != nil {
// 				fmt.Println(err)
// 				return
// 			}
// 			defer res.Body.Close()

// 			body, err := ioutil.ReadAll(res.Body)
// 			status := res.Status
// 			if err != nil {
// 				fmt.Println(err)
// 				return
// 			}

// 			It("It has response of data", func() {
// 				fmt.Println(string(body))
// 			})
// 			It("has status of code", func() {
// 				Expect(status).To(BeEquivalentTo("400 Bad Request"))
// 			})

// 		})

// 		Context("when a user registers with empty second name", func() {
// 			payload := map[string]string{"firstname": "Abenezer", "secondname": "", "lastname": "Angamo", "email": "thenewemail@email.com", "profile": "picture_", "password": "wordpass", "phonenumber": "0923002200"}
// 			values, _ := json.Marshal(payload)

// 			client := &http.Client{}
// 			req, err := http.NewRequest("POST", UserRegistrationurl, bytes.NewBuffer(values))

// 			if err != nil {
// 				fmt.Println(err)
// 				return
// 			}

// 			req.Header.Set("X-Custom-Header", "myvalue")
// 			req.Header.Add("Content-Type", "application/json")

// 			res, err := client.Do(req)
// 			if err != nil {
// 				fmt.Println(err)
// 				return
// 			}
// 			defer res.Body.Close()

// 			body, err := ioutil.ReadAll(res.Body)
// 			status := res.Status
// 			if err != nil {
// 				fmt.Println(err)
// 				return
// 			}

// 			It("It has response of data", func() {
// 				fmt.Println(string(body))
// 			})
// 			It("has status of code", func() {
// 				Expect(status).To(BeEquivalentTo("400 Bad Request"))
// 			})

// 		})

// 		Context("when a user registers with empty last name", func() {
// 			payload := map[string]string{"firstname": "Abenezer", "secondname": "Kebede", "lastname": "", "email": "thenewemail@email.com", "profile": "picture_", "password": "wordpass", "phonenumber": "0923002200"}
// 			values, _ := json.Marshal(payload)

// 			client := &http.Client{}
// 			req, err := http.NewRequest("POST", UserRegistrationurl, bytes.NewBuffer(values))

// 			if err != nil {
// 				fmt.Println(err)
// 				return
// 			}

// 			req.Header.Set("X-Custom-Header", "myvalue")
// 			req.Header.Add("Content-Type", "application/json")

// 			res, err := client.Do(req)
// 			if err != nil {
// 				fmt.Println(err)
// 				return
// 			}
// 			defer res.Body.Close()

// 			body, err := ioutil.ReadAll(res.Body)
// 			status := res.Status
// 			if err != nil {
// 				fmt.Println(err)
// 				return
// 			}

// 			It("It has response of data", func() {
// 				fmt.Println(string(body))
// 			})
// 			It("has status of code", func() {
// 				Expect(status).To(BeEquivalentTo("400 Bad Request"))
// 			})

// 		})

// 		Context("when a user registers with empty email", func() {
// 			payload := map[string]string{"firstname": "Abenezer", "secondname": "Kebede", "lastname": "Angamo", "email": "", "profile": "picture_", "password": "wordpass", "phonenumber": "0923002200"}
// 			values, _ := json.Marshal(payload)

// 			client := &http.Client{}
// 			req, err := http.NewRequest("POST", UserRegistrationurl, bytes.NewBuffer(values))

// 			if err != nil {
// 				fmt.Println(err)
// 				return
// 			}

// 			req.Header.Set("X-Custom-Header", "myvalue")
// 			req.Header.Add("Content-Type", "application/json")

// 			res, err := client.Do(req)
// 			if err != nil {
// 				fmt.Println(err)
// 				return
// 			}
// 			defer res.Body.Close()

// 			body, err := ioutil.ReadAll(res.Body)
// 			status := res.Status
// 			if err != nil {
// 				fmt.Println(err)
// 				return
// 			}

// 			It("It has response of data", func() {
// 				fmt.Println(string(body))
// 			})
// 			It("has status of code", func() {
// 				Expect(status).To(BeEquivalentTo("400 Bad Request"))
// 			})

// 		})

// 		Context("when a user registers with empty profile", func() {
// 			payload := map[string]string{"firstname": "Abenezer", "secondname": "Kebede", "lastname": "Angamo", "email": "thenewemail@email.com", "profile": "", "password": "wordpass", "phonenumber": "0923002200"}
// 			values, _ := json.Marshal(payload)

// 			client := &http.Client{}
// 			req, err := http.NewRequest("POST", UserRegistrationurl, bytes.NewBuffer(values))

// 			if err != nil {
// 				fmt.Println(err)
// 				return
// 			}

// 			req.Header.Set("X-Custom-Header", "myvalue")
// 			req.Header.Add("Content-Type", "application/json")

// 			res, err := client.Do(req)
// 			if err != nil {
// 				fmt.Println(err)
// 				return
// 			}
// 			defer res.Body.Close()

// 			body, err := ioutil.ReadAll(res.Body)
// 			status := res.Status
// 			if err != nil {
// 				fmt.Println(err)
// 				return
// 			}

// 			It("It has response of data", func() {
// 				fmt.Println(string(body))
// 			})
// 			It("has status of code", func() {
// 				Expect(status).To(BeEquivalentTo("400 Bad Request"))
// 			})

// 		})

// 		Context("when a user registers with empty password", func() {
// 			payload := map[string]string{"firstname": "Abenezer", "secondname": "Kebede", "lastname": "Angamo", "email": "thenewemail@email.com", "profile": "picture_", "password": "", "phonenumber": "0923002200"}
// 			values, _ := json.Marshal(payload)

// 			client := &http.Client{}
// 			req, err := http.NewRequest("POST", UserRegistrationurl, bytes.NewBuffer(values))

// 			if err != nil {
// 				fmt.Println(err)
// 				return
// 			}

// 			req.Header.Set("X-Custom-Header", "myvalue")
// 			req.Header.Add("Content-Type", "application/json")

// 			res, err := client.Do(req)
// 			if err != nil {
// 				fmt.Println(err)
// 				return
// 			}
// 			defer res.Body.Close()

// 			body, err := ioutil.ReadAll(res.Body)
// 			status := res.Status
// 			if err != nil {
// 				fmt.Println(err)
// 				return
// 			}

// 			It("It has response of data", func() {
// 				fmt.Println(string(body))
// 			})
// 			It("has status of code", func() {
// 				Expect(status).To(BeEquivalentTo("400 Bad Request"))
// 			})

// 		})

// 		Context("when a user registers with empty phone number", func() {
// 			payload := map[string]string{"firstname": "Abenezer", "secondname": "Kebede", "lastname": "Angamo", "email": "thenewemail@email.com", "profile": "picture_", "password": "wordpass", "phonenumber": ""}
// 			values, _ := json.Marshal(payload)

// 			client := &http.Client{}
// 			req, err := http.NewRequest("POST", url, bytes.NewBuffer(values))

// 			if err != nil {
// 				fmt.Println(err)
// 				return
// 			}

// 			req.Header.Set("X-Custom-Header", "myvalue")
// 			req.Header.Add("Content-Type", "application/json")

// 			res, err := client.Do(req)
// 			if err != nil {
// 				fmt.Println(err)
// 				return
// 			}
// 			defer res.Body.Close()

// 			body, err := ioutil.ReadAll(res.Body)
// 			status := res.Status
// 			if err != nil {
// 				fmt.Println(err)
// 				return
// 			}

// 			It("It has response of data", func() {
// 				fmt.Println(string(body))
// 			})
// 			It("has status of code", func() {
// 				Expect(status).To(BeEquivalentTo("400 Bad Request"))
// 			})

// 		})

// 	})

// })
// })
