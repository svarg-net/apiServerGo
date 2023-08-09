package main

import (
	"api_server/model"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
	"os"
	"strings"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.AllowContentType("application/json"))
	r.Get("/get-page/{slug}", func(w http.ResponseWriter, r *http.Request) {
		slugParam := strings.Replace(chi.URLParam(r, "slug"), ",", "/", -1)
		dataFile, _ := os.ReadFile("static/static.json")
		pages := model.TypeData{}
		json.Unmarshal(dataFile, &pages)
		page := model.TypePage{}
		fmt.Println("slugParam:", slugParam)
		for _, p := range pages.Pages {
			if p.Link == "/"+slugParam {
				fmt.Println("slugParam:", slugParam)
				page = p
				continue
			}
		}
		jsonPage, err := json.Marshal(page)
		if err != nil {
			fmt.Println(err)
		}
		w.Write(jsonPage)
	})
	r.Get("/getPagesSlug", func(w http.ResponseWriter, r *http.Request) {
		dataFile, _ := os.ReadFile("static/static.json")
		pages := model.TypeData{}
		json.Unmarshal(dataFile, &pages)

		jsonPage, err := json.Marshal(pages.ListPageSlug)
		if err != nil {
			fmt.Println(err)
		}
		w.Write(jsonPage)
	})
	http.ListenAndServe(":8000", r)
}
