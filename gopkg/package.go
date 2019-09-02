package gopkg

import (
	"path/filepath"
	"strings"
)

var (
	defaulType    = "git"
	defaultSchema = "https"
)

// Package contains all packages
type Package struct {
	Name      string               `yaml:"name"`
	PkgURL    string               `yaml:"pkg"`
	Type      string               `yaml:"type"`
	Schema    string               `yaml:"schema"`
	Path      string               `yaml:"path"`
	Sub       []*Package           `yaml:"sub"`
	Templates map[string]*Template `yaml:"templates"`

	// loaded files store here
	Files []*File

	URL string `yaml:"url"`

	// prvious package
	previous *Package
}

// Init create file
func (pkg *Package) Init(gopkg *Gopkg) error {

	if pkg.Schema == "" {
		pkg.Schema = "https"
	}

	if pkg.Type == "" {
		pkg.Type = "git"
	}

	if pkg.Path == "" {
		pkg.Path = pkg.Name
	} else if pkg.Name == "" {
		// cut from path
		var parts = strings.Split(pkg.Path, "/")
		pkg.Name = parts[len(parts)-1]
	} else {
		// return errors.New("name/path must need at least one")
	}

	// generate url and pkg url
	if pkg.previous != nil {
		// append if we are a child package
		if strings.Index(pkg.Path, "://") < 0 {
			pkg.Path = pkg.previous.Path + "/" + pkg.Path
		}

		if pkg.PkgURL == "" {
			pkg.PkgURL = filepath.Join(pkg.previous.PkgURL, pkg.Name)
		}

	} else {
		// build if we are a root package
		pkg.URL = pkg.Path
		if !strings.HasPrefix(pkg.URL, "http") {
			pkg.URL = pkg.Schema + "://" + pkg.Path
		}

		if pkg.PkgURL == "" {
			if gopkg.Base == "" {
				pkg.PkgURL = gopkg.Host + "/" + pkg.Name
			} else {
				pkg.PkgURL = gopkg.Host + "/" + gopkg.Base + "/" + pkg.Name
			}
		}
	}

	// load from global
	var tpls = make(map[string]*Template)

	// laod from global
	for k, t := range gopkg.PkgTemplates {
		tpls[k] = t
	}

	// load from self pacakge
	for k, t := range pkg.Templates {
		tpls[k] = t
	}

	// render files from templates
	for _, t := range tpls {
		// init template
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

	// call the sub package init
	for _, p := range pkg.Sub {
		// ignore sub package init error
		_ = p.Init(gopkg)
	}

	return nil
}
