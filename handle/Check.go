package handle

import (
	"data"
	"database/sql"
	"fmt"
)

func Check(username string, password string, db *sql.DB) (bool, error) {
	// Récupérer le hash du mot de passe enregistré dans la base de données pour cet utilisateur
	row := db.QueryRow("SELECT password FROM users WHERE username = ?", username)
	var passwordHash string
	err := row.Scan(&passwordHash)
	if err != nil {
		if err == sql.ErrNoRows {
			// L'utilisateur n'existe pas dans la base de données
			fmt.Println("user donsn't exist")
			return false, nil
		}
		// Erreur lors de l'exécution de la requête SQL
		fmt.Println("erreur de requête")
		return false, err
	}

	// Vérifier que le mot de passe fourni correspond au hash stocké dans la base de données
	ok := data.CheckPasswordHash(password, passwordHash)
	if ok == true {
		// Le mot de passe fourni ne correspond pas au hash stocké dans la base de données
		return true, nil
	}

	// Les identifiants fournis sont valides
	return false, nil
}
