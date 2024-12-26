package compress

import (
	"testing"
)

func TestJpegToWebP(t *testing.T) {
	JpegToWebP("./input/example.jpg", "./output/example_jpg.webp", 75)
}

func TestWebPToJpeg(t *testing.T) {
	WebPToJpeg("./input/example.webp", "./output/example_wepb.jpg", 75)
}

func TestCompressPNG(t *testing.T) {
	CompressPNG("./input/example.png", "./output/example_png_compress.png", 75)
}

func TestPngToWebP(t *testing.T) {
	PngToWebP("./input/example.png", "./output/example_png.webp", 75)
}

func TestWebPToPng(t *testing.T) {
	WebPToPng("./input/example.webp", "./output/example_webp.png")
}
