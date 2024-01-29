package controller

import (
	"html/template"
	"log"
	"net/http"
	"path"
)

func Dashboard(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		// return 404 not found page
		tmpl, err := template.ParseFiles(
			path.Join("views/pages", "404.html"), // halaman yang ingin ditampilkan
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
		return
	}

	tmpl, err := template.ParseFiles(
		path.Join("views/pages", "dashboard.html"), // halaman yang ingin ditampilkan
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
