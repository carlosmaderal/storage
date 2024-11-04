package routes

import (
	"bytes"
	"image"
	"image/jpeg"
	"github.com/nfnt/resize"
)

// Função para redimensionar a imagem usando o conteúdo em bytes e retornando o novo conteúdo
func resizeImage(content []byte, targetWidth, targetHeight uint) ([]byte, error) {
	// Decodifica a imagem a partir dos bytes
	img, _, err := image.Decode(bytes.NewReader(content))
	if err != nil {
		return nil, err
	}

	origWidth := uint(img.Bounds().Dx())
	origHeight := uint(img.Bounds().Dy())

	var newWidth, newHeight uint

	if targetWidth > 0 && targetHeight > 0 {
		newWidth = targetWidth
		newHeight = targetHeight
	} else if targetWidth > 0 {
		newWidth = targetWidth
		newHeight = uint(float64(origHeight) * (float64(newWidth) / float64(origWidth)))
	} else if targetHeight > 0 {
		newHeight = targetHeight
		newWidth = uint(float64(origWidth) * (float64(newHeight) / float64(origHeight)))
	} else {
		return content, nil
	}

	// Redimensiona a imagem
	newImage := resize.Resize(newWidth, newHeight, img, resize.Lanczos3)

	var buf bytes.Buffer
	err = jpeg.Encode(&buf, newImage, nil)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

