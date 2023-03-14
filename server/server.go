package main

import (
	_ "github.com/mattn/go-sqlite3"
	"handle"
	"log"
	"net/http"
)

func main() {
	// initialisation du fichier assets pour pouvoir afficher le css et les images en front
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", fs))
	// liste des routes http
	http.HandleFunc("/", handleSlash)
	http.HandleFunc("/accueil", handleAccueil)
	http.HandleFunc("/enregistrement", handleEnregistrement)
	http.HandleFunc("/connexion", handleConnexion)
	http.HandleFunc("/connecte", handleConnecte)
	http.HandleFunc("/deconnexion", handleDeconnexion)
	http.HandleFunc("/changeicon", handleChangeIcon)
	http.HandleFunc("/changemotdepasse", handleChangeMdp)
	http.HandleFunc("/creationposte", handleCreationPoste)
	http.HandleFunc("/connecterugby", handleConnecteRugby)
	http.HandleFunc("/connectetennis", handleConnecteTennis)
	http.HandleFunc("/connectebasket", handleConnecteBasket)
	http.HandleFunc("/connectefootball", handleConnectefootball)
	http.HandleFunc("/connecteformule1", handleConnecteformule1)
	http.HandleFunc("/connectehandball", handleConnecteHandball)
	http.HandleFunc("/like", handleLike)
	http.HandleFunc("/dislike", handleDislike)
	http.HandleFunc("/mespostes", handleMesPostes)
	// Écris dans le terminal, si le serveur a démarré, l'url du serveur avec le port
	log.Println("Serveur démarré sur http://localhost:8080")
	// Démarre le serveur sur le port 8080
	err := http.ListenAndServe(":8080", nil)
	// Si il y a une erreur
	if err != nil {
		// Stoppé le programme et écrire l'erreur dans le terminal
		log.Fatal(err)
	}
}

// Fonction handleSlash pour la route /
func handleSlash(w http.ResponseWriter, r *http.Request) {
	// redirection de l'utilisateur vers la route /accueil
	http.Redirect(w, r, "/accueil", http.StatusSeeOther)
}

// Fonction handleAccueil pour la route /accueil
func handleAccueil(w http.ResponseWriter, r *http.Request) {
	// appel de la fonction Accueil dans le dossier forum
	handle.Accueil(w, r)
}

// Fonction handleEnregistrement pour la route /enregistrement
func handleEnregistrement(w http.ResponseWriter, r *http.Request) {
	// appel de la fonction Enregistrement dans le dossier forum
	handle.Enregistrement(w, r)
}

// Fonction handleConnexion pour la route /connexion
func handleConnexion(w http.ResponseWriter, r *http.Request) {
	// appel de la fonction Connexion dans le dossier forum
	handle.Connexion(w, r)
}

// Fonction handleConnecte pour la route /connexion
func handleConnecte(w http.ResponseWriter, r *http.Request) {
	// appel de la fonction Connecte dans le dossier forum
	handle.Connecte(w, r)
}

// Fonction handleDeconnexion pour la route /connexion
func handleDeconnexion(w http.ResponseWriter, r *http.Request) {
	// appel de la fonction deconnexion dans le dossier forum
	handle.Deconnexion(w, r)
}

// Fonction handleChangeIcon pour la route /connexion
func handleChangeIcon(w http.ResponseWriter, r *http.Request) {
	// appel de la fonction changeicon dans le dossier forum
	handle.ChangeIcon(w, r)
}

// Fonction handleChangeMdp pour la route /connexion
func handleChangeMdp(w http.ResponseWriter, r *http.Request) {
	// appel de la fonction changemdp dans le dossier forum
	handle.ChangeMdp(w, r)
}

// Fonction handleCreationPoste pour la route /creationposte
func handleCreationPoste(w http.ResponseWriter, r *http.Request) {
	// appel de la fonction creationposte dans le dossier forum
	handle.CreationPost(w, r)
}

// Fonction handleConnecteRugby pour la route /connecterugby
func handleConnecteRugby(w http.ResponseWriter, r *http.Request) {
	// appel de la fonction creationposte dans le dossier forum
	handle.ConnecteRugby(w, r)
}

// Fonction handleConnecteTennis pour la route /connecterugby
func handleConnecteTennis(w http.ResponseWriter, r *http.Request) {
	// appel de la fonction creationposte dans le dossier forum
	handle.ConnecteTennis(w, r)
}

// Fonction handleConnecteTennis pour la route /connecterugby
func handleConnecteBasket(w http.ResponseWriter, r *http.Request) {
	// appel de la fonction creationposte dans le dossier forum
	handle.ConnecteBasket(w, r)
}

// Fonction handleConnecteTennis pour la route /connecterugby
func handleConnectefootball(w http.ResponseWriter, r *http.Request) {
	// appel de la fonction creationposte dans le dossier forum
	handle.Connectefootball(w, r)
}

// Fonction handleConnecteTennis pour la route /connecterugby
func handleConnecteformule1(w http.ResponseWriter, r *http.Request) {
	// appel de la fonction creationposte dans le dossier forum
	handle.Connecteformule1(w, r)
}

// Fonction handleConnecteTennis pour la route /connecterugby
func handleConnecteHandball(w http.ResponseWriter, r *http.Request) {
	// appel de la fonction creationposte dans le dossier forum
	handle.ConnecteHandball(w, r)
}

// Fonction handleConnecteTennis pour la route /connecterugby
func handleLike(w http.ResponseWriter, r *http.Request) {
	// appel de la fonction creationposte dans le dossier forum
	handle.Like(w, r)
}

// Fonction handleConnecteTennis pour la route /connecterugby
func handleDislike(w http.ResponseWriter, r *http.Request) {
	// appel de la fonction creationposte dans le dossier forum
	handle.Dislike(w, r)
}

// Fonction handleConnecteTennis pour la route /connecterugby
func handleMesPostes(w http.ResponseWriter, r *http.Request) {
	// appel de la fonction creationposte dans le dossier forum
	handle.MesPostes(w, r)
}
