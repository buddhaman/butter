package main

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}

func NewThing(fuckyou string) string {
	return fuckyou + "hi"
}
