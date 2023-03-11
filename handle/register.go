package handle

import (
	"data"
	"database/sql"
	"html/template"
	"log"
	"net/http"
)

type Message struct {
	ProcessMessage string
}

func Register(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/register.html"))
	db, err := sql.Open("sqlite3", "./database/db.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if r.Method == "POST" {
		// Récupérer les données du formulaire
		username := r.FormValue("username")
		password := r.FormValue("password")
		passwordcheck := r.FormValue("passwordcheck")
		firstname := r.FormValue("firstname")
		lastname := r.FormValue("lastname")
		mail := r.FormValue("mail")
		age := r.FormValue("age")
		IconColor := r.FormValue("icon")
		if username != "" || password != "" || passwordcheck != "" || firstname != "" || lastname != "" || mail != "" || age != "" {

			// Vérifier si l'utilisateur est déjà pris
			taken, err := isUsernameTaken(db, username)
			if err != nil {
				http.Error(w, "Une erreur est survenue.", http.StatusInternalServerError)
				return
			}
			if taken {
				messageerror := "Utilisateur deja utilisé !"
				Message := Message{
					ProcessMessage: messageerror,
				}
				tmpl.Execute(w, Message)
				return
			}
			if password != passwordcheck {
				messageerror := "mot de passe pas concordant !"
				Message := Message{
					ProcessMessage: messageerror,
				}
				tmpl.Execute(w, Message)
				return
			}
			hashpassword, err := data.HashPassword(password)
			_, err = db.Exec("INSERT INTO users (username, password, firstname, lastname, mail, age, icon) VALUES (?, ?, ?, ?, ?, ?, ?)", username, hashpassword, firstname, lastname, mail, age, IconColor)
			if err != nil {
				log.Fatal(err)
			}

			// Rediriger vers la page de connexion
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		} else {
			messageerror := "Erreur veuillez entre les informations !"
			Message := Message{
				ProcessMessage: messageerror,
			}
			tmpl.Execute(w, Message)
		}
	} else {
		messageerror := "Entrez bien toutes les informations demandé."
		Message := Message{
			ProcessMessage: messageerror,
		}
		tmpl.Execute(w, Message)

	}

}
