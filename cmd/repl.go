package main

import (
	"bufio"
	"fmt"
	"io"
)

const PROMPT = ">>"

func ReplStart(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(PROMPT)

		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()

		l := NewLexer(line)

		for tok := l.NextToken(); tok.Type != EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
