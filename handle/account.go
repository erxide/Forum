package handle

import (
	"data"
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"html/template"
	"net/http"
)

type User struct {
	ID        int
	Username  string
	Password  string
	Firstname string
	Lastname  string
	Mail      string
	Age       int
	Icon      string
}

func Account(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "./database/db.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	tmpl := template.Must(template.ParseFiles("./templates/account.html"))
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
	id, firstname, lastname, mail, age, icon, err := data.GetUserData(username, db)
	info := User{
		ID:        id,
		Username:  username,
		Firstname: firstname,
		Lastname:  lastname,
		Mail:      mail,
		Age:       age,
		Icon:      icon,
	}
	tmpl.Execute(w, info)
}
