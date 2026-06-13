// BookBinder — the record bound as a book, multi-page: each domain of
// the estate is one chapter, each chapter binds its own topics, the
// cover carries the table of contents. With -repo it is a direct
// repo-to-book converter: any repository's markdown, bound as found,
// content unaltered — the story is whatever the repo says. Generated,
// never hand-written; the prose stays canonical where it lives. A
// tool: it binds only when invoked.
package main

import (
	"flag"
	"fmt"
	"html"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type chapter struct {
	domain string
	title  string
	topics []string
}

// Each domain one chapter; each chapter, its topics in reading order.
var book = []chapter{
	{"unboxd.cloud", "The House", []string{
		"docs/FOUNDING-PRINCIPLES.md", "docs/PRIMITIVES.md",
		"docs/DOCTRINE-MAP.md", "docs/TOOLS.md",
	}},
	{"kubecontainer.xyz", "The Kube", []string{
		"docs/KUBE-SPEC.md", "docs/HEADLESS-DELIVERY.md",
		"docs/MEASUREMENT-STANDARD.md", "deploy/LEAPMICRO.md",
	}},
	{"agennext.com", "The Agents", []string{
		"docs/AGENT-PLATFORM.md", "docs/PERSONAL.md",
	}},
	{"agennext.space", "The Space", []string{
		"deploy/AGENT-STACK.md",
	}},
	{"openautonomyx.com", "The Platform", []string{
		"docs/CONTROL-PANEL.md", "docs/assessments/APACHE-ATTIC.md",
	}},
}

const style = `<style>body{font-family:Georgia,serif;max-width:46rem;margin:3rem auto;line-height:1.6;padding:0 1rem}
h1,h2,h3,h4{font-family:system-ui;line-height:1.2}pre{background:#f4f4f4;padding:.8rem;overflow-x:auto}
code{background:#f4f4f4;padding:0 .2rem}.cover{text-align:center;margin:5rem 0}
.toc a{display:block;padding:.25rem 0;text-decoration:none}.domain{color:#666;font-size:.9rem}
nav.foot{margin:4rem 0;display:flex;justify-content:space-between}hr{margin:3rem 0}</style>
`

func main() {
	repo := flag.String("repo", "", "bind this repository directory instead of the estate book")
	title := flag.String("title", "", "book title (repo mode; defaults to the directory name)")
	out := flag.String("out", "site/book", "output directory")
	flag.Parse()

	pages := book
	name := "The Book of Software"
	sub := "each domain one chapter; the prose stays canonical where it lives"
	if *repo != "" {
		pages = repoChapters(*repo)
		name = *title
		if name == "" {
			name = filepath.Base(*repo)
		}
		sub = "bound as found — content unaltered; the story is the repo's own"
	}
	if err := bind(pages, name, sub, *out); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// repoChapters reads a repository as a book: root markdown is the
// first chapter, each top-level directory holding markdown is a
// chapter of its own; README leads, the rest follow alphabetically.
func repoChapters(root string) []chapter {
	byDir := map[string][]string{}
	_ = filepath.WalkDir(root, func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		if d.IsDir() {
			n := d.Name()
			if n == ".git" || n == "node_modules" || n == "vendor" || n == "bin" || n == "dist" {
				return filepath.SkipDir
			}
			return nil
		}
		if !strings.HasSuffix(strings.ToLower(p), ".md") {
			return nil
		}
		rel, _ := filepath.Rel(root, p)
		dir := "."
		if before, _, found := strings.Cut(rel, "/"); found {
			dir = before
		}
		byDir[dir] = append(byDir[dir], p)
		return nil
	})
	dirs := make([]string, 0, len(byDir))
	for d := range byDir {
		if d != "." {
			dirs = append(dirs, d)
		}
	}
	sort.Strings(dirs)
	if _, ok := byDir["."]; ok {
		dirs = append([]string{"."}, dirs...)
	}
	chapters := make([]chapter, 0, len(dirs))
	for _, d := range dirs {
		topics := byDir[d]
		sort.Slice(topics, func(i, j int) bool {
			ri := strings.HasPrefix(strings.ToUpper(filepath.Base(topics[i])), "README")
			rj := strings.HasPrefix(strings.ToUpper(filepath.Base(topics[j])), "README")
			if ri != rj {
				return ri
			}
			return topics[i] < topics[j]
		})
		t := d
		if d == "." {
			t = "The Door"
		}
		chapters = append(chapters, chapter{domain: d, title: t, topics: topics})
	}
	return chapters
}

func bind(pages []chapter, name, sub, out string) error {
	if err := os.MkdirAll(out, 0o755); err != nil {
		return err
	}
	bound, missing := 0, 0
	var toc strings.Builder
	for i, ch := range pages {
		var b strings.Builder
		b.WriteString("<!doctype html><meta charset=\"utf-8\"><title>" +
			html.EscapeString(ch.title) + "</title>\n" + style)
		fmt.Fprintf(&b, "<p class=\"domain\">%s</p><h1>%d. %s</h1>\n",
			html.EscapeString(ch.domain), i+1, html.EscapeString(ch.title))
		for _, t := range ch.topics {
			raw, err := os.ReadFile(t)
			if err != nil {
				missing++
				continue
			}
			b.WriteString("<hr>\n" + render(string(raw)))
			bound++
		}
		b.WriteString(navFootN(i, len(pages)))
		file := fmt.Sprintf("%s/chapter-%d.html", out, i+1)
		if err := os.WriteFile(file, []byte(b.String()), 0o644); err != nil {
			return err
		}
		fmt.Fprintf(&toc, "<a href=\"chapter-%d.html\">%d. %s <span class=\"domain\">— %s</span></a>\n",
			i+1, i+1, html.EscapeString(ch.title), html.EscapeString(ch.domain))
	}
	cover := "<!doctype html><meta charset=\"utf-8\"><title>" + html.EscapeString(name) + "</title>\n" + style +
		"<div class=\"cover\"><h1>" + html.EscapeString(name) + "</h1>" +
		"<p>" + html.EscapeString(sub) + "</p></div>" +
		"<nav class=\"toc\"><h2>Contents</h2>\n" + toc.String() + "</nav>\n"
	if err := os.WriteFile(out+"/index.html", []byte(cover), 0o644); err != nil {
		return err
	}
	fmt.Printf("the book: %d chapter(s), %d topic(s) bound, %d missing -> %s/\n",
		len(pages), bound, missing, out)
	if missing > 0 {
		fmt.Println("verdict: a topic is missing — the binding is incomplete")
		os.Exit(1)
	}
	fmt.Println("verdict: the record is bound")
	return nil
}

func navFootN(i, n int) string {
	var b strings.Builder
	b.WriteString("<nav class=\"foot\">")
	if i > 0 {
		fmt.Fprintf(&b, "<a href=\"chapter-%d.html\">&larr; previous</a>", i)
	} else {
		b.WriteString("<span></span>")
	}
	b.WriteString("<a href=\"index.html\">contents</a>")
	if i+1 < n {
		fmt.Fprintf(&b, "<a href=\"chapter-%d.html\">next &rarr;</a>", i+2)
	} else {
		b.WriteString("<span></span>")
	}
	b.WriteString("</nav>\n")
	return b.String()
}

// render is a deliberately small markdown renderer: headings, code
// fences, lists, bold/italics/code, tables as preformatted rows. The
// book favors faithful text over typographic ambition.
func render(md string) string {
	var b strings.Builder
	inCode := false
	inList := false
	for line := range strings.SplitSeq(md, "\n") {
		if strings.HasPrefix(line, "```") {
			if inCode {
				b.WriteString("</pre>\n")
			} else {
				b.WriteString("<pre>")
			}
			inCode = !inCode
			continue
		}
		if inCode {
			b.WriteString(html.EscapeString(line) + "\n")
			continue
		}
		trimmed := strings.TrimSpace(line)
		isItem := strings.HasPrefix(trimmed, "- ") || strings.HasPrefix(trimmed, "* ")
		if inList && !isItem {
			b.WriteString("</ul>\n")
			inList = false
		}
		switch {
		case strings.HasPrefix(line, "#"):
			level := min(len(line)-len(strings.TrimLeft(line, "#")), 4)
			text := strings.TrimSpace(strings.TrimLeft(line, "#"))
			// Topic files carry their own h1; demote one level inside a chapter.
			fmt.Fprintf(&b, "<h%d>%s</h%d>\n", level+1, inline(text), level+1)
		case isItem:
			if !inList {
				b.WriteString("<ul>\n")
				inList = true
			}
			b.WriteString("<li>" + inline(trimmed[2:]) + "</li>\n")
		case strings.HasPrefix(trimmed, "|"):
			b.WriteString("<pre>" + html.EscapeString(line) + "</pre>\n")
		case trimmed == "":
			b.WriteString("<p></p>\n")
		default:
			b.WriteString(inline(line) + "\n")
		}
	}
	if inList {
		b.WriteString("</ul>\n")
	}
	if inCode {
		b.WriteString("</pre>\n")
	}
	return b.String()
}

func inline(s string) string {
	s = html.EscapeString(s)
	for _, m := range []struct{ tag, mark string }{{"strong", "**"}, {"code", "`"}, {"em", "*"}} {
		for {
			i := strings.Index(s, m.mark)
			if i < 0 {
				break
			}
			j := strings.Index(s[i+len(m.mark):], m.mark)
			if j < 0 {
				break
			}
			inner := s[i+len(m.mark) : i+len(m.mark)+j]
			s = s[:i] + "<" + m.tag + ">" + inner + "</" + m.tag + ">" + s[i+len(m.mark)+j+len(m.mark):]
		}
	}
	return s
}
