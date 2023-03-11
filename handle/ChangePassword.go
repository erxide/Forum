package handle

import (
	"data"
	"database/sql"
	"html/template"
	"log"
	"net/http"
)

func ChangePassword(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "./database/db.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	tmpl := template.Must(template.ParseFiles("./templates/ChangePassword.html"))
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
		OldPassword := r.FormValue("OldPassword")
		NewPassword := r.FormValue("Password")
		NewPasswordCheck := r.FormValue("PasswordCheck")
		if OldPassword != "" || NewPassword != "" || NewPasswordCheck != "" {
			test, err := Check(username, OldPassword, db)
			if err != nil {
				log.Fatal(err)
			}
			if test {
				// test ok bon ancien mot de passe
				if NewPassword == NewPasswordCheck {
					HashPassword, err := data.HashPassword(NewPassword)
					if err != nil {
						log.Fatal(err)
					}
					stmt, err := db.Prepare("UPDATE users SET password = ? WHERE username = ?")
					if err != nil {
						log.Fatal(err)
					}
					defer stmt.Close()
					_, err = stmt.Exec(HashPassword, username)
					if err != nil {
						log.Fatal(err)
					}
					http.Redirect(w, r, "/account", http.StatusFound)
				} else {
					errormessage := "Nouveau mot de passe et nouveau mot de passe ne sont pas les mêmes"
					message := Message{
						ProcessMessage: errormessage,
					}
					tmpl.Execute(w, message)
				}
			} else {
				// test faux message error
				errormessage := "Mauvais ancien mot de passe "
				message := Message{
					ProcessMessage: errormessage,
				}
				tmpl.Execute(w, message)
			}
		} else {
			errormessage := "Veuillez bien remplir tout les champs"
			message := Message{
				ProcessMessage: errormessage,
			}
			tmpl.Execute(w, message)
		}
	}
	errormessage := "Changer de mot de passe"
	message := Message{
		ProcessMessage: errormessage,
	}
	tmpl.Execute(w, message)
}
