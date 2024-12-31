package r

import "io"

func (a AdditionExpression) printSource(w io.Writer, v bool) {
	a.MultiplicationExpression.printSource(w, v)

	if a.AdditionType != AdditionNone && a.AdditionExpression != nil {
		if v {
			io.WriteString(w, " ")
		}

		a.AdditionType.printSource(w, v)

		if v {
			io.WriteString(w, " ")
		}

		a.AdditionExpression.printSource(w, v)
	}
}

func (a AndExpression) printSource(w io.Writer, v bool) {
	a.NotExpression.printSource(w, v)

	if a.AndType != AndNone && a.AndExpression != nil {
		if v {
			io.WriteString(w, " ")
		}

		a.AndType.printSource(w, v)

		if v {
			io.WriteString(w, " ")
		}

		a.AndExpression.printSource(w, v)
	}
}

func (a Arg) printSource(w io.Writer, v bool) {
	if a.Ellipsis != nil {
		io.WriteString(w, a.Ellipsis.Data)
	} else if a.QueryExpression != nil {
		a.QueryExpression.printSource(w, v)
	}
}

func (a ArgList) printSource(w io.Writer, v bool) {
	if len(a.Args) > 0 {
		a.Args[0].printSource(w, v)

		for _, arg := range a.Args[1:] {
			if v {
				io.WriteString(w, ", ")
			} else {
				io.WriteString(w, ",")
			}

			arg.printSource(w, v)
		}
	}
}

func (a Argument) printSource(w io.Writer, v bool) {
	if a.Identifier == nil {
		return
	}

	io.WriteString(w, a.Identifier.Data)

	if a.Identifier.Type == TokenIdentifier && a.Default != nil {
		if v {
			io.WriteString(w, " = ")
		} else {
			io.WriteString(w, "=")
		}

		a.Default.printSource(w, v)
	}
}

func (a AssignmentExpression) printSource(w io.Writer, v bool) {
	a.FormulaeExpression.printSource(w, v)

	if a.AssignmentType != AssignmentNone && a.AssignmentExpression != nil {
		if v {
			io.WriteString(w, " ")
		}

		a.AssignmentType.printSource(w, v)

		if v {
			io.WriteString(w, " ")
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
	}

	io.WriteString(w, ")")
}

func (c CompoundExpression) printSource(w io.Writer, v bool) {
	if len(c.Expressions) == 0 {
		return
	}

	ipp := indentPrinter{w}

	io.WriteString(&ipp, "{\n")
	c.Expressions[0].printSource(&ipp, v)

	for _, e := range c.Expressions[1:] {
		io.WriteString(&ipp, "\n")
		e.printSource(&ipp, v)
	}

	io.WriteString(w, "}")
}

func (e ExponentiationExpression) printSource(w io.Writer, v bool) {
	e.SubsetExpression.printSource(w, v)

	if e.ExponentiationExpression != nil {
		io.WriteString(w, "^")
		e.ExponentiationExpression.printSource(w, v)
	}
}

func (e Expression) printSource(w io.Writer, v bool) {
	if e.FlowControl != nil {
		e.FlowControl.printSource(w, v)
	} else if e.FunctionDefinition != nil {
		e.FunctionDefinition.printSource(w, v)
	} else if e.QueryExpression != nil {
		e.QueryExpression.printSource(w, v)
	}
}

func (f File) printSource(w io.Writer, v bool) {
	for _, e := range f.Statements {
		e.printSource(w, v)
		io.WriteString(w, "\n")
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
	if f.Var == nil || f.Var.Type != TokenIdentifier {
		return
	}

	if v {
		io.WriteString(w, "for (")
	} else {
		io.WriteString(w, "for(")
	}

	io.WriteString(w, f.Var.Data)
	io.WriteString(w, " in ")
	f.List.printSource(w, v)

	if v {
		io.WriteString(w, ") ")
	} else {
		io.WriteString(w, ")")
	}

	f.Expr.printSource(w, v)
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
		}

		f.FormulaeExpression.printSource(w, v)
	}
}

func (f FunctionDefinition) printSource(w io.Writer, v bool) {
	io.WriteString(w, "function(")
	f.ArgList.printSource(w, v)

	if v {
		io.WriteString(w, ") ")
	} else {
		io.WriteString(w, ")")
	}

	f.Body.printSource(w, v)
}

func (i IfControl) printSource(w io.Writer, v bool) {
	if v {
		io.WriteString(w, "if (")
	} else {
		io.WriteString(w, "if(")
	}

	i.Cond.printSource(w, v)

	if v {
		io.WriteString(w, ") ")
	} else {
		io.WriteString(w, ") ")
	}

	i.Expr.printSource(w, v)

	if i.Else != nil {
		io.WriteString(w, " else ")
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

func (i IndexOrCallExpression) printSource(w io.Writer, v bool) {
	if i.SimpleExpression != nil {
		i.SimpleExpression.printSource(w, v)
	} else if i.IndexOrCallExpression != nil {
		i.IndexOrCallExpression.printSource(w, v)

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
		}

		m.MultiplicationType.printSource(w, v)

		if v {
			io.WriteString(w, " ")
		}

		m.MultiplicationExpression.printSource(w, v)
	}
}

func (n NotExpression) printSource(w io.Writer, v bool) {
	for range n.Nots {
		io.WriteString(w, "!")
	}

	n.RelationalExpression.printSource(w, v)
}

func (o OrExpression) printSource(w io.Writer, v bool) {
	o.AndExpression.printSource(w, v)

	if o.OrType != OrNone && o.OrExpression != nil {
		o.OrType.printSource(w, v)
		o.OrExpression.printSource(w, v)
	}
}

func (p PipeOrSpecialExpression) printSource(w io.Writer, v bool) {
	p.SequenceExpression.printSource(w, v)

	if p.Operator != nil && p.PipeOrSpecialExpression != nil {
		if v {
			io.WriteString(w, " ")
		}

		io.WriteString(w, p.Operator.Data)

		if v {
			io.WriteString(w, " ")
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
		}

		io.WriteString(w, "?")

		if v {
			io.WriteString(w, " ")
		}

		q.QueryExpression.printSource(w, v)
	}
}

func (r RelationalExpression) printSource(w io.Writer, v bool) {
	r.AdditionExpression.printSource(w, v)

	if r.RelationalOperator != RelationalNone && r.RelationalExpression != nil {
		if v {
			io.WriteString(w, " ")
		}

		r.RelationalExpression.printSource(w, v)

		if v {
			io.WriteString(w, " ")
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
		io.WriteString(w, "::")
		s.ScopeExpression.printSource(w, v)
	}
}

func (s SequenceExpression) printSource(w io.Writer, v bool) {
	s.UnaryExpression.printSource(w, v)

	if s.SequenceExpression != nil {
		io.WriteString(w, ":")
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
		io.WriteString(w, "(")
		s.ParenthesizedExpression.printSource(w, v)
		io.WriteString(w, ")")
	} else if s.CompoundExpression != nil {
		s.CompoundExpression.printSource(w, v)
	}
}

func (s SubsetExpression) printSource(w io.Writer, v bool) {
	s.ScopeExpression.printSource(w, v)

	if s.SubsetExpression != nil && s.SubsetType != SubsetNone {
		if s.SubsetType == SubsetList {
			io.WriteString(w, "$")
		} else {
			io.WriteString(w, "@")
		}

		s.SubsetExpression.printSource(w, v)
	}
}

func (u UnaryExpression) printSource(w io.Writer, v bool) {
	for _, t := range u.UnaryType {
		switch t {
		case UnaryAdd:
			io.WriteString(w, "+")
		case UnaryMinus:
			io.WriteString(w, "-")
		}
	}

	u.ExponentiationExpression.printSource(w, v)
}

func (wc WhileControl) printSource(w io.Writer, v bool) {
	if v {
		io.WriteString(w, "while (")
	} else {
		io.WriteString(w, "while(")
	}

	wc.Cond.printSource(w, v)

	if v {
		io.WriteString(w, ") ")
	} else {
		io.WriteString(w, ")")
	}

	wc.Expr.printSource(w, v)
}
