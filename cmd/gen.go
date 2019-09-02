package cmd

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"

	"go.zoe.im/gopkg/gopkg"
)

var (
	defaultOutputTarget = "docs"
	defaultConfigFile   = "gopkg.yaml"
)

type generator struct {
	gopkg *gopkg.Gopkg

	target     string
	configPath string
	debug      bool
}

func newGenerator() *generator {
	return &generator{
		gopkg:      gopkg.NewGopkg(),
		target:     defaultOutputTarget,
		configPath: defaultConfigFile,
	}
}

func (gen *generator) init() error {
	// read data from yaml
	data, err := ioutil.ReadFile(gen.configPath)
	if err != nil {
		return err
	}

	// TODO: use x.Unmarshal() have not need to check if we are json or yaml
	err = yaml.Unmarshal(data, gen.gopkg)
	if err != nil {
		return err
	}

	return gen.gopkg.Init()
}

func (gen *generator) run() error {

	var err = gen.init()
	if err != nil {
		return err
	}

	for _, f := range gen.gopkg.Files {
		if gen.debug {
			log.Println("[gopkg] [gen] generate file:", f.Path)
		}
		err := f.Write(gen.target)
		if err != nil {
			log.Printf("[gopkg] [gen] generate file: %s, error: %s\n", f.Path, err)
		}
	}

	for _, pk := range gen.gopkg.Packages {
		for _, f := range pk.Files {
			if gen.debug {
				log.Println("[gopkg] [gen] generate file:", f.Path)
			}
			err = f.Write(gen.target)
			if err != nil {
				log.Printf("[gopkg] [gen] generate file: %s, error: %s\n", f.Path, err)
			}
		}
	}

	return nil
}
