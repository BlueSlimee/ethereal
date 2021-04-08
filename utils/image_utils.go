package utils

import(
	"image"
        _ "image/jpeg"
        _ "image/png"
	"net/http"
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
