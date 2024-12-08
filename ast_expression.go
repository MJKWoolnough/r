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

type IfControl struct {
	Cond   ConditionalExpression
	Expr   Expression
	Else   *Expression
	Tokens Tokens
}

func (i *IfControl) parse(r *rParser) error {
	r.AcceptToken(parser.Token{Type: TokenKeyword, Data: "if"})
	r.AcceptRunWhitespaceNoNewLine()

	if !r.AcceptToken(parser.Token{Type: TokenGrouping, Data: "("}) {
		return r.Error("IfControl", ErrMissingOpeningParen)
	}

	r.AcceptRunWhitespace()

	s := r.NewGoal()

	if err := i.Cond.parse(&s); err != nil {
		return r.Error("IfControl", err)
	}

	r.Score(s)
	r.AcceptRunWhitespaceNoNewLine()

	if !r.AcceptToken(parser.Token{Type: TokenGrouping, Data: ")"}) {
		return r.Error("IfControl", ErrMissingClosingParen)
	}

	r.AcceptRunWhitespaceNoNewLine()

	s = r.NewGoal()

	if err := i.Expr.parse(&s); err != nil {
		return r.Error("IfControl", err)
	}

	r.Score(s)

	s = r.NewGoal()

	s.AcceptRunWhitespaceNoNewLine()

	if s.AcceptToken(parser.Token{Type: TokenKeyword, Data: "else"}) {
		s.AcceptRunWhitespaceNoNewLine()

		t := s.NewGoal()
		i.Else = new(Expression)

		if err := i.Else.parse(&t); err != nil {
			return r.Error("IfControl", err)
		}

		s.Score(t)
		r.Score(t)
	}

	i.Tokens = r.ToTokens()

	return nil
}

type WhileControl struct {
	Cond   ConditionalExpression
	Expr   Expression
	Tokens Tokens
}

func (w *WhileControl) parse(r *rParser) error {
	r.AcceptToken(parser.Token{Type: TokenKeyword, Data: "while"})
	r.AcceptRunWhitespaceNoNewLine()

	if !r.AcceptToken(parser.Token{Type: TokenGrouping, Data: "("}) {
		return r.Error("WhileControl", ErrMissingOpeningParen)
	}

	r.AcceptRunWhitespace()

	s := r.NewGoal()

	if err := w.Cond.parse(&s); err != nil {
		return r.Error("WhileControl", err)
	}

	r.Score(s)
	r.AcceptRunWhitespaceNoNewLine()

	if !r.AcceptToken(parser.Token{Type: TokenGrouping, Data: ")"}) {
		return r.Error("WhileControl", ErrMissingClosingParen)
	}

	r.AcceptRunWhitespaceNoNewLine()

	s = r.NewGoal()

	if err := w.Expr.parse(&s); err != nil {
		return r.Error("WhileControl", err)
	}

	r.Score(s)

	w.Tokens = r.ToTokens()

	return nil
}

type RepeatControl struct {
	Cond   ConditionalExpression
	Tokens Tokens
}

func (rc *RepeatControl) parse(r *rParser) error {
	r.AcceptToken(parser.Token{Type: TokenKeyword, Data: "repeat"})
	r.AcceptRunWhitespaceNoNewLine()

	s := r.NewGoal()

	if err := rc.Cond.parse(&s); err != nil {
		return r.Error("RepeatControl", err)
	}

	r.Score(s)

	rc.Tokens = r.ToTokens()

	return nil
}

type ForControl struct {
	Var    Atom
	List   ConditionalExpression
	Expr   Expression
	Tokens Tokens
}

func (f *ForControl) parse(r *rParser) error {
	r.AcceptToken(parser.Token{Type: TokenKeyword, Data: "for"})
	r.AcceptRunWhitespaceNoNewLine()

	if !r.AcceptToken(parser.Token{Type: TokenGrouping, Data: "("}) {
		return r.Error("ForControl", ErrMissingOpeningParen)
	}

	r.AcceptRunWhitespace()

	s := r.NewGoal()

	if err := f.Var.parse(&s); err != nil {
		return r.Error("ForControl", err)
	}

	r.Score(s)
	r.AcceptRunWhitespaceNoNewLine()

	if !r.AcceptToken(parser.Token{Type: TokenKeyword, Data: "in"}) {
		return r.Error("ForControl", ErrMissingIn)
	}

	r.AcceptRunWhitespaceNoNewLine()

	s = r.NewGoal()

	if err := f.List.parse(&s); err != nil {
		return r.Error("ForControl", err)
	}

	r.Score(s)
	r.AcceptRunWhitespaceNoNewLine()

	if !r.AcceptToken(parser.Token{Type: TokenGrouping, Data: ")"}) {
		return r.Error("ForControl", ErrMissingClosingParen)
	}

	r.AcceptRunWhitespaceNoNewLine()

	s = r.NewGoal()

	if err := f.Expr.parse(&s); err != nil {
		return r.Error("ForControl", err)
	}

	r.Score(s)

	f.Tokens = r.ToTokens()

	return nil
}

type FunctionDefinition struct {
	ArgList ArgList
	Body    Expression
	Tokens  Tokens
}

func (f *FunctionDefinition) parse(r *rParser) error {
	r.AcceptToken(parser.Token{Type: TokenKeyword, Data: "function"})
	r.AcceptRunWhitespaceNoNewLine()

	if !r.AcceptToken(parser.Token{Type: TokenGrouping, Data: "("}) {
		return r.Error("FunctionDefinition", ErrMissingOpeningParen)
	}

	r.AcceptRunWhitespace()

	s := r.NewGoal()

	if err := f.ArgList.parse(&s); err != nil {
		return r.Error("FunctionDefinition", err)
	}

	r.Score(s)
	r.AcceptRunWhitespaceNoNewLine()

	if !r.AcceptToken(parser.Token{Type: TokenGrouping, Data: ")"}) {
		return r.Error("FunctionDefinition", ErrMissingClosingParen)
	}

	r.AcceptRunWhitespaceNoNewLine()

	s = r.NewGoal()

	if err := f.Body.parse(&s); err != nil {
		return r.Error("FunctionDefinition", err)
	}

	r.Score(s)

	f.Tokens = r.ToTokens()

	return nil
}

type ArgList struct {
	Args   []Argument
	Tokens Tokens
}

func (a *ArgList) parse(r *rParser) error {
	s := r.NewGoal()

	for {
		var arg Argument

		if err := arg.parse(&s); err != nil {
			return r.Error("ArgList", err)
		}

		r.Score(s)

		s = r.NewGoal()

		s.AcceptRunWhitespaceNoNewLine()

		if s.Peek() == (parser.Token{Type: TokenGrouping, Data: ")"}) {
			break
		} else if !s.AcceptToken(parser.Token{Type: TokenExpressionTerminator, Data: ","}) {
			return s.Error("ArgList", ErrMissingTerminator)
		}

		s.AcceptRunWhitespaceNoNewLine()
		r.Score(s)

		s = r.NewGoal()
	}

	a.Tokens = r.ToTokens()

	return nil
}

type Argument struct {
	Identifier *Token
	Default    *Expression
	Tokens     Tokens
}

func (a *Argument) parse(r *rParser) error {
	if !r.Accept(TokenIdentifier) && !r.AcceptToken(parser.Token{Type: TokenEllipsis, Data: "..."}) {
		return r.Error("Argument", ErrMissingIdentifier)
	}

	a.Identifier = r.GetLastToken()

	if a.Identifier.Type == TokenIdentifier {
		s := r.NewGoal()

		s.AcceptRunWhitespaceNoNewLine()

		if s.AcceptToken(parser.Token{Type: TokenOperator, Data: "="}) {
			s.AcceptRunWhitespaceNoNewLine()

			r.Score(s)

			s = r.NewGoal()
			a.Default = new(Expression)

			if err := a.Default.parse(&s); err != nil {
				return r.Error("Argument", err)
			}

			r.Score(s)
		}
	}

	a.Tokens = r.ToTokens()

	return nil
}

type AssignmentExpression struct {
	ConditionalExpression ConditionalExpression
	AssignmentExpression  *AssignmentExpression
	Tokens                Tokens
}

func (a *AssignmentExpression) parse(r *rParser) error {
	s := r.NewGoal()

	if err := a.ConditionalExpression.parse(&s); err != nil {
		return r.Error("AssignmentExpression", err)
	}

	r.Score(s)

	s = r.NewGoal()

	s.AcceptRunWhitespaceNoNewLine()

	if s.AcceptToken(parser.Token{Type: TokenOperator, Data: "="}) || s.AcceptToken(parser.Token{Type: TokenOperator, Data: "<-"}) || s.AcceptToken(parser.Token{Type: TokenOperator, Data: "<<-"}) || s.AcceptToken(parser.Token{Type: TokenOperator, Data: "->"}) || s.AcceptToken(parser.Token{Type: TokenOperator, Data: "->>"}) {
		s.AcceptRunWhitespaceNoNewLine()

		r.Score(s)

		s = r.NewGoal()
		a.AssignmentExpression = new(AssignmentExpression)

		if err := a.AssignmentExpression.parse(&s); err != nil {
			return r.Error("AssignmentExpression", err)
		}

		r.Score(s)
	}

	a.Tokens = r.ToTokens()

	return nil
}

type ConditionalExpression struct{}

func (c *ConditionalExpression) parse(r *rParser) error {
	return nil
}

type Atom struct{}

func (a *Atom) parse(r *rParser) error {
	return nil
}

var (
	ErrMissingTerminator   = errors.New("missing terminator")
	ErrMissingOpeningParen = errors.New("missing opening paren")
	ErrMissingClosingParen = errors.New("missing closing paren")
	ErrMissingIn           = errors.New("missing in keyword")
	ErrMissingIdentifier   = errors.New("missing identifier")
)
