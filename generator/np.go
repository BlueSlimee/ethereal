package generator

import (
	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"image"
	"lastgram.xyz/ethereal/utils"
)

func Np(track, album, artist, url string) image.Image {
	i := utils.GetImage("https://lastfm.freetls.fastly.net/i/u/500x500/" + url + ".png")
	if i == nil {
		i = utils.GetImage("https://lastgram.vercel.app/last/missingtrack.png")
	}
	op := NewPIO(1500, 500)
	// COVER

	op.RunTask(func(c chan TaskResponse) {
		c <- TaskResponse{
			Image:  *i,
			DrawAt: [2]int{0, 0},
		}
	})

	op.RunTask(func(c chan TaskResponse) {
		bg := imaging.Fill(*i, 1000, 500, imaging.Center, imaging.NearestNeighbor)
		bg = imaging.Blur(bg, 7.5)
		bg = imaging.AdjustBrightness(bg, -17)

		c <- TaskResponse{
			Image:  bg,
			DrawAt: [2]int{500, 0},
		}
	})
	alY := 250.0
	arY := 310.0
	trY := 190.0
	if album == track {
		alY = 0.0
		trY = 230.0
		arY = 285.0
	}
	op.RunTask(func(c chan TaskResponse) {
		dc := gg.NewContext(1000, 500)
		dc.SetRGB(1, 1, 1)
		utils.LoadAndUseFont(dc, "montserrat", "bold", 60)
		utils.DrawStringOutlined(dc, track, 15, trY)
		if alY != 0 {
			utils.LoadAndUseFont(dc, "montserrat", "medium", 55)
			utils.DrawStringOutlined(dc, album, 15, alY)
		}
		dc.SetRGB(0.7, 0.7, 0.7)
		utils.LoadAndUseFont(dc, "montserrat", "medium-italic", 55)
		utils.DrawStringOutlined(dc, artist, 15, arY)
		c <- TaskResponse{
			Image:  dc.Image(),
			DrawAt: [2]int{500, 0},
			Mask:   true, // this indicates that this should be the last shard to be drawn
		}
	})
	op.WaitAndDraw()
	return op.Image()
}
