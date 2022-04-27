package Delete_Store_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gin-exercise/db"
	"gin-exercise/entity"
	"gin-exercise/server"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Store Deletion", func() {

	var _ = Describe("Store Deletion", func() {
		HT := httptest.NewServer(server.TestingServer())
		url := HT.URL
		

		Context("Accepted Criteria", func() {
			var ID string
			BeforeEach(func() {

				loginPayload := map[string]string{"name": "shoa", "location": "Bole,Addis Ababa", "image": "pit"}
				values, _ := json.Marshal(loginPayload)

				HT := httptest.NewServer(server.TestingServer())
				url := HT.URL

				client := &http.Client{}

				req, eror := http.NewRequest("POST", url+"/stores", bytes.NewBuffer(values))

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
				body, err := ioutil.ReadAll(res.Body)
				if err !=nil{
					fmt.Println("the response has error")
				}
				var Savedstore entity.Store
				json.Unmarshal(body, &Savedstore)
				ID = Savedstore.ID
				Expect(status).To(BeEquivalentTo("200 OK"))
				res.Body.Close()
			})

			It("Delete registered store", func() {

				loginPayload := map[string]string{"id": ID}
				values, _ := json.Marshal(loginPayload)
				fmt.Println("the ID: ", ID)
				client := &http.Client{}
				req, err := http.NewRequest("DELETE", url+"/stores/"+ID, bytes.NewBuffer(values))

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
				if err != nil {

					return
				}

				Expect(status).To(BeEquivalentTo("200 OK"))
				//	})
			})

			AfterEach(func() {
				st, _ :=db.StoreSearch(ID, "id")
				fmt.Println("the deleted sesarch result: ",st)
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
