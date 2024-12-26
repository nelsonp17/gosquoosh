package gosquoosh

import (
	"github.com/nelsonp17/gosquoosh/appwrite"
	"github.com/nelsonp17/gosquoosh/config"
)

type Dir struct {
	INPUT_PATH  string `json:"input_path"`
	OUTPUT_PATH string `json:"output_path"`
}
type AppWrite struct {
	ENDPOINT   string `json:"endpoint"`
	API_SECRET string `json:"api_secret"`
}
type ServerConfig struct {
	PORT            string `json:"port"`
	HOST            string `json:"host"`
	ENDPOINT_EXPONE string `json:"endpoint_expone"`
}

type ImageConvertConfig struct {
	API_KEY  string       `json:"api_key"`
	Dir      Dir          `json:"dir"`
	AppWrite AppWrite     `json:"appwrite"`
	Server   ServerConfig `json:"server"`
}

func ImageConvert(imageConvertConfig ImageConvertConfig) {
	if imageConvertConfig.API_KEY != "" {
		//fmt.Println("API_KEY", imageConvertConfig.API_KEY)
		config.API_KEY_MICROSERVICE = imageConvertConfig.API_KEY
	}
	if imageConvertConfig.Dir.INPUT_PATH != "" {
		//fmt.Println("INPUT_PATH", imageConvertConfig.Dir.INPUT_PATH)
		config.INPUT_PATH_IMAGEN = imageConvertConfig.Dir.INPUT_PATH
	}
	if imageConvertConfig.Dir.OUTPUT_PATH != "" {
		//fmt.Println("OUTPUT_PATH", imageConvertConfig.Dir.OUTPUT_PATH)
		config.OUTPUT_PATH_IMAGEN = imageConvertConfig.Dir.OUTPUT_PATH
	}

	appwrite.Set(imageConvertConfig.AppWrite.ENDPOINT, imageConvertConfig.AppWrite.API_SECRET)

	if imageConvertConfig.Server.HOST != "" && imageConvertConfig.Server.PORT != "" && imageConvertConfig.Server.ENDPOINT_EXPONE != "" {
		server(imageConvertConfig.Server)
	} else {
		server(ServerConfig{
			PORT:            "3000",
			HOST:            "localhost",
			ENDPOINT_EXPONE: "/api/v1/image-convert",
		})
	}
}
