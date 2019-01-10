package gopkg // import "go.zoe.im/gopkg/gopkg"

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"
	"strings"
	"time"
)

var (
	STAT_OK = []byte("ok")
)

type RepoPrefix struct {
	Scheme string
	Host   string
	User   string
}

func (rp RepoPrefix) String() string {
	return fmt.Sprintf("%s://%s/%s", rp.Scheme, rp.Host, rp.User)
}

type Repo struct {
	RepoPrefix RepoPrefix
	Pkg        string
	Version    string
}

func (r Repo) String() string {
	return fmt.Sprintf("%s/%s", r.RepoPrefix.String(), r.Pkg)
}

type stat struct {
	Started int64  `json:"started"`
	Memory  string `json:"memory"`
	Threads int    `json:"threads"`
	GC      string `json:"gc_pause"`

	// TODO:
	last_refresh int64
}

func (s *stat) refresh() {

	var mstat runtime.MemStats
	runtime.ReadMemStats(&mstat)
	s.Threads = runtime.NumGoroutine()
	s.Memory = fmt.Sprintf("%.2fmb", float64(mstat.Alloc)/float64(1024*1024))
	s.GC = fmt.Sprintf("%.3fms", float64(mstat.PauseTotalNs)/(1000*1000))
}

// NewInstance ...
func NewInstance(opts ...Option) *Instance {
	inst := &Instance{
		stat: &stat{
			Started: time.Now().Unix(),
		},
	}

	for _, o := range opts {
		o(inst)
	}

	if inst.index == "" {
		inst.index = "https://zoe.im"
	}

	if len(inst.repoPrefix) == 0 {
		inst.repoPrefix = []RepoPrefix{{"https", "github.com", "jiusanzhou"}}
	}

	fmt.Printf("[gopkg] serve gopkg with prefix: %s\n", inst.repoPrefix)
	return inst
}

// Instance is a struct to store everything.
type Instance struct {
	stat *stat
	// static file or location to redirect
	index string
	// RepoPrefix
	repoPrefix []RepoPrefix
}

func (inst *Instance) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// handle index url
	if r.URL.Path == "/" {
		// TODO: root
		w.Header().Set("Location", inst.index)
		w.WriteHeader(http.StatusTemporaryRedirect)
		return
	}

	// handle favicon
	if r.URL.Path == "/favicon.ico" {
		// TODO:
		return
	}

	// handle runtime stat
	if r.URL.Path == "/_stat" {
		data, _ := json.Marshal(inst.stat)
		w.Write(data)
		w.Header().Set("Content-Type", "application/json")
		return
	}

	// handle health check
	if r.URL.Path == "/_health_check" {
		w.Write(STAT_OK)
		return
	}

	// we need to handle repo
	keys := strings.Split(r.URL.Path[1:], "/")

	// to handle more
	repo := &Repo{inst.repoPrefix[0], "", ""}

	// TODO handle with version
	switch len(keys) {
	case 1:
		repo.Pkg = keys[0]
	case 2:
		repo.RepoPrefix.User = keys[0]
		repo.Pkg = keys[1]
	default:
		// TODO handle with version
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Location", repo.String())
	w.WriteHeader(http.StatusTemporaryRedirect)
}
