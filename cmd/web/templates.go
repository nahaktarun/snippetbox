package main
import (
	"snippetbox.tarunnahak.com/internal/models"
	"html/template"
	"path/filepath"

)
// Define a templateData type to act as the holding structure for
// any dynamic data that we want to pass to our HTML templates.
// At the moment it only contains one field, but we'll add more
// to it as the build progresses.
type templateData struct {
Snippet *models.Snippet
Snippets []*models.Snippet
}

// caching the templates
func newTemplateCache()(map[string]*template.Template, error){

	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./ui/html/pages/*.tmpl")
	if err != nil{
		return nil, err
	}

	for _, page := range pages{
		name := filepath.Base(page)


		// parse the base template file into a template set
		ts, err := template.ParseFiles("./ui/html/base.tmpl")
		if err != nil{
			return nil, err
		}

		// call ParseGlob() *on this template set* to add any partials
		ts, err = ts.ParseGlob("./ui/html/partials/*.tmpl")

		if err != nil{
			return nil, err
		}

		// call ParseFiles() *on this template set* to add the page template
		ts, err = ts.ParseFiles(page)
		if err != nil{
			return nil, err
		}

		// files := []string{
		// 	"./ui/html/base.tmpl",
		// 	"./ui/html/partials/nav.tmpl",
		// 	page,
		// }

		// ts, err := template.ParseFiles(files...)
		// if err != nil{
		// 	return nil, err
		// }

		cache[name] = ts
	}
	return cache, nil
}