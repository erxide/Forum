package handle

import (
	"fmt"
	"forum/forum"
	"html/template"
	"log"
	"net/http"
)

func ChangeMdp(w http.ResponseWriter, r *http.Request) {
	page := template.Must(template.ParseFiles("./templates/changemdp.html"))
	session, err := forum.Store.Get(r, "forum")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Vérifier si l'utilisateur est connecté
	pseudo, ok := session.Values["pseudo"].(string)
	if !ok {
		// Rediriger l'utilisateur vers la page de connexion s'il n'est pas connecté
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	if r.Method == "POST" {
		AncienMdp := r.FormValue("AncienMdp")
		NouveauMdp := r.FormValue("NouveauMdp")
		NouveauMdpCheck := r.FormValue("NouveauMdpCheck")
		test, err := forum.Check(pseudo, AncienMdp)
		if err != nil {
			fmt.Println(err)
		}
		if test {
			// test ok bon ancien mot de passe
			if NouveauMdp == NouveauMdpCheck {
				HashMdp, err := forum.HashMdp(NouveauMdp)
				if err != nil {
					log.Fatal(err)
				}
				stmt, err := forum.Bd.Prepare("UPDATE Utilisateurs SET mdp = ? WHERE pseudo = ?")
				if err != nil {
					log.Fatal(err)
				}
				defer stmt.Close()
				_, err = stmt.Exec(HashMdp, pseudo)
				if err != nil {
					log.Fatal(err)
				}
				http.Redirect(w, r, "/account", http.StatusFound)
			} else {
				messageerreur := "Nouveau mot de passe et nouveau mot de passe ne sont pas les mêmes"
				message := forum.ErreurMessage{
					Message: messageerreur,
				}
				page.Execute(w, message)
			}
		} else {
			// test faux message error
			messageerreur := "Mauvais ancien mot de passe "
			message := forum.ErreurMessage{
				Message: messageerreur,
			}
			page.Execute(w, message)
		}
	} else {
		messageerreur := "Veuillez bien remplir tout les champs"
		message := forum.ErreurMessage{
			Message: messageerreur,
		}
		page.Execute(w, message)
	}

}
