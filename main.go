package main

import (
	"api_server/model"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
	"os"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.AllowContentType("application/json"))
	r.Get("/getpage/{slug}", func(w http.ResponseWriter, r *http.Request) {
		slugParam := chi.URLParam(r, "slug")
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
	http.ListenAndServe(":8000", r)
}
