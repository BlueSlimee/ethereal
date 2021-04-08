package utils

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"net/http"
	"github.com/fogleman/gg"
)

func GetImage(url string) *image.Image {
	res, err := http.Get(url)
	if err != nil {
		return nil
	}
	defer res.Body.Close()

	img, _, err := image.Decode(res.Body)
	if err != nil {
		return nil
	}
	return &img
}

func LoadAndUseFont(ctx *gg.Context, family, style string, size float64) {
  ctx.LoadFontFace("../public/fonts/"+ family +"/"+ style +".ttf", size)
}