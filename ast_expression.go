package r

import "vimagination.zapto.org/parser"

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

type CompoundExpression struct{}

func (c *CompoundExpression) parse(r *rParser) error {
	return nil
}

type FlowControl struct{}

func (f *FlowControl) parse(r *rParser) error {
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
