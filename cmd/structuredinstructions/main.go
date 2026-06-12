// StructuredInstructions — the flow's gate, one binary: declare a
// skeleton, validate a declaration against it, and refuse the
// undeclared shape. On a passing verdict with -submit, the declaration
// is appended to the registry — named, signed, on the record before
// the first act. A declaration missing a required field is not
// malformed; it is inadmissible.
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type skeleton struct {
	Skeleton      string   `json:"skeleton"`
	Required      []string `json:"required"`
	ContractTerms []string `json:"contractTerms"`
}

func lookup(doc map[string]any, dotted string) (any, bool) {
	cur := any(doc)
	for _, part := range strings.Split(dotted, ".") {
		m, ok := cur.(map[string]any)
		if !ok {
			return nil, false
		}
		cur, ok = m[part]
		if !ok {
			return nil, false
		}
	}
	return cur, true
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: structuredinstructions <declaration.json> [-submit]")
		os.Exit(2)
	}
	declPath := os.Args[1]
	submit := len(os.Args) > 2 && os.Args[2] == "-submit"

	skelRaw, err := os.ReadFile("registry/SKELETON.json")
	if err != nil {
		fmt.Println("[fail] no skeleton declared: ", err)
		os.Exit(1)
	}
	var skel skeleton
	if err := json.Unmarshal(skelRaw, &skel); err != nil {
		fmt.Println("[fail] skeleton unreadable:", err)
		os.Exit(1)
	}

	declRaw, err := os.ReadFile(declPath)
	if err != nil {
		fmt.Println("[fail] declaration unreadable:", err)
		os.Exit(1)
	}
	var decl map[string]any
	if err := json.Unmarshal(declRaw, &decl); err != nil {
		fmt.Println("[fail] not a structured instruction:", err)
		os.Exit(1)
	}

	inadmissible := 0
	if decl["skeleton"] != skel.Skeleton {
		fmt.Printf("[fail] undeclared shape: skeleton %q is not %q\n", decl["skeleton"], skel.Skeleton)
		inadmissible++
	}
	for _, field := range skel.Required {
		if v, ok := lookup(decl, field); !ok || v == "" {
			fmt.Printf("[fail] required field missing or empty: %s\n", field)
			inadmissible++
		}
	}
	terms, _ := decl["contractTerms"].([]any)
	have := map[string]bool{}
	for _, t := range terms {
		if s, ok := t.(string); ok {
			have[s] = true
		}
	}
	for _, t := range skel.ContractTerms {
		if !have[t] {
			fmt.Printf("[fail] contract term not acknowledged: %s\n", t)
			inadmissible++
		}
	}
	if inadmissible > 0 {
		fmt.Printf("verdict: inadmissible — %d gate(s) failed; the shape was not declared\n", inadmissible)
		os.Exit(1)
	}
	fmt.Println("verdict: admissible — the declaration matches its skeleton")

	if submit {
		name, _ := decl["name"].(string)
		dst := filepath.Join("registry", "agents", name+".json")
		if _, err := os.Stat(dst); err == nil {
			fmt.Printf("[fail] registry refuses duplicate: %s already registered (nothing removed, nothing duplicated)\n", name)
			os.Exit(1)
		}
		if err := os.WriteFile(dst, declRaw, 0o644); err != nil {
			fmt.Println("[fail] submit:", err)
			os.Exit(1)
		}
		fmt.Printf("submitted: %s — named, signed, on the record before the first act\n", dst)
	}
}
