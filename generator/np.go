package generator

import (
	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"image"
	"lastgram.xyz/ethereal/utils"
)

func Np(track, album, artist, scrobbles, url string) image.Image {
	dc := gg.NewContext(1500, 500)
	i := utils.GetImage("https://lastfm.freetls.fastly.net/i/u/500x500/" + url + ".png")
	if i == nil {
		i = utils.GetImage("https://lastgram.vercel.app/last/missingtrack.png")
	}

	// COVER
	{
		dc.DrawImage(*i, 0, 0)
	}
	// BLURRED BG
	{
		bg := imaging.Fill(*i, 1000, 500, imaging.Center, imaging.NearestNeighbor)
		bg = imaging.Blur(bg, 7.0)
		dc.DrawImage(bg, 500, 0)
		dc.SetRGBA(0, 0, 0, 90)
		dc.DrawRectangle(500, 0, 1000, 500)
		dc.Fill()
	}
  // TEXT
  blockBase := 120.0
  blockBaseX := 530.0
  {
    dc.SetRGB(1, 1, 1)
    utils.LoadAndUseFont(dc, "montserrat", "bold", 60)
    dc.DrawString(track, blockBaseX, blockBase)
    
    utils.LoadAndUseFont(dc, "montserrat", "medium", 55)
    dc.DrawString(album, blockBaseX, blockBase + 75)
    
    dc.SetRGB(0.7, 0.7, 0.7)
    utils.LoadAndUseFont(dc, "montserrat", "medium-italic", 55)
    dc.DrawString(artist, blockBaseX, blockBase + 140)
    
    utils.LoadAndUseFont(dc, "montserrat", "medium-italic", 48)
    dc.SetRGB(1, 1, 1)
    dc.DrawString(scrobbles, blockBaseX, blockBase + 325)
  }
	return dc.Image()
}
