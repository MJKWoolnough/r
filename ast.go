// Package r implements an R tokeniser and parser.
package r

import (
	"vimagination.zapto.org/parser"
)

// File represents a parsed R file.
type File struct {
	Statements []Expression
	Tokens     Tokens
}

// Parse parses R input into AST.
func Parse(t Tokeniser) (*File, error) {
	r, err := newRParser(t)
	if err != nil {
		return nil, err
	}

	f := new(File)
	if err = f.parse(&r); err != nil {
		return nil, err
	}

	return f, nil
}

func (f *File) parse(r *rParser) error {
	for r.AcceptRunWhitespaceNoComment() != parser.TokenDone {
		var e Expression

		s := r.NewGoal()

		if err := e.parse(&s); err != nil {
			return r.Error("File", err)
		}

		f.Statements = append(f.Statements, e)

		r.Score(s)
		r.AcceptRunWhitespaceNoNewLine()

		if !r.AcceptToken(parser.Token{Type: TokenExpressionTerminator, Data: ";"}) && !r.Accept(TokenLineTerminator) && r.Peek().Type != parser.TokenDone {
			return r.Error("File", ErrMissingStatementTerminator)
		}
	}

	f.Tokens = r.ToTokens()

	return nil
}
