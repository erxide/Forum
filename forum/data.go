package forum

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/sessions"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

// Var pour définir la base de donée
var Bd, err = OuvrirBaseDonnee("./data/db.sqlite")

// Store est var de la cle pour les cookies
var Store = sessions.NewCookieStore([]byte("Motdepassesupersecurisealamortquitu"))

// OuvrirBaseDonnee est une fonction pour ouvrir la connexion à la base de donnée
func OuvrirBaseDonnee(chemin string) (*sql.DB, error) {
	bd, err := sql.Open("sqlite3", chemin)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connexion a la base de donnée réussite")
	return bd, nil
}

// ObtenirInfoUtilisateur est une fonction pour avoir les informations de l'utilisateur demandé
func ObtenirInfoUtilisateur(NomUtilisateur string) (int, string, string, string, int, string, error) {
	// Initialisation de demande avec une requette SQL pour recuperer les informations de l'utilisateur
	demande := Bd.QueryRow("SELECT id, prenom, nom, mail, age, icon FROM Utilisateurs WHERE pseudo = ?", NomUtilisateur)
	// Initialisation des variables id, prenom, nom, mail, age et icon
	var id int
	var prenom string
	var nom string
	var mail string
	var age int
	var icon string
	// On rentre les informations de demande dans les variables
	err := demande.Scan(&id, &prenom, &nom, &mail, &age, &icon)
	// S'il y a une erreur renvoie des variables contenant rien plus l'erreur
	if err != nil {
		return 0, "", "", "", 0, "", err
	}
	// Sinon envoyer les variables
	return id, prenom, nom, mail, age, icon, nil
}
