// CodeCompiler — the first tool of the house: one binary that answers
// the counsel's standing question, "does the code compile, and does it
// conform?" It runs the admission gates in order and reports a verdict
// for each; it changes nothing, builds nothing into place, and acts
// only when invoked — a tool, which does nothing on its own.
package main

import (
	"fmt"
	"os"
	"os/exec"
)

type gate struct {
	name string
	cmd  []string
}

func main() {
	gates := []gate{
		{"compile", []string{"go", "build", "./..."}},
		{"vet", []string{"go", "vet", "./..."}},
		{"lint", []string{"make", "lint"}},
		{"vocabulary", []string{"./hack/check-vocabulary.sh"}},
		{"graph", []string{"go", "run", "./cmd/recordgraph"}},
		{"schema", []string{"go", "run", "./cmd/schemakeeper"}},
	}
	failed := 0
	for _, g := range gates {
		c := exec.Command(g.cmd[0], g.cmd[1:]...)
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		if err := c.Run(); err != nil {
			fmt.Printf("[fail] %s\n", g.name)
			failed++
			continue
		}
		fmt.Printf("[pass] %s\n", g.name)
	}
	if failed > 0 {
		fmt.Printf("verdict: %d gate(s) failed — the code does not conform\n", failed)
		os.Exit(1)
	}
	fmt.Println("verdict: the code compiles, and it conforms")
}
