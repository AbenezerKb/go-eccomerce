package service

import (
	//"fmt"
	// "net/http"
	"io"
	"mime/multipart"
	"os"
	"path"

	"github.com/gin-gonic/gin"
)

type ImageService interface {
	Save(*multipart.FileHeader, string)  error  //(*entity.User, *rest_error.RestErr)
	Display(ctx *gin.Context) //string
}

type imageService struct {
	image string
}

func NewImage() ImageService {
	return &imageService{}
}

func (i *imageService) Save(file *multipart.FileHeader, id string) error /*(*entity.User, *rest_error.RestErr)*/ {

	
	path.Join("./temp",id)
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(id)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	if err != nil {
		return err
	}
	return nil



	// //TODO handle it using gin
	// in, fileHeader, err := ctx.FormFile("image")
	// if err != nil {
	// 	rest_error.NewBadRequestError("image upload failed")
	// }
	// //fileHeader.Filename
	// //defer in.Close()

	// // filename, err := ctx.FormFile("image")
	// // if err !=nil{
	// // 	rest_error.NewBadRequestError("image upload failed")
	// // }
	// defer in.Close()

	// //path.Join("./static",image)
	// out, err := os.OpenFile(fileHeader.Filename, os.O_WRONLY, 0644)
	// if err != nil {
	// 	rest_error.NewInternalServerError("image saving failed")
	// }
	// defer out.Close()
	// io.Copy(out, in)
	// pw, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	// if err != nil {
	// 	return nil, rest_error.NewInternalServerError("password encryption falied")
	// }
	// user.Password = string(pw[:])
	// fmt.Println(user)
	// db.SaveUser(user)
	//	return &user, nil
}

func (i *imageService) Display(ctx *gin.Context) {
	// fn := ctx.Param("userid")
	// str, _ := os.Stat(filepath.Join("v1/image/", fn, ".jpg"))
	// static.Serve("/img", static.LocalFile("./img", true))
	// ctx.Static("/img", "./img")
	// ctx.
	// ctx.File(str.Name())

	// fmt.Println("herer,", filepath.Join("v1/image/", fn))
	// if _, err := os.Stat(filepath.Join("v1/image/", fn, ".jpg")); err == nil {
	// 	fmt.Println("file exists and returned", filepath.Join("v1/image/", fn))
	// 	ctx.JSON(200, filepath.Join("v1/image/", fn))
	// }

	//return "db.Userslist()"
}

// const MAX_UPLOAD_SIZE = 32

func saveFileHandler(file *multipart.FileHeader, id string) error {
	//	h,k:=c.MultipartForm()

	// var w http.ResponseWriter = c.Writer
	// c.Request.Body = http.MaxBytesReader(w, c.Request.Body, MAX_UPLOAD_SIZE)
	// c.Next()
	// file := incomingfile
	// fn := id
	// // The file cannot be received.
	// if err != nil {		
	// 	return err
	// }

	// Retrieve file information
	// newFileName := filepath.Ext(file.Filename)

	// newFile := id + newFileName



	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(id)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	if err != nil {
		return err
	}
	return nil







	// The file is received, so let's save it
	// if err := c.SaveUploadedFile(incomingfile, "C:\\Users\\Administrator\\Documents\\New folder (3)\\abenezer-dev-prep\\temp\\"+newFile); err != nil {		
	// 	return err
	// }

	// File saved successfully. Return proper result
	
}

