package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Pranaenae/lets-go/internal/models"
	"github.com/julienschmidt/httprouter"
)

// w is a r is a pointer to a struct
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// if r.URL.Path != "/" {
	// 	app.notFound(w)
	// 	return
	// }

	snippets, err := app.snippets.Latest()

	if err != nil {
		app.serverError(w, err)
		return
	}

	data := app.newTemplateData(r)
	data.Snippets = snippets

	app.render(w, http.StatusOK, "home.html", data)

}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {

	params := httprouter.ParamsFromContext(r.Context())

	// id, err := strconv.Atoi(r.URL.Query().Get("id"))
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	snippet, err := app.snippets.Get(id)

	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	data := app.newTemplateData(r)
	data.Snippet = snippet

	app.render(w, http.StatusOK, "view.html", data)
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display the form form for creating a new snippet..."))
}

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	// if r.Method != "POST" {
	// 	w.Header().Set("Allow", http.MethodPost)
	// 	app.clientError(w, http.StatusMethodNotAllowed)
	// 	return
	// }

	title := "0 Snail"
	content := "0 Snail\nClimb Mount Fuji,\nBut slowly,slowly!\n\n - Kobayashi Issa"
	expires := 7

	id, err := app.snippets.Insert(title, content, expires)

	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}
