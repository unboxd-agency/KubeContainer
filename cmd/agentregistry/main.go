// AgentRegistry — the registry's internet face: the append-only
// declarations this house keeps (skeletons, blueprints, agents,
// spaces) served read-only over HTTP, every declaration addressable
// by URL. The registry defines the contract; this binary only shows
// it — it admits nothing, changes nothing, and serves only when
// invoked. Writes stay where they always were: the record, by pull
// request, through the gates.
//
// Usage: agentregistry [-addr :8080] [-dir registry]
package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

func main() {
	addr := ":8080"
	dir := "registry"
	args := os.Args[1:]
	for i := 0; i+1 < len(args); i += 2 {
		switch args[i] {
		case "-addr":
			addr = args[i+1]
		case "-dir":
			dir = args[i+1]
		}
	}
	index, err := buildIndex(dir)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	page := template.Must(template.New("registry").Parse(pageHTML))
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		_ = page.Execute(w, index)
	})
	http.HandleFunc("/index.json", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(index)
	})
	http.HandleFunc("/registry/", func(w http.ResponseWriter, r *http.Request) {
		name := strings.TrimPrefix(r.URL.Path, "/registry/")
		clean := filepath.Clean(name)
		if clean != name || strings.Contains(clean, "..") || !strings.HasSuffix(clean, ".json") {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		http.ServeFile(w, r, filepath.Join(dir, clean))
	})
	fmt.Printf("agent registry: %d declaration(s) on %s\n", len(index.Declarations), addr)
	server := &http.Server{Addr: addr, ReadHeaderTimeout: 10 * time.Second}
	if err := server.ListenAndServe(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

type entry struct {
	Path     string `json:"path"`
	Kind     string `json:"kind"`
	Skeleton string `json:"skeleton,omitempty"`
	Name     string `json:"name,omitempty"`
}

type registryIndex struct {
	Note         string  `json:"note"`
	Declarations []entry `json:"declarations"`
}

func buildIndex(dir string) (*registryIndex, error) {
	idx := &registryIndex{
		Note: "Append-only; admission happens in the record, never here. This face is read-only.",
	}
	err := filepath.WalkDir(dir, func(p string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() || !strings.HasSuffix(p, ".json") {
			return err
		}
		rel, _ := filepath.Rel(dir, p)
		e := entry{Path: "registry/" + rel, Kind: kindOf(rel)}
		raw, err := os.ReadFile(p)
		if err == nil {
			var meta struct {
				Skeleton string `json:"skeleton"`
				Name     string `json:"name"`
			}
			_ = json.Unmarshal(raw, &meta)
			e.Skeleton = meta.Skeleton
			e.Name = meta.Name
		}
		idx.Declarations = append(idx.Declarations, e)
		return nil
	})
	sort.Slice(idx.Declarations, func(i, j int) bool {
		return idx.Declarations[i].Path < idx.Declarations[j].Path
	})
	return idx, err
}

func kindOf(rel string) string {
	switch {
	case strings.HasPrefix(rel, "agents/"):
		return "agent"
	case strings.HasPrefix(rel, "blueprints/"):
		return "blueprint"
	case strings.HasPrefix(rel, "spaces/"):
		return "space"
	default:
		return "skeleton"
	}
}

const pageHTML = `<!doctype html><meta charset="utf-8"><title>The Agent Registry</title>
<style>body{font-family:system-ui;margin:2rem auto;max-width:60rem}td,th{padding:.3rem .8rem;text-align:left}</style>
<h1>The Agent Registry</h1>
<p>{{len .Declarations}} declaration(s). {{.Note}}</p>
<table><tr><th>Kind</th><th>Name</th><th>Skeleton</th><th>Declaration</th></tr>
{{range .Declarations}}<tr><td>{{.Kind}}</td><td>{{.Name}}</td><td>{{.Skeleton}}</td>
<td><a href="/{{.Path}}">{{.Path}}</a></td></tr>{{end}}
</table><p><a href="/index.json">the index (JSON)</a></p>
`
