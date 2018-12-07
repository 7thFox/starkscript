package meta

type Meta string

const (
	StartEnd Meta = "```"
	Name     Meta = "NAME:"
	Helptext Meta = "HELPTEXT"

	StringMultiLineOpen  Meta = "["
	StringMultiLineClose Meta = "]"
)
