package sender

import (
	"fmt"
	"github.com/clanbeat/mailer/Godeps/_workspace/src/github.com/gin-gonic/gin/render"
	"html/template"
	"os"
	"path/filepath"
	"strings"
)

type (
	TemplateCache map[string]*template.Template

	Render struct {
		tmpl *template.Template
		data interface{}
	}

	fileInfo struct {
		Path         string
		TemplateName string
		Name         string
		Layout       string
	}
)

var currentCache TemplateCache

const layoutTemplate = "layout"

func BuildTemplateCache(path string) (TemplateCache, error) {
	if len(currentCache) > 0 {
		return currentCache, nil
	}

	tmplts := TemplateCache{}
	files, err := getFiles(path)
	if err != nil {
		return tmplts, err
	}
	for _, f := range files {
		if f.TemplateName != layoutTemplate {
			tmplts[f.TemplateName] = template.Must(
				template.New("").Funcs(funcMap).ParseFiles(f.Layout, f.Path),
			)
		}
	}
	currentCache = tmplts
	return tmplts, nil
}

func (t TemplateCache) Instance(name string, data interface{}) render.Render {
	return render.HTML{
		Template: t[name],
		Name:     layoutTemplate,
		Data:     data,
	}
}

func removeExtension(filename string) string {
	return strings.TrimSuffix(filename, filepath.Ext(filename))
}

func getFiles(searchDir string) ([]fileInfo, error) {
	var files []fileInfo
	err := filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !f.IsDir() {
			files = append(files, newFile(path, f))
		}
		return nil
	},
	)
	return files, err
}

func newFile(path string, f os.FileInfo) fileInfo {
	layout := fmt.Sprintf("%s%s.html", strings.TrimSuffix(path, f.Name()), layoutTemplate)
	return fileInfo{
		TemplateName: removeExtension(f.Name()),
		Name:         f.Name(),
		Path:         path,
		Layout:       layout,
	}
}
