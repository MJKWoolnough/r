package r

import "vimagination.zapto.org/parser"

const (
	whitespace      = "\t\v\f \xa0\ufeff" // Tab, Vertical Tab, Form Feed, Space, No-break space, ZeroWidth No-Break Space, https://262.ecma-international.org/11.0/#table-32
	lineTerminators = "\n\r\u2028\u2029"  // Line Feed, Carriage Return, Line Separator, Paragraph Separator, https://262.ecma-international.org/11.0/#table-33
)

type rTokeniser struct {
	tokenDepth []byte
}

func (r *rTokeniser) expression(t *parser.Tokeniser) (parser.Token, parser.TokenFunc) {
	return t.Done()
}
