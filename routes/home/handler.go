package home

import (
	"embed"
	"html/template"
	// "log"
	//
	// "github.com/niconiahi/gig.dance/packages/database"
	"github.com/niconiahi/gig.dance/packages/html"
)

type Loader struct {
	Message string
}

type Handler struct{}

type Data struct {
	Head   html.Head
	Loader Loader
}

func (h *Handler) GetData() Data {
	// d := database.GetDatabase()
	// defer d.Close()
	//
	// rows, err := d.Query("SELECT * FROM user")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer rows.Close()

	return Data{
		Head: html.Head{
			Title: "Home page",
		},
		Loader: Loader{
			Message: "Made in Go with only standard library",
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
