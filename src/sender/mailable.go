package sender

import (
	"bytes"
	"errors"
	"html/template"
)

type Mailable struct {
	To       string                 `json:"to"`
	From     string                 `json:"from"`
	Subject  string                 `json:"subject"`
	Message  map[string]interface{} `json:"message"`
	Template string                 `json:"template"`
}

func (m *Mailable) content(tmplts map[string]*template.Template) (string, error) {
	return m.getTemplate(tmplts)
}

func (m *Mailable) getTemplate(tmplts map[string]*template.Template) (string, error) {
	var layoutBuf bytes.Buffer

	t, exists := tmplts[m.Template]
	if !exists {
		return "", errors.New("no such template")
	}
	if err := t.ExecuteTemplate(&layoutBuf, "layout", m.Message); err != nil {
		return "", err
	}
	return layoutBuf.String(), nil
}

func (m *Mailable) validate() error {
	if len(m.Template) == 0 {
		return errors.New("template must be present")
	}
	if len(m.To) == 0 {
		return errors.New("to must be present")
	}
	if len(m.Subject) == 0 {
		return errors.New("subject must be present")
	}
	if len(m.From) == 0 {
		return errors.New("from must be present")
	}
	return nil
}
