package forum

type ErreurMessage struct {
	Message string
}

type Utilisateurs struct {
	ID     int
	Pseudo string
	Mdp    string
	Prenom string
	Nom    string
	Mail   string
	Age    int
	Icon   string
}

type Poste struct {
	ID          int
	Theme       string
	Titre       string
	Description string
	Creele      string
	CreePar     string
	Likes       int
	Dislikes    int
	Icon        string
	Coms        []Commentaire
}

type Commentaire struct {
	ID           int
	IdPost       int
	IdPseudo     int
	Contenu      string
	IconDuPseudo string
}

type Envoie struct {
	User Utilisateurs
	Post []Poste
}
