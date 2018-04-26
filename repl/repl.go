package repl

import (
	"io"
	"bufio"
	"fmt"
	"monkey/lexer"
	"monkey/parser"
)

const PROMPT = ">>"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned{
			return
		}
		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserError(out, p.Errors())
		}
		io.WriteString(out, program.String())
		io.WriteString(out, "\n")
	}
}
func printParserError(out io.Writer, errors []string){
	io.WriteString(out, "Woops! we ran into some monkey business here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors{
		io.WriteString(out, "\t" +msg+"\n")
	}
}
