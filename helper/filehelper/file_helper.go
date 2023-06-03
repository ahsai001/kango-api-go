package filehelper

import (
	"errors"
	"fmt"
	"image"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
)

func InitFolder(folderPath string) {
	if _, err := os.Stat(folderPath); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(folderPath, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}
}

func UploadResizeSingleFile(request *http.Request, keyName string, folderPath string) (filePath string, err error) {
	file, header, err := request.FormFile(keyName)
	if err != nil {
		log.Println(err)
		return
	}

	fileExt := filepath.Ext(header.Filename)
	originalFileName := strings.TrimSuffix(filepath.Base(header.Filename), filepath.Ext(header.Filename))
	now := time.Now()
	filename := strings.ReplaceAll(strings.ToLower(originalFileName), " ", "-") + "-" + fmt.Sprintf("%v", now.Unix()) + fileExt
	filePath = folderPath + filename

	imageFile, _, err := image.Decode(file)
	if err != nil {
		log.Println(err)
	}
	src := imaging.Resize(imageFile, 1000, 0, imaging.Lanczos)
	err = imaging.Save(src, fmt.Sprintf("%s%v", folderPath, filename))
	if err != nil {
		log.Printf("failed to save image: %v", err)
	}
	return
}

func UploadSingleFile(request *http.Request, keyName string, folderPath string) (filePath string, err error) {
	file, header, err := request.FormFile(keyName)
	if err != nil {
		log.Println(err)
		return
	}

	fileExt := filepath.Ext(header.Filename)
	originalFileName := strings.TrimSuffix(filepath.Base(header.Filename), filepath.Ext(header.Filename))
	now := time.Now()
	filename := strings.ReplaceAll(strings.ToLower(originalFileName), " ", "-") + "-" + fmt.Sprintf("%v", now.Unix()) + fileExt
	filePath = folderPath + filename

	out, err := os.Create(filePath)
	if err != nil {
		log.Println(err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Println(err)
	}
	return
}

func GetNewFileName(file *multipart.FileHeader) (fileName string) {
	fileExt := filepath.Ext(file.Filename)
	originalFileName := strings.TrimSuffix(filepath.Base(file.Filename), filepath.Ext(file.Filename))
	now := time.Now()
	fileName = strings.ReplaceAll(strings.ToLower(originalFileName), " ", "-") + "-" + fmt.Sprintf("%v", now.Unix()) + fileExt
	return
}

func UploadResizeMultipleFile(ctx *gin.Context, keyNames []string, folderPath string) (filePaths []string, errs []error) {
	form, _ := ctx.MultipartForm()
	for _, keyName := range keyNames {
		files := form.File[keyName]
		for _, file := range files {
			fileName := GetNewFileName(file)
			filePath := folderPath + fileName

			filePaths = append(filePaths, filePath)
			readerFile, _ := file.Open()
			imageFile, _, err := image.Decode(readerFile)
			if err != nil {
				log.Println(err)
			}
			src := imaging.Resize(imageFile, 1000, 0, imaging.Lanczos)
			err = imaging.Save(src, fmt.Sprintf("%s%v", folderPath, fileName))
			if err != nil {
				log.Printf("failed to save image: %v", err)
			}

			errs = append(errs, err)
		}
	}

	return
}
func UploadMultipleFile(ctx *gin.Context, keyNames []string, folderPath string) (filePaths []string, errs []error) {
	form, _ := ctx.MultipartForm()
	for _, keyName := range keyNames {
		files := form.File[keyName]
		for _, file := range files {
			fileName := GetNewFileName(file)
			filePath := folderPath + fileName

			filePaths = append(filePaths, filePath)
			out, err := os.Create(filePath)
			if err != nil {
				log.Println(err)
			}
			defer out.Close()

			readerFile, _ := file.Open()
			_, err = io.Copy(out, readerFile)
			if err != nil {
				log.Println(err)
			}

			errs = append(errs, err)

		}
	}

	return

}
