package apidoc

import (
	"fmt"
	"strings"

	"github.com/fulldump/golax"
)

type NodePrint struct {
	Api                 *golax.Api
	Node                *golax.Node
	Context             *golax.Context
	Path                string
	Level               int
	AllInterceptors     map[*golax.Interceptor]*golax.Interceptor
	CurrentInterceptors []*golax.Interceptor
	DeepInterceptors    []*golax.Interceptor
}

func (np *NodePrint) Println(s string) {
	fmt.Fprintln(np.Context.Response, s)
}

func (np *NodePrint) Printf(f string, a ...interface{}) {
	fmt.Fprintf(np.Context.Response, f, a...)
}

func md_link(l string) string {
	l = strings.ToLower(l)
	l = strings.Replace(l, " ", "-", -1)
	l = strings.Replace(l, "/", "", -1)
	l = strings.Replace(l, "{", "", -1)
	l = strings.Replace(l, "}", "", -1)
	return "#" + l
}

func md_count_tabs(d string) int {
	i := 0
	for _, c := range d {
		if c != '\t' {
			break
		}
		i++
	}

	return i
}

func md_crop_tabs(d string) string {
	// Split lines
	lines := strings.Split(d, "\n")

	first := 0
	last := len(lines)
	if len(lines) > 2 {
		first++
		last--
	}

	// Get min tabs
	min_tabs := 99999
	for _, line := range lines[first:last] {
		// if 0 == i {
		// 	continue
		// }
		if strings.TrimSpace(line) != "" {
			c := md_count_tabs(line)
			if min_tabs > c {
				min_tabs = c
			}
		}
	}

	// Prefix
	prefix := strings.Repeat("\t", min_tabs)

	// Do the work
	for i, line := range lines {
		lines[i] = strings.TrimPrefix(line, prefix)
	}

	return strings.Join(lines, "\n")
}

func md_description(d string) string {
	d = md_crop_tabs(d)
	d = strings.Replace(d, "\n´´´", "\n```", -1)
	//d = strings.Replace(d, "´", "`", -1)
	return d
}

func PrintApiMd(p NodePrint) {

	if p.Node.Documentation.Ommit {
		return
	}

	for _, i := range p.Node.Interceptors {
		p.AllInterceptors[i] = i
		p.CurrentInterceptors = append(p.CurrentInterceptors, i)
	}
	for _, i := range p.Node.InterceptorsDeep {
		p.AllInterceptors[i] = i
		p.DeepInterceptors = append([]*golax.Interceptor{i}, p.DeepInterceptors...)
	}

	is_root := 0 == p.Level
	p.Level++

	// Title
	if is_root {
		p.Println("# API Documentation")
	} else {
		p.Path += "/" + p.Node.GetPath()
		p.Println("\n## " + p.Path + "\n")
	}

	// Description
	p.Println(md_description(p.Node.Documentation.Description))

	// Applied interceptors
	if len(p.CurrentInterceptors) > 0 {
		interceptors := "\n**Interceptors chain:** "
		if is_root {
			interceptors = "\n**Interceptors applied to all API:** "
		}
		for _, v := range p.CurrentInterceptors {
			name := v.Documentation.Name
			link := md_link("Interceptor " + name)
			interceptors += " [`" + name + "`](" + link + ") "
		}
		for _, v := range p.DeepInterceptors {
			name := v.Documentation.Name
			link := md_link("Interceptor " + name)
			interceptors += " [`" + name + "`](" + link + ") "
		}
		p.Println(interceptors)
	}

	// Implemented methods
	if len(p.Node.Methods) > 0 {
		methods := "\n**Methods:** "
		for k, _ := range p.Node.Methods {
			link := md_link(k + " " + p.Path)
			methods += " [`" + k + "`](" + link + ") "
		}
		p.Println(methods)

		for k, _ := range p.Node.Methods {
			methods += " `" + k + "` "
			p.Println("\n### " + k + " " + p.Path + "\n")

			if d, e := p.Node.DocumentationMethods[k]; e {
				p.Println(md_description(d.Description))
			}
		}
	}

	// Document children
	for _, child := range p.Node.Children {
		p.Node = child
		PrintApiMd(p)
	}

	if is_root {
		p.Println("\n# Interceptors")

		for _, v := range p.AllInterceptors {
			p.Println("\n## Interceptor " + v.Documentation.Name)
			p.Println(md_description(v.Documentation.Description))
		}
	}
}
