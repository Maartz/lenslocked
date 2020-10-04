package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

type View struct {
	Template *template.Template
	Layout   string
}

var (
	LayoutDir   string = "views/layouts/"
	TemplateExt string = ".gohtml"
)

func NewView(layout string, files ...string) *View {
	files = append(files, layoutFiles()...)

	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View{
		Template: t,
		Layout:   layout,
	}
}

// Render is used to render the view with the predefined layout.
func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}

//layoutFiles return a slice of strings
//representing the layout files used in the application.
func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "*" + TemplateExt)
	if err != nil {
		panic(err)
	}
	return files
}
