package main

import (
	"html/template"
	"path/filepath"
	"time"

	"github.com/Pranaenae/lets-go/internal/models"
)

type templateData struct {
	Snippet     *models.Snippet
	Snippets    []*models.Snippet
	CurrentYear int
	Form        any
}

var functions = template.FuncMap{
	"humanDate": humanDate,
}

func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./ui/html/pages/*")

	if err != nil {
		return nil, err
	}

	for _, page := range pages {

		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles("./ui/html/base.html")

		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob("./ui/html/partials/*.html")

		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		cache[name] = ts

	}

	return cache, nil
}
