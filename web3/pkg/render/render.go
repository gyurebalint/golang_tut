package render

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
	models "web3/models"
	helper "web3/pkg/helpers"
)

var tmplCache = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter,
	t string, pd *models.PageData) {
	var tmpl *template.Template
	var err error

	_, inMap := tmplCache[t]
	if !inMap {
		err = makeTemplateCache(t)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println("Template in cache")
	}

	tmpl = tmplCache[t]
	err = tmpl.Execute(w, pd)
	if err != nil {
		fmt.Println(err)
	}
}

func makeTemplateCache(t string) error {

	specificTemplate := path.Join(helper.GetTemplatesDirectory(), t)
	baseTemplate := path.Join(helper.GetTemplatesDirectory(), "base.layout.tmpl")

	fmt.Println(specificTemplate)

	templates := []string{
		specificTemplate,
		baseTemplate,
	}
	fmt.Println()
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}
	tmplCache[t] = tmpl
	return nil
}
