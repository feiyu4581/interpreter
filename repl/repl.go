package repl

import (
	"bufio"
	"fmt"
	"interpreter/lexer"
	"interpreter/parser"
	"io"
	"strings"
)

const PROMPT = ">>"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		_, _ = fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.NewLexer(strings.NewReader(line))
		p := parser.NewParser(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		_, _ = io.WriteString(out, program.String())
		_, _ = io.WriteString(out, "\n")
	}

}

func printParserErrors(out io.Writer, errros []string) {
	for _, msg := range errros {
		_, _ = io.WriteString(out, "\t"+msg+"\n")
	}
}
