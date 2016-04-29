package sender

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"path/filepath"
)

type templateCache map[string]*template.Template

const layoutTemplate = "layout"

func buildTemplateCache(path string) (templateCache, error) {
	tmplts := make(templateCache)
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
