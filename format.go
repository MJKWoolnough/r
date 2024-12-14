package r

import "io"

func (a AdditionType) String() string {
	switch a {
	case AdditionNone:
		return "AdditionNone"
	case AdditionAdd:
		return "AdditionAdd"
	case AdditionSubtract:
		return "AdditionSubtract"
	default:
		return "Unknown"
	}
}

func (a AdditionType) printType(w io.Writer, _ bool) {
	io.WriteString(w, a.String())
}

func (a AndType) String() string {
	switch a {
	case AndNone:
		return "AndNone"
	case AndVectorized:
		return "AndVectorizes"
	case AndNotVectorized:
		return "AndNotVectorized"
	default:
		return "Unknown"
	}
}

func (a AndType) printType(w io.Writer, _ bool) {
	io.WriteString(w, a.String())
}

func (m MultiplicationType) String() string {
	switch m {
	case MultiplicationNone:
		return "MultiplicationNone"
	case MultiplicationMultiply:
		return "MultiplicationMultiply"
	case MultiplicationDivide:
		return "MultiplicationDivide"
	default:
		return "Unknown"
	}
}

func (m MultiplicationType) printType(w io.Writer, _ bool) {
	io.WriteString(w, m.String())
}

func (o OrType) String() string {
	switch o {
	case OrNone:
		return "OrNone"
	case OrVectorized:
		return "OrVectorized"
	case OrNotVectorized:
		return "OrNotVectorized"
	default:
		return "Unknown"
	}
}

func (o OrType) printType(w io.Writer, _ bool) {
	io.WriteString(w, o.String())
}

func (q QueryType) String() string {
	switch q {
	case QueryNone:
		return "QueryNone"
	case QueryUnary:
		return "QueryUnary"
	case QueryBinary:
		return "QueryBinary"
	default:
		return "Unknown"
	}
}

func (q QueryType) printType(w io.Writer, _ bool) {
	io.WriteString(w, q.String())
}

func (r RelationalOperator) String() string {
	switch r {
	case RelationalNone:
		return "RelationalNone"
	case RelationalGreaterThan:
		return "RelationalGreaterThan"
	case RelationalGreaterThanOrEqual:
		return "RelationalGreaterThanOrEqual"
	case RelationalLessThan:
		return "RelationalLessThan"
	case RelationalLessThanOrEqual:
		return "RelationalLessThanOrEqual"
	case RelationalEqual:
		return "RelationalEqual"
	case RelationalNotEqual:
		return "RelationalNotEqual"
	default:
		return "Unknown"
	}
}

func (r RelationalOperator) printType(w io.Writer, _ bool) {
	io.WriteString(w, r.String())
}
