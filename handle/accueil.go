package handle

import (
	"forum/forum"
	"html/template"
	"net/http"
)

func Accueil(w http.ResponseWriter, r *http.Request) {
	page := template.Must(template.ParseFiles("./templates/accueil.html"))
	session, err := forum.Store.Get(r, "forum")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Vérifier si l'utilisateur est connecté
	_, ok := session.Values["pseudo"].(string)
	if ok {
		// Rediriger l'utilisateur vers la page de connexion s'il n'est pas connecté
		http.Redirect(w, r, "/connecte", http.StatusSeeOther)
		return
	}
	page.Execute(w, r)
}
