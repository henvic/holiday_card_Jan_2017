package cardtemplate

import (
	"bytes"
	"html/template"
)

// Template structure
type Template struct {
	Subject  string
	template *template.Template
}

// GetSubject from a template
func (t *Template) GetSubject(name string) (string, error) {
	var buf bytes.Buffer
	var data = map[string]string{
		"name": name,
	}

	if t.template == nil {
		t.template = template.New("subject")
		t.template.Parse(t.Subject)
	}

	var err = t.template.Execute(&buf, data)
	return buf.String(), err
}
