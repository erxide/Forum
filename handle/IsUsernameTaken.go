package handle

import "database/sql"

func isUsernameTaken(db *sql.DB, username string) (bool, error) {
	// Requête SELECT pour récupérer les noms d'utilisateurs existants
	rows, err := db.Query("SELECT username FROM users")
	if err != nil {
		return false, err
	}
	defer rows.Close()

	// Parcours des résultats pour vérifier si le nom d'utilisateur est déjà utilisé
	for rows.Next() {
		var existingUsername string
		if err := rows.Scan(&existingUsername); err != nil {
			return false, err
		}
		if existingUsername == username {
			return true, nil
		}
	}

	// Si on est arrivé ici, c'est que le nom d'utilisateur n'est pas déjà pris
	return false, nil
}
