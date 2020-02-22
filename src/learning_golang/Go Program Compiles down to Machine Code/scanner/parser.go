package main

import (
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	src := []byte(`
	package main

	import "fmt"

	func main() {
		fmt.Println("Golang Again!")
	}
	`)
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		panic(err)
	}
	ast.Print(fset, file)
}
