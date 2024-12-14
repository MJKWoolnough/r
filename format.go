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
	switch a {
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
