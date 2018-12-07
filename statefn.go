package starkscript

import (
	"fmt"

	class "github.com/7thFox/startscript/characterclass"
	"github.com/7thFox/startscript/constants"
	"github.com/7thFox/startscript/meta"
	"github.com/7thFox/startscript/token"
)

type stateFn func(*lexer) stateFn

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

func lexStartEndMeta(l *lexer) stateFn {
	if l.ignoreString(meta.StartEnd) {
		l.ignoreRun(class.WhitespaceMultiLine)
		return lexName
	}
	return l.errorExpected(meta.StartEnd)
}

func lexName(l *lexer) stateFn {
	if !l.ignoreString(meta.Name) {
		return l.errorExpected(meta.Name)
	}
	l.ignoreRun(class.Whitespace)

	read := l.acceptRun(class.Lowercase)
	if read == 0 {
		return l.errorf("Invalid name character %c", l.next())
	} else if read < constants.NameMinimumLength {
		return l.errorf("Name must be at least %d characters", constants.NameMinimumLength)
	}

	l.emit(token.Name)
	l.ignoreRun(class.WhitespaceMultiLine)

	return lexHelptext
}

func lexHelptext(l *lexer) stateFn {
	if !l.ignoreString(meta.Helptext) {
		return l.errorExpected(meta.Helptext)
	}
	l.ignoreRun(class.Whitespace)
	if !l.ignoreString(meta.StringMultiLineOpen) {
		return l.errorExpected(meta.StringMultiLineOpen)
	}

	l.acceptRun(class.StringContentMultiLine)

	if !l.peekMatches(meta.StringMultiLineClose) {
		return l.errorExpected(meta.StringMultiLineClose)
	}
	l.emit(token.Helptext)

	return nil
}
