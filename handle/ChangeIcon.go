package handle

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
)

func ChangeIcon(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/changeicon.html"))
	db, err := sql.Open("sqlite3", "./database/db.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	session, err := store.Get(r, "forum")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Vérifier si l'utilisateur est connecté
	username, ok := session.Values["username"].(string)
	if !ok {
		// Rediriger l'utilisateur vers la page de connexion s'il n'est pas connecté
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	if r.Method == "POST" {
		IconColor := r.FormValue("icon")
		stmt, err := db.Prepare("UPDATE users SET icon = ? WHERE username = ?")
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
	tmpl.Execute(w, r)
}
