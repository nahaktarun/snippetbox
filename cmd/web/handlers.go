package main
import (
	"errors"
"fmt"
"html/template"

"net/http"
"strconv"

"snippetbox.tarunnahak.com/internal/models"

)
func (app *application)home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
		}

		snippets , err := app.snippets.Latest()
		if err != nil{
			app.serveError(w, err)
			return
		}

		for _ , snippet := range snippets {
			fmt.Fprintf(w, "%+v\n",snippet)
		}

		// Include the navigation partial in the template files.
		// files := []string{
		// "./ui/html/base.tmpl",
		// "./ui/html/partials/nav.tmpl",
		// "./ui/html/pages/home.tmpl",
		// }
		// ts, err := template.ParseFiles(files...)
		// if err != nil {
		// app.serveError(w, err)
		// http.Error(w, "Internal Server Error", 500)
		// return
		// }
		// err = ts.ExecuteTemplate(w, "base", nil)
		// if err != nil {
		// // app.errorLog.Print(err.Error())
		// app.serveError(w, err)
		// http.Error(w, "Internal Server Error", 500)
		// }


	// w.Write([]byte("Hello from Snippetbox"))
}
func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
	app.notFound(w)
	return
	}
	snippet, err := app.snippets.Get(id)
	if err != nil {
	if errors.Is(err, models.ErrNoRecords) {
	app.notFound(w)
	} else {
	app.serveError(w, err)
	}
	return
	}
	files := []string{
	"./ui/html/base.tmpl",
	"./ui/html/partials/nav.tmpl",
	"./ui/html/pages/view.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
	app.serveError(w, err)
	return
	}
	// Create an instance of a templateData struct holding the snippet data.
	data := &templateData{
	Snippet: snippet,
	}
	// Pass in the templateData struct when executing the template.
	err = ts.ExecuteTemplate(w, "base", data)
	if err != nil {
	app.serveError(w, err)
	}
	}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		// http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		app.clientError(w, http.StatusMethodNotAllowed)
	return
	}

	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
	expires := 7

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil{
		app.serveError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/view?id=%d",id),http.StatusSeeOther)
}