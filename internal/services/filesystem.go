package services

import (
	"compress/gzip"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

func NewFileSystem(uploadPath string) FS {
	return FS{UploadPath: uploadPath}
}

func (fs FS) Save(fileData *multipart.FileHeader) (string, error) {
	uniqFilename := uuid.New().String() + filepath.Ext(fileData.Filename)
	path := filepath.Join(fs.UploadPath, uniqFilename)
	src, err := fileData.Open()

	if err != nil {
		return "", err
	}

	defer src.Close()
	err = fs.compressFile(path, src)
	return path, err
}

func (fs FS) Open(path string) (*gzip.Reader, error) {
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

func (fs FS) deCompressFile(path string) (*gzip.Reader, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	return gzip.NewReader(file)
}
