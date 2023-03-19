package handle

import (
	"forum/forum"
	"net/http"
)

func SupPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Erreur lors de l'extraction des données de la requête", http.StatusInternalServerError)
		return
	}
	postID := r.FormValue("post_id")
	_, err = forum.Bd.Exec("DELETE FROM Postes WHERE id = ?", postID)
	if err != nil {
		http.Error(w, "Erreur lors de la mise à jour du nombre de likes", http.StatusInternalServerError)
		return
	}

	// Rediriger l'utilisateur vers la page précédente après avoir liké le poste
	http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
}
