package handle

import (
	"forum/forum"
	"html/template"
	"log"
	"net/http"
)

func ChangeIcon(w http.ResponseWriter, r *http.Request) {
	page := template.Must(template.ParseFiles("./templates/changeicon.html"))
	session, err := forum.Store.Get(r, "forum")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Vérifier si l'utilisateur est connecté
	username, ok := session.Values["pseudo"].(string)
	if !ok {
		// Rediriger l'utilisateur vers la page de connexion s'il n'est pas connecté
		http.Redirect(w, r, "/connexion", http.StatusSeeOther)
		return
	}
	if r.Method == "POST" {
		IconColor := r.FormValue("icon")
		stmt, err := forum.Bd.Prepare("UPDATE Utilisateurs SET icon = ? WHERE pseudo = ?")
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()
		_, err = stmt.Exec(IconColor, username)
		if err != nil {
			log.Fatal(err)
		}
		http.Redirect(w, r, "/account", http.StatusFound)
	}
	page.Execute(w, r)
}
