package r

import (
	"errors"

	"vimagination.zapto.org/parser"
)

type Expression struct {
	FlowControl          *FlowControl
	FunctionDefinition   *FunctionDefinition
	AssignmentExpression *AssignmentExpression
	Tokens               Tokens
}

func (e *Expression) parse(r *rParser) error {
	var err error

	s := r.NewGoal()

	switch tk := r.Peek(); tk {
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
	Cond   FormulaeExpression
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
	Cond   FormulaeExpression
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
	Cond   FormulaeExpression
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
	Var    SimpleExpression
	List   FormulaeExpression
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

type QueryType uint8

const (
	QueryNone QueryType = iota
	QueryUnary
	QueryBinary
)

type QueryExpression struct {
	QueryType            QueryType
	AssignmentExpression AssignmentExpression
	QueryExpression      *QueryExpression
	Tokens               Tokens
}

func (q *QueryExpression) parse(r *rParser) error {
	if r.AcceptToken(parser.Token{Type: TokenOperator, Data: "?"}) {
		q.QueryType = QueryUnary

		r.AcceptRunWhitespaceNoNewLine()

		s := r.NewGoal()
		q.QueryExpression = new(QueryExpression)

		if err := q.QueryExpression.parse(&s); err != nil {
			return r.Error("QueryExpression", err)
		}

		r.Score(s)
	} else {
		s := r.NewGoal()

		if err := q.AssignmentExpression.parse(&s); err != nil {
			return r.Error("QueryExpression", err)
		}

		r.Score(s)

		s = r.NewGoal()

		s.AcceptRunWhitespaceNoNewLine()

		if s.AcceptToken(parser.Token{Type: TokenOperator, Data: "?"}) {
			s.AcceptRunWhitespaceNoNewLine()

			r.Score(s)

			s = r.NewGoal()
			q.QueryType = QueryBinary
			q.QueryExpression = new(QueryExpression)

			if err := q.QueryExpression.parse(&s); err != nil {
				return r.Error("QueryExpression", err)
			}

			r.Score(s)
		}
	}

	q.Tokens = r.ToTokens()

	return nil
}

type AssignmentExpression struct {
	ConditionalExpression FormulaeExpression
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

type FormulaeExpression struct {
	OrExpression       *OrExpression
	FormulaeExpression *FormulaeExpression
	Tokens             Tokens
}

func (f *FormulaeExpression) parse(r *rParser) error {
	s := r.NewGoal()

	if s.Peek() != (parser.Token{Type: TokenOperator, Data: "~"}) {
		f.OrExpression = new(OrExpression)

		if err := f.OrExpression.parse(&s); err != nil {
			return r.Error("FormulaeExpression", err)
		}

		r.Score(s)

		s = r.NewGoal()

		s.AcceptRunWhitespaceNoNewLine()
	}

	if s.AcceptToken(parser.Token{Type: TokenOperator, Data: "~"}) {
		s.AcceptRunWhitespaceNoNewLine()

		r.Score(s)

		s = r.NewGoal()
		f.FormulaeExpression = new(FormulaeExpression)

		if err := f.FormulaeExpression.parse(&s); err != nil {
			return r.Error("FormulaeExpression", err)
		}

		r.Score(s)
	}

	f.Tokens = r.ToTokens()

	return nil
}

type OrType uint8

const (
	OrNone OrType = iota
	OrVectorized
	OrNotVectorized
)

type OrExpression struct {
	AndExpression AndExpression
	OrType        OrType
	OrExpression  *OrExpression
	Tokens        Tokens
}

func (o *OrExpression) parse(r *rParser) error {
	s := r.NewGoal()

	if err := o.AndExpression.parse(&s); err != nil {
		return r.Error("OrExpression", err)
	}

	r.Score(s)

	s = r.NewGoal()

	s.AcceptRunWhitespaceNoNewLine()

	if s.AcceptToken(parser.Token{Type: TokenOperator, Data: "|"}) {
		o.OrType = OrVectorized
	} else if s.AcceptToken(parser.Token{Type: TokenOperator, Data: "||"}) {
		o.OrType = OrNotVectorized
	}

	if o.OrType != OrNone {
		s.AcceptRunWhitespace()
		r.Score(s)

		s = r.NewGoal()
		o.OrExpression = new(OrExpression)

		if err := o.OrExpression.parse(&s); err != nil {
			return r.Error("OrExpression", err)
		}

		r.Score(s)
	}

	o.Tokens = r.ToTokens()

	return nil
}

type AndType uint8

const (
	AndNone AndType = iota
	AndVectorized
	AndNotVectorized
)

type AndExpression struct {
	NotExpression NotExpression
	AndType       AndType
	AndExpression *AndExpression
	Tokens        Tokens
}

func (a *AndExpression) parse(r *rParser) error {
	s := r.NewGoal()

	if err := a.NotExpression.parse(&s); err != nil {
		return r.Error("AndExpression", err)
	}

	r.Score(s)

	s = r.NewGoal()

	s.AcceptRunWhitespaceNoNewLine()

	if s.AcceptToken(parser.Token{Type: TokenOperator, Data: "&"}) {
		a.AndType = AndVectorized
	} else if s.AcceptToken(parser.Token{Type: TokenOperator, Data: "&&"}) {
		a.AndType = AndNotVectorized
	}

	if a.AndType != AndNone {
		s.AcceptRunWhitespace()
		r.Score(s)

		s = r.NewGoal()
		a.AndExpression = new(AndExpression)

		if err := a.NotExpression.parse(&s); err != nil {
			return r.Error("AndExpression", err)
		}

		r.Score(s)
	}

	a.Tokens = r.ToTokens()

	return nil
}

type NotExpression struct {
	Not                  bool
	RelationalExpression RelationalExpression
	Tokens               Tokens
}

func (n *NotExpression) parse(r *rParser) error {
	n.Not = r.AcceptToken(parser.Token{Type: TokenOperator, Data: "!"})

	r.AcceptRunWhitespaceNoNewLine()

	s := r.NewGoal()

	if err := n.RelationalExpression.parse(&s); err != nil {
		return r.Error("NotExpression", err)
	}

	r.Score(s)

	n.Tokens = r.ToTokens()

	return nil
}

type RelationalOperator uint8

const (
	RelationalNone RelationalOperator = iota
	RelationalGreaterThan
	RelationalGreaterThanOrEqual
	RelationalLessThan
	RelationalLessThanOrEqual
	RelationalEqual
	RelationalNotEqual
)

type RelationalExpression struct {
	AdditionExpression   AdditionExpression
	RelationalOperator   RelationalOperator
	ComparisonExpression *RelationalExpression
	Tokens               Tokens
}

func (re *RelationalExpression) parse(r *rParser) error {
	s := r.NewGoal()

	if err := re.AdditionExpression.parse(&s); err != nil {
		return r.Error("RelationalExpression", err)
	}

	r.Score(s)

	s = r.NewGoal()

	s.AcceptRunWhitespaceNoNewLine()

	if s.AcceptToken(parser.Token{Type: TokenOperator, Data: ">"}) {
		re.RelationalOperator = RelationalGreaterThan
	} else if s.AcceptToken(parser.Token{Type: TokenOperator, Data: ">="}) {
		re.RelationalOperator = RelationalGreaterThanOrEqual
	} else if s.AcceptToken(parser.Token{Type: TokenOperator, Data: "<"}) {
		re.RelationalOperator = RelationalLessThan
	} else if s.AcceptToken(parser.Token{Type: TokenOperator, Data: "<="}) {
		re.RelationalOperator = RelationalLessThanOrEqual
	} else if s.AcceptToken(parser.Token{Type: TokenOperator, Data: "=="}) {
		re.RelationalOperator = RelationalEqual
	} else if s.AcceptToken(parser.Token{Type: TokenOperator, Data: "!="}) {
		re.RelationalOperator = RelationalNotEqual
	}

	if re.RelationalOperator != RelationalNone {
		s.AcceptRunWhitespaceNoNewLine()
		r.Score(s)

		s = r.NewGoal()
		re.ComparisonExpression = new(RelationalExpression)

		if err := re.ComparisonExpression.parse(&s); err != nil {
			return r.Error("RelationalExpression", err)
		}

		r.Score(s)
	}

	re.Tokens = r.ToTokens()

	return nil
}

type AdditionType uint8

const (
	AdditionNone AdditionType = iota
	AdditionAdd
	AdditionSubtract
)

type AdditionExpression struct {
	MultiplicationExpression MultiplicationExpression
	AdditionType             AdditionType
	AdditionExpression       *AdditionExpression
	Tokens                   Tokens
}

func (a *AdditionExpression) parse(r *rParser) error {
	s := r.NewGoal()

	if err := a.MultiplicationExpression.parse(&s); err != nil {
		return r.Error("AdditionExpression", err)
	}

	r.Score(s)

	s = r.NewGoal()

	s.AcceptRunWhitespaceNoNewLine()

	if s.AcceptToken(parser.Token{Type: TokenOperator, Data: "+"}) {
		a.AdditionType = AdditionAdd
	} else if s.AcceptToken(parser.Token{Type: TokenOperator, Data: "-"}) {
		a.AdditionType = AdditionSubtract
	}

	if a.AdditionType != AdditionNone {
		s.AcceptRunWhitespaceNoNewLine()
		r.Score(s)

		s = r.NewGoal()
		a.AdditionExpression = new(AdditionExpression)

		if err := a.AdditionExpression.parse(&s); err != nil {
			return r.Error("AdditionExpression", err)
		}

		r.Score(s)
	}

	a.Tokens = r.ToTokens()

	return nil
}

type MultiplicationType uint8

const (
	MultiplicationNone MultiplicationType = iota
	MultiplicationMultiply
	MultiplicationDivide
)

type MultiplicationExpression struct {
	PipeOrSpecialExpression  PipeOrSpecialExpression
	MultiplicationType       MultiplicationType
	MultiplicationExpression *MultiplicationExpression
	Tokens                   Tokens
}

func (m *MultiplicationExpression) parse(r *rParser) error {
	s := r.NewGoal()

	if err := m.PipeOrSpecialExpression.parse(&s); err != nil {
		return r.Error("MultiplicationExpression", err)
	}

	r.Score(s)

	s = r.NewGoal()

	s.AcceptRunWhitespaceNoNewLine()

	if s.AcceptToken(parser.Token{Type: TokenOperator, Data: "*"}) {
		m.MultiplicationType = MultiplicationMultiply
	} else if s.AcceptToken(parser.Token{Type: TokenOperator, Data: "/"}) {
		m.MultiplicationType = MultiplicationDivide
	}

	if m.MultiplicationType != MultiplicationNone {
		s.AcceptRunWhitespaceNoNewLine()
		r.Score(s)

		s = r.NewGoal()
		m.MultiplicationExpression = new(MultiplicationExpression)

		if err := m.MultiplicationExpression.parse(&s); err != nil {
			return r.Error("MultiplicationExpression", err)
		}

		r.Score(s)
	}

	m.Tokens = r.ToTokens()

	return nil
}

type PipeOrSpecialExpression struct {
	SequenceExpression      SequenceExpression
	Operator                *Token
	PipeOrSpecialExpression *PipeOrSpecialExpression
	Tokens                  Tokens
}

func (p *PipeOrSpecialExpression) parse(r *rParser) error {
	s := r.NewGoal()

	if err := p.SequenceExpression.parse(&s); err != nil {
		return r.Error("PipeOrSpecialExpression", err)
	}

	r.Score(s)

	s = r.NewGoal()

	s.AcceptRunWhitespaceNoNewLine()

	if s.AcceptToken(parser.Token{}) || s.Accept(TokenSpecialOperator) {
		p.Operator = s.GetLastToken()

		s.AcceptRunWhitespaceNoNewLine()
		r.Score(s)

		s = r.NewGoal()
		p.PipeOrSpecialExpression = new(PipeOrSpecialExpression)

		if err := p.PipeOrSpecialExpression.parse(&s); err != nil {
			return r.Error("PipeOrSpecialExpression", err)
		}

		r.Score(s)
	}

	p.Tokens = r.ToTokens()

	return nil
}

type SequenceExpression struct {
	UnaryExpression    UnaryExpression
	SequenceExpression *SequenceExpression
	Tokens             Tokens
}

func (se *SequenceExpression) parse(r *rParser) error {
	s := r.NewGoal()

	if err := se.UnaryExpression.parse(&s); err != nil {
		return r.Error("SequenceExpression", err)
	}

	r.Score(s)

	s = r.NewGoal()

	s.AcceptRunWhitespaceNoNewLine()

	if s.AcceptToken(parser.Token{Type: TokenOperator, Data: ":"}) {
		s.AcceptRunWhitespaceNoNewLine()
		r.Score(s)

		s = r.NewGoal()
		se.SequenceExpression = new(SequenceExpression)

		if err := se.SequenceExpression.parse(&s); err != nil {
			return r.Error("SequenceExpression", err)
		}

		r.Score(s)
	}

	se.Tokens = r.ToTokens()

	return nil
}

type UnaryType uint8

const (
	UnaryNone UnaryType = iota
	UnaryAdd
	UnaryMinus
)

type UnaryExpression struct {
	UnaryType       UnaryType
	UnaryExpression *UnaryExpression
	Tokens          Tokens
}

func (u *UnaryExpression) parse(r *rParser) error {
	s := r.NewGoal()

	if s.AcceptToken(parser.Token{Type: TokenOperator, Data: "+"}) {
		u.UnaryType = UnaryAdd
		s.AcceptRunWhitespaceNoNewLine()
	} else if s.AcceptToken(parser.Token{Type: TokenOperator, Data: "-"}) {
		u.UnaryType = UnaryMinus
		s.AcceptRunWhitespaceNoNewLine()
	}

	r.Score(s)

	s = r.NewGoal()

	if err := u.UnaryExpression.parse(&s); err != nil {
		return r.Error("UnaryExpression", err)
	}

	r.Score(s)

	u.Tokens = r.ToTokens()

	return nil
}

type ExponentiationExpression struct {
	SubsetExpression         SubsetExpression
	ExponentiationExpression *ExponentiationExpression
	Tokens                   Tokens
}

func (e *ExponentiationExpression) parse(r *rParser) error {
	s := r.NewGoal()

	if err := e.SubsetExpression.parse(&s); err != nil {
		return r.Error("ExponentiationExpression", err)
	}

	r.Score(s)

	s = r.NewGoal()

	s.AcceptRunWhitespaceNoNewLine()

	if s.AcceptToken(parser.Token{Type: TokenOperator, Data: "^"}) {
		s.AcceptRunWhitespaceNoNewLine()
		r.Score(s)

		s = r.NewGoal()
		e.ExponentiationExpression = new(ExponentiationExpression)

		if err := e.ExponentiationExpression.parse(&s); err != nil {
			return r.Error("ExponentiationExpression", err)
		}

		r.Score(s)
	}

	e.Tokens = r.ToTokens()

	return nil
}

type SubsetType uint8

const (
	SubsetNone SubsetType = iota
	SubsetList
	SubsetStructure
)

type SubsetExpression struct {
	ScopeExpression  ScopeExpression
	SubsetType       SubsetType
	SubsetExpression *SubsetExpression
	Tokens           Tokens
}

func (se *SubsetExpression) parse(r *rParser) error {
	s := r.NewGoal()

	if err := se.ScopeExpression.parse(&s); err != nil {
		return r.Error("SubsetExpression", err)
	}

	r.Score(s)

	s = r.NewGoal()

	s.AcceptRunWhitespaceNoNewLine()

	if s.AcceptToken(parser.Token{Type: TokenOperator, Data: "$"}) {
		se.SubsetType = SubsetList
	} else if s.AcceptToken(parser.Token{Type: TokenOperator, Data: "@"}) {
		se.SubsetType = SubsetStructure
	}

	if se.SubsetType != SubsetNone {
		s.AcceptRunWhitespaceNoNewLine()
		r.Score(s)

		s = r.NewGoal()
		se.SubsetExpression = new(SubsetExpression)

		if err := se.SubsetExpression.parse(&s); err != nil {
			return r.Error("SubsetExpression", err)
		}

		r.Score(s)
	}

	se.Tokens = r.ToTokens()

	return nil
}

type ScopeExpression struct {
	IndexOrCallExpression IndexOrCallExpression
	ScopeExpression       *ScopeExpression
	Tokens                Tokens
}

func (se *ScopeExpression) parse(r *rParser) error {
	s := r.NewGoal()

	if err := se.IndexOrCallExpression.parse(&s); err != nil {
		return r.Error("ScopeExpression", err)
	}

	r.Score(s)

	s = r.NewGoal()

	s.AcceptRunWhitespaceNoNewLine()

	if s.AcceptToken(parser.Token{Type: TokenOperator, Data: "::"}) {
		s.AcceptRunWhitespaceNoNewLine()
		r.Score(s)

		s = r.NewGoal()
		se.ScopeExpression = new(ScopeExpression)

		if err := se.ScopeExpression.parse(&s); err != nil {
			return r.Error("ScopeExpression", err)
		}

		r.Score(s)
	}

	se.Tokens = r.ToTokens()

	return nil
}

type IndexOrCallExpression struct {
	SimpleExpression      *SimpleExpression
	IndexOrCallExpression *IndexOrCallExpression
	Index                 *Index
	Call                  *Call
	Tokens                Tokens
}

func (i *IndexOrCallExpression) parse(r *rParser) error {
	s := r.NewGoal()
	i.SimpleExpression = new(SimpleExpression)

	if err := i.SimpleExpression.parse(&s); err != nil {
		return r.Error("IndexOrCallExpression", err)
	}

	r.Score(s)

Loop:
	for {
		i.Tokens = r.ToTokens()
		s = r.NewGoal()

		s.AcceptRunWhitespaceNoNewLine()

		var (
			index *Index
			call  *Call
			err   error
		)

		switch s.Peek() {
		case parser.Token{Type: TokenGrouping, Data: "["}, parser.Token{Type: TokenGrouping, Data: "[["}:
			r.Score(s)

			s = r.NewGoal()
			index = new(Index)
			err = index.parse(&s)
		case parser.Token{Type: TokenGrouping, Data: "("}:
			r.Score(s)

			s = r.NewGoal()
			call = new(Call)
			err = call.parse(&s)
		default:

			break Loop
		}

		if err != nil {
			return r.Error("IndexOrCallExpression", err)
		}

		r.Score(s)

		i = &IndexOrCallExpression{
			IndexOrCallExpression: i,
			Index:                 index,
			Call:                  call,
		}
	}

	return nil
}

type SimpleExpression struct {
	Identifier              *Token
	Constant                *Token
	Ellipsis                *Token
	ParenthesizedExpression *Expression
	CompoundExpression      *CompoundExpression
	Tokens                  Tokens
}

func (a *SimpleExpression) parse(r *rParser) error {
	if r.Accept(TokenIdentifier) {
		a.Identifier = r.GetLastToken()
	} else if r.Accept(TokenStringLiteral, TokenNumericLiteral, TokenIntegerLiteral, TokenComplexLiteral, TokenBooleanLiteral, TokenNull, TokenNA) {
		a.Constant = r.GetLastToken()
	} else if r.Accept(TokenEllipsis) {
		a.Ellipsis = r.GetLastToken()
	} else if r.AcceptToken(parser.Token{Type: TokenGrouping, Data: "("}) {
		r.AcceptRunWhitespaceNoNewLine()

		s := r.NewGoal()
		a.ParenthesizedExpression = new(Expression)

		if err := a.ParenthesizedExpression.parse(&s); err != nil {
			return r.Error("SimpleExpression", err)
		}

		r.Score(s)
		r.AcceptRunWhitespaceNoNewLine()

		if !r.AcceptToken(parser.Token{Type: TokenGrouping, Data: "("}) {
			return r.Error("SimpleExpression", ErrMissingClosingParen)
		}
	} else if tk := r.Peek(); tk == (parser.Token{Type: TokenGrouping, Data: "{"}) {
		s := r.NewGoal()
		a.CompoundExpression = new(CompoundExpression)

		if err := a.CompoundExpression.parse(&s); err != nil {
			return r.Error("SimpleExpression", err)
		}
	} else {
		return r.Error("SimpleExpression", ErrInvalidAtom)
	}

	a.Tokens = r.ToTokens()

	return nil
}

type Index struct {
	Double bool
	Args   []QueryExpression
	Tokens Tokens
}

func (i *Index) parse(r *rParser) error {
	if r.AcceptToken(parser.Token{Type: TokenGrouping, Data: "[["}) {
		i.Double = true
	} else {
		r.AcceptToken(parser.Token{Type: TokenGrouping, Data: "["})
	}

	r.AcceptRunWhitespaceNoNewLine()

	if i.Double || !r.AcceptToken(parser.Token{Type: TokenGrouping, Data: "]"}) {
		for {
			s := r.NewGoal()

			var h QueryExpression

			if err := h.parse(&s); err != nil {
				return r.Error("Index", err)
			}

			i.Args = append(i.Args, h)

			r.Score(s)
			r.AcceptRunWhitespaceNoNewLine()

			if i.Double {
				if r.AcceptToken(parser.Token{Type: TokenGrouping, Data: "]]"}) {
					break
				}
			} else if r.AcceptToken(parser.Token{Type: TokenGrouping, Data: "]"}) {
				break
			} else if !r.AcceptToken(parser.Token{Type: TokenExpressionTerminator, Data: ","}) {
				return r.Error("Index", ErrMissingComma)
			}

			r.AcceptRunWhitespaceNoNewLine()
		}
	}

	i.Tokens = r.ToTokens()

	return nil
}

type Call struct {
	Args   []Arg
	Tokens Tokens
}

func (c *Call) parse(r *rParser) error {
	r.AcceptToken(parser.Token{Type: TokenGrouping, Data: "("})
	r.AcceptRunWhitespaceNoNewLine()

	if !r.AcceptToken(parser.Token{Type: TokenGrouping, Data: ")"}) {
		for {
			if tk := r.Peek(); tk == (parser.Token{Type: TokenExpressionTerminator, Data: ","}) || tk == (parser.Token{Type: TokenGrouping, Data: ")"}) {
				c.Args = append(c.Args, Arg{})
			} else {
				s := r.NewGoal()

				var a Arg

				if err := a.parse(&s); err != nil {
					return r.Error("Call", err)
				}

				r.Score(s)

				c.Args = append(c.Args, a)
			}

			if r.AcceptToken(parser.Token{Type: TokenGrouping, Data: ")"}) {
				break
			} else if !r.AcceptToken(parser.Token{Type: TokenExpressionTerminator, Data: ","}) {
				return r.Error("Call", ErrMissingComma)
			}

			r.AcceptRunWhitespaceNoNewLine()
		}
	}

	c.Tokens = r.ToTokens()

	return nil
}

type Arg struct {
	QueryExpression *QueryExpression
	Ellipsis        *Token
	Tokens          Tokens
}

func (a *Arg) parse(r *rParser) error {
	if r.Accept(TokenEllipsis) {
		a.Ellipsis = r.GetLastToken()
	} else {
		s := r.NewGoal()
		a.QueryExpression = new(QueryExpression)

		if err := a.QueryExpression.parse(&s); err != nil {
			return r.Error("Arg", err)
		}

		r.Score(s)
	}

	a.Tokens = r.ToTokens()

	return nil
}

var (
	ErrMissingTerminator   = errors.New("missing terminator")
	ErrMissingOpeningParen = errors.New("missing opening paren")
	ErrMissingClosingParen = errors.New("missing closing paren")
	ErrMissingIn           = errors.New("missing in keyword")
	ErrMissingIdentifier   = errors.New("missing identifier")
	ErrMissingComma        = errors.New("missing comma")
	ErrInvalidAtom         = errors.New("invalid atom")
)
