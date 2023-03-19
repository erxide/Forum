package handle

import (
	"fmt"
	"forum/forum"
	"net/http"
)

func Commentaire(w http.ResponseWriter, r *http.Request) {
	session, err := forum.Store.Get(r, "forum")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = r.ParseForm()
	if err != nil {
		http.Error(w, "Erreur lors de l'extraction des données de la requête", http.StatusInternalServerError)
		return
	}
	pseudo, _ := session.Values["pseudo"].(string)
	contenu := r.FormValue("contenu")
	idpost := r.FormValue("post_id")
	idpseudo, _, _, _, _, _, _ := forum.ObtenirInfoUtilisateur(pseudo)

	_, err = forum.Bd.Exec("INSERT INTO Commentaires (idPost, idPseudo, contenu) VALUES (?, ?, ?)", idpost, idpseudo, contenu)
	fmt.Println(err)
	// Rediriger l'utilisateur vers la page précédente après avoir liké le poste
	http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
}
