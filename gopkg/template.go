package gopkg

import (
	"bytes"
	"html/template"
	"os"
	"path/filepath"

	"go.zoe.im/x"
)

// Template for generate html
type Template struct {
	Name    string `yaml:"name"`
	Content string `yaml:"content"`

	tplName    *template.Template
	tplContent *template.Template
}

// Init template
func (tpl *Template) Init() error {
	var err error
	tpl.tplName, err = template.New(tpl.Name + "-name").Parse(tpl.Name)
	if err != nil {
		return err
	}

	tpl.tplContent, err = template.New(tpl.Name + "-content").Parse(tpl.Content)
	if err != nil {
		return err
	}

	return nil
}

// Render render data
func (tpl *Template) Render(data interface{}) (*File, error) {

	var buf bytes.Buffer
	var err error

	err = tpl.tplName.Execute(&buf, data)
	if err != nil {
		return nil, err
	}

	var fl = new(File)
	fl.Path = x.Bytes2Str(buf.Bytes())

	// reset the buffer
	var bufContent bytes.Buffer

	err = tpl.tplContent.Execute(&bufContent, data)
	if err != nil {
		return nil, err
	}

	fl.Data = bufContent.Bytes()
	return fl, nil
}

// File contains file name and bytes of fiole content
type File struct {
	Path string
	Data []byte
}

func (fl *File) Write(target string) error {
	var fp = filepath.Join(target, fl.Path)

	// to mkdir if not exits, ignore error
	_ = os.MkdirAll(filepath.Dir(fp), 0755)

	f, err := os.OpenFile(fp, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}

	_, err = f.Write(fl.Data)
	return err
}
