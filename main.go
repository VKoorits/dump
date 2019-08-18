package main

import (
	"fmt"
	"html/template"
	"net/http"
	"encoding/json"
)

type Post struct {
	Id uint64 `json:"post_id"`
	Publisher string `json:"publisher"`
	PublishTime string `json:"publish_time"`
	Text string `json:"text"`
}

type RootPageHandler struct {
	Tmpl *template.Template
}

func (h *RootPageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Tmpl.Execute(w, nil)
}

func HandlePostsApi(w http.ResponseWriter, r *http.Request) {
	posts := []Post{
		Post{1, "vkoorits", "16 hours ago", "Пустите меня на танцпол пьяным подвигаться"},
		Post{2, "vkoorits", "16 hours ago", "All those moments will be lost in time"},
		Post{3, "anonim", "17 hours ago", "Остновите, Вите надо выйти"},
	}
	postsJson, _ := json.Marshal(posts)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(postsJson))
}



func main() {
	rootHandler := &RootPageHandler{Tmpl: template.Must(template.ParseFiles("./template/index.html"))}
	http.Handle("/", rootHandler)
	http.HandleFunc("/api/posts", HandlePostsApi)
    http.Handle("/static/", http.StripPrefix(
        "/static/",
        http.FileServer(http.Dir("./static")),
    ))

	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", nil)
}
