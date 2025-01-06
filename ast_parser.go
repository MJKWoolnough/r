package r

import (
	"fmt"

	"vimagination.zapto.org/parser"
)

// Token represents a single parsed token with source positioning.
type Token struct {
	parser.Token
	Pos, Line, LinePos uint64
}

// Tokens is a collection of Token values.
type Tokens []Token

type rParser Tokens

type Comments []Token

// Tokeniser is an interface representing a tokeniser.
type Tokeniser interface {
	TokeniserState(parser.TokenFunc)
	Iter(func(parser.Token) bool)
	GetError() error
}

func newRParser(t Tokeniser) (rParser, error) {
	t.TokeniserState(new(rTokeniser).expression)

	var (
		tokens             rParser
		pos, line, linePos uint64
		err                error
	)

	for tk := range t.Iter {
		tokens = append(tokens, Token{
			Token:   tk,
			Pos:     pos,
			Line:    line,
			LinePos: linePos,
		})

		switch tk.Type {
		case parser.TokenError:
			err = Error{
				Err:     t.GetError(),
				Parsing: "Tokens",
				Token:   tokens[len(tokens)-1],
			}
		case TokenLineTerminator:
			var lastChar rune

			for _, c := range tk.Data {
				if lastChar != '\r' || c != '\n' {
					line++
				}

				lastChar = c
			}

			linePos = 0
		default:
			linePos += uint64(len(tk.Data))
		}

		pos += uint64(len(tk.Data))
	}

	return tokens[0:0:len(tokens)], err
}

func (r rParser) NewGoal() rParser {
	return r[len(r):]
}

func (r *rParser) Score(k rParser) {
	*r = (*r)[:len(*r)+len(k)]
}

func (r *rParser) next() Token {
	l := len(*r)
	*r = (*r)[:l+1]
	tk := (*r)[l]

	return tk
}

func (r *rParser) backup() {
	*r = (*r)[:len(*r)-1]
}

func (r *rParser) Peek() parser.Token {
	tk := r.next().Token

	r.backup()

	return tk
}

func (r *rParser) Accept(ts ...parser.TokenType) bool {
	tt := r.next().Type

	for _, pt := range ts {
		if pt == tt {
			return true
		}
	}

	r.backup()

	return false
}

func (r *rParser) AcceptRun(ts ...parser.TokenType) parser.TokenType {
Loop:
	for {
		tt := r.next().Type

		for _, pt := range ts {
			if pt == tt {
				continue Loop
			}
		}

		r.backup()

		return tt
	}
}

func (r *rParser) AcceptToken(tk parser.Token) bool {
	if r.next().Token == tk {
		return true
	}

	r.backup()

	return false
}

func (r *rParser) ToTokens() Tokens {
	return Tokens((*r)[:len(*r):len(*r)])
}

func (r *rParser) AcceptRunWhitespace() parser.TokenType {
	return r.AcceptRun(TokenWhitespace, TokenLineTerminator, TokenComment)
}

func (r *rParser) AcceptRunWhitespaceNoComment() parser.TokenType {
	return r.AcceptRun(TokenWhitespace, TokenLineTerminator)
}

func (r *rParser) AcceptRunWhitespaceComments() Comments {
	var c Comments

	for r.AcceptRun(TokenWhitespace, TokenLineTerminator) == TokenComment {
		c = append(c, r.next())
	}

	return c
}

func (r *rParser) AcceptRunWhitespaceNoNewLine() parser.TokenType {
	return r.AcceptRun(TokenWhitespace)
}

func (r *rParser) GetLastToken() *Token {
	return &(*r)[len(*r)-1]
}

// Error is a parsing error with trace details.
type Error struct {
	Err     error
	Parsing string
	Token   Token
}

// Error returns the error string.
func (e Error) Error() string {
	return fmt.Sprintf("%s: error at position %d (%d:%d):\n%s", e.Parsing, e.Token.Pos+1, e.Token.Line+1, e.Token.LinePos+1, e.Err)
}

// Unwrap returns the wrapped error.
func (e Error) Unwrap() error {
	return e.Err
}

func (r *rParser) Error(parsingFunc string, err error) error {
	tk := r.next()

	r.backup()

	return Error{
		Err:     err,
		Parsing: parsingFunc,
		Token:   tk,
	}
}
