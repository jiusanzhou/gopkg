package gopkg

import (
	"net/http"
)

var (
	defaultReadmeTpl = &Template{
		Name: "README.md",
		Content: `| PACKAGE | GODOC | SOURCE |
|---------|-------|--------|{{ range .Packages }}
| {{.Name}} | [![GoDoc](https://godoc.org/{{.Path}}?status.svg)](https://godoc.org/{{.Path}}) | [{{.PkgURL}}](https://{{.PkgURL}}) |{{ end }}

---

> Generated using [gopkg](https://go.zoe.im/gopkg) with :heart:.`,
	}
	defaultCnameTpl = &Template{
		Name:    "CNAME",
		Content: `{{.Host}}`,
	}

	defaultPkgIndexTpl = &Template{
		Name: "{{.Name}}/index.html",
		Content: `<!DOCTYPE html>
<html>
	<head>
		<meta name="go-import" content="{{.PkgURL}} {{.Type}} {{.URL}}" />
		<meta name="go-source" content="{{ .PkgURL }} {{.URL}} {{.URL}}/tree/master{/dir} {{.URL}}/tree/master{/dir}/{file}#L{line}">
		<meta http-equiv="refresh" content="0; url=https://{{.URL}}">
	</head>
	<body>
		Nothing to see here. Please <a href="{{.URL}}">move along</a>.
	</body>
</html>`,
	}

	defaultPkgReadmeTpl = &Template{
		Name:    "{{.Name}}/README.md",
		Content: `## {{.Name}}`,
	}
)

// Gopkg is the main struct for generator or server
type Gopkg struct {
	Name         string                 `yaml:"name"`
	Description  string                 `yaml:"description"`
	Host         string                 `yaml:"host"`
	Base         string                 `yaml:"base"`
	Theme        string                 `yaml:"theme"`
	Packages     []*Package             `yaml:"packages"`
	Templates    map[string]*Template   `yaml:"templates"`
	PkgTemplates map[string]*Template   `yaml:"pkg_templates"`
	Metadata     map[string]interface{} `yaml:"metadata"`

	// loaded files store here
	Files []*File
}

// NewGopkg return a new gopkg
func NewGopkg() *Gopkg {
	return &Gopkg{
		Templates: map[string]*Template{
			"readme": defaultReadmeTpl,
			"cname":  defaultCnameTpl,
		},
		PkgTemplates: map[string]*Template{
			"index":  defaultPkgIndexTpl,
			"readme": defaultPkgReadmeTpl,
		},
	}
}

// Init the package
func (pkg *Gopkg) Init() error {

	// load tempaltes from theme
	// TODO: load tempaltes from theme

	// init all of the packages
	for _, pk := range pkg.Packages {
		err := pk.Init(pkg)
		if err != nil {
			return err
		}
	}

	// init self
	if pkg.Templates == nil {
		pkg.Templates = make(map[string]*Template)
	}

	for _, t := range pkg.Templates {
		err := t.Init()
		if err != nil {
			return err
		}

		f, err := t.Render(pkg)
		if err != nil {
			return err
		}
		pkg.Files = append(pkg.Files, f)
	}

	return nil
}

// Generate htmls, TODO: move generate files at here
func (pkg *Gopkg) Generate(target string) error {

	return nil
}

// NewHandler return a handler
func (pkg *Gopkg) NewHandler(allowWild bool, next http.Handler) http.Handler {

	// TODO: init handler for package

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: handler request

		// parse the to

		// if allow wild we reurn any package
		// no matter exits or not
		if allowWild {

			return
		}

		// if not wild, we need to take argument from laoded

		if next != nil {
			next.ServeHTTP(w, r)
		}
	})
}
