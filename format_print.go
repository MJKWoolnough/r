package r

import "io"

func (a AdditionExpression) printSource(w io.Writer, v bool) {}

func (a AndExpression) printSource(w io.Writer, v bool) {}

func (a Arg) printSource(w io.Writer, v bool) {
	if a.Ellipsis != nil {
		io.WriteString(w, a.Ellipsis.Data)
	} else if a.QueryExpression != nil {
		a.QueryExpression.printSource(w, v)
	}
}

func (a ArgList) printSource(w io.Writer, v bool) {}

func (a Argument) printSource(w io.Writer, v bool) {}

func (a AssignmentExpression) printSource(w io.Writer, v bool) {}

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

func (e Expression) printSource(w io.Writer, v bool) {}

func (f FlowControl) printSource(w io.Writer, v bool) {}

func (f ForControl) printSource(w io.Writer, v bool) {}

func (f FormulaeExpression) printSource(w io.Writer, v bool) {}

func (f FunctionDefinition) printSource(w io.Writer, v bool) {}

func (i IfControl) printSource(w io.Writer, v bool) {}

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

func (m MultiplicationExpression) printSource(w io.Writer, v bool) {}

func (n NotExpression) printSource(w io.Writer, v bool) {}

func (o OrExpression) printSource(w io.Writer, v bool) {}

func (p PipeOrSpecialExpression) printSource(w io.Writer, v bool) {}

func (q QueryExpression) printSource(w io.Writer, v bool) {}

func (r RelationalExpression) printSource(w io.Writer, v bool) {}

func (r RepeatControl) printSource(w io.Writer, v bool) {}

func (s ScopeExpression) printSource(w io.Writer, v bool) {
	s.IndexOrCallExpression.printSource(w, v)

	if s.ScopeExpression != nil {
		io.WriteString(w, "::")
		s.ScopeExpression.printSource(w, v)
	}
}

func (s SequenceExpression) printSource(w io.Writer, v bool) {}

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

func (wc WhileControl) printSource(w io.Writer, v bool) {}
