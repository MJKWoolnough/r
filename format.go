package r

import (
	"fmt"
	"io"

	"vimagination.zapto.org/parser"
)

var indent = []byte{'\t'}

type indentPrinter struct {
	io.Writer
}

func (i *indentPrinter) Write(p []byte) (int, error) {
	var (
		total int
		last  int
	)

	for n, c := range p {
		if c == '\n' {
			m, err := i.Writer.Write(p[last : n+1])
			total += m

			if err != nil {
				return total, err
			}

			_, err = i.Writer.Write(indent)
			if err != nil {
				return total, err
			}

			last = n + 1
		}
	}

	if last != len(p) {
		m, err := i.Writer.Write(p[last:])
		total += m

		if err != nil {
			return total, err
		}
	}

	return total, nil
}

func (i *indentPrinter) Print(args ...interface{}) {
	fmt.Fprint(i, args...)
}

func (i *indentPrinter) Printf(format string, args ...interface{}) {
	fmt.Fprintf(i, format, args...)
}

func (i *indentPrinter) WriteString(s string) (int, error) {
	return i.Write([]byte(s))
}

func (t Token) printType(w io.Writer, v bool) {
	var typ string

	switch t.Type {
	case TokenWhitespace:
		typ = "Whitespace"
	case TokenLineTerminator:
		typ = "LineTerminator"
	case TokenExpressionTerminator:
		typ = "ExpressionTerminator"
	case TokenComment:
		typ = "Comment"
	case TokenStringLiteral:
		typ = "StringLiteral"
	case TokenNumericLiteral:
		typ = "NumericLiteral"
	case TokenIntegerLiteral:
		typ = "IntegerLiteral"
	case TokenComplexLiteral:
		typ = "ComplexLiteral"
	case TokenBooleanLiteral:
		typ = "BooleanLiteral"
	case TokenNull:
		typ = "Null"
	case TokenNA:
		typ = "NA"
	case TokenIdentifier:
		typ = "Identifier"
	case TokenKeyword:
		typ = "Keyword"
	case TokenEllipsis:
		typ = "Ellipsis"
	case TokenOperator:
		typ = "Operator"
	case TokenSpecialOperator:
		typ = "SpecialOperator"
	case TokenGrouping:
		typ = "Grouping"
	case parser.TokenDone:
		typ = "Done"
	case parser.TokenError:
		typ = "Error"
	default:
		typ = "Unknown"
	}

	fmt.Fprintf(w, "Type: %s - Data: %q", typ, t.Data)

	if v {
		fmt.Fprintf(w, " - Position: %d (%d: %d)", t.Pos, t.Line, t.LinePos)
	}
}

func (t Tokens) printType(w io.Writer, v bool) {
	if t == nil {
		io.WriteString(w, "nil")

		return
	}

	if len(t) == 0 {
		io.WriteString(w, "[]")

		return
	}

	io.WriteString(w, "[")

	ipp := indentPrinter{w}

	for n, t := range t {
		ipp.Printf("\n%d: ", n)
		t.printType(w, v)
	}

	io.WriteString(w, "\n]")
}

type formatter interface {
	printType(io.Writer, bool)
	printSource(io.Writer, bool)
}

func format(f formatter, s fmt.State, v rune) {
	switch v {
	case 'v':
		f.printType(s, s.Flag('+'))
	case 's':
		f.printSource(s, s.Flag('+'))
	}
}

// String implements the fmt.Stringer interface.
func (a AssignmentType) String() string {
	switch a {
	case AssignmentNone:
		return "AssignmentNone"
	case AssignmentEquals:
		return "AssignmentEquals"
	case AssignmentLeftAssign:
		return "AssignmentLeftAssign"
	case AssignmentRightAssign:
		return "AssignmentRightAssign"
	case AssignmentLeftParentAssign:
		return "AssignmentLeftParentAssign"
	case AssignmentRightParentAssign:
		return "AssignmentRightParentAssign"
	default:
		return "Unknown"
	}
}

func (a AssignmentType) printType(w io.Writer, _ bool) {
	io.WriteString(w, a.String())
}

func (a AssignmentType) printSource(w io.Writer, _ bool) {
	switch a {
	case AssignmentEquals:
		io.WriteString(w, "=")
	case AssignmentLeftAssign:
		io.WriteString(w, "<-")
	case AssignmentRightAssign:
		io.WriteString(w, "->")
	case AssignmentLeftParentAssign:
		io.WriteString(w, "<<-")
	case AssignmentRightParentAssign:
		io.WriteString(w, "->>")
	}
}

// String implements the fmt.Stringer interface.
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

func (a AdditionType) printSource(w io.Writer, _ bool) {
	switch a {
	case AdditionAdd:
		io.WriteString(w, "+")
	case AdditionSubtract:
		io.WriteString(w, "-")
	}
}

// String implements the fmt.Stringer interface.
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

func (a AndType) printSource(w io.Writer, _ bool) {
	switch a {
	case AndVectorized:
		io.WriteString(w, "&")
	case AndNotVectorized:
		io.WriteString(w, "&&")
	}
}

// String implements the fmt.Stringer interface.
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

func (m MultiplicationType) printSource(w io.Writer, _ bool) {
	switch m {
	case MultiplicationMultiply:
		io.WriteString(w, "*")
	case MultiplicationDivide:
		io.WriteString(w, "/")
	}
}

// String implements the fmt.Stringer interface.
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

func (o OrType) printSource(w io.Writer, _ bool) {
	switch o {
	case OrVectorized:
		io.WriteString(w, "|")
	case OrNotVectorized:
		io.WriteString(w, "||")
	}
}

// String implements the fmt.Stringer interface.
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

func (r RelationalOperator) printSource(w io.Writer, _ bool) {
	switch r {
	case RelationalGreaterThan:
		io.WriteString(w, ">")
	case RelationalGreaterThanOrEqual:
		io.WriteString(w, ">=")
	case RelationalLessThan:
		io.WriteString(w, "<")
	case RelationalLessThanOrEqual:
		io.WriteString(w, "<=")
	case RelationalEqual:
		io.WriteString(w, "==")
	case RelationalNotEqual:
		io.WriteString(w, "!=")
	}
}

// String implements the fmt.Stringer interface.
func (s SubsetType) String() string {
	switch s {
	case SubsetNone:
		return "SubsetNone"
	case SubsetList:
		return "SubsetList"
	case SubsetStructure:
		return "SubsetStructure"
	default:
		return "Unknown"
	}
}

func (s SubsetType) printType(w io.Writer, _ bool) {
	io.WriteString(w, s.String())
}

func (s SubsetType) printSource(w io.Writer, _ bool) {
	switch s {
	case SubsetList:
		io.WriteString(w, "$")
	case SubsetStructure:
		io.WriteString(w, "@")
	}
}

// String implements the fmt.Stringer interface.
func (u UnaryType) String() string {
	switch u {
	case UnaryAdd:
		return "UnaryAdd"
	case UnaryMinus:
		return "UnaryMinus"
	default:
		return "Unknown"
	}
}

func (u UnaryType) printType(w io.Writer, _ bool) {
	io.WriteString(w, u.String())
}

func (u UnaryType) printsource(w io.Writer, _ bool) {
	switch u {
	case UnaryAdd:
		io.WriteString(w, "+")
	case UnaryMinus:
		io.WriteString(w, "-")
	}
}
