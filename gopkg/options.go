package gopkg

import "strings"

var (
	defaultHost = "github.com"
)

type Option func(inst *Instance)

func Index(index string) Option {
	return func(inst *Instance) {
		inst.index = index
	}
}

func Prefix(hostwithuser string) Option {
	var host, user string
	keys := strings.Split(hostwithuser, "/")
	if len(keys) == 1 {
		host = defaultHost
		user = keys[0]
	} else {
		host = keys[0]
		user = keys[1]
	}
	return func(inst *Instance) {
		inst.repoPrefix = append(inst.repoPrefix, RepoPrefix{"https", host, user})
	}
}
