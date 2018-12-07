package lexer

import (
	"fmt"

	"github.com/7thFox/startscript/meta"
	"github.com/7thFox/startscript/token"
)

func (l *lexer) errorf(format string, args ...interface{}) stateFn {
	l.tokens <- token.Token{
		Typ: token.Error,
		Val: fmt.Sprintf(format, args...),
	}
	return nil
}

func (l *lexer) errorExpected(meta meta.Meta) stateFn {
	k := l.pos + len(meta)
	if k >= len(l.input) {
		k = len(l.input) - 1
	}
	return l.errorf("Expected `%s', got '%s'", meta, l.input[l.pos:k])
}
