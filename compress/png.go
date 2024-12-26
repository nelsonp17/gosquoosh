package compress

import (
	"bytes"
	"github.com/nelsonp17/gosquoosh/lib"
	"image"
	"image/png"
	"os"

	"github.com/kolesa-team/go-webp/decoder"
	"github.com/kolesa-team/go-webp/encoder"
	"github.com/kolesa-team/go-webp/webp"
)

// CompressPNG compresses a PNG image and saves it to the specified output path.
func CompressPNG(inputPath, outputPath string, quality int) (lib.ImageStruct, error) {
	// Open the input file
	r := lib.ImageStruct{}
	inputFile, err := os.Open(inputPath)
	if err != nil {
		return r, err
	}
	defer inputFile.Close()

	// Decode the image
	img, _, err := image.Decode(inputFile)
	if err != nil {
		return r, err
	}

	// Create a buffer to hold the compressed image
	var buffer bytes.Buffer

	// Encode the image to the buffer with the specified quality
	e := png.Encoder{CompressionLevel: png.CompressionLevel(quality)}
	if err := e.Encode(&buffer, img); err != nil {
		return r, err
	}

	// Write the compressed image to the output file
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return r, err
	}
	defer outputFile.Close()

	if _, err := buffer.WriteTo(outputFile); err != nil {
		return r, err
	}

	// Get file info for size
	fileInfo, err := outputFile.Stat()
	if err != nil {
		return r, err
	}
	r.Size = fileInfo.Size()
	r.SizeKb = lib.BytesToKB(fileInfo.Size())

	return r, nil
}

// PngToWebP encodes a PNG image to WebP format.
func PngToWebP(inputPath string, outputPath string, quality float32) (lib.ImageStruct, error) {
	// Open the input file
	r := lib.ImageStruct{}
	inputFile, err := os.Open(inputPath)
	if err != nil {
		return r, err
	}
	defer inputFile.Close()

	// Decode the PNG image
	img, err := png.Decode(inputFile)
	if err != nil {
		return r, err
	}

	// Create the output file
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return r, err
	}
	defer outputFile.Close()

	// Encode the image to WebP format
	options, err := encoder.NewLossyEncoderOptions(encoder.PresetDefault, quality)
	if err != nil {
		return r, err
	}
	if err := webp.Encode(outputFile, img, options); err != nil {
		return r, err
	}

	// Get file info for size
	fileInfo, err := outputFile.Stat()
	if err != nil {
		return r, err
	}
	r.Size = fileInfo.Size()
	r.SizeKb = lib.BytesToKB(fileInfo.Size())

	return r, nil
}

// WebPToPng decodes a WebP image to PNG format.
func WebPToPng(inputPath string, outputPath string) (lib.ImageStruct, error) {
	// Open the input file
	r := lib.ImageStruct{}
	inputFile, err := os.Open(inputPath)
	if err != nil {
		return r, err
	}
	defer inputFile.Close()

	// Decode the WebP image
	img, err := webp.Decode(inputFile, &decoder.Options{})
	if err != nil {
		return r, err
	}

	// Create the output file
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return r, err
	}
	defer outputFile.Close()

	// Encode the image to PNG format
	if err = png.Encode(outputFile, img); err != nil {
		return r, err
	}

	// Get file info for size
	fileInfo, err := outputFile.Stat()
	if err != nil {
		return r, err
	}
	r.Size = fileInfo.Size()
	r.SizeKb = lib.BytesToKB(fileInfo.Size())

	return r, nil
}
