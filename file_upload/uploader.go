package main

import (
	"fmt"
	"io"
	"log"
	"mime"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func main() {
	http.Handle("/avatar", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		err := request.ParseMultipartForm(10 * 1024 * 1024)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		var avatarFileHeader *multipart.FileHeader
		if avatars := request.MultipartForm.File["avatar"]; len(avatars) > 0 {
			avatarFileHeader = avatars[0]
		}
		if avatarFileHeader == nil {
			http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		ext := filepath.Ext(avatarFileHeader.Filename)
		if len(ext) == 0 {
			http.Error(writer, "file type not supported", http.StatusBadRequest)
			return
		}

		imageType := mime.TypeByExtension(ext)
		if len(imageType) == 0 {
			http.Error(writer, "file type not supported", http.StatusBadRequest)
			return
		}

		switch imageType {
		case "image/jpeg":
		case "image/png":
		default:
			http.Error(writer, "file type not supported", http.StatusBadRequest)
			return
		}

		avatarFile, err := avatarFileHeader.Open()
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		defer avatarFile.Close()

		destFilename := fmt.Sprintf("./avatars/avatar_%s%s", time.Now().Format("20060102150405"), ext)
		destFile, err := os.Create(destFilename)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		defer destFile.Close()

		_, err = io.Copy(destFile, avatarFile)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		writer.WriteHeader(http.StatusOK)
	}))
	log.Fatal(http.ListenAndServe(":12345", nil))
}
