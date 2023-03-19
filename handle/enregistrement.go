package handle

import (
	"fmt"
	"forum/forum"
	"html/template"
	"net/http"
)

// Enregistrement gere l'enregistrement des comptes dans la base de donnée
func Enregistrement(w http.ResponseWriter, r *http.Request) {
	// page est le fichier html a executer
	page := template.Must(template.ParseFiles("./templates/Inscription.html"))
	// recuperation de de la de la session utilisateur
	session, err := forum.Store.Get(r, "forum")
	// gestion de l'erreur
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Vérifier si l'utilisateur est connecté
	_, ok := session.Values["pseudo"].(string)
	if ok {
		// Rediriger l'utilisateur vers la page de connexion s'il n'est pas connecté
		http.Redirect(w, r, "/accueil", http.StatusSeeOther)
		return
	}
	// methode post
	if r.Method == "POST" {
		// Récupérer les données du formulaire
		pseudo := r.FormValue("pseudo")
		mdp := r.FormValue("mdp")
		testmdp := r.FormValue("testmdp")
		prenom := r.FormValue("prenom")
		nom := r.FormValue("nom")
		mail := r.FormValue("mail")
		age := r.FormValue("age")
		IconCouleur := r.FormValue("icon")

		// Vérifier si l'utilisateur est déjà pris
		taken, err := forum.PseudoCheck(pseudo)
		// gestion de l'erreur
		if err != nil {
			fmt.Println(err)
			return
		}
		if taken {
			// si l'utilisateur deja pris alors executer la page avec un message d'erreur
			messageerror := "Utilisateur deja utilisé !"
			Message := forum.ErreurMessage{
				Message: messageerror,
			}
			page.Execute(w, Message)
			return
		}
		if mdp != testmdp {
			// si les mot de passes ne concorde pas executer la page avec le message d'erreur
			messageerror := "mot de passe pas concordant !"
			Message := forum.ErreurMessage{
				Message: messageerror,
			}
			page.Execute(w, Message)
			return
		}
		// hashage du mot de passe
		hashmdp, _ := forum.HashMdp(mdp)
		// requette sql pour enregistrer un nouveau compte
		_, err = forum.Bd.Exec("INSERT INTO Utilisateurs (pseudo, mdp, prenom, nom, mail, age, icon) VALUES (?, ?, ?, ?, ?, ?, ?)", pseudo, hashmdp, prenom, nom, mail, age, IconCouleur)
		// gestion de l'erreur
		if err != nil {
			fmt.Println(err)
		}
		// ecrire dans le terminal quand un compte est creer
		fmt.Println("Nouveau compte : ", pseudo)
		// Rediriger vers la page d'accueil'
		http.Redirect(w, r, "/accueil", http.StatusSeeOther)
	} else {
		// si il y a erreur alors executer la page avec le message d'erreur
		messageerror := "Erreur veuillez entre les informations !"
		Message := forum.ErreurMessage{
			Message: messageerror,
		}
		page.Execute(w, Message)
	}
}
