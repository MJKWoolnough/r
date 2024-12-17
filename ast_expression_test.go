package r

import "testing"

func TestCompoundExpression(t *testing.T) {
	doTests(t, []sourceFn{
		{"{a}", func(t *test, tk Tokens) { // 1
			t.Output = CompoundExpression{
				Expressions: []Expression{
					{
						QueryExpression: WrapQuery(&SimpleExpression{
							Identifier: &tk[1],
							Tokens:     tk[1:2],
						}),
						Tokens: tk[1:2],
					},
				},
				Tokens: tk[:3],
			}
		}},
		{"{ a }", func(t *test, tk Tokens) { // 2
			t.Output = CompoundExpression{
				Expressions: []Expression{
					{
						QueryExpression: WrapQuery(&SimpleExpression{
							Identifier: &tk[2],
							Tokens:     tk[2:3],
						}),
						Tokens: tk[2:3],
					},
				},
				Tokens: tk[:5],
			}
		}},
		{"{a;b}", func(t *test, tk Tokens) { // 3
			t.Output = CompoundExpression{
				Expressions: []Expression{
					{
						QueryExpression: WrapQuery(&SimpleExpression{
							Identifier: &tk[1],
							Tokens:     tk[1:2],
						}),
						Tokens: tk[1:2],
					},
					{
						QueryExpression: WrapQuery(&SimpleExpression{
							Identifier: &tk[3],
							Tokens:     tk[3:4],
						}),
						Tokens: tk[3:4],
					},
				},
				Tokens: tk[:5],
			}
		}},
		{"{ a ; b }", func(t *test, tk Tokens) { // 4
			t.Output = CompoundExpression{
				Expressions: []Expression{
					{
						QueryExpression: WrapQuery(&SimpleExpression{
							Identifier: &tk[2],
							Tokens:     tk[2:3],
						}),
						Tokens: tk[2:3],
					},
					{
						QueryExpression: WrapQuery(&SimpleExpression{
							Identifier: &tk[6],
							Tokens:     tk[6:7],
						}),
						Tokens: tk[6:7],
					},
				},
				Tokens: tk[:9],
			}
		}},
		{"{a\nb}", func(t *test, tk Tokens) { // 5
			t.Output = CompoundExpression{
				Expressions: []Expression{
					{
						QueryExpression: WrapQuery(&SimpleExpression{
							Identifier: &tk[1],
							Tokens:     tk[1:2],
						}),
						Tokens: tk[1:2],
					},
					{
						QueryExpression: WrapQuery(&SimpleExpression{
							Identifier: &tk[3],
							Tokens:     tk[3:4],
						}),
						Tokens: tk[3:4],
					},
				},
				Tokens: tk[:5],
			}
		}},
		{"{\n\ta\n\tb\n}", func(t *test, tk Tokens) { // 6
			t.Output = CompoundExpression{
				Expressions: []Expression{
					{
						QueryExpression: WrapQuery(&SimpleExpression{
							Identifier: &tk[3],
							Tokens:     tk[3:4],
						}),
						Tokens: tk[3:4],
					},
					{
						QueryExpression: WrapQuery(&SimpleExpression{
							Identifier: &tk[6],
							Tokens:     tk[6:7],
						}),
						Tokens: tk[6:7],
					},
				},
				Tokens: tk[:9],
			}
		}},
		{"{in}", func(t *test, tk Tokens) { // 7
			t.Err = wrapQueryExpressionError(Error{
				Err:     ErrInvalidSimpleExpression,
				Parsing: "SimpleExpression",
				Token:   tk[1],
			})
		}},
		{"{a a}", func(t *test, tk Tokens) { // 8
			t.Err = Error{
				Err:     ErrMissingTerminator,
				Parsing: "CompoundExpression",
				Token:   tk[2],
			}
		}},
	}, func(t *test) (Type, error) {
		var ce CompoundExpression

		err := ce.parse(&t.Tokens)

		return ce, err
	})
}

func TestSimpleExpression(t *testing.T) {
	doTests(t, []sourceFn{
		{"a", func(t *test, tk Tokens) { // 1
			t.Output = SimpleExpression{
				Identifier: &tk[0],
				Tokens:     tk[:1],
			}
		}},
		{"1", func(t *test, tk Tokens) { // 2
			t.Output = SimpleExpression{
				Constant: &tk[0],
				Tokens:   tk[:1],
			}
		}},
		{"2L", func(t *test, tk Tokens) { // 3
			t.Output = SimpleExpression{
				Constant: &tk[0],
				Tokens:   tk[:1],
			}
		}},
		{"3i", func(t *test, tk Tokens) { // 4
			t.Output = SimpleExpression{
				Constant: &tk[0],
				Tokens:   tk[:1],
			}
		}},
		{"TRUE", func(t *test, tk Tokens) { // 5
			t.Output = SimpleExpression{
				Constant: &tk[0],
				Tokens:   tk[:1],
			}
		}},
		{"FALSE", func(t *test, tk Tokens) { // 6
			t.Output = SimpleExpression{
				Constant: &tk[0],
				Tokens:   tk[:1],
			}
		}},
		{"\"abc\"", func(t *test, tk Tokens) { // 7
			t.Output = SimpleExpression{
				Constant: &tk[0],
				Tokens:   tk[:1],
			}
		}},
		{"NULL", func(t *test, tk Tokens) { // 8
			t.Output = SimpleExpression{
				Constant: &tk[0],
				Tokens:   tk[:1],
			}
		}},
		{"NA", func(t *test, tk Tokens) { // 9
			t.Output = SimpleExpression{
				Constant: &tk[0],
				Tokens:   tk[:1],
			}
		}},
		{"...", func(t *test, tk Tokens) { // 10
			t.Output = SimpleExpression{
				Ellipsis: &tk[0],
				Tokens:   tk[:1],
			}
		}},
		{"..1", func(t *test, tk Tokens) { // 11
			t.Output = SimpleExpression{
				Ellipsis: &tk[0],
				Tokens:   tk[:1],
			}
		}},
		{"(a)", func(t *test, tk Tokens) { // 12
			t.Output = SimpleExpression{
				ParenthesizedExpression: &Expression{
					QueryExpression: WrapQuery(&SimpleExpression{
						Identifier: &tk[1],
						Tokens:     tk[1:2],
					}),
					Tokens: tk[1:2],
				},
				Tokens: tk[:3],
			}
		}},
		{"{a}", func(t *test, tk Tokens) { // 13
			t.Output = SimpleExpression{
				CompoundExpression: &CompoundExpression{
					Expressions: []Expression{
						{
							QueryExpression: WrapQuery(&SimpleExpression{
								Identifier: &tk[1],
								Tokens:     tk[1:2],
							}),
							Tokens: tk[1:2],
						},
					},
					Tokens: tk[:3],
				},
				Tokens: tk[:3],
			}
		}},
		{"in", func(t *test, tk Tokens) { // 14
			t.Err = Error{
				Err:     ErrInvalidSimpleExpression,
				Parsing: "SimpleExpression",
				Token:   tk[0],
			}
		}},
	}, func(t *test) (Type, error) {
		var se SimpleExpression

		err := se.parse(&t.Tokens)

		return se, err
	})
}
