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
