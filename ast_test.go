package r

import (
	"errors"
	"reflect"
	"testing"

	"vimagination.zapto.org/parser"
)

type sourceFn struct {
	Source string
	Fn     func(*test, Tokens)
}

type test struct {
	Tokens rParser
	Output Type
	Err    error
}

func makeTokeniser(tk parser.Tokeniser) *parser.Tokeniser {
	return &tk
}

func doTests(t *testing.T, tests []sourceFn, fn func(*test) (Type, error)) {
	t.Helper()

	var err error

	for n, tt := range tests {
		var ts test

		if ts.Tokens, err = newRParser(makeTokeniser(parser.NewStringTokeniser(tt.Source))); err != nil {
			t.Errorf("test %d: unexpected error: %s", n+1, err)

			continue
		}

		tt.Fn(&ts, Tokens(ts.Tokens[:cap(ts.Tokens)]))

		if output, err := fn(&ts); !errors.Is(err, ts.Err) {
			t.Errorf("test %d: expecting error: %v, got %v", n+1, ts.Err, err)
		} else if ts.Output != nil && !reflect.DeepEqual(output, ts.Output) {
			t.Errorf("test %d: expecting \n%+v\n...got...\n%+v", n+1, ts.Output, output)
		}
	}
}

func wrapQueryExpressionError(err Error) Error {
	switch err.Parsing {
	case "CompoundExpression":
		err = Error{
			Err:     err,
			Parsing: "SimpleExpression",
			Token:   err.Token,
		}

		fallthrough
	case "SimpleExpression":
		err = Error{
			Err:     err,
			Parsing: "IndexOrCallExpression",
			Token:   err.Token,
		}

		fallthrough
	case "IndexOrCallExpression":
		err = Error{
			Err:     err,
			Parsing: "ScopeExpression",
			Token:   err.Token,
		}

		fallthrough
	case "ScopeExpression":
		err = Error{
			Err:     err,
			Parsing: "SubsetExpression",
			Token:   err.Token,
		}

		fallthrough
	case "SubsetExpression":
		err = Error{
			Err:     err,
			Parsing: "ExponentiationExpression",
			Token:   err.Token,
		}

		fallthrough
	case "ExponentiationExpression":
		err = Error{
			Err:     err,
			Parsing: "UnaryExpression",
			Token:   err.Token,
		}

		fallthrough
	case "UnaryExpression":
		err = Error{
			Err:     err,
			Parsing: "SequenceExpression",
			Token:   err.Token,
		}

		fallthrough
	case "SequenceExpression":
		err = Error{
			Err:     err,
			Parsing: "PipeOrSpecialExpression",
			Token:   err.Token,
		}

		fallthrough
	case "PipeOrSpecialExpression":
		err = Error{
			Err:     err,
			Parsing: "MultiplicationExpression",
			Token:   err.Token,
		}

		fallthrough
	case "MultiplicationExpression":
		err = Error{
			Err:     err,
			Parsing: "AdditionExpression",
			Token:   err.Token,
		}

		fallthrough
	case "AdditionExpression":
		err = Error{
			Err:     err,
			Parsing: "RelationalExpression",
			Token:   err.Token,
		}

		fallthrough
	case "RelationalExpression":
		err = Error{
			Err:     err,
			Parsing: "NotExpression",
			Token:   err.Token,
		}

		fallthrough
	case "NotExpression":
		err = Error{
			Err:     err,
			Parsing: "AndExpression",
			Token:   err.Token,
		}

		fallthrough
	case "AndExpression":
		err = Error{
			Err:     err,
			Parsing: "OrExpression",
			Token:   err.Token,
		}

		fallthrough
	case "OrExpression":
		err = Error{
			Err:     err,
			Parsing: "FormulaeExpression",
			Token:   err.Token,
		}
		fallthrough
	case "FormulaeExpression":
		err = Error{
			Err:     err,
			Parsing: "AssignmentExpression",
			Token:   err.Token,
		}
		fallthrough
	case "AssignmentExpression":
		err = Error{
			Err:     err,
			Parsing: "QueryExpression",
			Token:   err.Token,
		}
	}

	return err
}

func TestFile(t *testing.T) {
	doTests(t, []sourceFn{
		{"a", func(t *test, tk Tokens) { // 1
			t.Output = File{
				Statements: []Expression{
					{
						QueryExpression: WrapQuery(&SimpleExpression{
							Identifier: &tk[0],
							Tokens:     tk[:1],
						}),
						Tokens: tk[:1],
					},
				},
				Tokens: tk[:1],
			}
		}},
		{"a;b", func(t *test, tk Tokens) { // 2
			t.Output = File{
				Statements: []Expression{
					{
						QueryExpression: WrapQuery(&SimpleExpression{
							Identifier: &tk[0],
							Tokens:     tk[:1],
						}),
						Tokens: tk[:1],
					},
					{
						QueryExpression: WrapQuery(&SimpleExpression{
							Identifier: &tk[2],
							Tokens:     tk[2:3],
						}),
						Tokens: tk[2:3],
					},
				},
				Tokens: tk[:3],
			}
		}},
		{"a\nb", func(t *test, tk Tokens) { // 3
			t.Output = File{
				Statements: []Expression{
					{
						QueryExpression: WrapQuery(&SimpleExpression{
							Identifier: &tk[0],
							Tokens:     tk[:1],
						}),
						Tokens: tk[:1],
					},
					{
						QueryExpression: WrapQuery(&SimpleExpression{
							Identifier: &tk[2],
							Tokens:     tk[2:3],
						}),
						Tokens: tk[2:3],
					},
				},
				Tokens: tk[:3],
			}
		}},
	}, func(t *test) (Type, error) {
		var f File

		err := f.parse(&t.Tokens)

		return f, err
	})
}
