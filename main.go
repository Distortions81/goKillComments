package main

import (
	"bytes"
	"fmt"
	"go/format"
	"go/parser"
	"go/token"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <file.go>\n", os.Args[0])
		os.Exit(1)
	}
	filename := os.Args[1]

	// Create a new token.FileSet for position information.
	fset := token.NewFileSet()

	// Parse the file including comments.
	file, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing file: %v\n", err)
		os.Exit(1)
	}

	// Remove all comments by setting the Comments field to nil.
	file.Comments = nil

	// Format the AST back into source code.
	var buf bytes.Buffer
	if err := format.Node(&buf, fset, file); err != nil {
		fmt.Fprintf(os.Stderr, "Error formatting source: %v\n", err)
		os.Exit(1)
	}

	// Print the resulting source code to stdout.
	fmt.Print(buf.String())
}
