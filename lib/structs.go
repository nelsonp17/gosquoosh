package lib

import "github.com/appwrite/sdk-for-go/models"

type ImageStruct struct {
	Url      string  `json:"url"`
	Download string  `json:"download"`
	Size     int64   `json:"size"`
	SizeKb   float64 `json:"size_kb"`
	Width    int     `json:"width"`
	Height   int     `json:"height"`
	Mimetype string  `json:"mimetype"`
	Filename string  `json:"filename"`
}
type ResponseCompress struct {
	Original   ImageStruct `json:"original"`
	Compressed ImageStruct `json:"compressed"`
}
type ResponseConvert struct {
	Original  ImageStruct `json:"original"`
	Converted ImageStruct `json:"converted"`
}
type ResponseAppWrite struct {
	File    *models.File `json:"file"`
	Preview string       `json:"preview"`
}
type ResponseAppWriteLocal struct {
	AppWrite  ResponseAppWrite `json:"appwrite"`
	Converted ImageStruct      `json:"converted"`
	Original  ImageStruct      `json:"original"`
}
