package handle

import (
	"forum/forum"
	"net/http"
)

func Like(w http.ResponseWriter, r *http.Request) {
	// Extraire l'ID du poste à liker à partir des données de la requête POST
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Erreur lors de l'extraction des données de la requête", http.StatusInternalServerError)
		return
	}
	postID := r.FormValue("post_id")
	// Mettre à jour le nombre de likes dans la base de données
	_, err = forum.Bd.Exec("UPDATE Postes SET likes = likes + 1 WHERE id = ?", postID)
	if err != nil {
		http.Error(w, "Erreur lors de la mise à jour du nombre de likes", http.StatusInternalServerError)
		return
	}

	// Rediriger l'utilisateur vers la page précédente après avoir liké le poste
	http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
}
