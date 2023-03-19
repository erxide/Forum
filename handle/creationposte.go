package handle

import (
	"fmt"
	"forum/forum"
	"html/template"
	"net/http"
	"time"
)

// CreationPost gere la page pour creer un post
func CreationPost(w http.ResponseWriter, r *http.Request) {
	// page est le fichier html a executer
	page := template.Must(template.ParseFiles("./templates/creationdeposte.html"))
	// recuperation de de la de la session utilisateur
	session, err := forum.Store.Get(r, "forum")
	// gestion de l'erreur
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Vérifier si l'utilisateur est connecté
	pseudo, ok := session.Values["pseudo"].(string)
	if !ok {
		// Rediriger l'utilisateur vers la page de connexion s'il n'est pas connecté
		http.Redirect(w, r, "/connexion", http.StatusSeeOther)
		return
	}
	// methode post
	if r.Method == "POST" {
		// recupation des valeurs
		theme := r.FormValue("theme")
		titre := r.FormValue("titre")
		description := r.FormValue("description")
		date := time.Now().Format("15:04:05 02-01-2006")
		// requette sql pour creer le poste
		_, err := forum.Bd.Exec("INSERT INTO Postes (theme, titre, description, cree_le, cree_par, likes) VALUES (?, ?, ?, ?, ?, 0)", theme, titre, description, date, pseudo)
		// gestion de l'erreur
		if err != nil {
			fmt.Println(err)
		}
		// Rediriger vers la page d'acceuil
		http.Redirect(w, r, "/accueil", http.StatusSeeOther)
	}
	_, _, _, _, _, Icon, err := forum.ObtenirInfoUtilisateur(pseudo)
	Utilisateur := forum.Utilisateurs{
		Pseudo: pseudo,
		Icon:   Icon,
	}

	// exexcution de la page
	page.Execute(w, Utilisateur)
}
