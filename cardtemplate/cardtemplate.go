package cardtemplate

import (
	"bytes"
	"html/template"
)

// Template structure
type Template struct {
	template *template.Template
}

// New creates a new card template
func (t Template) New(subject string) {
	t.template = template.New("subject")
}

// GetSubject from a template
func (t *Template) GetSubject(name string) (string, error) {
	var buf bytes.Buffer
	var data = map[string]string{
		"foo": "bar",
	}
	var err = t.template.Execute(&buf, data)
	return buf.String(), err
}
