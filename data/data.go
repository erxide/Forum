package data

import "database/sql"

func GetUserData(username string, db *sql.DB) (int, string, string, string, int, string, error) {
	row := db.QueryRow("SELECT id, firstname, lastname, mail, age, icon FROM users WHERE username = ?", username)
	var id int
	var firstname string
	var lastname string
	var mail string
	var age int
	var icon string
	err := row.Scan(&id, &firstname, &lastname, &mail, &age, &icon)
	if err != nil {
		return 0, "", "", "", 0, "", err
	}
	return id, firstname, lastname, mail, age, icon, nil
}
