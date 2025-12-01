package common

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
)

type Uploader interface {
	UploadImage(file *multipart.FileHeader, path string) (string, error)
	UploadFile(file *multipart.FileHeader, path string) (string, error)
}

type uploader struct {
}

var FileUploader Uploader

func NewFileUploader() Uploader {
	FileUploader = &uploader{}
	return FileUploader
}

// UPloadFile implements Uploader.
func (ic *uploader) UploadFile(file *multipart.FileHeader, path string) (string, error) {
	Logger.LogError().Interface("UploadFile(", file).Msg("Uploader")

	src, err := file.Open()
	if err != nil {
		return "", err
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			return "", err
		}
	}
	defer src.Close()

	timeNow := time.Now().Format("2006-0102T15_0405_070000")
	fileName := timeNow + uuid.New().String() + filepath.Ext(file.Filename)
	fullPath := filepath.Join(path, fileName)

	// destination file
	dst, err := os.Create(fullPath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	// copy file to destination
	if _, err := io.Copy(dst, src); err != nil {
		return "", err
	}
	return fullPath, nil
}


func (ic *uploader) UploadImage(file *multipart.FileHeader, path string) (string, error) {
	Logger.LogInfo().Interface("UploadFile(", file).Msg("Uploader")

	// source
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			return "", err
		}
	}
	defer src.Close()

	timeNow := time.Now().Format("2006-0102T15_0405_070000")
	fileName := timeNow + uuid.New().String() + filepath.Ext(file.Filename)
	fullPath := filepath.Join(path, fileName)

	// destination file
	dst, err := os.Create(fullPath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	// copy file to destination
	if _, err := io.Copy(dst, src); err != nil {
		return "", err
	}
	return fullPath, nil
}