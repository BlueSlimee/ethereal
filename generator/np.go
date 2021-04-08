package generator

import(
	"image"
	"github.com/fogleman/gg"
	"github.com/disintegration/imaging"
	"lastgram.xyz/ethereal/utils"
)

func Np(url string) image.Image {
	dc := gg.NewContext(450, 150)
	i := utils.GetImage("https://lastfm.freetls.fastly.net/i/u/500x500/"+ url +".jpg")
	if i == nil {
		i = utils.GetImage("https://lastgram.vercel.app/last/missingtrack.png")
	}
	
	// COVER
	{
		c := imaging.Resize(*i, 150, 150, imaging.Linear)
		dc.DrawImage(c, 0, 0)
	}
	// BLURRED BG
	{
		bg := imaging.Fill(*i, 300, 150, imaging.Center, imaging.NearestNeighbor)
		bg = imaging.Blur(bg, 4.0)
		dc.DrawImage(bg, 150, 0)
		dc.SetRGBA(0, 0, 0, 100)
		dc.DrawRectangle(150, 0, 300, 150)
		dc.Fill()
	}

	return dc.Image()
}
