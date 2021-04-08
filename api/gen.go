package api

import(
	"lastgram.xyz/ethereal/utils"
	"lastgram.xyz/ethereal/generator"
	"net/http"
	"net/url"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	if r := params.Get("resource"); r != "" {
		Check(w, params)
		return
	}
	http.Error(w, "Missing parameters", http.StatusBadRequest)
	return
}

func Check(w http.ResponseWriter, p url.Values) {
	switch res := p.Get("resource"); res {
	case "np":
		if i := p.Get("i"); i == "" {
			http.Error(w, "Missing parameters", http.StatusBadRequest)
			return
		}
		img := generator.Np(p.Get("i"))
		utils.ReplyWithImage(w, img)
		return
	}
}
