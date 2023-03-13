package handle

import (
	"fmt"
	"forum/forum"
	"html/template"
	"net/http"
)

func Enregistrement(w http.ResponseWriter, r *http.Request) {
	page := template.Must(template.ParseFiles("./templates/enregistrement.html"))
	session, err := forum.Store.Get(r, "forum")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Vérifier si l'utilisateur est connecté
	_, ok := session.Values["username"].(string)
	if ok {
		// Rediriger l'utilisateur vers la page de connexion s'il n'est pas connecté
		http.Redirect(w, r, "/account", http.StatusSeeOther)
		return
	}
	if r.Method == "POST" {
		// Récupérer les données du formulaire
		var age string
		pseudo := r.FormValue("pseudo")
		mdp := r.FormValue("mdp")
		testmdp := r.FormValue("testmdp")
		prenom := r.FormValue("prenom")
		nom := r.FormValue("nom")
		mail := r.FormValue("mail")
		age = r.FormValue("age")
		IconCouleur := r.FormValue("icon")
		if pseudo != "" || mdp != "" || testmdp != "" || prenom != "" || nom != "" || mail != "" || age != "" {

			// Vérifier si l'utilisateur est déjà pris
			taken, err := forum.PseudoCheck(pseudo)
			if err != nil {
				fmt.Println(err)
				return
			}
			if taken {
				messageerror := "Utilisateur deja utilisé !"
				Message := forum.ErreurMessage{
					Message: messageerror,
				}
				page.Execute(w, Message)
				return
			}
			if mdp != testmdp {
				messageerror := "mot de passe pas concordant !"
				Message := forum.ErreurMessage{
					Message: messageerror,
				}
				page.Execute(w, Message)
				return
			}
			hashmdp, _ := forum.HashMdp(mdp)
			_, err = forum.Bd.Exec("INSERT INTO Utilisateurs (pseudo, mdp, prenom, nom, mail, age, icon) VALUES (?, ?, ?, ?, ?, ?, ?)", pseudo, hashmdp, prenom, nom, mail, age, IconCouleur)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("Nouveau compte : ", pseudo)
			// Rediriger vers la page de connexion
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		} else {
			messageerror := "Erreur veuillez entre les informations !"
			Message := forum.ErreurMessage{
				Message: messageerror,
			}
			page.Execute(w, Message)
		}
	} else {
		messageerror := "Entrez bien toutes les informations demandé."
		Message := forum.ErreurMessage{
			Message: messageerror,
		}
		page.Execute(w, Message)
	}
}
