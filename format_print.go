package r

import (
	"io"
)

func (a AdditionExpression) printSource(w io.Writer, v bool) {
	a.MultiplicationExpression.printSource(w, v)

	if a.AdditionType != AdditionNone && a.AdditionExpression != nil {
		if v {
			io.WriteString(w, " ")
			a.Comments[0].printSource(w, v)
		}

		a.AdditionType.printSource(w, v)

		if v {
			io.WriteString(w, " ")
			a.Comments[1].printSource(w, v)
		}

		a.AdditionExpression.printSource(w, v)
	}
}

func (a AndExpression) printSource(w io.Writer, v bool) {
	a.NotExpression.printSource(w, v)

	if a.AndType != AndNone && a.AndExpression != nil {
		if v {
			io.WriteString(w, " ")
			a.Comments[0].printSource(w, v)
		}

		a.AndType.printSource(w, v)

		if v {
			io.WriteString(w, " ")
			a.Comments[1].printSource(w, v)
		}

		a.AndExpression.printSource(w, v)
	}
}

func (a Arg) printSource(w io.Writer, v bool) {
	if v {
		a.Comments[0].printSource(w, v)
	}

	if a.Ellipsis != nil {
		io.WriteString(w, a.Ellipsis.Data)
	} else if a.QueryExpression != nil {
		a.QueryExpression.printSource(w, v)
	}

	if v && len(a.Comments[1]) > 0 {
		io.WriteString(w, " ")
		a.Comments[1].printSource(w, v)
	}
}

func (a ArgList) printSource(w io.Writer, v bool) {
	if len(a.Args) > 0 {
		ipp := indentPrinter{w}

		for n, arg := range a.Args {
			if n > 0 {
				if v {
					io.WriteString(w, ", ")
				} else {
					io.WriteString(w, ",")
				}
			}

			arg.printSource(&ipp, v)

			if v && (arg.Default == nil && len(arg.Comments[1]) > 0 || arg.Default != nil && len(arg.Default.Comments[1]) > 0) {
				if n == len(a.Args)-1 {
					io.WriteString(w, "\n")
				} else {
					ipp.WriteString("\n")
				}
			}
		}
	} else if v && len(a.Comments) > 0 {
		ipp := indentPrinter{w}

		io.WriteString(&ipp, "\n")
		a.Comments.printSource(&ipp, false)
		io.WriteString(w, "\n")
	}
}

func (a Argument) printSource(w io.Writer, v bool) {
	if a.Identifier != nil {
		if v {
			a.Comments[0].printSource(w, v)
		}

		io.WriteString(w, a.Identifier.Data)

		if v && len(a.Comments[1]) > 0 {
			io.WriteString(w, " ")
			a.Comments[1].printSource(w, false)
		}

		if a.Identifier.Type == TokenIdentifier && a.Default != nil {
			if v {
				if len(a.Comments[1]) > 0 {
					io.WriteString(w, "\n= ")
				} else {
					io.WriteString(w, " = ")
				}
			} else {
				io.WriteString(w, "=")
			}

			a.Default.printSource(w, v)
		}
	}
}

func (a AssignmentExpression) printSource(w io.Writer, v bool) {
	a.FormulaeExpression.printSource(w, v)

	if a.AssignmentType != AssignmentNone && a.AssignmentExpression != nil {
		if v {
			io.WriteString(w, " ")
			a.Comments[0].printSource(w, v)
		}

		a.AssignmentType.printSource(w, v)

		if v {
			io.WriteString(w, " ")
			a.Comments[1].printSource(w, v)
		}

		a.AssignmentExpression.printSource(w, v)
	}
}

func (c Call) printSource(w io.Writer, v bool) {
	io.WriteString(w, "(")

	if len(c.Args) > 0 {
		c.Args[0].printSource(w, v)

		for _, a := range c.Args[1:] {
			if v {
				io.WriteString(w, ", ")
			} else {
				io.WriteString(w, ",")
			}

			a.printSource(w, v)
		}
	} else if v && len(c.Comments) > 0 {
		ipp := indentPrinter{w}

		io.WriteString(&ipp, "\n")
		c.Comments.printSource(&ipp, false)
		io.WriteString(w, "\n")
	}

	io.WriteString(w, ")")
}

func (c CompoundExpression) printSource(w io.Writer, v bool) {
	io.WriteString(w, "{")

	if len(c.Expressions) > 0 {
		ipp := indentPrinter{w}

		for _, e := range c.Expressions {
			io.WriteString(&ipp, "\n")
			e.printSource(&ipp, v)
		}

		io.WriteString(w, "\n")
	}

	if v && len(c.Comments) > 0 {
		ipp := indentPrinter{w}

		io.WriteString(&ipp, "\n")
		c.Comments.printSource(&ipp, false)
		io.WriteString(w, "\n")
	}

	io.WriteString(w, "}")
}

func (e ExponentiationExpression) printSource(w io.Writer, v bool) {
	e.SubsetExpression.printSource(w, v)

	if e.ExponentiationExpression != nil {
		if v && len(e.Comments[0]) > 0 {
			io.WriteString(w, " ")
			e.Comments[0].printSource(w, v)
		}

		io.WriteString(w, "^")

		if v && len(e.Comments[1]) > 0 {
			io.WriteString(w, " ")
			e.Comments[1].printSource(w, v)
		}

		e.ExponentiationExpression.printSource(w, v)
	}
}

func (e Expression) printSource(w io.Writer, v bool) {
	if v {
		e.Comments[0].printSource(w, v)
	}

	if e.FlowControl != nil {
		e.FlowControl.printSource(w, v)
	} else if e.FunctionDefinition != nil {
		e.FunctionDefinition.printSource(w, v)
	} else if e.QueryExpression != nil {
		e.QueryExpression.printSource(w, v)
	}

	if v && e.Comments[1] != nil {
		io.WriteString(w, " ")
		e.Comments[1].printSource(w, false)
	}
}

func (f File) printSource(w io.Writer, v bool) {
	for _, e := range f.Statements {
		e.printSource(w, v)
		io.WriteString(w, "\n")
	}

	if v && len(f.Comments) > 0 {
		io.WriteString(w, "\n")
		f.Comments.printSource(w, v)
	}
}

func (f FlowControl) printSource(w io.Writer, v bool) {
	if f.IfControl != nil {
		f.IfControl.printSource(w, v)
	} else if f.WhileControl != nil {
		f.WhileControl.printSource(w, v)
	} else if f.RepeatControl != nil {
		f.RepeatControl.printSource(w, v)
	} else if f.ForControl != nil {
		f.ForControl.printSource(w, v)
	}
}

func (f ForControl) printSource(w io.Writer, v bool) {
	if f.Var != nil && f.Var.Type == TokenIdentifier {
		ipp := indentPrinter{w}

		if v {
			io.WriteString(w, "for ")

			f.Comments[0].printSource(w, v)
			io.WriteString(w, "(")
		} else {
			io.WriteString(w, "for(")
		}

		if v {
			f.Comments[1].printSource(&ipp, v)
		}

		io.WriteString(w, f.Var.Data)
		io.WriteString(w, " ")

		if v {
			f.Comments[2].printSource(&ipp, v)
		}

		io.WriteString(w, "in")
		io.WriteString(w, " ")

		if v {
			f.Comments[3].printSource(&ipp, v)
		}

		f.List.printSource(&ipp, v)

		if v && len(f.Comments[4]) > 0 {
			io.WriteString(w, " ")
			f.Comments[4].printSource(&ipp, false)
			io.WriteString(w, "\n")
		}

		if v {
			io.WriteString(w, ") ")
		} else {
			io.WriteString(w, ")")
		}

		f.Expr.printSource(w, v)
	}
}

func (f FormulaeExpression) printSource(w io.Writer, v bool) {
	if f.OrExpression != nil {
		f.OrExpression.printSource(w, v)

		if v && f.FormulaeExpression != nil {
			io.WriteString(w, " ")
		}
	}

	if f.FormulaeExpression != nil {
		io.WriteString(w, "~")

		if v {
			io.WriteString(w, " ")
			f.Comments.printSource(w, v)
		}

		f.FormulaeExpression.printSource(w, v)
	}
}

func (f FunctionDefinition) printSource(w io.Writer, v bool) {
	io.WriteString(w, "function")

	if v && len(f.Comments) > 0 {
		ipp := indentPrinter{w}

		io.WriteString(w, " ")
		f.Comments.printSource(&ipp, false)
		io.WriteString(w, "\n")
	}

	io.WriteString(w, "(")

	f.ArgList.printSource(w, v)

	if v {
		io.WriteString(w, ") ")
	} else {
		io.WriteString(w, ")")
	}

	f.Body.printSource(w, v)
}

func (i IfControl) printSource(w io.Writer, v bool) {
	ipp := indentPrinter{w}

	if v {
		io.WriteString(w, "if ")

		if len(i.Comments[0]) > 0 {
			i.Comments[0].printSource(&ipp, false)
			io.WriteString(w, "\n(")
		} else {
			io.WriteString(w, "(")
		}

		i.Comments[1].printSource(&ipp, v)
	} else {
		io.WriteString(w, "if(")
	}

	i.Cond.printSource(w, v)

	if v {
		if len(i.Comments[2]) > 0 {
			i.Comments[2].printSource(&ipp, false)
			io.WriteString(w, "\n")
		}

		io.WriteString(w, ") ")
	} else {
		io.WriteString(w, ")")
	}

	i.Expr.printSource(w, v)

	if i.Else != nil {
		if v && (len(i.Expr.Comments[1]) > 0 || len(i.Comments[3]) > 0) {
			io.WriteString(w, "\n")
		}

		if v && len(i.Comments[3]) > 0 {
			io.WriteString(w, "\n")
			i.Comments[3].printSource(w, v)
			io.WriteString(w, "else ")
		} else {
			io.WriteString(w, " else ")
		}

		i.Else.printSource(w, v)
	}
}

func (i Index) printSource(w io.Writer, v bool) {
	if !i.Double {
		io.WriteString(w, "[")

		if len(i.Args) > 0 {
			i.Args[0].printSource(w, v)

			for _, a := range i.Args[1:] {
				if v {
					io.WriteString(w, ", ")
				} else {
					io.WriteString(w, ",")
				}

				a.printSource(w, v)
			}
		}

		io.WriteString(w, "]")
	} else if len(i.Args) == 1 {
		io.WriteString(w, "[[")
		i.Args[0].printSource(w, v)
		io.WriteString(w, "]]")
	}
}

func (i IndexExpression) printSource(w io.Writer, v bool) {
	if v {
		i.Comments[0].printSource(w, v)
	}

	i.QueryExpression.printSource(w, v)

	if v && len(i.Comments[1]) > 0 {
		io.WriteString(w, " ")
		i.Comments[1].printSource(w, v)
	}
}

func (i IndexOrCallExpression) printSource(w io.Writer, v bool) {
	if i.SimpleExpression != nil {
		i.SimpleExpression.printSource(w, v)
	} else if i.IndexOrCallExpression != nil {
		i.IndexOrCallExpression.printSource(w, v)

		if v && len(i.Comments) > 0 {
			io.WriteString(w, " ")
			i.Comments.printSource(w, v)
		}

		if i.Index != nil {
			i.Index.printSource(w, v)
		} else if i.Call != nil {
			i.Call.printSource(w, v)
		}
	}
}

func (m MultiplicationExpression) printSource(w io.Writer, v bool) {
	m.PipeOrSpecialExpression.printSource(w, v)

	if m.MultiplicationType != MultiplicationNone && m.MultiplicationExpression != nil {
		if v {
			io.WriteString(w, " ")
			m.Comments[0].printSource(w, v)
		}

		m.MultiplicationType.printSource(w, v)

		if v {
			io.WriteString(w, " ")
			m.Comments[1].printSource(w, v)
		}

		m.MultiplicationExpression.printSource(w, v)
	}
}

func (n NotExpression) printSource(w io.Writer, v bool) {
	for m := range n.Nots {
		io.WriteString(w, "!")

		if v && uint(len(n.Comments)) > m && len(n.Comments[m]) > 0 {
			io.WriteString(w, " ")
			n.Comments[m].printSource(w, v)
		}
	}

	n.RelationalExpression.printSource(w, v)
}

func (o OrExpression) printSource(w io.Writer, v bool) {
	o.AndExpression.printSource(w, v)

	if o.OrType != OrNone && o.OrExpression != nil {
		if v {
			io.WriteString(w, " ")
			o.Comments[0].printSource(w, v)
		}

		o.OrType.printSource(w, v)

		if v {
			io.WriteString(w, " ")
			o.Comments[1].printSource(w, v)
		}

		o.OrExpression.printSource(w, v)
	}
}

func (p ParenthesizedExpression) printSource(w io.Writer, v bool) {
	io.WriteString(w, "(")
	p.Expression.printSource(w, v)

	if v && len(p.Expression.Comments[1]) > 0 {
		io.WriteString(w, "\n")
	}

	io.WriteString(w, ")")
}

func (p PipeOrSpecialExpression) printSource(w io.Writer, v bool) {
	p.SequenceExpression.printSource(w, v)

	if p.Operator != nil && p.PipeOrSpecialExpression != nil {
		if v {
			io.WriteString(w, " ")
			p.Comments[0].printSource(w, v)
		}

		io.WriteString(w, p.Operator.Data)

		if v {
			io.WriteString(w, " ")
			p.Comments[1].printSource(w, v)
		}

		p.PipeOrSpecialExpression.printSource(w, v)
	}
}

func (q QueryExpression) printSource(w io.Writer, v bool) {
	if q.AssignmentExpression != nil {
		q.AssignmentExpression.printSource(w, v)
	}

	if q.QueryExpression != nil {
		if v && q.AssignmentExpression != nil {
			io.WriteString(w, " ")

			q.Comments[0].printSource(w, v)
		}

		io.WriteString(w, "?")

		if v {
			if q.QueryExpression.AssignmentExpression != nil {
				io.WriteString(w, " ")
			}

			if len(q.Comments) > 0 {
				q.Comments[1].printSource(w, v)
			}
		}

		q.QueryExpression.printSource(w, v)
	}
}

func (r RelationalExpression) printSource(w io.Writer, v bool) {
	r.AdditionExpression.printSource(w, v)

	if r.RelationalOperator != RelationalNone && r.RelationalExpression != nil {
		if v {
			io.WriteString(w, " ")
			r.Comments[0].printSource(w, v)
		}

		r.RelationalOperator.printSource(w, v)

		if v {
			io.WriteString(w, " ")
			r.Comments[1].printSource(w, v)
		}

		r.RelationalExpression.printSource(w, v)
	}
}

func (r RepeatControl) printSource(w io.Writer, v bool) {
	io.WriteString(w, "repeat ")
	r.Expr.printSource(w, v)
}

func (s ScopeExpression) printSource(w io.Writer, v bool) {
	s.IndexOrCallExpression.printSource(w, v)

	if s.ScopeExpression != nil {
		if v && len(s.Comments[0]) > 0 {
			io.WriteString(w, " ")
			s.Comments[0].printSource(w, v)
		}

		io.WriteString(w, "::")

		if v && len(s.Comments[1]) > 0 {
			io.WriteString(w, " ")
			s.Comments[1].printSource(w, v)
		}

		s.ScopeExpression.printSource(w, v)
	}
}

func (s SequenceExpression) printSource(w io.Writer, v bool) {
	s.UnaryExpression.printSource(w, v)

	if s.SequenceExpression != nil {
		if v && len(s.Comments[0]) > 0 {
			io.WriteString(w, " ")
			s.Comments[0].printSource(w, v)
		}

		io.WriteString(w, ":")

		if v && len(s.Comments[1]) > 0 {
			io.WriteString(w, " ")
			s.Comments[1].printSource(w, v)
		}

		s.SequenceExpression.printSource(w, v)
	}
}

func (s SimpleExpression) printSource(w io.Writer, v bool) {
	if s.Identifier != nil {
		io.WriteString(w, s.Identifier.Data)
	} else if s.Constant != nil {
		io.WriteString(w, s.Constant.Data)
	} else if s.Ellipsis != nil {
		io.WriteString(w, s.Ellipsis.Data)
	} else if s.ParenthesizedExpression != nil {
		s.ParenthesizedExpression.printSource(w, v)
	} else if s.CompoundExpression != nil {
		s.CompoundExpression.printSource(w, v)
	}
}

func (s SubsetExpression) printSource(w io.Writer, v bool) {
	s.ScopeExpression.printSource(w, v)

	if s.SubsetExpression != nil && s.SubsetType != SubsetNone {
		if v && len(s.Comments[0]) > 0 {
			io.WriteString(w, " ")
			s.Comments[0].printSource(w, v)
		}

		if s.SubsetType == SubsetList {
			io.WriteString(w, "$")
		} else {
			io.WriteString(w, "@")
		}

		if v && len(s.Comments[1]) > 0 {
			io.WriteString(w, " ")
			s.Comments[1].printSource(w, v)
		}

		s.SubsetExpression.printSource(w, v)
	}
}

func (u UnaryExpression) printSource(w io.Writer, v bool) {
	for n, t := range u.UnaryType {
		switch t {
		case UnaryAdd:
			io.WriteString(w, "+")
		case UnaryMinus:
			io.WriteString(w, "-")
		}

		if v && len(u.Comments) > n && len(u.Comments[n]) > 0 {
			io.WriteString(w, " ")
			u.Comments[n].printSource(w, v)
		}
	}

	u.ExponentiationExpression.printSource(w, v)
}

func (wc WhileControl) printSource(w io.Writer, v bool) {
	ipp := indentPrinter{w}

	if v {
		if len(wc.Comments[0]) > 0 {
			io.WriteString(w, "while ")
			wc.Comments[0].printSource(w, true)
			io.WriteString(w, "(")
		} else {
			io.WriteString(w, "while (")
		}

		wc.Comments[1].printSource(&ipp, v)
	} else {
		io.WriteString(w, "while(")
	}

	wc.Cond.printSource(w, v)

	if v && len(wc.Comments[2]) > 0 {
		io.WriteString(w, " ")
		wc.Comments[2].printSource(&ipp, false)
		io.WriteString(w, "\n")
	}

	if v {
		io.WriteString(w, ") ")
	} else {
		io.WriteString(w, ")")
	}

	wc.Expr.printSource(w, v)
}
