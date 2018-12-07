package characterclass

type CharacterClass string

const (
	Whitespace          CharacterClass = " \t"
	Newline             CharacterClass = "\r\n"
	WhitespaceMultiLine CharacterClass = Whitespace + Newline

	Digit CharacterClass = "0123456789"
	Hex   CharacterClass = Digit + "ABCDEF"

	Lowercase    CharacterClass = "abcdefghijklmnopqrstuvwxyz"
	Uppercase    CharacterClass = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Alpha        CharacterClass = Lowercase + Uppercase
	Alphanumeric CharacterClass = Alpha + Digit
	Special      CharacterClass = "!@#$%^&*()+-.,<>/?';:{}\\|`~"
	SpecialAll   CharacterClass = Special + "[]\""

	StringContentSingleLine CharacterClass = Alphanumeric + Special + "[]" + Whitespace
	StringContentMultiLine  CharacterClass = Alphanumeric + Special + "\"" + WhitespaceMultiLine
)
