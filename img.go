package imgediting

import (
	"bytes"
	"image"
	"image/jpeg"
	"io"
	"mime/multipart"
	"path/filepath"

	"github.com/ilyasiv2003/imgediting/code.google.com/p/graphics-go/graphics"
	"github.com/pkg/errors"
)

// Returns buffer with bytes of image and image itself. They both are converted to jpg
func MakeJpgImage(file multipart.File, header *multipart.FileHeader, img image.Image) (*bytes.Buffer, image.Image, error) {
	jpgImgBuf := bytes.NewBuffer(nil)
	if _, err := io.Copy(jpgImgBuf, file); err != nil {
		return nil, nil, errors.Wrap(err, "failed to copy file bytes")
	}

	extension := filepath.Ext(header.Filename)
	if extension == ".png" {
		jpegImgRaw, err := convertToJPG(img)
		if err != nil {
			return nil, nil, errors.Wrap(err, "failed to convert to jpg")
		}

		jpgImgBuf = bytes.NewBuffer(jpegImgRaw)

		img, err = jpeg.Decode(jpgImgBuf)
		if err != nil {
			return nil, nil, errors.Wrap(err, "failed to make image from buffer")
		}

		jpgImgBuf = bytes.NewBuffer(jpegImgRaw)
	}

	return jpgImgBuf, img, nil
}

func convertToJPG(origImg image.Image) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	if err := jpeg.Encode(buf, origImg, nil); err != nil {
		return nil, errors.Wrap(err, "failed to encode to jpg")
	}

	return buf.Bytes(), nil
}

// MakeThumbnail returns bytes of thumbnail from image with saved proportions
func MakeThumbnail(img image.Image, config image.Config) ([]byte, error) {
	proportion := float32(config.Width) / float32(config.Height)
	dstImage := image.NewRGBA(image.Rect(0, 0, int(200*proportion), 200))

	if err := graphics.Thumbnail(dstImage, img); err != nil {
		return nil, errors.Wrap(err, "failed to create thumbnail")
	}

	thumbnail := bytes.NewBuffer(nil)
	_ = jpeg.Encode(thumbnail, dstImage, nil)

	return thumbnail.Bytes(), nil
}
