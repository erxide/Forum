package handle

import (
	"fmt"
	"forum/forum"
	"html/template"
	"net/http"
)

// Formule1 gere l'affichage sur la page d'acceuil avec les posts dont le theme est formule1
func Formule1(w http.ResponseWriter, r *http.Request) {
	// definision de posts et coms qui sont des tableaux de structure
	var posts []forum.Poste
	var coms []forum.Commentaire
	// definision des page a executer
	pageconnecte := template.Must(template.ParseFiles("./templates/accueilco.html"))
	pagenonconnecte := template.Must(template.ParseFiles("./templates/accueil.html"))
	// recuperation de de la de la session utilisateur
	session, err := forum.Store.Get(r, "forum")
	// gestion de l'erreur
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// recuperation des postes  avec le theme formule 1 depuis la base de donnée via une requette sql
	POSTES, err := forum.Bd.Query("SELECT id, theme, titre, description, cree_le, cree_par, likes, dislikes FROM Postes WHERE theme = 'formule1' ORDER BY id DESC;")
	defer POSTES.Close()
	// gestion de l'erreur
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des posts", http.StatusInternalServerError)
		return
	}
	// Vérifier si l'utilisateur est connecté
	pseudo, ok := session.Values["pseudo"].(string)
	// recuperation des information de l'utilisateur stocké dans le session
	idPseudo, prenom, nom, mail, age, icon, err := forum.ObtenirInfoUtilisateur(pseudo)
	// entré les valeur dans la structure utilisateur
	utilisateurs := forum.Utilisateurs{
		ID:     idPseudo,
		Pseudo: pseudo,
		Prenom: prenom,
		Nom:    nom,
		Mail:   mail,
		Age:    age,
		Icon:   icon,
	}
	// gestion des postes
	for POSTES.Next() {
		// déclaration de variables
		var id int
		var theme, titre, description, cree_le, cree_par string
		var likes, dislikes int
		// rentrer les informations des postes dans les variables
		err := POSTES.Scan(&id, &theme, &titre, &description, &cree_le, &cree_par, &likes, &dislikes)
		// gestion de l'erreur
		if err != nil {
			http.Error(w, "Erreur lors de la récupération des posts", http.StatusInternalServerError)
			return
		}
		// recuperation des information de createur du poste
		idpseudo, _, _, _, _, icon, err := forum.ObtenirInfoUtilisateur(cree_par)
		// recuperation des commentaire par rapport au poste
		for i := range posts {
			// requette sql pour recuperer les commentaire de la base de donnée
			COM, err := forum.Bd.Query("SELECT id, contenu, idPost, idPseudo FROM Commentaires WHERE idPost=?", posts[i].ID)
			// gestion de l'erreur
			if err != nil {
				http.Error(w, "Erreur lors de la récupération des commentaires", http.StatusInternalServerError)
				return
			}
			defer COM.Close()
			// declaration de coms, tableau de la structure Commentaire
			var coms []forum.Commentaire
			// gestion des commentaires
			for COM.Next() {
				// declaration de variables
				var id int
				var contenu, cree_le, cree_par string
				// rentrais des infos des commentaires dans les variables
				err := COM.Scan(&id, &contenu, &cree_le, &cree_par)
				// gestion de l'erreur
				if err != nil {
					fmt.Println(err)
					return
				}
				// recuperation de l'auteur du commentaire
				_, _, _, _, _, icon, err := forum.ObtenirInfoUtilisateurID(idpseudo)
				// entré les valeurs dans la structure commentaire
				com := forum.Commentaire{
					ID:           id,
					Contenu:      contenu,
					IdPost:       posts[i].ID,
					IconDuPseudo: icon,
				}
				//ajout du commantaire dans le tableau coms
				coms = append(coms, com)
			}
			posts[i].Coms = coms
		}
		// entré les valeurs dans la structure poste
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
		// ajout du post dans le tableau posts
		posts = append(posts, post)

	}
	// entré les valeurs dans la structure envoie
	envoie := forum.Envoie{
		User: utilisateurs,
		Post: posts,
	}
	// si l'utilisateur connecté
	if ok {
		// executer la page connecte
		pageconnecte.Execute(w, envoie)
	}
	// si l'utilisateur pas connecté
	if !ok {
		// executer la page non connecté
		pagenonconnecte.Execute(w, envoie)
	}
}
