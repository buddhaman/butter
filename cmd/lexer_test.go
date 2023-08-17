package main

import (
	"fmt"
	"testing"
)

func TestLexer(t *testing.T) {
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
