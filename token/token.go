package token

import "fmt"

const (
	Error TokenType = iota
	StartEndMeta
	Name
	Helptext
	Usage
)

type Token struct {
	Typ TokenType
	Val string
}

type TokenType int

func (t Token) String() string {
	switch t.Typ {
	case Error:
		return t.Val
	}
	if len(t.Val) > 10 {
		return fmt.Sprintf("%.10q...", t.Val)
	}
	return fmt.Sprintf("%q", t.Val)
}
