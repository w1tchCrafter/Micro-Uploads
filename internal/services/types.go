package services

type FS struct {
	UploadPath string
}

type Auth struct{}
type Front struct{}

// this struct formats the uploaded files data to the user
type FileResponseData struct {
	Filename string
	StrSize  string
	Link     string
}
