package utils

import (
	"bytes"
	"image"
	"image/jpeg"
	"net/http"
	"strconv"
)

func ReplyWithImage(w http.ResponseWriter, img image.Image) {
	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, img, &jpeg.Options{Quality:95}); err != nil {
		http.Error(w, "Failure encoding image to JPEG buffer", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	// Tell Vercel to cache this response for 31 days.
	w.Header().Set("Cache-Control", "max-age=2678400, public")
	if _, err := w.Write(buffer.Bytes()); err != nil {
		http.Error(w, "Failure writing JPEG buffer to HTTP response writer", http.StatusInternalServerError)
	}
}
