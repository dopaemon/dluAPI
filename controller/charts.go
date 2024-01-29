package controller

import (
	"html/template"
	"log"
	"net/http"
	"path"
)

func Charts(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		path.Join("views/pages", "charts.html"), // halaman yang ingin ditampilkan
		path.Join("views", "layout.html"),
		path.Join("views/includes", "sidebar.html"),
		path.Join("views/includes", "navbar.html"),
		path.Join("views/includes", "footer.html"),
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
