package image_manipulation

import (
	"errors"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"mime"
	"os"
	"path/filepath"
)

const (
	GrayscaleAverage      = 0
	GrayscaleLuma         = 1
	GrayscaleDesaturation = 2
)

var ErrUnknownImageType = errors.New("unknown image type")

type Pixel struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

type Image struct {
	Pixels [][]Pixel
	Width  uint32
	Height uint32
	_Rect  image.Rectangle
	_Image image.Image
}

func NewImage(path string) (*Image, error) {
	ext := filepath.Ext(path)
	if len(ext) == 0 {
		return nil, ErrUnknownImageType
	}
	mimeType := mime.TypeByExtension(ext)
	if len(mimeType) == 0 {
		return nil, ErrUnknownImageType
	}
	switch mimeType {
	case "image/jpeg":
		image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
	case "image/png":
		image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
	default:
		return nil, ErrUnknownImageType
	}

	fileReader, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer fileReader.Close()

	img, _, err := image.Decode(fileReader)
	if err != nil {
		return nil, err
	}

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	var pixels [][]Pixel
	for y := 0; y < height; y++ {
		var row []Pixel
		for x := 0; x < width; x++ {
			pixel := colorToPixel(img.At(x, y))
			row = append(row, pixel)
		}
		pixels = append(pixels, row)
	}

	return &Image{
		Pixels: pixels,
		Width:  uint32(width),
		Height: uint32(height),
		_Rect:  bounds,
		_Image: img,
	}, nil
}

func colorToPixel(c color.Color) Pixel {
	r, g, b, a := c.RGBA()
	return Pixel{
		R: uint8(r >> 8),
		G: uint8(g >> 8),
		B: uint8(b >> 8),
		A: uint8(a >> 8),
	}
}

func (img *Image) WriteToFile(path string) error {
	ext := filepath.Ext(path)
	if len(ext) == 0 {
		return ErrUnknownImageType
	}
	mimeType := mime.TypeByExtension(ext)
	if len(mimeType) == 0 {
		return ErrUnknownImageType
	}
	switch mimeType {
	case "image/jpeg":
	case "image/png":
	default:
		return ErrUnknownImageType
	}

	nImg := image.NewRGBA(img._Rect)
	draw.Draw(nImg, img._Rect, img._Image, image.Point{}, draw.Over)
	for y := 0; y < int(img.Height); y++ {
		for x := 0; x < int(img.Width); x++ {
			pixel := img.Pixels[y][x]
			nImg.Set(x, y, color.RGBA{
				R: pixel.R,
				G: pixel.G,
				B: pixel.B,
				A: pixel.A,
			})
		}
	}

	fileWriter, err := os.Create(path)
	if err != nil {
		return err
	}
	defer fileWriter.Close()

	switch mimeType {
	case "image/jpeg":
		jpeg.Encode(fileWriter, nImg, nil)
	case "image/png":
		png.Encode(fileWriter, nImg)
	default:
		panic(errors.New("this shouldn't happen"))
	}
	return nil
}

func (img *Image) Grayscale(algorithm int) *Image {
	algorithmFun := grayscaleAlgorithm(algorithm)
	for y := 0; y < int(img.Height); y++ {
		for x := 0; x < int(img.Width); x++ {
			pixel := img.Pixels[y][x]
			gray := algorithmFun(pixel)
			pixel.R = gray
			pixel.G = gray
			pixel.B = gray
			img.Pixels[y][x] = pixel
		}
	}
	return img
}

func grayscaleAlgorithm(algorithm int) func(pixel Pixel) uint8 {
	switch algorithm {
	case GrayscaleLuma:
		return func(pixel Pixel) uint8 {
			return uint8(float32(pixel.R)*0.2126 + float32(pixel.G)*0.7152 + float32(pixel.B)*0.0722)
		}
	case GrayscaleDesaturation:
		return func(pixel Pixel) uint8 {
			return (Max(pixel.R, pixel.G, pixel.B) + Min(pixel.R, pixel.G, pixel.B)) / 2
		}
	case GrayscaleAverage:
		fallthrough
	default:
		return func(pixel Pixel) uint8 {
			return (pixel.R + pixel.G + pixel.B) / 3
		}
	}
}
