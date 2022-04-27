package User_Update_test

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



var _ = Describe("User Registration", func() {
	
	var ID string
	
	Context("Accepted Tests", func() {
		BeforeEach(func() {
			// payload := &bytes.Buffer{}
			// writer := multipart.NewWriter(payload)
			// file, errFile1 := os.Open("C:/Users/Administrator/Downloads/index1.jpg")
			// if errFile1 != nil {
			// 	fmt.Println("image path1", errFile1)
			// 	return
			// }

			// defer file.Close()

			// part1, errFile2 := writer.CreateFormFile("file", filepath.Base("C:/Users/Administrator/Downloads/index1.jpg"))
			// if errFile2 != nil {
			// 	fmt.Println("image path2", errFile2)
			// 	return
			// }

			// _, errFile3 := io.Copy(part1, file)
			// if errFile3 != nil {
			// 	fmt.Println("image path3", errFile3)
			// 	return
			// }
			// err := writer.Close()
			// if err != nil {
			// 	fmt.Println("image copying: ", err)
			// 	return
			// }

			// Imgclient := &http.Client{}
			// Imgreq, err := http.NewRequest("POST", url+"/image", payload)

			// if err != nil {
			// 	fmt.Println(err)
			// 	return
			// }
			// Imgreq.Header.Set("Content-Type", writer.FormDataContentType())
			// Imgres, err := Imgclient.Do(Imgreq)
			// if err != nil {
			// 	fmt.Println("image client error", err)
			// 	return
			// }

			// fmt.Println("the image response: ",Imgres)
			// defer Imgres.Body.Close()

			// //Imgbody, err := ioutil.ReadAll(Imgres.Body)
			// if err != nil {
			// 	fmt.Println(err)
			// }

			
			loginPayload := map[string]string{"firstname": "john", "secondname": "smith", "lastname": "steven", "email": "newemailtesting@gmail.com", "password": "wordpass", "phonenumber": "0911111122", "imageurl": `C:\Users\Administrator\Documents\E-commerce\abenezer-dev-prep\temp\3.jpg`}
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
			var Saveduser entity.User
			json.Unmarshal(body, &Saveduser)
			ID = Saveduser.Id

			fmt.Println("the status: ", status)
			fmt.Println("the ID: ", ID)
			Expect(status).To(BeEquivalentTo("200 OK"))

		})

		It("Update registered user profile", func() {

			loginPayload := map[string]string{"firstname": "Bob", "secondname": "jack"}
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
			fmt.Println("the updated user: ",string(body))
			Expect(status).To(BeEquivalentTo("200 OK"))

		})



		AfterEach(func() {


//			Expect(db.Delete("newemailtesting@gmail.com", "email")).To(BeTrue())

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
