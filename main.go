package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Post struct {
	Id uint64
	Publisher string
	PublishTime string
	Text string 
}


func main() {
	tmpl := template.Must(template.ParseFiles("./template/index.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			posts := []Post{ // TODO change to API
				Post{1, "vkoorits", "16 hours ago", "Пустите меня на танцпол пьяным подвигаться"},
				Post{2, "vkoorits", "16 hours ago", "All those moments will be lost in time"},
				Post{3, "anonim", "17 hours ago", "Остновите, Вите надо выйти"},
			}

			tmpl.Execute(w, struct {
					Posts []Post
				}{
					posts,
				})
	})

    http.Handle("/static/", http.StripPrefix(
        "/static/",
        http.FileServer(http.Dir("./static")),
    ))

	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", nil)
}
