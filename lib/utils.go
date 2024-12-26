package lib

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// CreateFolder A new folder is created at the root of the project.
func CreateFolder(dirname string) error {
	_, err := os.Stat(dirname)
	if os.IsNotExist(err) {
		errDir := os.MkdirAll(dirname, 0755)
		if errDir != nil {
			return errors.New("Error al crear la carpeta de destino")
		}
	}
	return nil
}

// ClearFolder deletes all contents of the specified directory
func ClearFolder(dirname string) error {
	files, err := os.ReadDir(dirname)
	if err != nil {
		return err
	}

	for _, file := range files {
		err = os.RemoveAll(filepath.Join(dirname, file.Name()))
		if err != nil {
			return err
		}
	}

	return nil
}

func GenerateAPIKey() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func UploadFile(c *fiber.Ctx, file *multipart.FileHeader, filename string) (string, error) {
	// 2. Crear la carpeta de destino si no existe
	//uploadPath := dir                           // Ruta de la carpeta de destino (relativa al ejecutable)
	//err := os.MkdirAll(uploadPath, os.ModePerm) // Crea la carpeta y las subcarpetas necesarias
	//if err != nil {
	//	return "", errors.New("Error al crear la carpeta de destino")
	//}

	// 3. Generar un nombre de archivo único (opcional pero recomendado para evitar colisiones)
	// filename := fmt.Sprintf("%d-%s", time.Now().UnixNano(), file.Filename) // Usando timestamp
	//filename := filepath.Join(uploadPath, file.Filename) // Usando el nombre original (menos seguro)

	if err := c.SaveFile(file, filename); err != nil {
		return "", errors.New(fmt.Sprintf("Error al guardar el archivo: %v", err))
	}

	//fmt.Println("Archivo subido correctamente: %s", file.Filename)

	return filename, nil
}

// DeleteFile deletes a file given its path.
func DeleteFile(filePath string) error {
	err := os.Remove(filePath)
	if err != nil {
		return fmt.Errorf("failed to delete file: %w", err)
	}
	fmt.Println("File deleted successfully:", filePath)
	return nil
}

// image/png
func GetMimeType(fileHeader *multipart.FileHeader) (string, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Read the first 512 bytes to detect the file type
	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		return "", err
	}

	// Detect the content type
	contentType := http.DetectContentType(buffer)

	return contentType, nil
}

// ReplaceSpaces replaces spaces in the filename with underscores.
func ReplaceSpaces(filename string) string {
	return strings.ReplaceAll(filename, " ", "_")
}

// BytesToKB converts bytes to kilobytes (KB)
func BytesToKB(bytes int64) float64 {
	return float64(bytes) / 1024
}

// ReplaceToWebP replaces .jpg or .jpeg extension with .webp
func ReplaceToWebP(filename string) string {
	if strings.HasSuffix(filename, ".jpg") {
		return strings.Replace(filename, ".jpg", ".webp", 1)
	} else if strings.HasSuffix(filename, ".jpeg") {
		return strings.Replace(filename, ".jpeg", ".webp", 1)
	}
	return filename
}

func ReplaceMimeType(filename string, mimeTypeFrom string, mimeTypeTo string) string {
	mimetypeFrom := strings.Replace(mimeTypeFrom, "image/", "", 1)
	mimeTypeTo = strings.Replace(mimeTypeTo, "image/", "", 1)

	if strings.HasSuffix(filename, "."+mimetypeFrom) {
		return strings.Replace(filename, "."+mimetypeFrom, "."+mimeTypeTo, 1)
	}
	if mimeTypeTo == "webp" {
		if mimetypeFrom == "jpeg" || mimetypeFrom == "jpg" {
			return ReplaceToWebP(filename)
		}
	}

	return filename
}

func GetDigitFromRule(rule string) (int, error) {
	parts := strings.Split(rule, ":")
	if len(parts) != 2 {
		return 0, fmt.Errorf("invalid rule format")
	}
	digit, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, fmt.Errorf("invalid digit: %v", err)
	}
	return digit, nil
}
