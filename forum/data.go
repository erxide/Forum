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
	_, err = bd.Exec("SELECT * FROM Utilisateurs")
	if err != nil {
		// Si la table n'existe pas, la créer
		_, err := bd.Exec(`CREATE TABLE Utilisateurs (
    id INTEGER PRIMARY KEY,
    pseudo TEXT NOT NULL,
    mdp TEXT NOT NULL,
    prenom TEXT NOT NULL,
    nom TEXT NOT NULL,
    mail TEXT NOT NULL,
    age INTEGER NOT NULL,
    icon TEXT
);`)
		if err != nil {
			fmt.Println(err)
			return bd, err
		}
		fmt.Println("Table Utilisateurs créée avec succès.")
	} else {
		fmt.Println("La table Utilisateurs existe déjà.")
	}
	_, err = bd.Exec("SELECT * FROM Postes")
	if err != nil {
		// Si la table n'existe pas, la créer
		_, err := bd.Exec(`CREATE TABLE Postes (
  id INTEGER PRIMARY KEY,
  theme TEXT NOT NULL,
  titre TEXT NOT NULL,
  description TEXT NOT NULL,
  cree_le TEXT NOT NULL,
  cree_par TEXT NOT NULL,
  likes INT DEFAULT 0,
  dislikes INT DEFAULT 0
  )`)
		if err != nil {
			fmt.Println(err)
			return bd, err
		}
		fmt.Println("Table Postes créée avec succès.")
	} else {
		fmt.Println("La table Postes existe déjà.")
	}

	return bd, err
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

func ObtenirInfoUtilisateurID(idUtilisateur int) (int, string, string, string, int, string, error) {
	// Initialisation de demande avec une requette SQL pour recuperer les informations de l'utilisateur
	demande := Bd.QueryRow("SELECT id, prenom, nom, mail, age, icon FROM Utilisateurs WHERE id = ?", idUtilisateur)
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

func ObtenirInfoPoste(ID string) (string, string, string, string, string, int, int, error) {
	// Initialisation de demande avec une requette SQL pour recuperer les informations de l'utilisateur
	demande := Bd.QueryRow("SELECT theme, titre, description, cree_le, cree_par, likes, dislikes FROM Postes WHERE id = ?", ID)
	// Initialisation des variables id, prenom, nom, mail, age et icon
	var theme, titre, description, cree_le, cree_par string
	var likes, dislikes int
	// On rentre les informations de demande dans les variables
	err := demande.Scan(&theme, &titre, &description, &cree_le, &cree_par, likes, dislikes)
	// S'il y a une erreur renvoie des variables contenant rien plus l'erreur
	if err != nil {
		return "", "", "", "", "", 0, 0, err
	}
	// Sinon envoyer les variables
	return theme, titre, description, cree_le, cree_par, likes, dislikes, nil
}
