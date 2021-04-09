package utils

import (
	"github.com/fogleman/gg"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"net/http"
	"os"
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
	ctx.LoadFontFace(GetPath()+"/_files/"+family+"/"+style+".ttf", size)
}

func GetPath() string {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err) // if this goes wrong oop
	}
	return pwd
}
