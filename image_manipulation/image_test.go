package image_manipulation

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNewImage(t *testing.T) {
	image, err := NewImage("./assets/gopher.jpg")
	assert.Nil(t, err)
	assert.Equal(t, image.Width, uint32(1000))
	assert.Equal(t, image.Height, uint32(920))
	assert.Equal(t, len(image.Pixels), 920)
	assert.Equal(t, len(image.Pixels[0]), 1000)
}

func TestWriteToFile(t *testing.T) {
	image, err := NewImage("./assets/gopher.jpg")
	assert.Nil(t, err)
	err = image.WriteToFile("./assets/gopher_new.png")
	assert.Nil(t, err)
	_, err = os.Stat("./assets/gopher_new.png")
	assert.Nil(t, err)
}

func TestGrayscale(t *testing.T) {
	image, err := NewImage("./assets/gopher.jpg")
	assert.Nil(t, err)
	err = image.Grayscale(GrayscaleLuma).WriteToFile("./assets/gopher_gray.jpg")
	assert.Nil(t, err)
	_, err = os.Stat("./assets/gopher_gray.jpg")
	assert.Nil(t, err)
}
