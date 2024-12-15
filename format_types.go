package r

// File automatically generated with format.sh.

import "io"

func (f *AdditionExpression) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("AdditionExpression {")

	pp.Print("\nMultiplicationExpression: ")
	f.MultiplicationExpression.printType(&pp, v)

	pp.Print("\nAdditionType: ")
	f.AdditionType.printType(&pp, v)

	if f.AdditionExpression != nil {
		pp.Print("\nAdditionExpression: ")
		f.AdditionExpression.printType(&pp, v)
	} else if v {
		pp.Print("\nAdditionExpression: nil")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *AndExpression) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("AndExpression {")

	pp.Print("\nNotExpression: ")
	f.NotExpression.printType(&pp, v)

	pp.Print("\nAndType: ")
	f.AndType.printType(&pp, v)

	if f.AndExpression != nil {
		pp.Print("\nAndExpression: ")
		f.AndExpression.printType(&pp, v)
	} else if v {
		pp.Print("\nAndExpression: nil")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *Arg) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("Arg {")

	if f.QueryExpression != nil {
		pp.Print("\nQueryExpression: ")
		f.QueryExpression.printType(&pp, v)
	} else if v {
		pp.Print("\nQueryExpression: nil")
	}

	if f.Ellipsis != nil {
		pp.Print("\nEllipsis: ")
		f.Ellipsis.printType(&pp, v)
	} else if v {
		pp.Print("\nEllipsis: nil")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *ArgList) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("ArgList {")

	if f.Args == nil {
		pp.Print("\nArgs: nil")
	} else if len(f.Args) > 0 {
		pp.Print("\nArgs: [")

		ipp := indentPrinter{&pp}

		for n, e := range f.Args {
			ipp.Printf("\n%d: ", n)
			e.printType(&ipp, v)
		}

		pp.Print("\n]")
	} else if v {
		pp.Print("\nArgs: []")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *Argument) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("Argument {")

	if f.Identifier != nil {
		pp.Print("\nIdentifier: ")
		f.Identifier.printType(&pp, v)
	} else if v {
		pp.Print("\nIdentifier: nil")
	}

	if f.Default != nil {
		pp.Print("\nDefault: ")
		f.Default.printType(&pp, v)
	} else if v {
		pp.Print("\nDefault: nil")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *AssignmentExpression) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("AssignmentExpression {")

	pp.Print("\nConditionalExpression: ")
	f.ConditionalExpression.printType(&pp, v)

	if f.AssignmentExpression != nil {
		pp.Print("\nAssignmentExpression: ")
		f.AssignmentExpression.printType(&pp, v)
	} else if v {
		pp.Print("\nAssignmentExpression: nil")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *Call) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("Call {")

	if f.Args == nil {
		pp.Print("\nArgs: nil")
	} else if len(f.Args) > 0 {
		pp.Print("\nArgs: [")

		ipp := indentPrinter{&pp}

		for n, e := range f.Args {
			ipp.Printf("\n%d: ", n)
			e.printType(&ipp, v)
		}

		pp.Print("\n]")
	} else if v {
		pp.Print("\nArgs: []")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *CompoundExpression) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("CompoundExpression {")

	if f.Expressions == nil {
		pp.Print("\nExpressions: nil")
	} else if len(f.Expressions) > 0 {
		pp.Print("\nExpressions: [")

		ipp := indentPrinter{&pp}

		for n, e := range f.Expressions {
			ipp.Printf("\n%d: ", n)
			e.printType(&ipp, v)
		}

		pp.Print("\n]")
	} else if v {
		pp.Print("\nExpressions: []")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *ExponentiationExpression) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("ExponentiationExpression {")

	pp.Print("\nSubsetExpression: ")
	f.SubsetExpression.printType(&pp, v)

	if f.ExponentiationExpression != nil {
		pp.Print("\nExponentiationExpression: ")
		f.ExponentiationExpression.printType(&pp, v)
	} else if v {
		pp.Print("\nExponentiationExpression: nil")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *Expression) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("Expression {")

	if f.CompoundExpression != nil {
		pp.Print("\nCompoundExpression: ")
		f.CompoundExpression.printType(&pp, v)
	} else if v {
		pp.Print("\nCompoundExpression: nil")
	}

	if f.FlowControl != nil {
		pp.Print("\nFlowControl: ")
		f.FlowControl.printType(&pp, v)
	} else if v {
		pp.Print("\nFlowControl: nil")
	}

	if f.FunctionDefinition != nil {
		pp.Print("\nFunctionDefinition: ")
		f.FunctionDefinition.printType(&pp, v)
	} else if v {
		pp.Print("\nFunctionDefinition: nil")
	}

	if f.AssignmentExpression != nil {
		pp.Print("\nAssignmentExpression: ")
		f.AssignmentExpression.printType(&pp, v)
	} else if v {
		pp.Print("\nAssignmentExpression: nil")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *FlowControl) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("FlowControl {")

	if f.IfControl != nil {
		pp.Print("\nIfControl: ")
		f.IfControl.printType(&pp, v)
	} else if v {
		pp.Print("\nIfControl: nil")
	}

	if f.WhileControl != nil {
		pp.Print("\nWhileControl: ")
		f.WhileControl.printType(&pp, v)
	} else if v {
		pp.Print("\nWhileControl: nil")
	}

	if f.RepeatControl != nil {
		pp.Print("\nRepeatControl: ")
		f.RepeatControl.printType(&pp, v)
	} else if v {
		pp.Print("\nRepeatControl: nil")
	}

	if f.ForControl != nil {
		pp.Print("\nForControl: ")
		f.ForControl.printType(&pp, v)
	} else if v {
		pp.Print("\nForControl: nil")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *ForControl) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("ForControl {")

	pp.Print("\nVar: ")
	f.Var.printType(&pp, v)

	pp.Print("\nList: ")
	f.List.printType(&pp, v)

	pp.Print("\nExpr: ")
	f.Expr.printType(&pp, v)

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *FormulaeExpression) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("FormulaeExpression {")

	if f.OrExpression != nil {
		pp.Print("\nOrExpression: ")
		f.OrExpression.printType(&pp, v)
	} else if v {
		pp.Print("\nOrExpression: nil")
	}

	if f.FormulaeExpression != nil {
		pp.Print("\nFormulaeExpression: ")
		f.FormulaeExpression.printType(&pp, v)
	} else if v {
		pp.Print("\nFormulaeExpression: nil")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *FunctionDefinition) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("FunctionDefinition {")

	pp.Print("\nArgList: ")
	f.ArgList.printType(&pp, v)

	pp.Print("\nBody: ")
	f.Body.printType(&pp, v)

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *IfControl) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("IfControl {")

	pp.Print("\nCond: ")
	f.Cond.printType(&pp, v)

	pp.Print("\nExpr: ")
	f.Expr.printType(&pp, v)

	if f.Else != nil {
		pp.Print("\nElse: ")
		f.Else.printType(&pp, v)
	} else if v {
		pp.Print("\nElse: nil")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *Index) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("Index {")

	if f.Double || v {
		pp.Printf("\nDouble: %v", f.Double)
	}

	if f.Args == nil {
		pp.Print("\nArgs: nil")
	} else if len(f.Args) > 0 {
		pp.Print("\nArgs: [")

		ipp := indentPrinter{&pp}

		for n, e := range f.Args {
			ipp.Printf("\n%d: ", n)
			e.printType(&ipp, v)
		}

		pp.Print("\n]")
	} else if v {
		pp.Print("\nArgs: []")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *IndexOrCallExpression) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("IndexOrCallExpression {")

	if f.Atom != nil {
		pp.Print("\nAtom: ")
		f.Atom.printType(&pp, v)
	} else if v {
		pp.Print("\nAtom: nil")
	}

	if f.IndexOrCallExpression != nil {
		pp.Print("\nIndexOrCallExpression: ")
		f.IndexOrCallExpression.printType(&pp, v)
	} else if v {
		pp.Print("\nIndexOrCallExpression: nil")
	}

	if f.Index != nil {
		pp.Print("\nIndex: ")
		f.Index.printType(&pp, v)
	} else if v {
		pp.Print("\nIndex: nil")
	}

	if f.Call != nil {
		pp.Print("\nCall: ")
		f.Call.printType(&pp, v)
	} else if v {
		pp.Print("\nCall: nil")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *MultiplicationExpression) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("MultiplicationExpression {")

	pp.Print("\nPipeOrSpecialExpression: ")
	f.PipeOrSpecialExpression.printType(&pp, v)

	pp.Print("\nMultiplicationType: ")
	f.MultiplicationType.printType(&pp, v)

	if f.MultiplicationExpression != nil {
		pp.Print("\nMultiplicationExpression: ")
		f.MultiplicationExpression.printType(&pp, v)
	} else if v {
		pp.Print("\nMultiplicationExpression: nil")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *NotExpression) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("NotExpression {")

	if f.Not || v {
		pp.Printf("\nNot: %v", f.Not)
	}

	pp.Print("\nComparisonExpression: ")
	f.ComparisonExpression.printType(&pp, v)

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *OrExpression) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("OrExpression {")

	pp.Print("\nAndExpression: ")
	f.AndExpression.printType(&pp, v)

	pp.Print("\nOrType: ")
	f.OrType.printType(&pp, v)

	if f.OrExpression != nil {
		pp.Print("\nOrExpression: ")
		f.OrExpression.printType(&pp, v)
	} else if v {
		pp.Print("\nOrExpression: nil")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *PipeOrSpecialExpression) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("PipeOrSpecialExpression {")

	pp.Print("\nSequenceExpression: ")
	f.SequenceExpression.printType(&pp, v)

	if f.Operator != nil {
		pp.Print("\nOperator: ")
		f.Operator.printType(&pp, v)
	} else if v {
		pp.Print("\nOperator: nil")
	}

	if f.PipeOrSpecialExpression != nil {
		pp.Print("\nPipeOrSpecialExpression: ")
		f.PipeOrSpecialExpression.printType(&pp, v)
	} else if v {
		pp.Print("\nPipeOrSpecialExpression: nil")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *QueryExpression) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("QueryExpression {")

	pp.Print("\nQueryType: ")
	f.QueryType.printType(&pp, v)

	pp.Print("\nAssignmentExpression: ")
	f.AssignmentExpression.printType(&pp, v)

	if f.QueryExpression != nil {
		pp.Print("\nQueryExpression: ")
		f.QueryExpression.printType(&pp, v)
	} else if v {
		pp.Print("\nQueryExpression: nil")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *RelationalExpression) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("RelationalExpression {")

	pp.Print("\nAdditionExpression: ")
	f.AdditionExpression.printType(&pp, v)

	pp.Print("\nRelationalOperator: ")
	f.RelationalOperator.printType(&pp, v)

	if f.ComparisonExpression != nil {
		pp.Print("\nComparisonExpression: ")
		f.ComparisonExpression.printType(&pp, v)
	} else if v {
		pp.Print("\nComparisonExpression: nil")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *RepeatControl) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("RepeatControl {")

	pp.Print("\nCond: ")
	f.Cond.printType(&pp, v)

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *ScopeExpression) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("ScopeExpression {")

	pp.Print("\nIndexOrCallExpression: ")
	f.IndexOrCallExpression.printType(&pp, v)

	if f.ScopeExpression != nil {
		pp.Print("\nScopeExpression: ")
		f.ScopeExpression.printType(&pp, v)
	} else if v {
		pp.Print("\nScopeExpression: nil")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *SequenceExpression) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("SequenceExpression {")

	pp.Print("\nUnaryExpression: ")
	f.UnaryExpression.printType(&pp, v)

	if f.SequenceExpression != nil {
		pp.Print("\nSequenceExpression: ")
		f.SequenceExpression.printType(&pp, v)
	} else if v {
		pp.Print("\nSequenceExpression: nil")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *SimpleExpression) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("SimpleExpression {")

	if f.Identifier != nil {
		pp.Print("\nIdentifier: ")
		f.Identifier.printType(&pp, v)
	} else if v {
		pp.Print("\nIdentifier: nil")
	}

	if f.Constant != nil {
		pp.Print("\nConstant: ")
		f.Constant.printType(&pp, v)
	} else if v {
		pp.Print("\nConstant: nil")
	}

	if f.Ellipsis != nil {
		pp.Print("\nEllipsis: ")
		f.Ellipsis.printType(&pp, v)
	} else if v {
		pp.Print("\nEllipsis: nil")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *SubsetExpression) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("SubsetExpression {")

	pp.Print("\nScopeExpression: ")
	f.ScopeExpression.printType(&pp, v)

	pp.Print("\nSubsetType: ")
	f.SubsetType.printType(&pp, v)

	if f.SubsetExpression != nil {
		pp.Print("\nSubsetExpression: ")
		f.SubsetExpression.printType(&pp, v)
	} else if v {
		pp.Print("\nSubsetExpression: nil")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *UnaryExpression) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("UnaryExpression {")

	pp.Print("\nUnaryType: ")
	f.UnaryType.printType(&pp, v)

	if f.UnaryExpression != nil {
		pp.Print("\nUnaryExpression: ")
		f.UnaryExpression.printType(&pp, v)
	} else if v {
		pp.Print("\nUnaryExpression: nil")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *WhileControl) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("WhileControl {")

	pp.Print("\nCond: ")
	f.Cond.printType(&pp, v)

	pp.Print("\nExpr: ")
	f.Expr.printType(&pp, v)

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}
