package main

import (
	"context"
	"html/template"
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
)

var templates *template.Template
var client *redis.Client
var ctx = context.Background()

func main() {
	templates = template.Must(templates.ParseGlob("templates/*.html"))
	client = redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // default redis address
	})
	r := mux.NewRouter()
	r.HandleFunc("/", indexPage).Methods("GET")
	http.Handle("/", r)
	http.ListenAndServe(":8000", nil)

}

func indexPage(w http.ResponseWriter, r *http.Request) {
	comments, err := client.LRange(ctx, "comments", 0, 20).Result()
	if err != nil {
		return
	}
	templates.ExecuteTemplate(w, "index.html", comments)
}
