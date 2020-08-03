package parse

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"strings"
)

// Parser reads the service names from given Go structures.
type Parser struct {
	Services []string `json:"services"`
}

// Parse ...
func (p *Parser) Parse() error {
	// Parse file content from go:generate environment.
	filename := os.Getenv("GOFILE")
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	src := string(bytes)
	// Create an AST from source.
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", src, parser.AllErrors)
	if err != nil {
		return err
	}
	// Inspect the Syntax-Tree and read the service names by using the suffix Request and Response.
	services := make(map[string]int)
	ast.Inspect(f, func(n ast.Node) bool {
		switch v := n.(type) {
		case *ast.TypeSpec:
			name := v.Name.Name
			name = strings.Replace(name, "Request", "", -1)
			name = strings.Replace(name, "Response", "", -1)
			services[name]++
		}
		return true
	})
	// Add Services only if there is a request and response (num = 2).
	p.Services = make([]string, 0)
	for name, num := range services {
		if num == 2 {
			p.Services = append(p.Services, name)
		}
	}
	return nil
}

// NewParser ...
func NewParser() *Parser {
	return &Parser{
		Services: []string{},
	}
}
