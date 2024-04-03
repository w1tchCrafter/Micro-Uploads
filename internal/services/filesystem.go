package services

import (
	"compress/gzip"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

func NewFileSystem(uploadPath string) FS {
	return FS{UploadPath: uploadPath}
}

func (fs FS) Save(fileData *multipart.FileHeader) (string, error) {
	src, err := fileData.Open()

	if err != nil {
		return "", err
	}

	defer src.Close()
	uniqFilename := uuid.New().String() + filepath.Ext(fileData.Filename)
	content_type := fileData.Header.Get("Content-Type")
	path := filepath.Join(fs.UploadPath, uniqFilename)

	if fs.IsMidia(content_type) {
		src, err := fileData.Open()

		if err != nil {
			return "", err
		}

		file, err := os.Create(path)

		if err != nil {
			return "", err
		}

		defer file.Close()
		_, err = io.Copy(file, src)
		return path, err
	}

	err = fs.compressFile(path, src)
	return path, err
}

func (fs FS) Open(path string, midia bool) (io.ReadCloser, error) {
	if midia {
		return os.Open(path)
	}

	return fs.deCompressFile(path)
}

func (fs FS) compressFile(path string, src multipart.File) error {
	out, err := os.Create(path)

	if err != nil {
		return err
	}

	gw := gzip.NewWriter(out)
	defer out.Close()
	defer gw.Close()

	_, err = io.Copy(gw, src)
	return err
}

func (fs FS) deCompressFile(path string) (io.ReadCloser, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	return gzip.NewReader(file)
}

func (fs FS) IsMidia(content_type string) bool {
	return strings.HasPrefix(content_type, "image/") ||
		strings.HasPrefix(content_type, "document/") ||
		strings.HasPrefix(content_type, "audio/") ||
		strings.HasPrefix(content_type, "video/")
}
