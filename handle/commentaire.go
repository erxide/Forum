package handle

import (
	"fmt"
	"forum/forum"
	"net/http"
)

// Commentaire gere la creation des commantaires
func Commentaire(w http.ResponseWriter, r *http.Request) {
	// recuperation de de la de la session utilisateur
	session, err := forum.Store.Get(r, "forum")
	// gestion de l'erreur
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// gestion de l'erreur
	err = r.ParseForm()
	if err != nil {
		http.Error(w, "Erreur lors de l'extraction des données de la requête", http.StatusInternalServerError)
		return
	}
	// recuperation des valeur pour creer le post
	pseudo, _ := session.Values["pseudo"].(string)
	contenu := r.FormValue("contenu")
	idpost := r.FormValue("post_id")
	idpseudo, _, _, _, _, _, _ := forum.ObtenirInfoUtilisateur(pseudo)
	// resqette sql pour creer le commentaire
	_, err = forum.Bd.Exec("INSERT INTO Commentaires (idPost, idPseudo, contenu) VALUES (?, ?, ?)", idpost, idpseudo, contenu)
	// gestion de l'erreur
	if err != nil {
		fmt.Println(err)
	}

	// Rediriger l'utilisateur vers la page précédente après avoir commenté le poste
	http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
}
