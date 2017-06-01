package typing

import (
	"github.com/rhysd/gocaml/alpha"
	"github.com/rhysd/gocaml/lexer"
	"github.com/rhysd/gocaml/parser"
	"github.com/rhysd/gocaml/token"
	"path/filepath"
)

func Example() {
	file := filepath.FromSlash("../testdata/from-mincaml/ack.ml")
	src, err := token.NewSourceFromFile(file)
	if err != nil {
		// File not found
		panic(err)
	}

	lex := lexer.NewLexer(src)
	go lex.Lex()

	ast, err := parser.Parse(lex.Tokens)
	if err != nil {
		// When parse failed
		panic(err)
	}

	if err = alpha.Transform(ast.Root); err != nil {
		// When some some duplicates found
		panic(err)
	}

	// Create new type analysis environment
	// (symbol table and external variables table)
	env := NewEnv()

	// Apply type inference. After this, all symbols in AST should have exact
	// types. It also checks types are valid and all types are determined by
	// inference
	if err := env.ApplyTypeAnalysis(ast.Root); err != nil {
		// Type error detected
		panic(err)
	}

	// You can dump the type table
	env.Dump()
}
