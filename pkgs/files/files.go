package files

import (
	"io"
	"os"
	"log"
	"net/http"
	"mime/multipart"
	"path/filepath"
	"errors"
)

func SaveFile(path string, file multipart.File, handler *multipart.FileHeader) error {
	f, err := os.OpenFile(path+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		return err
	}

	io.Copy(f, file)
	
	return nil
}

func DeleteFile(filepath string) error {
	return os.Remove(filepath)
}

func SaveUploadedCsv(r *http.Request) (string, error) {
	//parsing as multipart/form-data
	r.ParseMultipartForm(32 << 20)

	file, handler, err := r.FormFile("fileupload")

	if err != nil {
		return "", err
	}

	defer file.Close()

	var extension = filepath.Ext(handler.Filename)

	if extension != ".csv" {
		return handler.Filename, errors.New("Invalid extension. File must be a CSV file, ending with .csv") 
	}

	erro := SaveFile("upload-files/", file, handler)

	if erro != nil {
		log.Fatal(erro)
		return "", nil
	}
	return handler.Filename, nil
}