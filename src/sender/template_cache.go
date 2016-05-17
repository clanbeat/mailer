package sender

import (
	"fmt"
	"github.com/clanbeat/mailer/Godeps/_workspace/src/github.com/gin-gonic/gin/render"
	"html/template"
	"io/ioutil"
	"path/filepath"
)

type TemplateCache map[string]*template.Template

type Render struct {
	tmpl *template.Template
	data interface{}
}

const layoutTemplate = "layout"

func BuildTemplateCache(path string) (TemplateCache, error) {
	tmplts := make(TemplateCache)
	fnames, err := files(path)
	if err != nil {
		return tmplts, err
	}

	for _, fname := range fnames {
		if fname != layoutTemplate {
			tmplts[fname] = template.Must(parseTemplate(path, fname))
		}
	}
	return tmplts, nil
}

func (t TemplateCache) Instance(name string, data interface{}) render.Render {
	return render.HTML{
		Template: t[name],
		Name:     layoutTemplate,
		Data:     data,
	}
}

func parseTemplate(path, fname string) (*template.Template, error) {
	layout := fmt.Sprintf("%s/layout.html", path)
	content := fmt.Sprintf("%s/%s.html", path, fname)
	return template.ParseFiles(layout, content)
}

func files(path string) ([]string, error) {
	var names []string
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return names, err
	}
	for _, f := range files {
		names = append(names, fileWithoutExtension(f.Name()))
	}
	return names, nil
}

func fileWithoutExtension(filename string) string {
	fName := filepath.Base(filename)
	extName := filepath.Ext(filename)
	return fName[:len(fName)-len(extName)]
}
