package generator

import (
	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"image"
	"lastgram.xyz/ethereal/utils"
	"strings"
)

func Collage(list []string, w, h int) image.Image {
	op := NewPIO(w*400, h*400)
	x := 0
	y := 0

	for _, i := range list {
		a := strings.Split(i, "-")
		op.RunTaskWithData(func(channel chan TaskResponse, b interface{}) {
			dc := gg.NewContext(400, 400)
			i := utils.GetImage("https://lastfm.freetls.fastly.net/i/u/500x500/" + a[0] + ".png")
			if i == nil {
				i = utils.GetImage("https://lastgram.vercel.app/last/missingtrack.png")
			}
			c := imaging.Resize(*i, 400, 400, imaging.Linear)
			dc.DrawImage(c, 0, 0)
			if len(a) > 1 {
				utils.LoadAndUseFont(dc, "montserrat", "bold", 20)
				dc.SetRGB(1, 1, 1)
				utils.DrawStringOutlined(dc, a[1], 10, 350)
			}
			scrb := 370.0
			posi := 2
			if len(a) == 4 {
				posi = 3
				scrb = 390.0
				utils.LoadAndUseFont(dc, "montserrat", "semi-bold", 18)
				dc.SetRGB(1, 1, 1)
				utils.DrawStringOutlined(dc, a[2], 10, 370)
			}

			if len(a) <= 4 && len(a) > 2 {
				utils.LoadAndUseFont(dc, "montserrat", "medium-italic", 18)
				dc.SetRGB(0.9, 0.9, 0.9)
				j := "scrobbles"
				item := a[posi]
				if item == "1" {
					j = "scrobble"
				}
				utils.DrawStringOutlined(dc, item+" "+j, 10, scrb)
			}

			channel <- TaskResponse{
				Image:  dc.Image(),
				DrawAt: b.([2]int),
			}
		}, [2]int{x, y})
		x = x + 400
		if x >= w*400 {
			x = 0
			y = y + 400
		}
	}
	op.WaitAndDraw()

	return op.Image()
}
