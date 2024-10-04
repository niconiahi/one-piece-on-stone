package home

import (
	"embed"
	"fmt"
	"html/template"
	"log"

	"github.com/niconiahi/one-piece-on-stone/packages/database"
	"github.com/niconiahi/one-piece-on-stone/packages/html"
)

type Loader struct {
	Message string
	Users   []User
}

type Handler struct{}

type Data struct {
	Head   html.Head
	Loader Loader
}

type User struct {
	ID      int
	Name    string
	Surname string
}

func (h *Handler) GetData() Data {
	d := database.GetDatabase()
	defer d.Close()

	user, err := d.Exec("INSERT INTO user (name, surname) VALUES (?, ?)", "nono", "lorenzo")
	if err != nil {
		log.Fatal(err)
	}
	_, err = user.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	rows, err := d.Query("SELECT * FROM user")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		err := rows.Scan(&u.ID, &u.Name, &u.Surname)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, u)
	}
	fmt.Printf("users amount: %d\n", len(users))

	return Data{
		Head: html.Head{
			Title: "Home page",
		},
		Loader: Loader{
			Message: "Made in Go with only standard library",
			Users:   users,
		},
	}
}

func (h *Handler) GetFiles() []string {
	return []string{
		"root.html",
		"route.html",
	}
}

//go:embed route.html root.html
var files embed.FS

func (h *Handler) GetTemplate() (*template.Template, error) {
	return template.ParseFS(files, h.GetFiles()...)
}
