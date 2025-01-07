package r

import (
	"io"
	"strconv"

	"vimagination.zapto.org/parser"
)

const (
	whitespace      = "\t\v\f \u00a0\ufeff" // Tab, Vertical Tab, Form Feed, Space, No-break space, ZeroWidth No-Break Space, https://262.ecma-international.org/11.0/#table-32
	lineTerminators = "\n\r\u2028\u2029"    // Semi-Colon, Line Feed, Carriage Return, Line Separator, Paragraph Separator, https://262.ecma-international.org/11.0/#table-33
	octalDigit      = "01234567"
	decimalDigit    = "0123456789"
	hexDigit        = "0123456789abcdefABCDEF"
	identifierStart = "ABCDEFGHIJKLMNOPQRSTUVQXYZabcdefghijklmnopqrstuvwxyz"
	identifierCont  = "_.0123456789ABCDEFGHIJKLMNOPQRSTUVQXYZabcdefghijklmnopqrstuvwxyz"
)

// TokenType IDs
const (
	TokenWhitespace parser.TokenType = iota
	TokenLineTerminator
	TokenExpressionTerminator
	TokenComment
	TokenStringLiteral
	TokenNumericLiteral
	TokenIntegerLiteral
	TokenComplexLiteral
	TokenBooleanLiteral
	TokenNull
	TokenNA
	TokenIdentifier
	TokenKeyword
	TokenEllipsis
	TokenOperator
	TokenSpecialOperator
	TokenGrouping
)

// SetTokeniser sets the initial tokeniser state of a parser.Tokeniser.
//
// Used if you want to manually tokeniser R source code.
func SetTokeniser(t *parser.Tokeniser) *parser.Tokeniser {
	t.TokeniserState(new(rTokeniser).expression)

	return t
}

type rTokeniser struct {
	tokenDepth []byte
}

func (r *rTokeniser) lastDepth() byte {
	if len(r.tokenDepth) == 0 {
		return 0
	}

	d := r.tokenDepth[len(r.tokenDepth)-1]

	r.tokenDepth = r.tokenDepth[:len(r.tokenDepth)-1]

	return d
}

func (r *rTokeniser) expression(t *parser.Tokeniser) (parser.Token, parser.TokenFunc) {
	if t.Peek() == -1 {
		if len(r.tokenDepth) != 0 {
			return t.ReturnError(io.ErrUnexpectedEOF)
		}

		return t.Done()
	}

	if t.Accept(whitespace) {
		t.AcceptRun(whitespace)

		return t.Return(TokenWhitespace, r.expression)
	}

	if t.Accept(lineTerminators) {
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

	return r.operator(t)
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
			return r.ellipsisOrIdentifier(t)
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

func (r *rTokeniser) ellipsisOrIdentifier(t *parser.Tokeniser) (parser.Token, parser.TokenFunc) {
	if t.Accept(".") {
		if t.Accept(".") {
			if !t.Accept(identifierCont) {
				return t.Return(TokenEllipsis, r.expression)
			}
		} else if t.Accept(decimalDigit) {
			t.AcceptRun(decimalDigit)

			if !t.Accept(identifierCont) {
				return t.Return(TokenEllipsis, r.expression)
			}
		}
	}

	return r.identifier(t)
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
	case "if", "else", "repeat", "while", "function", "for", "in", "next", "break":
		tk.Type = TokenKeyword
	}

	return tk, r.expression
}

func (r *rTokeniser) operator(t *parser.Tokeniser) (parser.Token, parser.TokenFunc) {
	if t.Accept(";,") {
		return t.Return(TokenExpressionTerminator, r.expression)
	} else if t.Accept("+*/^~$@?") {
	} else if t.Accept(">=!") {
		t.Accept("=")
	} else if t.Accept("-") {
		t.Accept(">")
		t.Accept(">")
	} else if t.Accept("<") {
		if t.Accept("<") && !t.Accept("-") {
			return t.ReturnError(ErrInvalidOperator)
		} else {
			t.Accept("=-")
		}
	} else if t.Accept("%") {
		for ; t.Peek() != '%' && strconv.IsPrint(t.Peek()); t.Next() {
		}

		if !t.Accept("%") {
			return t.ReturnError(ErrInvalidOperator)
		}

		return t.Return(TokenSpecialOperator, r.expression)
	} else if t.Accept(":") {
		t.Accept(":")
	} else if t.Accept("&") {
		t.Accept("&")
	} else if t.Accept("|") {
		t.Accept("|>")
	} else if t.Accept("[") {
		if t.Accept("[") {
			r.tokenDepth = append(r.tokenDepth, '#')
		} else {
			r.tokenDepth = append(r.tokenDepth, ']')
		}

		return t.Return(TokenGrouping, r.expression)
	} else if t.Accept("(") {
		r.tokenDepth = append(r.tokenDepth, ')')

		return t.Return(TokenGrouping, r.expression)
	} else if t.Accept("{") {
		r.tokenDepth = append(r.tokenDepth, '}')

		return t.Return(TokenGrouping, r.expression)
	} else if t.Accept("]") {
		if g := r.lastDepth(); g == '#' {
			if !t.Accept("]") {
				return t.ReturnError(ErrInvalidCharacter)
			}
		} else if g != ']' {
			return t.ReturnError(ErrInvalidCharacter)
		}

		return t.Return(TokenGrouping, r.expression)
	} else if t.Accept(")") {
		if r.lastDepth() != ')' {
			return t.ReturnError(ErrInvalidCharacter)
		}

		return t.Return(TokenGrouping, r.expression)
	} else if t.Accept("}") {
		if r.lastDepth() != '}' {
			return t.ReturnError(ErrInvalidCharacter)
		}

		return t.Return(TokenGrouping, r.expression)
	} else {
		return t.ReturnError(ErrInvalidCharacter)
	}

	return t.Return(TokenOperator, r.expression)
}
