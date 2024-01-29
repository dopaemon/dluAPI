package controller

import (
	"html/template"
	"log"
	"net/http"
	"path"
)

func Register(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		path.Join("views/pages", "register.html"), // halaman yang ingin ditampilkan
		path.Join("views", "pages.html"),
		path.Join("views/includes", "scripts.html"),
	)

	if err != nil {
		log.Println(err)
		http.Error(w, "Error is hapenning, keep calm", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is hapenning, keep calm", http.StatusInternalServerError)
		return
	}
}
