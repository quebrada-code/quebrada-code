package resources

import (
	"embed"
	"html/template"
	"io/fs"
)

type TemplateEmail string

const (
	WelcomeTemplate           TemplateEmail = "welcome.html"
	VerificationEmailTemplate TemplateEmail = "verification_email.html"
)

var (
	//go:embed verification_email.html welcome.html
	files     embed.FS
	templates map[string]*template.Template
)

func LoadTemplates() error {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}
	tmplFiles, err := fs.ReadDir(files, ".")
	if err != nil {
		return err
	}

	for _, tmpl := range tmplFiles {
		if tmpl.IsDir() {
			continue
		}

		pt, err := template.ParseFS(files, tmpl.Name())
		if err != nil {
			return err
		}

		templates[tmpl.Name()] = pt
	}
	return nil
}

func (t TemplateEmail) GetTemplate() *template.Template {
	return templates[string(t)]
}
