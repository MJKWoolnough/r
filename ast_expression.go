package r

import (
	"errors"

	"vimagination.zapto.org/parser"
)

type Expression struct {
	CompoundExpression   *CompoundExpression
	FlowControl          *FlowControl
	FunctionDefinition   *FunctionDefinition
	AssignmentExpression *AssignmentExpression
	Tokens               Tokens
}

func (e *Expression) parse(r *rParser) error {
	var err error

	s := r.NewGoal()

	switch tk := r.Peek(); tk {
	case parser.Token{Type: TokenGrouping, Data: "{"}:
		err = e.CompoundExpression.parse(&s)
	case parser.Token{Type: TokenKeyword, Data: "if"}, parser.Token{Type: TokenKeyword, Data: "while"}, parser.Token{Type: TokenKeyword, Data: "repeat"}, parser.Token{Type: TokenKeyword, Data: "for"}:
		err = e.FlowControl.parse(&s)
	case parser.Token{Type: TokenKeyword, Data: "function"}:
		err = e.FunctionDefinition.parse(&s)
	default:
		err = e.AssignmentExpression.parse(&s)
	}

	if err != nil {
		return r.Error("Expression", err)
	}

	r.Score(s)

	e.Tokens = r.ToTokens()

	return nil
}

type CompoundExpression struct {
	Expressions []Expression
	Tokens      Tokens
}

func (c *CompoundExpression) parse(r *rParser) error {
	r.AcceptToken(parser.Token{Type: TokenGrouping, Data: "{"})
	r.AcceptRunWhitespace()

	for !r.AcceptToken(parser.Token{Type: TokenGrouping, Data: "}"}) {
		s := r.NewGoal()

		var e Expression

		if err := e.parse(&s); err != nil {
			return s.Error("CompoundExpression", err)
		}

		c.Expressions = append(c.Expressions, e)

		r.Score(s)
		r.AcceptRunWhitespaceNoNewLine()

		if s.AcceptToken(parser.Token{Type: TokenGrouping, Data: "}"}) {
			break
		} else if !s.Accept(TokenLineTerminator, TokenExpressionTerminator) {
			return s.Error("CompoundExpression", ErrMissingTerminator)
		}

		r.Score(s)
		r.AcceptRunWhitespace()
	}

	c.Tokens = r.ToTokens()

	return nil
}

type FlowControl struct {
	IfControl     *IfControl
	WhileControl  *WhileControl
	RepeatControl *RepeatControl
	ForControl    *ForControl
	Tokens        Tokens
}

func (f *FlowControl) parse(r *rParser) error {
	var err error

	s := r.NewGoal()
	switch r.Peek() {
	case parser.Token{Type: TokenKeyword, Data: "if"}:
		f.IfControl = new(IfControl)

		err = f.IfControl.parse(&s)
	case parser.Token{Type: TokenKeyword, Data: "while"}:
		f.WhileControl = new(WhileControl)

		err = f.WhileControl.parse(&s)
	case parser.Token{Type: TokenKeyword, Data: "repeat"}:
		f.RepeatControl = new(RepeatControl)

		err = f.RepeatControl.parse(&s)
	case parser.Token{Type: TokenKeyword, Data: "for"}:
		f.ForControl = new(ForControl)

		err = f.ForControl.parse(&s)
	}

	if err != nil {
		return r.Error("FlowControl", err)
	}

	r.Score(s)

	f.Tokens = r.ToTokens()

	return nil
}

type IfControl struct{}

func (i *IfControl) parse(r *rParser) error {
	return nil
}

type WhileControl struct{}

func (w *WhileControl) parse(r *rParser) error {
	return nil
}

type RepeatControl struct{}

func (rc *RepeatControl) parse(r *rParser) error {
	return nil
}

type ForControl struct{}

func (f *ForControl) parse(r *rParser) error {
	return nil
}

type FunctionDefinition struct{}

func (f *FunctionDefinition) parse(r *rParser) error {
	return nil
}

type AssignmentExpression struct{}

func (a *AssignmentExpression) parse(r *rParser) error {
	return nil
}

var ErrMissingTerminator = errors.New("missing terminator")
