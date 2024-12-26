package routers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// ImageCompress microservice de compresión de imágenes y conversión de formatos, estas rutas estan protegidas por un API Key
func ImageCompress(image fiber.Router) {
	compress := image.Group("/compress", APIKeyAuth)  // <- /api/v1/image/compress
	convert := image.Group("/convert", APIKeyAuth)    // <- /api/v1/image/convert
	cache := image.Group("/cache", APIKeyAuth)        // <- /api/v1/image/cache
	upload := image.Group("/upload", APIKeyAuth)      // <- /api/v1/image/upload
	appwrite := upload.Group("/appwrite", APIKeyAuth) // <- /api/v1/image/upload/appwrite

	cache.Add(http.MethodPost, "/clear", ClearCacheWebp)

	appwrite.Add(http.MethodPost, "/", UploadAppWrite)
	appwrite.Add(http.MethodPost, "/convert", UploadAndConvertLocal)

	compress.Add(http.MethodPost, "/", CompressImage)
	convert.Add(http.MethodPost, "/", ConvertFormat) // <- /api/v1/image/convert/webp

	image.Get("/visor/:path/:filename", GetUrl).Name("get_image")
	image.Get("/download/:path/:filename", DownloadImage).Name("download_image")

}
