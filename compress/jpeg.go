package compress

import (
	"github.com/nelsonp17/gosquoosh/lib"
	"image/jpeg"
	"os"

	"github.com/kolesa-team/go-webp/decoder"
	"github.com/kolesa-team/go-webp/encoder"
	"github.com/kolesa-team/go-webp/webp"
)

// EncodeToWebP encodes a JPEG image to WebP format.
func JpegToWebP(inputPath string, outputPath string, quality float32) (lib.ImageStruct, error) {
	// Open the input file
	r := lib.ImageStruct{}
	inputFile, err := os.Open(inputPath)
	if err != nil {
		return r, err
	}
	defer inputFile.Close()

	// Decode the JPEG image
	img, err := jpeg.Decode(inputFile)
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

// DecodeFromWebP decodes a WebP image to JPEG format.
func WebPToJpeg(inputPath string, outputPath string, quality int) (lib.ImageStruct, error) {
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

	// Encode the image to JPEG format
	if err = jpeg.Encode(outputFile, img, &jpeg.Options{Quality: quality}); err != nil {
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
