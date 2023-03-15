package main
import (
"fmt"
"html/template"

"net/http"
"strconv"

)
func (app *application)home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
		}
		// Include the navigation partial in the template files.
		files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/partials/nav.tmpl",
		"./ui/html/pages/home.tmpl",
		}
		ts, err := template.ParseFiles(files...)
		if err != nil {
		app.serveError(w, err)
		http.Error(w, "Internal Server Error", 500)
		return
		}
		err = ts.ExecuteTemplate(w, "base", nil)
		if err != nil {
		// app.errorLog.Print(err.Error())
		app.serveError(w, err)
		http.Error(w, "Internal Server Error", 500)
		}


	// w.Write([]byte("Hello from Snippetbox"))
}
func (app *application)snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		// http.NotFound(w, r)
		app.notFound(w)
	return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
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