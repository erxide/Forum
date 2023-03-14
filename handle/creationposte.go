package handle

import (
	"fmt"
	"forum/forum"
	"html/template"
	"net/http"
	"time"
)

func CreationPost(w http.ResponseWriter, r *http.Request) {
	page := template.Must(template.ParseFiles("./templates/creationposte.html"))
	session, err := forum.Store.Get(r, "forum")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Vérifier si l'utilisateur est connecté
	pseudo, ok := session.Values["pseudo"].(string)
	if !ok {
		// Rediriger l'utilisateur vers la page de connexion s'il n'est pas connecté
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	if r.Method == "POST" {
		theme := r.FormValue("theme")
		titre := r.FormValue("titre")
		description := r.FormValue("description")
		date := time.Now().Format("15:04:05 02-01-2006")
		_, err := forum.Bd.Exec("INSERT INTO Postes (theme, titre, description, cree_le, cree_par, likes) VALUES (?, ?, ?, ?, ?, 0)", theme, titre, description, date, pseudo)
		if err != nil {
			fmt.Println(err)
		}
		// Rediriger vers la page se connecte
		http.Redirect(w, r, "/connecte", http.StatusSeeOther)
	}
	page.Execute(w, r)
}
