package main

import (
	"handle"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
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
	http.HandleFunc("/deconnexion", handleDeconnexion)
	http.HandleFunc("/changeicon", handleChangeIcon)
	http.HandleFunc("/changemotdepasse", handleChangeMdp)
	http.HandleFunc("/creationposte", handleCreationPoste)
	http.HandleFunc("/rugby", handleRugby)
	http.HandleFunc("/tennis", handleTennis)
	http.HandleFunc("/basket", handleBasket)
	http.HandleFunc("/football", handlefootball)
	http.HandleFunc("/formule1", handleformule1)
	http.HandleFunc("/handball", handleHandball)
	http.HandleFunc("/like", handleLike)
	http.HandleFunc("/dislike", handleDislike)
	http.HandleFunc("/mespostes", handleMesPostes)
	http.HandleFunc("/suppost", handleSupPost)
	http.HandleFunc("/commentaire", handleCommentaire)
	http.HandleFunc("/test", handleTest)
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
//func handleConnecte(w http.ResponseWriter, r *http.Request) {
// appel de la fonction Connecte dans le dossier forum
//handle.Connecte(w, r)
//}

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
func handleRugby(w http.ResponseWriter, r *http.Request) {
	// appel de la fonction creationposte dans le dossier forum
	handle.Rugby(w, r)
}

// Fonction handleConnecteTennis pour la route /connecterugby
func handleTennis(w http.ResponseWriter, r *http.Request) {
	// appel de la fonction creationposte dans le dossier forum
	handle.Tennis(w, r)
}

// Fonction handleConnecteTennis pour la route /connecterugby
func handleBasket(w http.ResponseWriter, r *http.Request) {
	// appel de la fonction creationposte dans le dossier forum
	handle.Basket(w, r)
}

// Fonction handleConnecteTennis pour la route /connecterugby
func handlefootball(w http.ResponseWriter, r *http.Request) {
	// appel de la fonction creationposte dans le dossier forum
	handle.Football(w, r)
}

// Fonction handleConnecteTennis pour la route /connecterugby
func handleformule1(w http.ResponseWriter, r *http.Request) {
	// appel de la fonction creationposte dans le dossier forum
	handle.Formule1(w, r)
}

// Fonction handleConnecteTennis pour la route /connecterugby
func handleHandball(w http.ResponseWriter, r *http.Request) {
	// appel de la fonction creationposte dans le dossier forum
	handle.Handball(w, r)
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

func handleSupPost(w http.ResponseWriter, r *http.Request) {
	handle.SupPost(w, r)
}

func handleCommentaire(w http.ResponseWriter, r *http.Request) {
	handle.Commentaire(w, r)
}

func handleTest(w http.ResponseWriter, r *http.Request) {
	handle.Test(w, r)
}
