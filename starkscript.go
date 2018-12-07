package starkscript

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/7thFox/starkscript/lexer"
	"github.com/7thFox/starkscript/token"
)

func Lex(name string, input string) {
	lexer.Lex(name, input)
	fmt.Println("Foobar")
}

func main() {
	scriptString, _ := ioutil.ReadFile("./example.stark")

	_, c := lexer.Lex("foobar", string(scriptString))
	for itm := range c {
		switch itm.Typ {
		case token.Name:
			fmt.Printf("Name: \"%s\"\n", itm.Val)
		case token.Helptext:
			parsed := ""
			previousNewline := true
			for _, line := range strings.Split(itm.Val, "\n") {
				line = strings.TrimSpace(line)
				if line == "" {
					if !previousNewline {
						parsed += "\n"
					}
					previousNewline = true
				} else {
					if !previousNewline {
						parsed += " "
					}
					parsed += line
					previousNewline = false
				}
			}
			itm.Val = parsed
			fmt.Printf("Helptext: \n%v\n", itm.Val)
		}
	}
}
