// BookBinder — the record bound as a book, multi-page: each domain of
// the estate is one chapter, each chapter binds its own topics, the
// cover carries the table of contents. Generated from the corpus and
// never hand-written; the prose stays canonical where it lives. A
// tool: it binds only when invoked.
package main

import (
	"fmt"
	"html"
	"os"
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
	if err := os.MkdirAll("site/book", 0o755); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	bound, missing := 0, 0
	var toc strings.Builder
	for i, ch := range book {
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
		b.WriteString(navFoot(i))
		name := fmt.Sprintf("site/book/chapter-%d.html", i+1)
		if err := os.WriteFile(name, []byte(b.String()), 0o644); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Fprintf(&toc, "<a href=\"chapter-%d.html\">%d. %s <span class=\"domain\">— %s</span></a>\n",
			i+1, i+1, html.EscapeString(ch.title), html.EscapeString(ch.domain))
	}
	cover := "<!doctype html><meta charset=\"utf-8\"><title>The Book of Software</title>\n" + style +
		"<div class=\"cover\"><h1>The Book of Software</h1>" +
		"<p>each domain one chapter; the prose stays canonical where it lives</p></div>" +
		"<nav class=\"toc\"><h2>Contents</h2>\n" + toc.String() + "</nav>\n"
	if err := os.WriteFile("site/book/index.html", []byte(cover), 0o644); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("the book: %d chapter(s), %d topic(s) bound, %d missing -> site/book/\n",
		len(book), bound, missing)
	if missing > 0 {
		fmt.Println("verdict: a topic is missing — the binding is incomplete")
		os.Exit(1)
	}
	fmt.Println("verdict: the record is bound")
}

func navFoot(i int) string {
	var b strings.Builder
	b.WriteString("<nav class=\"foot\">")
	if i > 0 {
		fmt.Fprintf(&b, "<a href=\"chapter-%d.html\">&larr; previous</a>", i)
	} else {
		b.WriteString("<span></span>")
	}
	b.WriteString("<a href=\"index.html\">contents</a>")
	if i+1 < len(book) {
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
