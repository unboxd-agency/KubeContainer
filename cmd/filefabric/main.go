// FileFabric — the VDFS seat, built fresh in this house: one virtual
// file fabric across every device of a principal. Every file is known
// by its content (SHA-256); the catalog is one JSON-LD document
// (schema.org MediaObject) that devices merge into — the same file on
// two devices is one object with two locations, which is the whole
// trick. Prior art: Spacedrive's VDFS (paused upstream); this is the
// capability rebuilt, not the code adopted. A tool: it indexes and
// serves only when invoked, and never acts on its own.
//
// Usage:
//
//	filefabric index -dir <path> -device <name> [-catalog fabric.jsonld]
//	filefabric serve [-addr :8080] [-catalog fabric.jsonld]
package main

import (
	"crypto/sha256"
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"slices"
	"sort"
	"time"
)

type location struct {
	Device string `json:"device"`
	Path   string `json:"path"`
}

type object struct {
	Type      string     `json:"@type"`
	ID        string     `json:"@id"`
	Name      string     `json:"name"`
	Size      int64      `json:"contentSize"`
	Locations []location `json:"locations"`
	IndexedAt string     `json:"dateModified"`
}

type catalog struct {
	Context string    `json:"@context"`
	Graph   []*object `json:"@graph"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: filefabric index|serve [flags]")
		os.Exit(2)
	}
	var err error
	switch os.Args[1] {
	case "index":
		err = runIndex(os.Args[2:])
	case "serve":
		err = runServe(os.Args[2:])
	default:
		err = fmt.Errorf("unknown command %q", os.Args[1])
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func runIndex(args []string) error {
	fl := flag.NewFlagSet("index", flag.ExitOnError)
	dir := fl.String("dir", ".", "directory to index")
	device := fl.String("device", "", "device name (required)")
	path := fl.String("catalog", "fabric.jsonld", "catalog file to merge into")
	if err := fl.Parse(args); err != nil {
		return err
	}
	if *device == "" {
		return fmt.Errorf("-device is required: the fabric records where a file lives")
	}
	cat := load(*path)
	byID := make(map[string]*object, len(cat.Graph))
	for _, o := range cat.Graph {
		byID[o.ID] = o
	}
	files, dupes := 0, 0
	err := filepath.WalkDir(*dir, func(p string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return err
		}
		sum, size, err := hashFile(p)
		if err != nil {
			return err
		}
		files++
		loc := location{Device: *device, Path: p}
		if o, ok := byID[sum]; ok {
			if slices.Contains(o.Locations, loc) {
				return nil
			}
			o.Locations = append(o.Locations, loc)
			dupes++
			return nil
		}
		o := &object{
			Type: "MediaObject", ID: sum, Name: d.Name(), Size: size,
			Locations: []location{loc}, IndexedAt: time.Now().UTC().Format(time.RFC3339),
		}
		byID[sum] = o
		cat.Graph = append(cat.Graph, o)
		return nil
	})
	if err != nil {
		return err
	}
	sort.Slice(cat.Graph, func(i, j int) bool { return cat.Graph[i].Name < cat.Graph[j].Name })
	if err := save(*path, cat); err != nil {
		return err
	}
	fmt.Printf("fabric: %d file(s) indexed on %q, %d already known by content, %d object(s) total -> %s\n",
		files, *device, dupes, len(cat.Graph), *path)
	return nil
}

func runServe(args []string) error {
	fl := flag.NewFlagSet("serve", flag.ExitOnError)
	addr := fl.String("addr", ":8080", "listen address")
	path := fl.String("catalog", "fabric.jsonld", "catalog file to serve")
	if err := fl.Parse(args); err != nil {
		return err
	}
	cat := load(*path)
	byID := make(map[string]*object, len(cat.Graph))
	for _, o := range cat.Graph {
		byID[o.ID] = o
	}
	page := template.Must(template.New("fabric").Parse(pageHTML))
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		_ = page.Execute(w, cat)
	})
	http.HandleFunc("/catalog", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/ld+json")
		_ = json.NewEncoder(w).Encode(cat)
	})
	http.HandleFunc("/f/", func(w http.ResponseWriter, r *http.Request) {
		o, ok := byID["sha256:"+r.URL.Path[len("/f/"):]]
		if !ok {
			http.NotFound(w, r)
			return
		}
		for _, l := range o.Locations {
			if _, err := os.Stat(l.Path); err == nil {
				http.ServeFile(w, r, l.Path)
				return
			}
		}
		http.Error(w, "object known to the fabric but not present on this device", http.StatusBadGateway)
	})
	fmt.Printf("fabric: serving %d object(s) on %s\n", len(cat.Graph), *addr)
	server := &http.Server{Addr: *addr, ReadHeaderTimeout: 10 * time.Second}
	return server.ListenAndServe()
}

func hashFile(p string) (string, int64, error) {
	f, err := os.Open(p)
	if err != nil {
		return "", 0, err
	}
	defer func() { _ = f.Close() }()
	h := sha256.New()
	n, err := io.Copy(h, f)
	if err != nil {
		return "", 0, err
	}
	return fmt.Sprintf("sha256:%x", h.Sum(nil)), n, nil
}

func load(path string) *catalog {
	cat := &catalog{Context: "https://schema.org"}
	raw, err := os.ReadFile(path)
	if err == nil {
		_ = json.Unmarshal(raw, cat)
	}
	return cat
}

func save(path string, cat *catalog) error {
	raw, err := json.MarshalIndent(cat, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, append(raw, '\n'), 0o644)
}

const pageHTML = `<!doctype html><meta charset="utf-8"><title>FileFabric</title>
<style>body{font-family:system-ui;margin:2rem auto;max-width:60rem}td,th{padding:.3rem .8rem;text-align:left}</style>
<h1>FileFabric</h1><p>{{len .Graph}} object(s), known by content.</p>
<table><tr><th>Name</th><th>Size</th><th>Locations</th><th>Open</th></tr>
{{range .Graph}}<tr><td>{{.Name}}</td><td>{{.Size}}</td>
<td>{{range .Locations}}{{.Device}}:{{.Path}}<br>{{end}}</td>
<td><a href="/f/{{slice .ID 7}}">by content</a></td></tr>{{end}}
</table><p><a href="/catalog">the catalog (JSON-LD)</a></p>
`
