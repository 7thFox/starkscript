package lexer

import (
	"strings"
	"unicode/utf8"

	class "github.com/7thFox/startscript/characterclass"
	"github.com/7thFox/startscript/meta"
	"github.com/7thFox/startscript/token"
)

type lexer struct {
	name   string
	input  string
	start  int
	pos    int
	width  int
	tokens chan token.Token

	whitespace map[rune]bool
}

func Lex(name string, input string) (*lexer, chan token.Token) {
	l := &lexer{
		name:   name,
		input:  input,
		tokens: make(chan token.Token),
	}
	go l.run()
	return l, l.tokens
}

func (l *lexer) run() {
	l.ignoreRun(class.WhitespaceMultiLine)
	for state := lexStartEndMeta; state != nil; {
		state = state(l)
	}
	close(l.tokens)
}

func (l *lexer) emit(t token.TokenType) {
	l.tokens <- token.Token{Typ: t, Val: l.input[l.start:l.pos]}
	l.start = l.pos
}

func (l *lexer) next() (r rune) {
	if l.pos >= len(l.input) {
		l.width = 0
		return 0
	}
	r, l.width = utf8.DecodeRuneInString(l.input[l.pos:])
	l.pos += l.width
	return r
}

func (l *lexer) backup() {
	l.pos -= l.width
}

func (l *lexer) ignore() {
	l.start = l.pos
}

func (l *lexer) ignoreRun(valid class.CharacterClass) (count int) {
	for strings.IndexRune(string(valid), l.next()) >= 0 {
		count++
		l.ignore()
	}
	l.backup()
	return
}

func (l *lexer) matches(valid meta.Meta) bool {
	if l.peekMatches(valid) {
		l.pos += len(valid)
		return true
	}
	return false
}

func (l *lexer) peekMatches(valid meta.Meta) bool {
	return strings.HasPrefix(l.input[l.pos:], string(valid))
}

func (l *lexer) ignoreString(valid meta.Meta) bool {
	if l.matches(valid) {
		l.ignore()
		return true
	}
	return false
}

func (l *lexer) accept(valid class.CharacterClass) bool {
	if strings.IndexRune(string(valid), l.next()) >= 0 {
		return true
	}
	l.backup()
	return false
}

func (l *lexer) acceptRun(valid class.CharacterClass) (count int) {
	for strings.IndexRune(string(valid), l.next()) >= 0 {
		count++
	}
	l.backup()
	return
}
