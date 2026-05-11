package walk

import (
	"errors"
	"reflect"
	"testing"

	"vimagination.zapto.org/parser"
	"vimagination.zapto.org/r"
)

var (
	sentinel = errors.New("")
	nilErr   = errors.New("nil received")
	nilRet   = func(_ *r.File) r.Type { return nil }
)

type walker struct {
	end   r.Type
	level []string
}

func (w *walker) Handle(t r.Type) error {
	if reflect.ValueOf(t).IsNil() {
		return nilErr
	}

	if t == w.end {
		w.level = append(w.level, reflect.TypeOf(t).Elem().Name())

		return sentinel
	}

	err := Walk(t, w)
	if err != nil {
		w.level = append(w.level, reflect.TypeOf(t).Elem().Name())
	}

	return err
}

func TestWalk(t *testing.T) {
	for n, test := range [...]struct {
		Input string
		End   func(f *r.File) r.Type
		Level []string
	}{
		{ // 1
			"",
			nilRet,
			nil,
		},
		{ // 2
			"a;b",
			func(r *r.File) r.Type { return r },
			[]string{"File"},
		},
		{ // 3
			"a;b",
			func(r *r.File) r.Type { return &r.Statements[0] },
			[]string{"File", "Expression"},
		},
		{ // 4
			"a;b",
			func(r *r.File) r.Type { return &r.Statements[1] },
			[]string{"File", "Expression"},
		},
		{ // 5
			"repeat {}",
			func(r *r.File) r.Type { return r.Statements[0].FlowControl },
			[]string{"File", "Expression", "FlowControl"},
		},
		{ // 6
			"function() a",
			func(r *r.File) r.Type { return r.Statements[0].FunctionDefinition },
			[]string{"File", "Expression", "FunctionDefinition"},
		},
		{ // 7
			"a",
			func(r *r.File) r.Type { return r.Statements[0].QueryExpression },
			[]string{"File", "Expression", "QueryExpression"},
		},
		{ // 8
			"if(a) b",
			func(r *r.File) r.Type { return r.Statements[0].FlowControl.IfControl },
			[]string{"File", "Expression", "FlowControl", "IfControl"},
		},
		{ // 9
			"while (a) b",
			func(r *r.File) r.Type { return r.Statements[0].FlowControl.WhileControl },
			[]string{"File", "Expression", "FlowControl", "WhileControl"},
		},
		{ // 10
			"repeat a",
			func(r *r.File) r.Type { return r.Statements[0].FlowControl.RepeatControl },
			[]string{"File", "Expression", "FlowControl", "RepeatControl"},
		},
		{ // 11
			"for (a in b) c",
			func(r *r.File) r.Type { return r.Statements[0].FlowControl.ForControl },
			[]string{"File", "Expression", "FlowControl", "ForControl"},
		},
		{ // 12
			"if(a) b\nelse c",
			func(r *r.File) r.Type { return &r.Statements[0].FlowControl.IfControl.Cond },
			[]string{"File", "Expression", "FlowControl", "IfControl", "FormulaeExpression"},
		},
		{ // 13
			"if(a) b\nelse c",
			func(r *r.File) r.Type { return &r.Statements[0].FlowControl.IfControl.Expr },
			[]string{"File", "Expression", "FlowControl", "IfControl", "Expression"},
		},
		{ // 14
			"if(a) b\nelse c",
			func(r *r.File) r.Type { return r.Statements[0].FlowControl.IfControl.Else },
			[]string{"File", "Expression", "FlowControl", "IfControl", "Expression"},
		},
		{ // 15
			"if(a) b",
			func(r *r.File) r.Type { return r.Statements[0].FlowControl.IfControl.Else },
			nil,
		},
		{ // 16
			"while (a) b",
			func(r *r.File) r.Type { return &r.Statements[0].FlowControl.WhileControl.Cond },
			[]string{"File", "Expression", "FlowControl", "WhileControl", "FormulaeExpression"},
		},
		{ // 17
			"while (a) b",
			func(r *r.File) r.Type { return &r.Statements[0].FlowControl.WhileControl.Expr },
			[]string{"File", "Expression", "FlowControl", "WhileControl", "Expression"},
		},
	} {
		tk := parser.NewStringTokeniser(test.Input)

		m, err := r.Parse(&tk)
		if err != nil {
			t.Errorf("test %d: unexpected error parsing file: %s", n+1, err)
		} else {
			w := walker{end: test.End(m)}

			if err := w.Handle(m); err == nil && test.Level != nil {
				t.Errorf("test %d: expected to recieve sentinel error, but didn't", n+1)
			} else if err != nil && test.Level == nil {
				t.Errorf("test %d: expected no error, but recieved %v", n+1, err)
			} else if len(w.level) != len(test.Level) {
				t.Errorf("test %d: expected to have %d levels, got %d", n+1, len(test.Level), len(w.level))
			} else {
				for m, l := range w.level {
					if e := test.Level[len(test.Level)-m-1]; e != l {
						t.Errorf("test %d.%d: expected to read level %s, got %s", n+1, m+1, e, l)
					}
				}
			}
		}
	}
}
