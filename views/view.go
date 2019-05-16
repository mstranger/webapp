package views

import (
	"html/template"
	"path/filepath"
)

var (
	LayoutDir   string = "views/layouts/"
	TemplateExt string = ".gohtml"
)

func NewView(layout string, files ...string) *View {
	// files = append(files,
	// 	"views/layouts/footer.gohtml",
	// 	"views/layouts/bootstrap.gohtml",
	// 	"views/layouts/navbar.gohtml",
	// )
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

type View struct {
	Template *template.Template
	Layout   string
}

// Render is used to render the view with the predefined layout.
func (v *View) Render(w http.ResponseWriter, data {}interface) error {
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}

// layoutFiles returs a slice of strings representing
// the layout files used in our application.
func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "*" + TemplateExt)
	if err != nil {
		panic(err)
	}
	return files
}
