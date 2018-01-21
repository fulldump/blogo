package apidoc

import "github.com/fulldump/golax"

type NodeJson struct {
	Api                 *golax.Api
	Node                *golax.Node
	Context             *golax.Context
	Path                string
	Level               int
	AllInterceptors     map[*golax.Interceptor]*golax.Interceptor
	CurrentInterceptors []*golax.Interceptor
	DeepInterceptors    []*golax.Interceptor
	JsonDoc             *JsonDoc
}

type JsonDoc struct {
	Endpoints    map[string]JsonEndpoint `json:"endpoints"`
	Interceptors map[string]string       `json:"interceptors"`
	Title        string                  `json:"title"`
	Subtitle     string                  `json:"subtitle"`
}

type JsonEndpoint struct {
	Description  string                `json:"description"`
	Interceptors []string              `json:"interceptors"`
	Methods      map[string]JsonMethod `json:"methods"`
}

type JsonMethod struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func NewNodeJson(api *golax.Api) NodeJson {
	return NodeJson{
		Api:                 api,
		Path:                api.Prefix,
		Node:                api.Root,
		Level:               0,
		AllInterceptors:     map[*golax.Interceptor]*golax.Interceptor{},
		CurrentInterceptors: []*golax.Interceptor{},
		DeepInterceptors:    []*golax.Interceptor{},
		JsonDoc: &JsonDoc{
			Endpoints:    map[string]JsonEndpoint{},
			Interceptors: map[string]string{},
		},
	}
}

func RunJsonDoc(p NodeJson) {

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
	if !is_root {
		p.Path += "/" + p.Node.GetPath()
	}

	endpoint := JsonEndpoint{
		Description:  md_description(p.Node.Documentation.Description),
		Interceptors: []string{},
		Methods:      map[string]JsonMethod{},
	}

	// Add interceptors
	for _, v := range p.CurrentInterceptors {
		name := v.Documentation.Name
		endpoint.Interceptors = append(endpoint.Interceptors, name)
	}
	for _, v := range p.DeepInterceptors {
		name := v.Documentation.Name
		endpoint.Interceptors = append(endpoint.Interceptors, name)
	}

	// Add methods
	for method, _ := range p.Node.Methods {
		if d, e := p.Node.DocumentationMethods[method]; e {
			endpoint.Methods[method] = JsonMethod{
				Name:        d.Name,
				Description: md_description(d.Description),
			}
		} else {
			endpoint.Methods[method] = JsonMethod{
				Description: "",
			}
		}
	}

	p.JsonDoc.Endpoints[p.Path] = endpoint

	for _, operation := range p.Node.Operations {
		RunJsonDocOperation(p, operation)
	}

	// Document children
	for _, child := range p.Node.Children {
		p.Node = child
		RunJsonDoc(p)
	}

	if is_root {
		for _, v := range p.AllInterceptors {
			name := v.Documentation.Name
			description := md_description(v.Documentation.Description)
			p.JsonDoc.Interceptors[name] = description
		}
	}
}

/*
type JsonDoc struct {
	Endpoints    map[string]JsonEndpoint `json:"endpoints"`
	Interceptors map[string]string       `json:"interceptors"`
	Title        string                  `json:"title"`
	Subtitle     string                  `json:"subtitle"`
}
*/

func RunJsonDocOperation(p NodeJson, o *golax.Operation) {
	e := JsonEndpoint{
		Description:  "", // TODO: add a documentation structure :S
		Interceptors: []string{},
		Methods:      map[string]JsonMethod{},
	}

	p.Path += ":" + o.Path

	for _, v := range p.CurrentInterceptors {
		name := v.Documentation.Name
		e.Interceptors = append(e.Interceptors, name)
	}

	// Add methods
	for method, _ := range o.Methods {
		e.Methods[method] = JsonMethod{
			Name:        o.Path,
			Description: "operation description!",
		}
	}

	p.JsonDoc.Endpoints[p.Path] = e

}
