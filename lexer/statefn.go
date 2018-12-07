package lexer

import (
	class "github.com/7thFox/starkscript/characterclass"
	"github.com/7thFox/starkscript/constants"
	"github.com/7thFox/starkscript/meta"
	"github.com/7thFox/starkscript/token"
)

type stateFn func(*lexer) stateFn

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

	return nil //lexUsage
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
