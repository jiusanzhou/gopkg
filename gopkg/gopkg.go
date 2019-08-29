package gopkg

import (
	"net/http"
)

var (
	defaultPkgIndexTpl = &Template{
		Name: "{{.Name}}/index.html",
		Content: `<!DOCTYPE html>
<html lang="en">
	<head>
		<meta charset="utf-8" />
		<meta name="go-import" content="{{.PkgURL}} {{.Type}} {{.URL}}" />
		<meta http-equiv="refresh" content="0; URL='{{.URL}}'" />
	</head>
</html>`,
	}

	defaultPkgReadmeTpl = &Template{
		Name:    "{{.Name}}/README.md",
		Content: `## {{.Name}}`,
	}
)

// Gopkg is the main struct for generator or server
type Gopkg struct {
	Name        string               `yaml:"name"`
	Description string               `yaml:"description"`
	Host        string               `yaml:"host"`
	Base        string               `yaml:"base"`
	Theme       string               `yaml:"theme"`
	Packages    []*Package           `yaml:"packages"`
	Templates   map[string]*Template `yaml:"templates"`

	// for self template
}

// NewGopkg return a new gopkg
func NewGopkg() *Gopkg {
	return &Gopkg{
		Templates: map[string]*Template{
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

	return nil
}

// Generate htmls
func (pkg *Gopkg) Generate(target string) error {

	// take files from packages and write to file
	for _, pk := range pkg.Packages {
		for _, f := range pk.Files {
			err := f.Write(target)
			if err != nil {
				return err
			}
		}
	}

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
