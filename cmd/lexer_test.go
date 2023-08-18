package main

import (
	"fmt"
	"testing"
)

func TestLexer(t *testing.T) {
	t.Skip("Not testing lexer now.")
	path := "tok_test.butter"
	fmt.Println("Start reading file", path)

	/*
		fileBytes, err := os.ReadFile(path)
		if err != nil {
			fmt.Print("Could not read file.")
		}
	*/

	file := `
	(){}=;+
	let hello = 5;
	let largeNumber = 123123891
	
	// this is bullshit but that doesnt matter.
	let print = fn(something) {
		<< something;
	}

	let n = 4*hello;

	let somebool = true

	if hello != 4 {
		print()
	}

	if hello == 5 {
		sdfsdf
	} else {
		print(hello)
	}
	`

	l := NewLexer(file)

	var tok Token

	for {
		tok = l.NextToken()
		name := TokStr(tok.Type)
		fmt.Println("Just read", name, tok.Literal)
		if tok.Type == EOF {
			break
		}
	}
}

func TestParser(t *testing.T) {

	file := `
	let x = 5;
	return lskdfj dskjlflskdjf ls;
	`

	l := NewLexer(file)
	parser := NewParser(l)
	program := parser.ParseProgram()
	for i, s := range program.Statements {
		fmt.Printf("%d: %+v\n", i, s)
	}

	fmt.Println("\nErrors:")

	for i, e := range parser.Errors() {
		fmt.Printf("%d: %s\n", i, e)
	}

	fmt.Println("\nAST Output:")
	fmt.Println(program.String())
}
