package handle

import (
	"fmt"
	"forum/forum"
	"html/template"
	"log"
	"net/http"
)

func Connexion(w http.ResponseWriter, r *http.Request) {
	page := template.Must(template.ParseFiles("./templates/connexion.html"))
	if r.Method == "POST" {
		pseudo := r.FormValue("pseudo")
		mdp := r.FormValue("mdp")
		// Vérifier les identifiants de l'utilisateur
		test, err := forum.Check(pseudo, mdp)
		if err != nil {
			log.Fatal(err)
		}
		if test {
			// Authentification réussie : créer une session pour l'utilisateur
			fmt.Println("connexion reussite de ", pseudo)
			session, err := forum.Store.Get(r, "forum")
			if err != nil {
				fmt.Println(err)
			}
			session.Values["pseudo"] = pseudo
			session.Save(r, w)
			http.Redirect(w, r, "/connecte", http.StatusFound)
		} else {
			// Authentification échouée : afficher un message d'erreur
			message := "mauvais identifiant"
			Message := forum.ErreurMessage{
				Message: message,
			}
			page.Execute(w, Message)
		}
	} else {
		message := "entrez identifiant"
		Message := forum.ErreurMessage{
			Message: message,
		}
		page.Execute(w, Message)

	}
}
