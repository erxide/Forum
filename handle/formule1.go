package handle

import (
	"fmt"
	"forum/forum"
	"html/template"
	"net/http"
)

func Formule1(w http.ResponseWriter, r *http.Request) {
	var posts []forum.Poste
	var coms []forum.Commentaire
	pageconnecte := template.Must(template.ParseFiles("./templates/connecte.html"))
	pagenonconnecte := template.Must(template.ParseFiles("./templates/accueil.html"))
	session, err := forum.Store.Get(r, "forum")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	POSTES, err := forum.Bd.Query("SELECT id, theme, titre, description, cree_le, cree_par, likes, dislikes FROM Postes WHERE theme = 'formule1' ORDER BY id DESC;")
	defer POSTES.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	// Vérifier si l'utilisateur est connecté
	pseudo, ok := session.Values["pseudo"].(string)
	idPseudo, prenom, nom, mail, age, icon, err := forum.ObtenirInfoUtilisateur(pseudo)
	utilisateurs := forum.Utilisateurs{
		ID:     idPseudo,
		Pseudo: pseudo,
		Prenom: prenom,
		Nom:    nom,
		Mail:   mail,
		Age:    age,
		Icon:   icon,
	}

	for POSTES.Next() {
		var id int
		var theme, titre, description, cree_le, cree_par string
		var likes, dislikes int

		err := POSTES.Scan(&id, &theme, &titre, &description, &cree_le, &cree_par, &likes, &dislikes)
		if err != nil {
			http.Error(w, "Erreur lors de la récupération des posts", http.StatusInternalServerError)
			return
		}
		idpseudo, _, _, _, _, icon, err := forum.ObtenirInfoUtilisateur(cree_par)
		for i, _ := range posts {
			COM, err := forum.Bd.Query("SELECT id, contenu, idPost, idPseudo FROM Commentaires WHERE idPost=?", posts[i].ID)
			if err != nil {
				http.Error(w, "Erreur lors de la récupération des commentaires", http.StatusInternalServerError)
				return
			}
			defer COM.Close()
			var coms []forum.Commentaire
			for COM.Next() {
				var id int
				var contenu, cree_le, cree_par string

				err := COM.Scan(&id, &contenu, &cree_le, &cree_par)
				if err != nil {
					fmt.Println(err)
					return
				}
				_, _, _, _, _, icon, err := forum.ObtenirInfoUtilisateurID(idpseudo)
				com := forum.Commentaire{
					ID:           id,
					Contenu:      contenu,
					IdPseudo:     idpseudo,
					IdPost:       posts[i].ID,
					IconDuPseudo: icon,
				}
				coms = append(coms, com)
			}
			posts[i].Coms = coms
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
			Icon:        icon,
			Coms:        coms,
		}
		posts = append(posts, post)

	}

	envoie := forum.Envoie{
		User: utilisateurs,
		Post: posts,
	}
	if ok {
		pageconnecte.Execute(w, envoie)
	} else {
		pagenonconnecte.Execute(w, envoie)
	}
}
