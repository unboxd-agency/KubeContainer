// RecordGraph — the corpus as a graph, the DBpedia pattern applied to
// the record: the prose stays canonical, the graph is extracted from
// it, never hand-written. Nodes are the corpus files; an edge is a
// reference one file makes to another. The tool emits the generated
// graph (eval/graph.txt) and returns a verdict: a reference to a path
// that does not exist is a broken edge, and broken edges fail the
// gate. It changes nothing else and acts only when invoked.
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

var topDirs = map[string]bool{
	"api": true, "cmd": true, "config": true, "deploy": true,
	"docs": true, "eval": true, "hack": true, "internal": true,
	"registry": true, "release": true, "site": true, "test": true,
	".github": true, ".githooks": true,
}

// Generated at build time; their absence in the tree is not a broken edge.
var generated = []string{"dist/", "bin/"}

var spanRe = regexp.MustCompile("`([^`\n]+)`")
var jsonRe = regexp.MustCompile(`[A-Za-z0-9_.-]+(?:/[A-Za-z0-9_.-]+)+`)

func main() {
	planned := plannedRefs()
	var nodes []string
	_ = filepath.WalkDir(".", func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		if d.IsDir() {
			name := d.Name()
			if name == ".git" || name == "bin" || name == "dist" || name == "node_modules" {
				return filepath.SkipDir
			}
			return nil
		}
		if strings.HasSuffix(path, ".md") ||
			(strings.HasPrefix(path, "registry/") && strings.HasSuffix(path, ".json")) {
			nodes = append(nodes, path)
		}
		return nil
	})
	sort.Strings(nodes)

	edges := map[string]bool{}
	var broken []string
	for _, n := range nodes {
		raw, err := os.ReadFile(n)
		if err != nil {
			continue
		}
		for _, t := range tokens(string(raw), strings.HasSuffix(n, ".json")) {
			if t == n {
				continue
			}
			if isGenerated(t) || planned[t] {
				continue
			}
			if _, err := os.Stat(t); err != nil {
				broken = append(broken, fmt.Sprintf("%s -> %s", n, t))
				continue
			}
			edges[fmt.Sprintf("%s -> %s", n, t)] = true
		}
	}

	var lines []string
	for e := range edges {
		lines = append(lines, e)
	}
	sort.Strings(lines)
	out := "# RecordGraph — generated from the corpus; do not edit. Rebuild: make graph\n" +
		fmt.Sprintf("# nodes=%d edges=%d\n", len(nodes), len(lines)) +
		strings.Join(lines, "\n") + "\n"
	if err := os.WriteFile("eval/graph.txt", []byte(out), 0o644); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if err := writeJSONLD(nodes, lines); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Printf("record graph: %d nodes, %d edges -> eval/graph.txt, eval/graph.jsonld\n", len(nodes), len(lines))
	if len(broken) > 0 {
		for _, b := range broken {
			fmt.Printf("[broken] %s\n", b)
		}
		fmt.Printf("verdict: %d broken reference(s) — the record does not resolve\n", len(broken))
		os.Exit(1)
	}
	fmt.Println("verdict: every reference resolves")
}

// tokens extracts path-like references. In markdown only backticked
// spans count (prose mentions a path by quoting it); in JSON any
// path-shaped string counts.
func tokens(content string, isJSON bool) []string {
	var spans []string
	if isJSON {
		spans = jsonRe.FindAllString(content, -1)
	} else {
		for _, m := range spanRe.FindAllStringSubmatch(content, -1) {
			// A span may list several files: `docs/a.md, b.md` — later
			// bare names inherit the directory of the first.
			dir := ""
			for _, part := range strings.Split(m[1], ",") {
				part = strings.TrimSpace(part)
				if !strings.Contains(part, "/") && dir != "" {
					part = dir + part
				}
				spans = append(spans, part)
				if i := strings.LastIndex(part, "/"); i >= 0 {
					dir = part[:i+1]
				}
			}
		}
	}
	var out []string
	lineRef := regexp.MustCompile(`:\d+$`)
	for _, s := range spans {
		s = lineRef.ReplaceAllString(s, "")
		s = strings.TrimSuffix(s, "/")
		if s == "" || strings.ContainsAny(s, " *${}()§\"'") || !strings.Contains(s, "/") {
			continue
		}
		if !topDirs[strings.SplitN(s, "/", 2)[0]] {
			continue
		}
		out = append(out, s)
	}
	return out
}

// writeJSONLD renders the same graph as schema.org JSON-LD: every
// corpus file a CreativeWork, every edge a citation — standard linked
// data, the DBpedia family's interchange format.
func writeJSONLD(nodes, edges []string) error {
	cites := map[string][]map[string]string{}
	all := map[string]bool{}
	for _, n := range nodes {
		all[n] = true
	}
	for _, e := range edges {
		parts := strings.SplitN(e, " -> ", 2)
		cites[parts[0]] = append(cites[parts[0]], map[string]string{"@id": parts[1]})
		all[parts[1]] = true
	}
	var ids []string
	for id := range all {
		ids = append(ids, id)
	}
	sort.Strings(ids)
	graph := make([]map[string]any, 0, len(ids))
	for _, n := range ids {
		entry := map[string]any{"@id": n, "@type": "CreativeWork", "name": n}
		if c := cites[n]; len(c) > 0 {
			entry["citation"] = c
		}
		graph = append(graph, entry)
	}
	doc := map[string]any{"@context": "https://schema.org", "@graph": graph}
	raw, err := json.MarshalIndent(doc, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile("eval/graph.jsonld", append(raw, '\n'), 0o644)
}

// plannedRefs reads eval/graph-planned.txt: declared-future paths the
// roadmap names before they exist. Shrink-only, like the vocabulary
// baseline — a planned path that lands leaves the file.
func plannedRefs() map[string]bool {
	out := map[string]bool{}
	raw, err := os.ReadFile("eval/graph-planned.txt")
	if err != nil {
		return out
	}
	for _, l := range strings.Split(string(raw), "\n") {
		l = strings.TrimSpace(l)
		if l != "" && !strings.HasPrefix(l, "#") {
			out[l] = true
		}
	}
	return out
}

func isGenerated(t string) bool {
	for _, g := range generated {
		if strings.HasPrefix(t, g) {
			return true
		}
	}
	return false
}
