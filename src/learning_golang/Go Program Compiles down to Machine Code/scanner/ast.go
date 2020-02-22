package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
)

func main() {
	src := []byte(`
	package main

	import "fmt"

	func main() {
		fmt.Println("Golang Ast!")
	}
	`)
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		panic(err)
	}
	ast.Inspect(file, func(n ast.Node) bool {
		call, ok := n.(*ast.CallExpr)
		if !ok {
			return true
		}
		printer.Fprint(os.Stdout, fset, call.Fun)
		fmt.Println()
		return false
	})
}
