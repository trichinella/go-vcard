package vcard

import (
	"encoding/base64"
	"net/http"
	"os"
)

type Photo interface {
	String() string
}

type FilePhoto struct {
	content []byte
}

func NewFilePhoto(path string) (*FilePhoto, error) {
	content, err := getContentFile(path)
	if err != nil {
		return nil, err
	}

	return &FilePhoto{
		content: content,
	}, nil
}

func (p *FilePhoto) String() string {
	return `PHOTO;` + "ENCODING=b;TYPE=" + getImageType(p.content) + ":" + encodingPhoto(p.content) + "\n"
}

func getContentFile(path string) ([]byte, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return content, err
}

func getImageType(content []byte) string {
	switch http.DetectContentType(content) {
	case "image/png":
		return "PNG"
	case "image/jpeg":
		return "JPEG"
	default:
		return "JPEG"
	}
}

func encodingPhoto(content []byte) string {
	return base64.StdEncoding.EncodeToString(content)
}
