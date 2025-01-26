package r

// File automatically generated with format.sh.

import "fmt"

// Format implements the fmt.Formatter interface
func (f AdditionExpression) Format(s fmt.State, v rune) {
	if v == 'v' && s.Flag('#') {
		type X = AdditionExpression
		type AdditionExpression X

		fmt.Fprintf(s, "%#v", AdditionExpression(f))
	} else {
		format(&f, s, v)
	}
}

// Format implements the fmt.Formatter interface
func (f AndExpression) Format(s fmt.State, v rune) {
	if v == 'v' && s.Flag('#') {
		type X = AndExpression
		type AndExpression X

		fmt.Fprintf(s, "%#v", AndExpression(f))
	} else {
		format(&f, s, v)
	}
}

// Format implements the fmt.Formatter interface
func (f Arg) Format(s fmt.State, v rune) {
	if v == 'v' && s.Flag('#') {
		type X = Arg
		type Arg X

		fmt.Fprintf(s, "%#v", Arg(f))
	} else {
		format(&f, s, v)
	}
}

// Format implements the fmt.Formatter interface
func (f ArgList) Format(s fmt.State, v rune) {
	if v == 'v' && s.Flag('#') {
		type X = ArgList
		type ArgList X

		fmt.Fprintf(s, "%#v", ArgList(f))
	} else {
		format(&f, s, v)
	}
}

// Format implements the fmt.Formatter interface
func (f Argument) Format(s fmt.State, v rune) {
	if v == 'v' && s.Flag('#') {
		type X = Argument
		type Argument X

		fmt.Fprintf(s, "%#v", Argument(f))
	} else {
		format(&f, s, v)
	}
}

// Format implements the fmt.Formatter interface
func (f AssignmentExpression) Format(s fmt.State, v rune) {
	if v == 'v' && s.Flag('#') {
		type X = AssignmentExpression
		type AssignmentExpression X

		fmt.Fprintf(s, "%#v", AssignmentExpression(f))
	} else {
		format(&f, s, v)
	}
}

// Format implements the fmt.Formatter interface
func (f Call) Format(s fmt.State, v rune) {
	if v == 'v' && s.Flag('#') {
		type X = Call
		type Call X

		fmt.Fprintf(s, "%#v", Call(f))
	} else {
		format(&f, s, v)
	}
}

// Format implements the fmt.Formatter interface
func (f CompoundExpression) Format(s fmt.State, v rune) {
	if v == 'v' && s.Flag('#') {
		type X = CompoundExpression
		type CompoundExpression X

		fmt.Fprintf(s, "%#v", CompoundExpression(f))
	} else {
		format(&f, s, v)
	}
}

// Format implements the fmt.Formatter interface
func (f ExponentiationExpression) Format(s fmt.State, v rune) {
	if v == 'v' && s.Flag('#') {
		type X = ExponentiationExpression
		type ExponentiationExpression X

		fmt.Fprintf(s, "%#v", ExponentiationExpression(f))
	} else {
		format(&f, s, v)
	}
}

// Format implements the fmt.Formatter interface
func (f Expression) Format(s fmt.State, v rune) {
	if v == 'v' && s.Flag('#') {
		type X = Expression
		type Expression X

		fmt.Fprintf(s, "%#v", Expression(f))
	} else {
		format(&f, s, v)
	}
}

// Format implements the fmt.Formatter interface
func (f File) Format(s fmt.State, v rune) {
	if v == 'v' && s.Flag('#') {
		type X = File
		type File X

		fmt.Fprintf(s, "%#v", File(f))
	} else {
		format(&f, s, v)
	}
}

// Format implements the fmt.Formatter interface
func (f FlowControl) Format(s fmt.State, v rune) {
	if v == 'v' && s.Flag('#') {
		type X = FlowControl
		type FlowControl X

		fmt.Fprintf(s, "%#v", FlowControl(f))
	} else {
		format(&f, s, v)
	}
}

// Format implements the fmt.Formatter interface
func (f ForControl) Format(s fmt.State, v rune) {
	if v == 'v' && s.Flag('#') {
		type X = ForControl
		type ForControl X

		fmt.Fprintf(s, "%#v", ForControl(f))
	} else {
		format(&f, s, v)
	}
}

// Format implements the fmt.Formatter interface
func (f FormulaeExpression) Format(s fmt.State, v rune) {
	if v == 'v' && s.Flag('#') {
		type X = FormulaeExpression
		type FormulaeExpression X

		fmt.Fprintf(s, "%#v", FormulaeExpression(f))
	} else {
		format(&f, s, v)
	}
}

// Format implements the fmt.Formatter interface
func (f FunctionDefinition) Format(s fmt.State, v rune) {
	if v == 'v' && s.Flag('#') {
		type X = FunctionDefinition
		type FunctionDefinition X

		fmt.Fprintf(s, "%#v", FunctionDefinition(f))
	} else {
		format(&f, s, v)
	}
}

// Format implements the fmt.Formatter interface
func (f IfControl) Format(s fmt.State, v rune) {
	if v == 'v' && s.Flag('#') {
		type X = IfControl
		type IfControl X

		fmt.Fprintf(s, "%#v", IfControl(f))
	} else {
		format(&f, s, v)
	}
}

// Format implements the fmt.Formatter interface
func (f Index) Format(s fmt.State, v rune) {
	if v == 'v' && s.Flag('#') {
		type X = Index
		type Index X

		fmt.Fprintf(s, "%#v", Index(f))
	} else {
		format(&f, s, v)
	}
}

// Format implements the fmt.Formatter interface
func (f IndexExpression) Format(s fmt.State, v rune) {
	if v == 'v' && s.Flag('#') {
		type X = IndexExpression
		type IndexExpression X

		fmt.Fprintf(s, "%#v", IndexExpression(f))
	} else {
		format(&f, s, v)
	}
}

// Format implements the fmt.Formatter interface
func (f IndexOrCallExpression) Format(s fmt.State, v rune) {
	if v == 'v' && s.Flag('#') {
		type X = IndexOrCallExpression
		type IndexOrCallExpression X

		fmt.Fprintf(s, "%#v", IndexOrCallExpression(f))
	} else {
		format(&f, s, v)
	}
}

// Format implements the fmt.Formatter interface
func (f MultiplicationExpression) Format(s fmt.State, v rune) {
	if v == 'v' && s.Flag('#') {
		type X = MultiplicationExpression
		type MultiplicationExpression X

		fmt.Fprintf(s, "%#v", MultiplicationExpression(f))
	} else {
		format(&f, s, v)
	}
}

// Format implements the fmt.Formatter interface
func (f NotExpression) Format(s fmt.State, v rune) {
	if v == 'v' && s.Flag('#') {
		type X = NotExpression
		type NotExpression X

		fmt.Fprintf(s, "%#v", NotExpression(f))
	} else {
		format(&f, s, v)
	}
}

// Format implements the fmt.Formatter interface
func (f OrExpression) Format(s fmt.State, v rune) {
	if v == 'v' && s.Flag('#') {
		type X = OrExpression
		type OrExpression X

		fmt.Fprintf(s, "%#v", OrExpression(f))
	} else {
		format(&f, s, v)
	}
}

// Format implements the fmt.Formatter interface
func (f ParenthesizedExpression) Format(s fmt.State, v rune) {
	if v == 'v' && s.Flag('#') {
		type X = ParenthesizedExpression
		type ParenthesizedExpression X

		fmt.Fprintf(s, "%#v", ParenthesizedExpression(f))
	} else {
		format(&f, s, v)
	}
}

// Format implements the fmt.Formatter interface
func (f PipeOrSpecialExpression) Format(s fmt.State, v rune) {
	if v == 'v' && s.Flag('#') {
		type X = PipeOrSpecialExpression
		type PipeOrSpecialExpression X

		fmt.Fprintf(s, "%#v", PipeOrSpecialExpression(f))
	} else {
		format(&f, s, v)
	}
}

// Format implements the fmt.Formatter interface
func (f QueryExpression) Format(s fmt.State, v rune) {
	if v == 'v' && s.Flag('#') {
		type X = QueryExpression
		type QueryExpression X

		fmt.Fprintf(s, "%#v", QueryExpression(f))
	} else {
		format(&f, s, v)
	}
}

// Format implements the fmt.Formatter interface
func (f RelationalExpression) Format(s fmt.State, v rune) {
	if v == 'v' && s.Flag('#') {
		type X = RelationalExpression
		type RelationalExpression X

		fmt.Fprintf(s, "%#v", RelationalExpression(f))
	} else {
		format(&f, s, v)
	}
}

// Format implements the fmt.Formatter interface
func (f RepeatControl) Format(s fmt.State, v rune) {
	if v == 'v' && s.Flag('#') {
		type X = RepeatControl
		type RepeatControl X

		fmt.Fprintf(s, "%#v", RepeatControl(f))
	} else {
		format(&f, s, v)
	}
}

// Format implements the fmt.Formatter interface
func (f ScopeExpression) Format(s fmt.State, v rune) {
	if v == 'v' && s.Flag('#') {
		type X = ScopeExpression
		type ScopeExpression X

		fmt.Fprintf(s, "%#v", ScopeExpression(f))
	} else {
		format(&f, s, v)
	}
}

// Format implements the fmt.Formatter interface
func (f SequenceExpression) Format(s fmt.State, v rune) {
	if v == 'v' && s.Flag('#') {
		type X = SequenceExpression
		type SequenceExpression X

		fmt.Fprintf(s, "%#v", SequenceExpression(f))
	} else {
		format(&f, s, v)
	}
}

// Format implements the fmt.Formatter interface
func (f SimpleExpression) Format(s fmt.State, v rune) {
	if v == 'v' && s.Flag('#') {
		type X = SimpleExpression
		type SimpleExpression X

		fmt.Fprintf(s, "%#v", SimpleExpression(f))
	} else {
		format(&f, s, v)
	}
}

// Format implements the fmt.Formatter interface
func (f SubsetExpression) Format(s fmt.State, v rune) {
	if v == 'v' && s.Flag('#') {
		type X = SubsetExpression
		type SubsetExpression X

		fmt.Fprintf(s, "%#v", SubsetExpression(f))
	} else {
		format(&f, s, v)
	}
}

// Format implements the fmt.Formatter interface
func (f UnaryExpression) Format(s fmt.State, v rune) {
	if v == 'v' && s.Flag('#') {
		type X = UnaryExpression
		type UnaryExpression X

		fmt.Fprintf(s, "%#v", UnaryExpression(f))
	} else {
		format(&f, s, v)
	}
}

// Format implements the fmt.Formatter interface
func (f WhileControl) Format(s fmt.State, v rune) {
	if v == 'v' && s.Flag('#') {
		type X = WhileControl
		type WhileControl X

		fmt.Fprintf(s, "%#v", WhileControl(f))
	} else {
		format(&f, s, v)
	}
}
