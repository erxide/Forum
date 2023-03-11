package main

import (
	"database/sql"
	"handle/handle"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

// User représente un utilisateur
type User struct {
	ID           int
	Username     string
	Password     string
	PasswordHash string
}

func main() {
	db, err := sql.Open("sqlite3", "./database/db.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", fs))
	http.HandleFunc("/", handleSlash)
	http.HandleFunc("/accueil", handleHome)
	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/register", handleRegister)
	http.HandleFunc("/logout", handleLogout)
	http.HandleFunc("/account", handleAccount)
	http.HandleFunc("/changepassword", handleChangePassword)
	http.HandleFunc("/changeicon", handleChangeIcon)
	log.Println("Serveur démarré sur http://localhost:8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	handle.Home(w, r)
}

// handleLogin gère la page de connexion
func handleLogin(w http.ResponseWriter, r *http.Request) {
	handle.Login(w, r)
}

func handleRegister(w http.ResponseWriter, r *http.Request) {
	handle.Register(w, r)
}

func handleSlash(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/accueil", http.StatusSeeOther)
}

func handleLogout(w http.ResponseWriter, r *http.Request) {
	handle.Logout(w, r)
}

func handleAccount(w http.ResponseWriter, r *http.Request) {
	handle.Account(w, r)
}

func handleChangePassword(w http.ResponseWriter, r *http.Request) {
	handle.ChangePassword(w, r)
}

func handleChangeIcon(w http.ResponseWriter, r *http.Request) {
	handle.ChangeIcon(w, r)
}
