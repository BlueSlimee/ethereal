package utils

import(
	"net/http"
	"image"
	"image/png"
	"bytes"
	"strconv"
)

func ReplyWithImage(w http.ResponseWriter, img image.Image) {
	buffer := new(bytes.Buffer)
	if err := png.Encode(buffer, img); err != nil {
		http.Error(w, "Failure encoding image to PNG buffer", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	
	if _, err := w.Write(buffer.Bytes()); err != nil {
		http.Error(w, "Failure writing PNG buffer to HTTP response writer", http.StatusInternalServerError)
        }
}
