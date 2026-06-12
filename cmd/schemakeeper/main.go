// SchemaKeeper — the schema as a tool of the house: the machine face
// of the record speaks schema.org JSON-LD, and this binary keeps it
// true. It validates eval/graph.jsonld against the pinned schema:
// the context must be the pinned vocabulary, every node's type must
// be in the pinned type set, and every citation must resolve to a
// node present in the graph — no dangling nodes, the failure DBpedia
// tolerates at scale and this record refuses at the desk. It changes
// nothing and acts only when invoked.
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

const pinnedContext = "https://schema.org"

var pinnedTypes = map[string]bool{"CreativeWork": true}

type node struct {
	ID       string `json:"@id"`
	Type     string `json:"@type"`
	Citation []struct {
		ID string `json:"@id"`
	} `json:"citation"`
}

type doc struct {
	Context string `json:"@context"`
	Graph   []node `json:"@graph"`
}

func main() {
	raw, err := os.ReadFile("eval/graph.jsonld")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	var d doc
	if err := json.Unmarshal(raw, &d); err != nil {
		fmt.Fprintf(os.Stderr, "not valid JSON: %v\n", err)
		os.Exit(1)
	}
	failed := 0
	if d.Context != pinnedContext {
		fmt.Printf("[fail] context %q is not the pinned %q\n", d.Context, pinnedContext)
		failed++
	}
	ids := map[string]bool{}
	for _, n := range d.Graph {
		ids[n.ID] = true
	}
	for _, n := range d.Graph {
		if !pinnedTypes[n.Type] {
			fmt.Printf("[fail] %s: type %q is not in the pinned set\n", n.ID, n.Type)
			failed++
		}
		for _, c := range n.Citation {
			if !ids[c.ID] {
				fmt.Printf("[fail] %s cites %s — a node not in the graph\n", n.ID, c.ID)
				failed++
			}
		}
	}
	if failed > 0 {
		fmt.Printf("verdict: %d schema violation(s) — the graph does not keep its schema\n", failed)
		os.Exit(1)
	}
	fmt.Printf("schema kept: %d nodes, context %s, all types pinned, every citation resolves\n",
		len(d.Graph), pinnedContext)
}
