package api

import (
	"lastgram.xyz/ethereal/generator"
	"lastgram.xyz/ethereal/utils"
	"net/http"
	"net/url"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	if r := params.Get("r"); r != "" {
		Check(w, params)
		return
	}
	http.Error(w, "Missing parameters", http.StatusBadRequest)
	return
}

func Check(w http.ResponseWriter, p url.Values) {
	switch res := p.Get("r"); res {
	case "np":
		if i := p.Get("im"); i == "" {
			http.Error(w, "Missing parameters", http.StatusBadRequest)
			return
		}
		if i := p.Get("ar"); i == "" {
			http.Error(w, "Missing parameters", http.StatusBadRequest)
			return
		}
		if i := p.Get("al"); i == "" {
			http.Error(w, "Missing parameters", http.StatusBadRequest)
			return
		}
		if i := p.Get("tr"); i == "" {
			http.Error(w, "Missing parameters", http.StatusBadRequest)
			return
		}
		if i := p.Get("sc"); i == "" {
			http.Error(w, "Missing parameters", http.StatusBadRequest)
			return
		}
		img := generator.Np(p.Get("tr"), p.Get("al"), p.Get("ar"), p.Get("sc"), p.Get("im"))
		utils.ReplyWithImage(w, img)
		return
	}
}
