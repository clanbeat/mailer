package sender

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"math"
)

type Mailable struct {
	To       string                 `json:"to"`
	From     string                 `json:"from"`
	Subject  string                 `json:"subject"`
	Message  map[string]interface{} `json:"message"`
	Template string                 `json:"template"`
}

var funcMap = template.FuncMap{
	"addReviewSeparator": addReviewSeparator,
	"isMiddle":           isMiddle,
	"hasMiddle":          hasMiddle,
	"colorForList":       colorForList,
	"lessOrMore":         lessOrMore,
}

func addReviewSeparator(i int) bool {
	return i != 0
}

func hasMiddle(s int, max int) bool {
	return s == max
}

func isMiddle(i int, total int) bool {
	return i == total/2
}

func colorForList(i int, hasMiddle bool) string {
	green := "#7ed220"
	red := "#f8503b"

	if !hasMiddle || i < 3 {
		return green
	}
	return red
}

func lessOrMore(d float64) string {
	if d > 0 {
		return fmt.Sprintf("%d more", int64(d))
	}
	return fmt.Sprintf("%d less", int64(math.Abs(d)))
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
