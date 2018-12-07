package starkscript

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/7thFox/startscript/token"
)

func TestMain(t *testing.T) {
	scriptString, _ := ioutil.ReadFile("./docs/example.stark")

	_, c := lex("foobar", string(scriptString))
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
