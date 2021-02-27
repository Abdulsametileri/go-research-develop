package main

import (
	"embed"
	_ "embed"
	"fmt"
	"github.com/Abdulsametileri/ingilizce-kelime-go/config"
	"github.com/gorilla/mux"
	"net/http"
)

//go:embed assets/*
var assets embed.FS

//go:embed index.html
var websiteIndexFile []byte

func main() {
	config.Setup()

	r := mux.NewRouter()

	r.HandleFunc("/home-page", homePageHandler)

	r.PathPrefix("/assets/").Handler(http.FileServer(http.FS(assets)))

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html")
		w.Write(websiteIndexFile)
	})

	_ = http.ListenAndServe(":8080", r)
	// services.Setup(&services.Telegram{})
	// database.SetupDB(database.MongoClient{})
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprint(w, "Bismillah")
}
