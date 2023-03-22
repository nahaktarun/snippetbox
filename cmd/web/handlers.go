package main
import (
	"errors"
"fmt"
// "html/template"

"net/http"
"strconv"

"snippetbox.tarunnahak.com/internal/models"

"github.com/julienschmidt/httprouter"
"strings"
"unicode/utf8"

)
func (app *application)home(w http.ResponseWriter, r *http.Request) {
	
		// panic("OOps, somthing went wrong");

		snippets , err := app.snippets.Latest()
		if err != nil{
			app.serveError(w, err)
			return
		}

		// for _ , snippet := range snippets {
		// 	fmt.Fprintf(w, "%+v\n",snippet)
		// }

		// // Include the navigation partial in the template files.
		// files := []string{
		// "./ui/html/base.tmpl",
		// "./ui/html/partials/nav.tmpl",
		// "./ui/html/pages/home.tmpl",
		// }
		// ts, err := template.ParseFiles(files...)
		// if err != nil {
		// app.serveError(w, err)
		// // http.Error(w, "Internal Server Error", 500)
		// return
		// }

		// // snippets
		// data := &templateData{
		// 	Snippets: snippets,
		// }

		// err = ts.ExecuteTemplate(w, "base", data)
		// if err != nil {
		// // app.errorLog.Print(err.Error())
		// app.serveError(w, err)
		// // http.Error(w, "Internal Server Error", 500)
		// }

		data := app.newTemplateData(r)
		data.Snippets = snippets

		app.render(w, http.StatusOK, "home.tmpl",data)


	// w.Write([]byte("Hello from Snippetbox"))
}
func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	
	
	params := httprouter.ParamsFromContext(r.Context())
	
	id, err := strconv.Atoi(params.ByName("id"))
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
	// files := []string{
	// "./ui/html/base.tmpl",
	// "./ui/html/partials/nav.tmpl",
	// "./ui/html/pages/view.tmpl",
	// }
	// ts, err := template.ParseFiles(files...)
	// if err != nil {
	// app.serveError(w, err)
	// return
	// }
	// // Create an instance of a templateData struct holding the snippet data.
	// data := &templateData{
	// Snippet: snippet,
	// }
	// // Pass in the templateData struct when executing the template.
	// err = ts.ExecuteTemplate(w, "base", data)
	// if err != nil {
	// app.serveError(w, err)
	// }


	data := app.newTemplateData(r)
	data.Snippet = snippet

	// use the new render helper
	app.render(w, http.StatusOK, "view.tmpl",data)

	}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {


	data := app.newTemplateData(r)

	app.render(w, http.StatusOK, "create.tmpl",data)

}

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request){
	

	// error handler
	err := r.ParseForm()
	if err != nil{
		app.clientError(w, http.StatusBadRequest)
		return
	}

	title := r.PostForm.Get("title")
	content := r.PostForm.Get("content")


	expires , err := strconv.Atoi(r.PostForm.Get("expires"))
	if err != nil{
		app.clientError(w, http.StatusBadRequest)
		return
	}

	fieldErrors := make(map[string]string)

	if strings.TrimSpace(title) == "" {
		fieldErrors["title"] = "This field cannot not be blank"
	}else if utf8.RuneCountInString(title) > 100{
		fieldErrors["title"] = "This field cannot be more than 100 characters long"
	}

	// check that the content value isn't blank
	if strings.TrimSpace(content) == ""{
		fieldErrors["content"] = "This field cannot be blank"
	}

	if expires != 1 && expires != 7 && expires != 365 {
		fieldErrors["expires"] = "This field must equal 1, 7 or 365"
	}
	if len(fieldErrors ) > 0{
		fmt.Fprint(w,fieldErrors)
		return
	}

	id, err := app.snippets.Insert(title, content,expires)
	if err != nil{
		app.serveError(w, err)
		return
	}



	
	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d",id),http.StatusSeeOther)
}