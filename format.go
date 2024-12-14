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
