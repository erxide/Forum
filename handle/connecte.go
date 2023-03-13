package handle

import (
	"forum/forum"
	"html/template"
	"net/http"
)

func Connecte(w http.ResponseWriter, r *http.Request) {
	page := template.Must(template.ParseFiles("./templates/connecte.html"))
	session, err := forum.Store.Get(r, "forum")
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
	id, prenom, nom, mail, age, icon, err := forum.ObtenirInfoUtilisateur(pseudo)
	//fmt.Println("test :", id, prenom, nom, mail, age, icon, err, pseudo)
	utilisateurs := forum.Utilisateurs{
		ID:     id,
		Pseudo: pseudo,
		Prenom: prenom,
		Nom:    nom,
		Mail:   mail,
		Age:    age,
		Icon:   icon,
	}
	page.Execute(w, utilisateurs)
}
