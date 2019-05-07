package img2ascii

import (
	"github.com/nfnt/resize"

	"bytes"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"reflect"
)

var ASCIISTR = "MND8OZ$7I?+=~:,.."

func ScaleImage(img image.Image, w int) (image.Image, int, int) {
	sz := img.Bounds()
	h := (sz.Max.Y * w * 10) / (sz.Max.X * 16)
	img = resize.Resize(uint(w), uint(h), img, resize.Lanczos3)
	return img, w, h
}

func Img2Ascii(img image.Image, w, h int) []byte {
	table := []byte(ASCIISTR)
	buf := new(bytes.Buffer)

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			g := color.GrayModel.Convert(img.At(j, i))
			y := reflect.ValueOf(g).FieldByName("Y").Uint()
			pos := int(y * 16 / 255)
			_ = buf.WriteByte(table[pos])
		}
		_ = buf.WriteByte('\n')
	}
	return buf.Bytes()
}

func File2Ascii(fpath string, width int) ([]byte, error) {
	f, err := os.Open(fpath)
	if err != nil {
		return nil, err
	}

	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	f.Close()
	return Img2Ascii(ScaleImage(img, width)), nil
}

func MustFile2Ascii(fpath string, width int) []byte {
	a, err := File2Ascii(fpath, width)
	if err != nil {
		panic(err)
	}
	return a
}
