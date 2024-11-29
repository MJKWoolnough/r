package r

import (
	"errors"

	"vimagination.zapto.org/parser"
)

const (
	whitespace      = "\t\v\f \xa0\ufeff" // Tab, Vertical Tab, Form Feed, Space, No-break space, ZeroWidth No-Break Space, https://262.ecma-international.org/11.0/#table-32
	lineTerminators = "\n\r\u2028\u2029"  // Line Feed, Carriage Return, Line Separator, Paragraph Separator, https://262.ecma-international.org/11.0/#table-33
	octalDigit      = "01234567"
	decimalDigit    = "0123456789"
	hexDigit        = "0123456789abcdefABCDEF"
	identifierStart = "ABCDEFGHIJKLMNOPQRSTUVQXYZabcdefghijklmnopqrstuvwxyz"
	identifierCont  = "_.0123456789ABCDEFGHIJKLMNOPQRSTUVQXYZabcdefghijklmnopqrstuvwxyz"
)

const (
	TokenWhitespace parser.TokenType = iota
	TokenLineTerminator
	TokenComment
	TokenStringLiteral
	TokenNumericLiteral
	TokenIntegerLiteral
	TokenComplexLiteral
	TokenBooleanLiteral
	TokenNull
	TokenNA
	TokenIdentifier
)

type rTokeniser struct {
	tokenDepth []byte
}

func (r *rTokeniser) expression(t *parser.Tokeniser) (parser.Token, parser.TokenFunc) {
	if t.Accept(whitespace) {
		t.AcceptRun(whitespace)

		return t.Return(TokenWhitespace, r.expression)
	}

	if t.Accept(lineTerminators) {
		t.AcceptRun(lineTerminators)

		return t.Return(TokenLineTerminator, r.expression)
	}

	if t.Accept("#") {
		t.ExceptRun(lineTerminators)

		return t.Return(TokenComment, r.expression)
	}

	if t.Accept(identifierStart) {
		return r.identifier(t)
	}

	if c := t.Peek(); c == '"' || c == '\'' {
		return r.string(t)
	} else if c == '.' || c >= '0' && c <= '9' {
		return r.number(t)
	}

	return t.Done()
}

func (r *rTokeniser) string(t *parser.Tokeniser) (parser.Token, parser.TokenFunc) {
	quote := t.Next()

	chars := string(quote) + "\\"

	for {
		switch t.ExceptRun(chars) {
		case quote:
			t.Next()

			return t.Return(TokenStringLiteral, r.expression)
		case '\\':
			t.Next()

			switch t.Peek() {
			case '\'', '"', 'n', 'r', 't', 'b', 'a', 'f', 'v', '\\':
				t.Next()
			case '0', '1', '2', '3', '4', '5', '6', '7':
				t.Next()

				if !t.Accept(octalDigit) || !t.Accept(octalDigit) {
					return t.ReturnError(ErrInvalidString)
				}
			case 'x':
				t.Next()

				if !t.Accept(hexDigit) || !t.Accept(hexDigit) {
					return t.ReturnError(ErrInvalidString)
				}
			case 'u':
				t.Next()

				brace := t.Accept("{")

				if !t.Accept(hexDigit) {
					return t.ReturnError(ErrInvalidString)
				}

				if brace {

					t.Accept(hexDigit)
					t.Accept(hexDigit)
					t.Accept(hexDigit)

					if !t.Accept("}") {
						return t.ReturnError(ErrInvalidString)
					}
				} else if !t.Accept(hexDigit) || !t.Accept(hexDigit) || !t.Accept(hexDigit) {
					return t.ReturnError(ErrInvalidString)
				}
			case 'U':
				t.Next()

				brace := t.Accept("{")

				if !t.Accept(hexDigit) {
					return t.ReturnError(ErrInvalidString)
				}

				if brace {
					t.Accept(hexDigit)
					t.Accept(hexDigit)
					t.Accept(hexDigit)
					t.Accept(hexDigit)
					t.Accept(hexDigit)
					t.Accept(hexDigit)
					t.Accept(hexDigit)

					if !t.Accept("}") {
						return t.ReturnError(ErrInvalidString)
					}
				} else if !t.Accept(hexDigit) || !t.Accept(hexDigit) || !t.Accept(hexDigit) || !t.Accept(hexDigit) || !t.Accept(hexDigit) || !t.Accept(hexDigit) || !t.Accept(hexDigit) {
					return t.ReturnError(ErrInvalidString)
				}
			default:
				return t.ReturnError(ErrInvalidString)
			}
		case -1:
			return t.ReturnError(ErrInvalidString)
		}
	}
}

func (r *rTokeniser) number(t *parser.Tokeniser) (parser.Token, parser.TokenFunc) {
	if t.Accept(".") {
		if !t.Accept(decimalDigit) {
			return r.identifier(t)
		}

		return r.float(t, decimalDigit)
	}

	digits := decimalDigit

	if t.Accept("0") {
		if t.Accept("x") {
			digits = hexDigit

			if !t.Accept(digits) {
				return t.ReturnError(ErrInvalidNumber)
			}
		}
	}

	t.AcceptRun(digits)

	if t.Accept("L") {
		return t.Return(TokenIntegerLiteral, r.expression)
	} else if t.Accept("i") {
		return t.Return(TokenComplexLiteral, r.expression)
	} else if t.Accept(".") {
		return r.float(t, digits)
	}

	return r.exponential(t, digits)
}

func (r *rTokeniser) float(t *parser.Tokeniser, digits string) (parser.Token, parser.TokenFunc) {
	if digits == hexDigit && !t.Accept(digits) {
		return t.ReturnError(ErrInvalidNumber)
	}

	t.AcceptRun(digits)

	if t.Accept("L") {
		return t.Return(TokenIntegerLiteral, r.expression)
	} else if t.Accept("i") {
		return t.Return(TokenComplexLiteral, r.expression)
	}

	return r.exponential(t, digits)
}

func (r *rTokeniser) exponential(t *parser.Tokeniser, digits string) (parser.Token, parser.TokenFunc) {
	e := "eE"

	if digits == hexDigit {
		e = "pP"
	}

	if t.Accept(e) {
		t.Accept("+-")

		if !t.Accept(digits) {
			return t.ReturnError(ErrInvalidNumber)
		}

		t.AcceptRun(digits)
	}

	if t.Accept("L") {
		return t.Return(TokenIntegerLiteral, r.expression)
	} else if t.Accept("i") {
		return t.Return(TokenComplexLiteral, r.expression)
	}

	return t.Return(TokenNumericLiteral, r.expression)
}

func (r *rTokeniser) identifier(t *parser.Tokeniser) (parser.Token, parser.TokenFunc) {
	t.AcceptRun(identifierCont)

	tk := parser.Token{Type: TokenIdentifier, Data: t.Get()}

	switch tk.Data {
	case "NULL":
		tk.Type = TokenNull
	case "TRUE", "FALSE":
		tk.Type = TokenBooleanLiteral
	case "Inf", "NaN":
		tk.Type = TokenNumericLiteral
	case "NA", "NA_character_", "NA_integer_", "NA_real_", "NA_complex_":
		tk.Type = TokenNA
	}

	return tk, r.expression
}

var (
	ErrInvalidString = errors.New("invalid string")
	ErrInvalidNumber = errors.New("invalid number")
)
