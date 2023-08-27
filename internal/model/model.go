package model

type User struct {
	ID       int    `json:"-" db:"id"`
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Note struct {
	ID    int    `json:"id" db:"id"`
	Title string `json:"title" db:"title" binding:"required"`
	Text  string `json:"text" db:"text"`
}
