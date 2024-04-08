package services

import (
	"fmt"
	"micro_uploads/internal/models"
	"path/filepath"
	"strings"
)

func (f Front) IsLogged(username string) bool {
	return username != ""
}

func (f Front) SetData(filesData ...models.FileModel) []FileResponseData {
	newData := make([]FileResponseData, 0)
	var filename, size string
	const KB = 1 << 10
	const MB = 1 << 20

	for _, v := range filesData {
		if len(v.OriginalName) >= 12 {
			ext := filepath.Ext(v.OriginalName)
			filename = v.OriginalName[0:12] + "... " + ext
		} else {
			filename = v.OriginalName
		}

		switch {
		case v.Size >= MB:
			size = fmt.Sprint(int(float64(v.Size))/MB, "MB")
		case v.Size >= KB:
			size = fmt.Sprint(int(float64(v.Size))/KB, "KB")
		default:
			size = fmt.Sprint(v.Size, "bytes")
		}

		newData = append(newData, FileResponseData{
			Filename: filename,
			StrSize:  size,
			Link:     "/api/v1/uploads/" + strings.Split(v.Filename, "/")[1],
		})
	}

	return newData
}
