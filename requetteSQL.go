package main

import (
	"data"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type User struct {
	ID        int
	Username  string
	Password  string
	Firstname string
	Lastname  string
	Email     string
	Age       int
}

func Check(username string, password string, db *sql.DB) (bool, error) {
	/*// Récupérer le hash du mot de passe enregistré dans la base de données pour cet utilisateur
	row := db.QueryRow("SELECT password FROM users WHERE username = ?", username)
	var passwordHash string
	err := row.Scan(&passwordHash)
	if err != nil {
		if err == sql.ErrNoRows {
			// L'utilisateur n'existe pas dans la base de données
			return false, nil
		}
		// Erreur lors de l'exécution de la requête SQL
		return false, err
	}

	// Vérifier que le mot de passe fourni correspond au hash stocké dans la base de données
	ok := data.CheckPasswordHash(password, passwordHash)
	if ok {
		// Le mot de passe fourni ne correspond pas au hash stocké dans la base de données
		return true, nil
	}

	// Les identifiants fournis sont valides
	return false, nil*/
	fmt.Println(username, password)
	row := db.QueryRow("SELECT password FROM users WHERE username = ?", username)
	var passwordHash string
	err := row.Scan(&passwordHash)
	fmt.Println("passwordhash = ", passwordHash)
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
		fmt.Println("bon mdp")
		return true, nil
	}

	// Les identifiants fournis sont valides
	return false, nil
}

func GetUserData(username string, db *sql.DB) (int, string, string, string, int, error) {
	row := db.QueryRow("SELECT id, firstname, lastname, mail, age FROM users WHERE username = ?", username)
	var id int
	var firstname string
	var lastname string
	var mail string
	var age int
	err := row.Scan(&id, &firstname, &lastname, &mail, &age)
	if err != nil {
		return 0, "", "", "", 0, err
	}
	return id, firstname, lastname, mail, age, nil
}

func main() {
	db, err := sql.Open("sqlite3", "./database/db.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	id, firstname, lastname, email, age, err := GetUserData("erwan", db)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(id, firstname, lastname, email, age)
	/*result, err := Check("test24", "test", db)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)*/
}

// "SELECT info FROM users WHERE username = ?", info, username

/*
tmpl := template.Must(template.ParseFiles("./templates/inscription.html"))
	db, err := sql.Open("sqlite3", "./database/db.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// Vérifier si l'utilisateur est déjà connecté
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if _, ok := session.Values["username"].(string); ok {
		// Rediriger l'utilisateur vers la page de compte s'il est déjà connecté
		http.Redirect(w, r, "/account", http.StatusSeeOther)
		return
	}
	if r.Method != "POST" {
		message := "Entrez vos identifiants !"
		Message := Message{
			ProcessMessage: message,
		}
		tmpl.Execute(w, Message)
	}
	// Récupération des données du formulaire
	username := r.FormValue("username")
	password := r.FormValue("password")
	taken, err := Check(username, password)
	fmt.Println(taken)
	if err != nil {
		//gestion de l'erreur
		fmt.Println(err)
	}
	if taken == false {
		// Mauvais password
		messageerror := "Mauvais username ou password"
		Message := Message{
			ProcessMessage: messageerror,
		}
		tmpl.Execute(w, Message)

		return
	} else {
		//bon password
		session.Values["username"] = username
		session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/account", http.StatusSeeOther)
	}
	// Vérification de la méthode HTTP

// Exécute la première requête
	rows, err := db.Query("SELECT * FROM user WHERE age > ?", 20)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	// Parcours les résultats de la première requête
	for rows.Next() {
		var id int
		var name string
		var age int
		err = rows.Scan(&id, &name, &age)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", id, name, age)
	}

	// Exécute la deuxième requête
	rows2, err := db.Query("SELECT COUNT(*) FROM user WHERE age > ?", 20)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows2.Close()

	// Parcours les résultats de la deuxième requête
	for rows2.Next() {
		var count int
		err = rows2.Scan(&count)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Count: %d\n", count)
	}*/
/*
	username := "alice"
	password := "25"
	_, err = db.Exec("INSERT INTO user (name, age) VALUES (?, ?)", username, password)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Utilisateur ajouté avec succès !")
*/
