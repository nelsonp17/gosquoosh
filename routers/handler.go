package routers

import (
	"errors"
	"fmt"
	"github.com/appwrite/sdk-for-go/file"
	"github.com/appwrite/sdk-for-go/id"
	"github.com/nelsonp17/gosquoosh/appwrite"
	"github.com/nelsonp17/gosquoosh/compress"
	"github.com/nelsonp17/gosquoosh/config"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/nelsonp17/gofiber_request"
	"github.com/nelsonp17/gosquoosh/lib"
	"net/http"
	"path/filepath"
)

func Init(image fiber.Router) {
	// limpia la cache
	ClearCache()

	// crea las carpetas de destino
	err := lib.CreateFolder(config.OUTPUT_PATH_IMAGEN)
	if err != nil {
		fmt.Println("Error al crear la carpeta de destino")
	}
	err = lib.CreateFolder(config.INPUT_PATH_IMAGEN)
	if err != nil {
		fmt.Println("Error al crear la carpeta de destino")
	}

	// rutas de la API
	ImageCompress(image)
}

func ClearCache() error {
	err := lib.ClearFolder(config.OUTPUT_PATH_IMAGEN)
	if err != nil {
		return errors.New("Error al limpiar la carpeta de destino")
	}
	err = lib.ClearFolder(config.INPUT_PATH_IMAGEN)
	if err != nil {
		return errors.New("Error al limpiar la carpeta de destino")
	}
	return nil
}

func ClearCacheWebp(c *fiber.Ctx) error {
	if ClearCache() != nil {
		return c.Status(500).JSON(config.Response{Error: "Error al limpiar la carpeta de destino"})
	}
	return c.JSON(config.Response{Ok: "Cache limpiada correctamente"})
}

func GetImage(c *fiber.Ctx, mimeType string) (lib.ImageStruct, error) {
	image := lib.ImageStruct{}
	formFile, err := c.FormFile("image")
	if err != nil {
		return image, c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to read image"})
	}
	src, err := formFile.Open()
	if err != nil {
		return image, c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to open image"})
	}
	defer src.Close()

	buffer := make([]byte, 512)
	_, err = src.Read(buffer)
	if err != nil {
		return image, c.Status(500).JSON(fiber.Map{"error": "Failed to read image"})
	}

	contentType := http.DetectContentType(buffer)
	if contentType == "image/jpeg" || contentType == "image/jpg" {
		if mimeType != "image/jpeg" && mimeType != "image/jpg" {
			return image, errors.New("Invalid image format 1")
		}
	} else if contentType != mimeType {
		return image, errors.New("Invalid image format")
	}

	filename := lib.ReplaceSpaces(formFile.Filename)
	filenamePath := filepath.Join(config.INPUT_PATH_IMAGEN, filename)
	_, err = lib.UploadFile(c, formFile, filenamePath)
	if err != nil {
		return image, err
	}

	urlOutput, _ := c.GetRouteURL("get_image", fiber.Map{"filename": filename, "path": "input"})
	urlDownload, _ := c.GetRouteURL("download_image", fiber.Map{"filename": filename, "path": "input"})
	return lib.ImageStruct{
		Url:      c.BaseURL() + urlOutput,
		Download: c.BaseURL() + urlDownload,
		Size:     formFile.Size,
		SizeKb:   lib.BytesToKB(formFile.Size),
		Mimetype: contentType,
		Filename: filename,
	}, nil
}

func validateImageFormat(valor string) error {
	validFormats := []string{"image/jpeg", "image/jpg", "image/png", "image/webp"}
	for _, validFormat := range validFormats {
		if valor == validFormat {
			return nil
		}
	}
	return errors.New("Invalid format")
}

func ProcessImage(c *fiber.Ctx, opc string) (config.Response, int) {
	Request := request.Request{Fields: map[string]string{
		"quality": "required|integer",
		"to":      "required", // image/jpeg, image/jpg, image/png, image/webp
		"from":    "required", // image/jpeg, image/jpg, image/png, image/webp
	}}
	err := Request.Start(c)
	if err != nil {
		return config.Response{Error: err.Error()}, 400
	}

	// Validate the request
	if !Request.Validated() {
		return config.Response{Error: Request.Errors}, 400
	}
	if Request.GetString("quality") == "" {
		Request.Form["quality"] = "80"
	}
	err = validateImageFormat(Request.GetString("to"))
	if err != nil {
		return config.Response{Error: err.Error()}, 400
	}
	err = validateImageFormat(Request.GetString("from"))
	if err != nil {
		return config.Response{Error: err.Error()}, 400
	}

	// obtener la imagen del request
	mimeTypeFrom := Request.GetString("from")
	mimeTypeTo := Request.GetString("to")
	imageOriginal, err := GetImage(c, mimeTypeFrom)
	if err != nil {
		return config.Response{Error: err.Error()}, 500
	}

	dirInput := filepath.Join(config.INPUT_PATH_IMAGEN, imageOriginal.Filename)
	filenameOutput := lib.ReplaceMimeType(imageOriginal.Filename, mimeTypeFrom, mimeTypeTo)
	urlOutput, _ := c.GetRouteURL("get_image", fiber.Map{"filename": filenameOutput, "path": "output"})
	dirOutput := filepath.Join(config.OUTPUT_PATH_IMAGEN, filenameOutput)

	imageProcess := lib.ImageStruct{}

	if opc == "convert" {
		if mimeTypeFrom == "image/webp" && mimeTypeTo == "image/png" {
			imageProcess, err = compress.WebPToPng(dirInput, dirOutput)
		}
		if mimeTypeFrom == "image/webp" && (mimeTypeTo == "image/jpeg" || mimeTypeTo == "image/jpg") {
			quality := Request.GetInt("quality")
			imageProcess, err = compress.WebPToJpeg(dirInput, dirOutput, quality)
		}
		if mimeTypeFrom == "image/png" && mimeTypeTo == "image/webp" {
			qualityFloat32 := Request.GetFloat32("quality")
			imageProcess, err = compress.PngToWebP(dirInput, dirOutput, qualityFloat32)
		}
		if (mimeTypeFrom == "image/jpg" || mimeTypeFrom == "image/jpeg") && mimeTypeTo == "image/webp" {
			qualityFloat32 := Request.GetFloat32("quality")
			imageProcess, err = compress.JpegToWebP(dirInput, dirOutput, qualityFloat32)
		}
	} else if opc == "compress" {
		if mimeTypeFrom == "image/png" && mimeTypeTo == "image/png" {
			quality := Request.GetInt("quality")
			imageProcess, err = compress.CompressPNG(dirInput, dirOutput, quality)
		}
	}

	if err != nil {
		return config.Response{Error: err.Error()}, 500
	}
	imageProcess.Filename = filenameOutput
	imageProcess.Url = c.BaseURL() + urlOutput
	urlDownload, _ := c.GetRouteURL("download_image", fiber.Map{"filename": filenameOutput, "path": "output"})
	imageProcess.Download = c.BaseURL() + urlDownload
	imageProcess.Mimetype = mimeTypeTo

	if imageProcess.Filename == "" {
		return config.Response{Error: "Error al procesar la imagen"}, 500
	}

	if opc == "compress" {
		return config.Response{
			Data: lib.ResponseCompress{
				Original:   imageOriginal,
				Compressed: imageProcess,
			},
		}, 200
	}

	return config.Response{
		Data: lib.ResponseConvert{
			Original:  imageOriginal,
			Converted: imageProcess,
		},
	}, 200
}

func CreateFileAppWrite(c *fiber.Ctx, filename string) (config.Response, int) {
	Request := request.Request{
		Fields: map[string]string{
			"project_id": "required",
			"bucket":     "required",
			"dir":        "required",
			"id":         "optional",
			"endpoint":   "optional",
			"api_secret": "optional",
		},
	}
	err := Request.Start(c)
	if err != nil {
		return config.Response{Error: err.Error()}, 400
	}

	// Validate the request
	if !Request.Validated() {
		return config.Response{Error: Request.Errors}, 400
	}
	if Request.Form["api_secret"] != "" && Request.Form["endpoint"] == "" {
		return config.Response{Error: "Invalid endpoint"}, 400
	}
	if Request.Form["api_secret"] == "" && Request.Form["endpoint"] != "" {
		return config.Response{Error: "Invalid api secret"}, 400
	}

	// inicio el cliente de appwrite
	projectId := Request.GetString("project_id")
	apiSecret := Request.GetString("api_secret")
	endpoint := Request.GetString("endpoint")
	bucket := Request.GetString("bucket")
	_id := Request.GetString("id")
	dir := Request.GetString("dir")

	service := appwrite.NewInstanceService(endpoint, projectId, apiSecret)

	fileId := _id
	if fileId == "" {
		fileId = id.Unique()
	}

	var path string
	if dir == "compress" {
		path = filepath.Join(config.OUTPUT_PATH_IMAGEN, filename)
	} else if dir == "convert" {
		path = filepath.Join(config.OUTPUT_PATH_IMAGEN, filename)
	}

	if path == "" {
		return config.Response{Error: "Invalid path"}, 400
	}
	response, err := service.CreateFile(
		bucket,
		fileId,
		file.NewInputFile(path, filename),
		service.WithCreateFilePermissions([]string{"read(\"any\")"}), // Optional
	)

	if err != nil {
		fmt.Println("Error al crear el archivo", err)
		return config.Response{Error: err.Error()}, 500
	}

	err = lib.DeleteFile(path)
	if err != nil {
		fmt.Println("Error al eliminar el archivo", err)
	}
	url := appwrite.GetPathFileFull(endpoint, projectId, bucket, fileId)

	responseAppWrite := lib.ResponseAppWrite{
		Preview: url,
		File:    response,
	}
	return config.Response{Data: responseAppWrite}, 200
}

func CompressImage(c *fiber.Ctx) error {
	process, code := ProcessImage(c, "compress")
	return c.Status(code).JSON(process)
}

func ConvertFormat(c *fiber.Ctx) error {
	process, code := ProcessImage(c, "convert")
	return c.Status(code).JSON(process)
}

func UploadAppWrite(c *fiber.Ctx) error {
	Request := request.Request{
		Fields: map[string]string{
			"filename": "required",
		},
	}
	err := Request.Start(c)
	if err != nil {
		return c.Status(400).JSON(config.Response{Error: err.Error()})
	}
	if !Request.Validated() {
		return c.Status(400).JSON(Request.Errors)
	}
	upload, codeUpload := CreateFileAppWrite(c, Request.GetString("filename"))
	if codeUpload != 200 {
		return c.Status(codeUpload).JSON(upload)
	}
	return c.Status(codeUpload).JSON(
		config.Response{
			Data: upload.Data,
		})
}

func UploadAndConvertLocal(c *fiber.Ctx) error {
	process, code := ProcessImage(c, "convert")
	if code != 200 {
		return c.Status(code).JSON(process)
	}
	filename := process.Data.(lib.ResponseConvert).Converted.Filename
	upload, codeUpload := CreateFileAppWrite(c, filename)
	if codeUpload != 200 {
		return c.Status(codeUpload).JSON(upload)
	}
	err := lib.DeleteFile(filepath.Join(config.INPUT_PATH_IMAGEN, process.Data.(lib.ResponseConvert).Original.Filename))
	if err != nil {
		fmt.Println("Error al eliminar el archivo:", err)
	}
	process.Data = lib.ResponseConvert{
		Original: lib.ImageStruct{
			Url:      "",
			Download: "",
			Size:     process.Data.(lib.ResponseConvert).Original.Size,
			SizeKb:   process.Data.(lib.ResponseConvert).Original.SizeKb,
			Mimetype: process.Data.(lib.ResponseConvert).Original.Mimetype,
			Filename: process.Data.(lib.ResponseConvert).Original.Filename,
		},
		Converted: lib.ImageStruct{
			Url:      upload.Data.(lib.ResponseAppWrite).Preview,
			Download: "",
			Size:     process.Data.(lib.ResponseConvert).Converted.Size,
			SizeKb:   process.Data.(lib.ResponseConvert).Converted.SizeKb,
			Mimetype: process.Data.(lib.ResponseConvert).Converted.Mimetype,
			Filename: process.Data.(lib.ResponseConvert).Converted.Filename,
		},
	}
	return c.Status(codeUpload).JSON(
		config.Response{
			Data: lib.ResponseAppWriteLocal{
				AppWrite:  upload.Data.(lib.ResponseAppWrite),
				Converted: process.Data.(lib.ResponseConvert).Converted,
				Original:  process.Data.(lib.ResponseConvert).Original,
			},
		})
}

func GetUrl(c *fiber.Ctx) error {
	filename := c.Params("filename")
	path := c.Params("path")
	if filename == "" {
		return c.Status(400).JSON(config.Response{Error: "Invalid filename"})
	}
	if path == "" || (path != "input" && path != "output") {
		return c.Status(400).JSON(config.Response{Error: "Invalid path"})
	}

	if path == "input" {
		return c.SendFile(filepath.Join(config.INPUT_PATH_IMAGEN, filename))
	}

	return c.SendFile(filepath.Join(config.OUTPUT_PATH_IMAGEN, filename))
}

func DownloadImage(c *fiber.Ctx) error {
	filename := c.Params("filename")
	path := c.Params("path")
	if filename == "" {
		return c.Status(400).JSON(config.Response{Error: "Invalid filename"})
	}
	if path == "" || (path != "input" && path != "output") {
		return c.Status(400).JSON(config.Response{Error: "Invalid path"})
	}

	filePath := ""
	if path == "input" {
		filePath = filepath.Join(config.INPUT_PATH_IMAGEN, filename)
	} else {
		filePath = filepath.Join(config.OUTPUT_PATH_IMAGEN, filename)
	}

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return c.Status(404).JSON(config.Response{Error: "File not found"})
	}

	c.Set("Content-Disposition", "attachment; filename="+filename)
	return c.SendFile(filePath)
}
