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

func DrawStringOutlined(dc *gg.Context, s string, xa, ya float64) {
	dc.SetRGB(0, 0, 0)
	n := 4
	for dy := -n; dy <= n; dy++ {
		for dx := -n; dx <= n; dx++ {
			if dx*dx+dy*dy >= n*n {
				// give it rounded corners
				continue
			}
			x := xa + float64(dx)
			y := ya + float64(dy)
			dc.DrawString(s, x, y)
		}
	}
	dc.SetRGB(1, 1, 1)
	dc.DrawString(s, xa, ya)
}

func GetPath() string {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err) // if this goes wrong oop
	}
	return pwd
}
