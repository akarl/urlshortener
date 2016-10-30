package handlers

import (
	"log"
	"net/http"

	"urlshortener/models"

	"github.com/julienschmidt/httprouter"
)

func redirectToURL(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	code := ps.ByName("code")
	log.Println(code)

	if url, err := models.GetURL(code); err == nil {
		log.Printf("The url: %+v\n", url)
		http.Redirect(w, r, url.URL, http.StatusFound)
	} else {
		http.NotFound(w, r)
	}
}

func listUrls(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	urls, _ := models.GetURLs()

	data := struct {
		URLs []models.URL
	}{
		urls,
	}

	renderTemplate(w, "list_urls", data)
}

func createURL(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	urlstring := r.FormValue("url")

	models.CreateURL(urlstring)

	http.Redirect(w, r, "/", http.StatusFound)
}

func deleteURL(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	code := ps.ByName("code")

	models.DeleteURL(code)

	http.Redirect(w, r, "/", http.StatusFound)
}
