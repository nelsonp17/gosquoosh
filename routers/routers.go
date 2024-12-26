package routers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// ImageCompress microservice de compresión de imágenes y conversión de formatos, estas rutas estan protegidas por un API Key
func ImageCompress(image fiber.Router) {

	image.Get("/robots.txt", func(c *fiber.Ctx) error {
		return c.SendFile("./robots.txt")
	})
	// image.Static("/robots.txt", "../robots.txt")

	compress := image.Group("/compress", APIKeyAuth, NoIndexMiddleware)  // <- /api/v1/image/compress
	convert := image.Group("/convert", APIKeyAuth, NoIndexMiddleware)    // <- /api/v1/image/convert
	cache := image.Group("/cache", APIKeyAuth, NoIndexMiddleware)        // <- /api/v1/image/cache
	upload := image.Group("/upload", APIKeyAuth, NoIndexMiddleware)      // <- /api/v1/image/upload
	appwrite := upload.Group("/appwrite", APIKeyAuth, NoIndexMiddleware) // <- /api/v1/image/upload/appwrite

	cache.Add(http.MethodPost, "/clear", ClearCacheWebp)

	appwrite.Add(http.MethodPost, "/", UploadAppWrite)
	appwrite.Add(http.MethodPost, "/convert", UploadAndConvertLocal)

	compress.Add(http.MethodPost, "/", CompressImage)
	convert.Add(http.MethodPost, "/", ConvertFormat) // <- /api/v1/image/convert/webp

	image.Get("/visor/:path/:filename", GetUrl).Name("get_image")
	image.Get("/download/:path/:filename", DownloadImage).Name("download_image")

}
