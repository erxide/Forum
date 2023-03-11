package handle

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/sessions"
	_ "github.com/mattn/go-sqlite3"
	"html/template"
	"log"
	"net/http"
)

var store = sessions.NewCookieStore([]byte("cle"))
var test1 string

func Login(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/login.html"))
	db, err := sql.Open("sqlite3", "./database/db.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")
		// Vérifier les identifiants de l'utilisateur
		test, err := Check(username, password, db)
		if err != nil {
			log.Fatal(err)
		}
		if test {
			// Authentification réussie : créer une session pour l'utilisateur
			fmt.Println("connexion reussite de ", username)
			session, err := store.Get(r, "forum")
			if err != nil {
				fmt.Println(err)
			}
			session.Values["username"] = username
			session.Save(r, w)
			// fmt.Println("w = ", w, ",\n r = ", r, "\n http.Statusfound = ", http.StatusFound)
			// Rediriger l'utilisateur vers la page de compte
			http.Redirect(w, r, "/account", http.StatusFound)
		} else {
			// Authentification échouée : afficher un message d'erreur
			message := "mauvais identifiant"
			Message := Message{
				ProcessMessage: message,
			}
			tmpl.Execute(w, Message)
			return
		}
	}
	message := "entrez identifiant"
	Message := Message{
		ProcessMessage: message,
	}
	tmpl.Execute(w, Message)
}
