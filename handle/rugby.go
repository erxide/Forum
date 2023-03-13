package handle

import (
	"forum/forum"
	"html/template"
	"net/http"
)

func ConnecteRugby(w http.ResponseWriter, r *http.Request) {
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
	// fmt.Println("test :", id, prenom, nom, mail, age, icon, err, pseudo)
	utilisateurs := forum.Utilisateurs{
		ID:     id,
		Pseudo: pseudo,
		Prenom: prenom,
		Nom:    nom,
		Mail:   mail,
		Age:    age,
		Icon:   icon,
	}
	// Récupérer les données des posts dans la base de données
	rows, err := forum.Bd.Query("SELECT id, theme, titre, description, cree_le, cree_par, likes, dislikes FROM Postes WHERE theme = 'rugby' ORDER BY id DESC;")

	if err != nil {
		http.Error(w, "Erreur lors de la récupération des posts", http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	var posts []forum.Poste
	for rows.Next() {
		var id int
		var theme, titre, description, cree_le, cree_par string
		var likes, dislikes int

		err := rows.Scan(&id, &theme, &titre, &description, &cree_le, &cree_par, &likes, &dislikes)
		if err != nil {
			http.Error(w, "Erreur lors de la récupération des posts", http.StatusInternalServerError)
			return
		}

		post := forum.Poste{
			ID:          id,
			Titre:       titre,
			Theme:       theme,
			Description: description,
			Creele:      cree_le,
			CreePar:     cree_par,
			Likes:       likes,
			Dislikes:    dislikes,
		}

		posts = append(posts, post)
	}
	envoie := forum.Envoie{
		User: utilisateurs,
		Post: posts,
	}

	page.Execute(w, envoie)
}
