package r

import "testing"

func TestExpression(t *testing.T) {
	doTests(t, []sourceFn{
		{"if(a)b", func(t *test, tk Tokens) { // 1
			t.Output = Expression{
				FlowControl: &FlowControl{
					IfControl: &IfControl{
						Cond: WrapQuery(&SimpleExpression{
							Identifier: &tk[2],
							Tokens:     tk[2:3],
						}).AssignmentExpression.FormulaeExpression,
						Expr: Expression{
							QueryExpression: WrapQuery(&SimpleExpression{
								Identifier: &tk[4],
								Tokens:     tk[4:5],
							}),
							Tokens: tk[4:5],
						},
						Tokens: tk[:5],
					},
					Tokens: tk[:5],
				},
				Tokens: tk[:5],
			}
		}},
		{"function()a", func(t *test, tk Tokens) { // 2
			t.Output = Expression{
				FunctionDefinition: &FunctionDefinition{
					ArgList: ArgList{
						Tokens: tk[2:2],
					},
					Body: Expression{
						QueryExpression: WrapQuery(&SimpleExpression{
							Identifier: &tk[3],
							Tokens:     tk[3:4],
						}),
						Tokens: tk[3:4],
					},
					Tokens: tk[:4],
				},
				Tokens: tk[:4],
			}
		}},
		{"a", func(t *test, tk Tokens) { // 3
			t.Output = Expression{
				QueryExpression: WrapQuery(&SimpleExpression{
					Identifier: &tk[0],
					Tokens:     tk[:1],
				}),
				Tokens: tk[:1],
			}
		}},
		{"if a", func(t *test, tk Tokens) { // 4
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err:     ErrMissingOpeningParen,
						Parsing: "IfControl",
						Token:   tk[2],
					},
					Parsing: "FlowControl",
					Token:   tk[0],
				},
				Parsing: "Expression",
				Token:   tk[0],
			}
		}},
		{"function a", func(t *test, tk Tokens) { // 5
			t.Err = Error{
				Err: Error{
					Err:     ErrMissingOpeningParen,
					Parsing: "FunctionDefinition",
					Token:   tk[2],
				},
				Parsing: "Expression",
				Token:   tk[0],
			}
		}},
		{"in", func(t *test, tk Tokens) { // 6
			t.Err = Error{
				Err: wrapQueryExpressionError(Error{
					Err:     ErrInvalidSimpleExpression,
					Parsing: "SimpleExpression",
					Token:   tk[0],
				}),
				Parsing: "Expression",
				Token:   tk[0],
			}
		}},
		{"#a comment\na", func(t *test, tk Tokens) { // 7
			t.Output = Expression{
				QueryExpression: WrapQuery(&SimpleExpression{
					Identifier: &tk[2],
					Tokens:     tk[2:3],
				}),
				Comments: [2]Comments{{tk[0]}, nil},
				Tokens:   tk[:3],
			}
		}},
		{"#a comment\n# Another Comment\na", func(t *test, tk Tokens) { // 8
			t.Output = Expression{
				QueryExpression: WrapQuery(&SimpleExpression{
					Identifier: &tk[4],
					Tokens:     tk[4:5],
				}),
				Comments: [2]Comments{{tk[0], tk[2]}, nil},
				Tokens:   tk[:5],
			}
		}},
		{"#a comment\na # attached comment\n#another attached", func(t *test, tk Tokens) { // 9
			t.Output = Expression{
				QueryExpression: WrapQuery(&SimpleExpression{
					Identifier: &tk[2],
					Tokens:     tk[2:3],
				}),
				Comments: [2]Comments{{tk[0]}, {tk[4], tk[6]}},
				Tokens:   tk[:7],
			}
		}},
	}, func(t *test) (Type, error) {
		var e Expression

		err := e.parse(&t.Tokens)

		return e, err
	})
}

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
		{"{# abc\na #def\n\n#ghi\n}", func(t *test, tk Tokens) { // 9
			t.Output = CompoundExpression{
				Expressions: []Expression{
					{
						QueryExpression: WrapQuery(&SimpleExpression{
							Identifier: &tk[3],
							Tokens:     tk[3:4],
						}),
						Comments: [2]Comments{{tk[1]}, {tk[5]}},
						Tokens:   tk[1:6],
					},
				},
				Comments: Comments{tk[8]},
				Tokens:   tk[:11],
			}
		}},
	}, func(t *test) (Type, error) {
		var ce CompoundExpression

		err := ce.parse(&t.Tokens)

		return ce, err
	})
}

func TestFlowControl(t *testing.T) {
	doTests(t, []sourceFn{
		{"if(a)b", func(t *test, tk Tokens) { // 1
			t.Output = FlowControl{
				IfControl: &IfControl{
					Cond: WrapQuery(&SimpleExpression{
						Identifier: &tk[2],
						Tokens:     tk[2:3],
					}).AssignmentExpression.FormulaeExpression,
					Expr: Expression{
						QueryExpression: WrapQuery(&SimpleExpression{
							Identifier: &tk[4],
							Tokens:     tk[4:5],
						}),
						Tokens: tk[4:5],
					},
					Tokens: tk[:5],
				},
				Tokens: tk[:5],
			}
		}},
		{"while(a)b", func(t *test, tk Tokens) { // 2
			t.Output = FlowControl{
				WhileControl: &WhileControl{
					Cond: WrapQuery(&SimpleExpression{
						Identifier: &tk[2],
						Tokens:     tk[2:3],
					}).AssignmentExpression.FormulaeExpression,
					Expr: Expression{
						QueryExpression: WrapQuery(&SimpleExpression{
							Identifier: &tk[4],
							Tokens:     tk[4:5],
						}),
						Tokens: tk[4:5],
					},
					Tokens: tk[:5],
				},
				Tokens: tk[:5],
			}
		}},
		{"repeat a", func(t *test, tk Tokens) { // 3
			t.Output = FlowControl{
				RepeatControl: &RepeatControl{
					Expr: Expression{
						QueryExpression: WrapQuery(&SimpleExpression{
							Identifier: &tk[2],
							Tokens:     tk[2:3],
						}),
						Tokens: tk[2:3],
					},
					Tokens: tk[:3],
				},
				Tokens: tk[:3],
			}
		}},
		{"for(a in b)c", func(t *test, tk Tokens) { // 4
			t.Output = FlowControl{
				ForControl: &ForControl{
					Var: &tk[2],
					List: WrapQuery(&SimpleExpression{
						Identifier: &tk[6],
						Tokens:     tk[6:7],
					}).AssignmentExpression.FormulaeExpression,
					Expr: Expression{
						QueryExpression: WrapQuery(&SimpleExpression{
							Identifier: &tk[8],
							Tokens:     tk[8:9],
						}),
						Tokens: tk[8:9],
					},
					Tokens: tk[:9],
				},
				Tokens: tk[:9],
			}
		}},
		{"if a", func(t *test, tk Tokens) { // 5
			t.Err = Error{
				Err: Error{
					Err:     ErrMissingOpeningParen,
					Parsing: "IfControl",
					Token:   tk[2],
				},
				Parsing: "FlowControl",
				Token:   tk[0],
			}
		}},
		{"while a", func(t *test, tk Tokens) { // 6
			t.Err = Error{
				Err: Error{
					Err:     ErrMissingOpeningParen,
					Parsing: "WhileControl",
					Token:   tk[2],
				},
				Parsing: "FlowControl",
				Token:   tk[0],
			}
		}},
		{"repeat in", func(t *test, tk Tokens) { // 7
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: wrapQueryExpressionError(Error{
							Err:     ErrInvalidSimpleExpression,
							Parsing: "SimpleExpression",
							Token:   tk[2],
						}),
						Parsing: "Expression",
						Token:   tk[2],
					},
					Parsing: "RepeatControl",
					Token:   tk[2],
				},
				Parsing: "FlowControl",
				Token:   tk[0],
			}
		}},
		{"for a", func(t *test, tk Tokens) { // 8
			t.Err = Error{
				Err: Error{
					Err:     ErrMissingOpeningParen,
					Parsing: "ForControl",
					Token:   tk[2],
				},
				Parsing: "FlowControl",
				Token:   tk[0],
			}
		}},
	}, func(t *test) (Type, error) {
		var fc FlowControl

		err := fc.parse(&t.Tokens)

		return fc, err
	})
}

func TestIfControl(t *testing.T) {
	doTests(t, []sourceFn{
		{"if(a)b", func(t *test, tk Tokens) { // 1
			t.Output = IfControl{
				Cond: WrapQuery(&SimpleExpression{
					Identifier: &tk[2],
					Tokens:     tk[2:3],
				}).AssignmentExpression.FormulaeExpression,
				Expr: Expression{
					QueryExpression: WrapQuery(&SimpleExpression{
						Identifier: &tk[4],
						Tokens:     tk[4:5],
					}),
					Tokens: tk[4:5],
				},
				Tokens: tk[:5],
			}
		}},
		{"if ( a ) b", func(t *test, tk Tokens) { // 2
			t.Output = IfControl{
				Cond: WrapQuery(&SimpleExpression{
					Identifier: &tk[4],
					Tokens:     tk[4:5],
				}).AssignmentExpression.FormulaeExpression,
				Expr: Expression{
					QueryExpression: WrapQuery(&SimpleExpression{
						Identifier: &tk[8],
						Tokens:     tk[8:9],
					}),
					Tokens: tk[8:9],
				},
				Tokens: tk[:9],
			}
		}},
		{"if ( a ) b else c", func(t *test, tk Tokens) { // 3
			t.Output = IfControl{
				Cond: WrapQuery(&SimpleExpression{
					Identifier: &tk[4],
					Tokens:     tk[4:5],
				}).AssignmentExpression.FormulaeExpression,
				Expr: Expression{
					QueryExpression: WrapQuery(&SimpleExpression{
						Identifier: &tk[8],
						Tokens:     tk[8:9],
					}),
					Tokens: tk[8:9],
				},
				Else: &Expression{
					QueryExpression: WrapQuery(&SimpleExpression{
						Identifier: &tk[12],
						Tokens:     tk[12:13],
					}),
					Tokens: tk[12:13],
				},
				Tokens: tk[:13],
			}
		}},
		{"if a", func(t *test, tk Tokens) { // 4
			t.Err = Error{
				Err:     ErrMissingOpeningParen,
				Parsing: "IfControl",
				Token:   tk[2],
			}
		}},
		{"if(in)", func(t *test, tk Tokens) { // 5
			t.Err = Error{
				Err: wrapQueryExpressionError(Error{
					Err:     ErrInvalidSimpleExpression,
					Parsing: "SimpleExpression",
					Token:   tk[2],
				}).Err.(Error).Err,
				Parsing: "IfControl",
				Token:   tk[2],
			}
		}},
		{"if(a b)", func(t *test, tk Tokens) { // 6
			t.Err = Error{
				Err:     ErrMissingClosingParen,
				Parsing: "IfControl",
				Token:   tk[4],
			}
		}},
		{"if(a)in", func(t *test, tk Tokens) { // 7
			t.Err = Error{
				Err: Error{
					Err: wrapQueryExpressionError(Error{
						Err:     ErrInvalidSimpleExpression,
						Parsing: "SimpleExpression",
						Token:   tk[4],
					}),
					Parsing: "Expression",
					Token:   tk[4],
				},
				Parsing: "IfControl",
				Token:   tk[4],
			}
		}},
		{"if(a)b else in", func(t *test, tk Tokens) { // 8
			t.Err = Error{
				Err: Error{
					Err: wrapQueryExpressionError(Error{
						Err:     ErrInvalidSimpleExpression,
						Parsing: "SimpleExpression",
						Token:   tk[8],
					}),
					Parsing: "Expression",
					Token:   tk[8],
				},
				Parsing: "IfControl",
				Token:   tk[8],
			}
		}},
	}, func(t *test) (Type, error) {
		var ic IfControl

		err := ic.parse(&t.Tokens)

		return ic, err
	})
}

func TestWhileControl(t *testing.T) {
	doTests(t, []sourceFn{
		{"while(a)b", func(t *test, tk Tokens) { // 1
			t.Output = WhileControl{
				Cond: WrapQuery(&SimpleExpression{
					Identifier: &tk[2],
					Tokens:     tk[2:3],
				}).AssignmentExpression.FormulaeExpression,
				Expr: Expression{
					QueryExpression: WrapQuery(&SimpleExpression{
						Identifier: &tk[4],
						Tokens:     tk[4:5],
					}),
					Tokens: tk[4:5],
				},
				Tokens: tk[:5],
			}
		}},
		{"while ( a ) b", func(t *test, tk Tokens) { // 2
			t.Output = WhileControl{
				Cond: WrapQuery(&SimpleExpression{
					Identifier: &tk[4],
					Tokens:     tk[4:5],
				}).AssignmentExpression.FormulaeExpression,
				Expr: Expression{
					QueryExpression: WrapQuery(&SimpleExpression{
						Identifier: &tk[8],
						Tokens:     tk[8:9],
					}),
					Tokens: tk[8:9],
				},
				Tokens: tk[:9],
			}
		}},
		{"while a", func(t *test, tk Tokens) { // 3
			t.Err = Error{
				Err:     ErrMissingOpeningParen,
				Parsing: "WhileControl",
				Token:   tk[2],
			}
		}},
		{"while(in)", func(t *test, tk Tokens) { // 4
			t.Err = Error{
				Err: wrapQueryExpressionError(Error{
					Err:     ErrInvalidSimpleExpression,
					Parsing: "SimpleExpression",
					Token:   tk[2],
				}).Err.(Error).Err,
				Parsing: "WhileControl",
				Token:   tk[2],
			}
		}},
		{"while(a b)", func(t *test, tk Tokens) { // 5
			t.Err = Error{
				Err:     ErrMissingClosingParen,
				Parsing: "WhileControl",
				Token:   tk[4],
			}
		}},
		{"while(a)in", func(t *test, tk Tokens) { // 6
			t.Err = Error{
				Err: Error{
					Err: wrapQueryExpressionError(Error{
						Err:     ErrInvalidSimpleExpression,
						Parsing: "SimpleExpression",
						Token:   tk[4],
					}),
					Parsing: "Expression",
					Token:   tk[4],
				},
				Parsing: "WhileControl",
				Token:   tk[4],
			}
		}},
	}, func(t *test) (Type, error) {
		var wc WhileControl

		err := wc.parse(&t.Tokens)

		return wc, err
	})
}

func TestRepeatControl(t *testing.T) {
	doTests(t, []sourceFn{
		{"repeat a", func(t *test, tk Tokens) { // 1
			t.Output = RepeatControl{
				Expr: Expression{
					QueryExpression: WrapQuery(&SimpleExpression{
						Identifier: &tk[2],
						Tokens:     tk[2:3],
					}),
					Tokens: tk[2:3],
				},
				Tokens: tk[:3],
			}
		}},
		{"repeat in", func(t *test, tk Tokens) { // 2
			t.Err = Error{
				Err: Error{
					Err: wrapQueryExpressionError(Error{
						Err:     ErrInvalidSimpleExpression,
						Parsing: "SimpleExpression",
						Token:   tk[2],
					}),
					Parsing: "Expression",
					Token:   tk[2],
				},
				Parsing: "RepeatControl",
				Token:   tk[2],
			}
		}},
	}, func(t *test) (Type, error) {
		var rc RepeatControl

		err := rc.parse(&t.Tokens)

		return rc, err
	})
}

func TestForControl(t *testing.T) {
	doTests(t, []sourceFn{
		{"for(a in b)c", func(t *test, tk Tokens) { // 1
			t.Output = ForControl{
				Var: &tk[2],
				List: WrapQuery(&SimpleExpression{
					Identifier: &tk[6],
					Tokens:     tk[6:7],
				}).AssignmentExpression.FormulaeExpression,
				Expr: Expression{
					QueryExpression: WrapQuery(&SimpleExpression{
						Identifier: &tk[8],
						Tokens:     tk[8:9],
					}),
					Tokens: tk[8:9],
				},
				Tokens: tk[:9],
			}
		}},
		{"for(a in b) c", func(t *test, tk Tokens) { // 2
			t.Output = ForControl{
				Var: &tk[2],
				List: WrapQuery(&SimpleExpression{
					Identifier: &tk[6],
					Tokens:     tk[6:7],
				}).AssignmentExpression.FormulaeExpression,
				Expr: Expression{
					QueryExpression: WrapQuery(&SimpleExpression{
						Identifier: &tk[9],
						Tokens:     tk[9:10],
					}),
					Tokens: tk[9:10],
				},
				Tokens: tk[:10],
			}
		}},
		{"for a", func(t *test, tk Tokens) { // 3
			t.Err = Error{
				Err:     ErrMissingOpeningParen,
				Parsing: "ForControl",
				Token:   tk[2],
			}
		}},
		{"for(1)", func(t *test, tk Tokens) { // 4
			t.Err = Error{
				Err:     ErrMissingIdentifier,
				Parsing: "ForControl",
				Token:   tk[2],
			}
		}},
		{"for(a b)", func(t *test, tk Tokens) { // 5
			t.Err = Error{
				Err:     ErrMissingIn,
				Parsing: "ForControl",
				Token:   tk[4],
			}
		}},
		{"for(a in in)", func(t *test, tk Tokens) { // 6
			t.Err = Error{
				Err: wrapQueryExpressionError(Error{
					Err:     ErrInvalidSimpleExpression,
					Parsing: "SimpleExpression",
					Token:   tk[6],
				}).Err.(Error).Err,
				Parsing: "ForControl",
				Token:   tk[6],
			}
		}},
		{"for(a in b c)", func(t *test, tk Tokens) { // 7
			t.Err = Error{
				Err:     ErrMissingClosingParen,
				Parsing: "ForControl",
				Token:   tk[8],
			}
		}},
		{"for(a in b)in", func(t *test, tk Tokens) { // 8
			t.Err = Error{
				Err: Error{
					Err: wrapQueryExpressionError(Error{
						Err:     ErrInvalidSimpleExpression,
						Parsing: "SimpleExpression",
						Token:   tk[8],
					}),
					Parsing: "Expression",
					Token:   tk[8],
				},
				Parsing: "ForControl",
				Token:   tk[8],
			}
		}},
	}, func(t *test) (Type, error) {
		var fc ForControl

		err := fc.parse(&t.Tokens)

		return fc, err
	})
}

func TestFunctionDefinition(t *testing.T) {
	doTests(t, []sourceFn{
		{"function()a", func(t *test, tk Tokens) { // 1
			t.Output = FunctionDefinition{
				ArgList: ArgList{
					Tokens: tk[2:2],
				},
				Body: Expression{
					QueryExpression: WrapQuery(&SimpleExpression{
						Identifier: &tk[3],
						Tokens:     tk[3:4],
					}),
					Tokens: tk[3:4],
				},
				Tokens: tk[:4],
			}
		}},
		{"function ( ) a", func(t *test, tk Tokens) { // 2
			t.Output = FunctionDefinition{
				ArgList: ArgList{
					Tokens: tk[3:3],
				},
				Body: Expression{
					QueryExpression: WrapQuery(&SimpleExpression{
						Identifier: &tk[6],
						Tokens:     tk[6:7],
					}),
					Tokens: tk[6:7],
				},
				Tokens: tk[:7],
			}
		}},
		{"function(a)b", func(t *test, tk Tokens) { // 3
			t.Output = FunctionDefinition{
				ArgList: ArgList{
					Args: []Argument{
						{
							Identifier: &tk[2],
							Tokens:     tk[2:3],
						},
					},
					Tokens: tk[2:3],
				},
				Body: Expression{
					QueryExpression: WrapQuery(&SimpleExpression{
						Identifier: &tk[4],
						Tokens:     tk[4:5],
					}),
					Tokens: tk[4:5],
				},
				Tokens: tk[:5],
			}
		}},
		{"function(a, b){}", func(t *test, tk Tokens) { // 4
			t.Output = FunctionDefinition{
				ArgList: ArgList{
					Args: []Argument{
						{
							Identifier: &tk[2],
							Tokens:     tk[2:3],
						},
						{
							Identifier: &tk[5],
							Tokens:     tk[5:6],
						},
					},
					Tokens: tk[2:6],
				},
				Body: Expression{
					QueryExpression: WrapQuery(&CompoundExpression{
						Tokens: tk[7:9],
					}),
					Tokens: tk[7:9],
				},
				Tokens: tk[:9],
			}
		}},
		{"function a", func(t *test, tk Tokens) { // 5
			t.Err = Error{
				Err:     ErrMissingOpeningParen,
				Parsing: "FunctionDefinition",
				Token:   tk[2],
			}
		}},
		{"function(in)a", func(t *test, tk Tokens) { // 6
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err:     ErrMissingIdentifier,
						Parsing: "Argument",
						Token:   tk[2],
					},
					Parsing: "ArgList",
					Token:   tk[2],
				},
				Parsing: "FunctionDefinition",
				Token:   tk[2],
			}
		}},
		{"function()in", func(t *test, tk Tokens) { // 7
			t.Err = Error{
				Err: Error{
					Err: wrapQueryExpressionError(Error{
						Err:     ErrInvalidSimpleExpression,
						Parsing: "SimpleExpression",
						Token:   tk[3],
					}),
					Parsing: "Expression",
					Token:   tk[3],
				},
				Parsing: "FunctionDefinition",
				Token:   tk[3],
			}
		}},
	}, func(t *test) (Type, error) {
		var fd FunctionDefinition

		err := fd.parse(&t.Tokens)

		return fd, err
	})
}

func TestArgList(t *testing.T) {
	doTests(t, []sourceFn{
		{"a", func(t *test, tk Tokens) { // 1
			t.Output = ArgList{
				Args: []Argument{
					{
						Identifier: &tk[0],
						Tokens:     tk[:1],
					},
				},
				Tokens: tk[:1],
			}
		}},
		{"a,b", func(t *test, tk Tokens) { // 2
			t.Output = ArgList{
				Args: []Argument{
					{
						Identifier: &tk[0],
						Tokens:     tk[:1],
					},
					{
						Identifier: &tk[2],
						Tokens:     tk[2:3],
					},
				},
				Tokens: tk[:3],
			}
		}},
		{"a , b", func(t *test, tk Tokens) { // 3
			t.Output = ArgList{
				Args: []Argument{
					{
						Identifier: &tk[0],
						Tokens:     tk[:1],
					},
					{
						Identifier: &tk[4],
						Tokens:     tk[4:5],
					},
				},
				Tokens: tk[:5],
			}
		}},
		{"a, b, c", func(t *test, tk Tokens) { // 4
			t.Output = ArgList{
				Args: []Argument{
					{
						Identifier: &tk[0],
						Tokens:     tk[:1],
					},
					{
						Identifier: &tk[3],
						Tokens:     tk[3:4],
					},
					{
						Identifier: &tk[6],
						Tokens:     tk[6:7],
					},
				},
				Tokens: tk[:7],
			}
		}},
		{"", func(t *test, tk Tokens) { // 5
			t.Output = ArgList{
				Tokens: tk[:0],
			}
		}},
		{"in", func(t *test, tk Tokens) { // 6
			t.Err = Error{
				Err: Error{
					Err:     ErrMissingIdentifier,
					Parsing: "Argument",
					Token:   tk[0],
				},
				Parsing: "ArgList",
				Token:   tk[0],
			}
		}},
		{"a b", func(t *test, tk Tokens) { // 7
			t.Err = Error{
				Err:     ErrMissingTerminator,
				Parsing: "ArgList",
				Token:   tk[2],
			}
		}},
	}, func(t *test) (Type, error) {
		var al ArgList

		err := al.parse(&t.Tokens)

		return al, err
	})
}

func TestArgument(t *testing.T) {
	doTests(t, []sourceFn{
		{"a", func(t *test, tk Tokens) { // 1
			t.Output = Argument{
				Identifier: &tk[0],
				Tokens:     tk[:1],
			}
		}},
		{"a=b", func(t *test, tk Tokens) { // 2
			t.Output = Argument{
				Identifier: &tk[0],
				Default: &Expression{
					QueryExpression: WrapQuery(&SimpleExpression{
						Identifier: &tk[2],
						Tokens:     tk[2:3],
					}),
					Tokens: tk[2:3],
				},
				Tokens: tk[:3],
			}
		}},
		{"a = b", func(t *test, tk Tokens) { // 3
			t.Output = Argument{
				Identifier: &tk[0],
				Default: &Expression{
					QueryExpression: WrapQuery(&SimpleExpression{
						Identifier: &tk[4],
						Tokens:     tk[4:5],
					}),
					Tokens: tk[4:5],
				},
				Tokens: tk[:5],
			}
		}},
		{"...", func(t *test, tk Tokens) { // 4
			t.Output = Argument{
				Identifier: &tk[0],
				Tokens:     tk[:1],
			}
		}},
		{"in", func(t *test, tk Tokens) { // 5
			t.Err = Error{
				Err:     ErrMissingIdentifier,
				Parsing: "Argument",
				Token:   tk[0],
			}
		}},
		{"a=in", func(t *test, tk Tokens) { // 6
			t.Err = Error{
				Err: Error{
					Err: wrapQueryExpressionError(Error{
						Err:     ErrInvalidSimpleExpression,
						Parsing: "SimpleExpression",
						Token:   tk[2],
					}),
					Parsing: "Expression",
					Token:   tk[2],
				},
				Parsing: "Argument",
				Token:   tk[2],
			}
		}},
	}, func(t *test) (Type, error) {
		var a Argument

		err := a.parse(&t.Tokens)

		return a, err
	})
}

func TestQueryExpression(t *testing.T) {
	doTests(t, []sourceFn{
		{"a", func(t *test, tk Tokens) { // 1
			t.Output = QueryExpression{
				AssignmentExpression: &AssignmentExpression{
					FormulaeExpression: FormulaeExpression{
						OrExpression: &OrExpression{
							AndExpression: AndExpression{
								NotExpression: NotExpression{
									RelationalExpression: RelationalExpression{
										AdditionExpression: AdditionExpression{
											MultiplicationExpression: MultiplicationExpression{
												PipeOrSpecialExpression: PipeOrSpecialExpression{
													SequenceExpression: SequenceExpression{
														UnaryExpression: UnaryExpression{
															ExponentiationExpression: ExponentiationExpression{
																SubsetExpression: SubsetExpression{
																	ScopeExpression: ScopeExpression{
																		IndexOrCallExpression: IndexOrCallExpression{
																			SimpleExpression: &SimpleExpression{
																				Identifier: &tk[0],
																				Tokens:     tk[:1],
																			},
																			Tokens: tk[:1],
																		},
																		Tokens: tk[:1],
																	},
																	Tokens: tk[:1],
																},
																Tokens: tk[:1],
															},
															Tokens: tk[:1],
														},
														Tokens: tk[:1],
													},
													Tokens: tk[:1],
												},
												Tokens: tk[:1],
											},
											Tokens: tk[:1],
										},
										Tokens: tk[:1],
									},
									Tokens: tk[:1],
								},
								Tokens: tk[:1],
							},
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				Tokens: tk[:1],
			}
		}},
		{"?a", func(t *test, tk Tokens) { // 2
			t.Output = QueryExpression{
				QueryExpression: &QueryExpression{
					AssignmentExpression: &AssignmentExpression{
						FormulaeExpression: FormulaeExpression{
							OrExpression: &OrExpression{
								AndExpression: AndExpression{
									NotExpression: NotExpression{
										RelationalExpression: RelationalExpression{
											AdditionExpression: AdditionExpression{
												MultiplicationExpression: MultiplicationExpression{
													PipeOrSpecialExpression: PipeOrSpecialExpression{
														SequenceExpression: SequenceExpression{
															UnaryExpression: UnaryExpression{
																ExponentiationExpression: ExponentiationExpression{
																	SubsetExpression: SubsetExpression{
																		ScopeExpression: ScopeExpression{
																			IndexOrCallExpression: IndexOrCallExpression{
																				SimpleExpression: &SimpleExpression{
																					Identifier: &tk[1],
																					Tokens:     tk[1:2],
																				},
																				Tokens: tk[1:2],
																			},
																			Tokens: tk[1:2],
																		},
																		Tokens: tk[1:2],
																	},
																	Tokens: tk[1:2],
																},
																Tokens: tk[1:2],
															},
															Tokens: tk[1:2],
														},
														Tokens: tk[1:2],
													},
													Tokens: tk[1:2],
												},
												Tokens: tk[1:2],
											},
											Tokens: tk[1:2],
										},
										Tokens: tk[1:2],
									},
									Tokens: tk[1:2],
								},
								Tokens: tk[1:2],
							},
							Tokens: tk[1:2],
						},
						Tokens: tk[1:2],
					},
					Tokens: tk[1:2],
				},
				Tokens: tk[:2],
			}
		}},
		{"? a", func(t *test, tk Tokens) { // 3
			t.Output = QueryExpression{
				QueryExpression: &QueryExpression{
					AssignmentExpression: &AssignmentExpression{
						FormulaeExpression: FormulaeExpression{
							OrExpression: &OrExpression{
								AndExpression: AndExpression{
									NotExpression: NotExpression{
										RelationalExpression: RelationalExpression{
											AdditionExpression: AdditionExpression{
												MultiplicationExpression: MultiplicationExpression{
													PipeOrSpecialExpression: PipeOrSpecialExpression{
														SequenceExpression: SequenceExpression{
															UnaryExpression: UnaryExpression{
																ExponentiationExpression: ExponentiationExpression{
																	SubsetExpression: SubsetExpression{
																		ScopeExpression: ScopeExpression{
																			IndexOrCallExpression: IndexOrCallExpression{
																				SimpleExpression: &SimpleExpression{
																					Identifier: &tk[2],
																					Tokens:     tk[2:3],
																				},
																				Tokens: tk[2:3],
																			},
																			Tokens: tk[2:3],
																		},
																		Tokens: tk[2:3],
																	},
																	Tokens: tk[2:3],
																},
																Tokens: tk[2:3],
															},
															Tokens: tk[2:3],
														},
														Tokens: tk[2:3],
													},
													Tokens: tk[2:3],
												},
												Tokens: tk[2:3],
											},
											Tokens: tk[2:3],
										},
										Tokens: tk[2:3],
									},
									Tokens: tk[2:3],
								},
								Tokens: tk[2:3],
							},
							Tokens: tk[2:3],
						},
						Tokens: tk[2:3],
					},
					Tokens: tk[2:3],
				},
				Tokens: tk[:3],
			}
		}},
		{"a?b", func(t *test, tk Tokens) { // 4
			t.Output = QueryExpression{
				AssignmentExpression: &AssignmentExpression{
					FormulaeExpression: FormulaeExpression{
						OrExpression: &OrExpression{
							AndExpression: AndExpression{
								NotExpression: NotExpression{
									RelationalExpression: RelationalExpression{
										AdditionExpression: AdditionExpression{
											MultiplicationExpression: MultiplicationExpression{
												PipeOrSpecialExpression: PipeOrSpecialExpression{
													SequenceExpression: SequenceExpression{
														UnaryExpression: UnaryExpression{
															ExponentiationExpression: ExponentiationExpression{
																SubsetExpression: SubsetExpression{
																	ScopeExpression: ScopeExpression{
																		IndexOrCallExpression: IndexOrCallExpression{
																			SimpleExpression: &SimpleExpression{
																				Identifier: &tk[0],
																				Tokens:     tk[:1],
																			},
																			Tokens: tk[:1],
																		},
																		Tokens: tk[:1],
																	},
																	Tokens: tk[:1],
																},
																Tokens: tk[:1],
															},
															Tokens: tk[:1],
														},
														Tokens: tk[:1],
													},
													Tokens: tk[:1],
												},
												Tokens: tk[:1],
											},
											Tokens: tk[:1],
										},
										Tokens: tk[:1],
									},
									Tokens: tk[:1],
								},
								Tokens: tk[:1],
							},
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				QueryExpression: &QueryExpression{
					AssignmentExpression: &AssignmentExpression{
						FormulaeExpression: FormulaeExpression{
							OrExpression: &OrExpression{
								AndExpression: AndExpression{
									NotExpression: NotExpression{
										RelationalExpression: RelationalExpression{
											AdditionExpression: AdditionExpression{
												MultiplicationExpression: MultiplicationExpression{
													PipeOrSpecialExpression: PipeOrSpecialExpression{
														SequenceExpression: SequenceExpression{
															UnaryExpression: UnaryExpression{
																ExponentiationExpression: ExponentiationExpression{
																	SubsetExpression: SubsetExpression{
																		ScopeExpression: ScopeExpression{
																			IndexOrCallExpression: IndexOrCallExpression{
																				SimpleExpression: &SimpleExpression{
																					Identifier: &tk[2],
																					Tokens:     tk[2:3],
																				},
																				Tokens: tk[2:3],
																			},
																			Tokens: tk[2:3],
																		},
																		Tokens: tk[2:3],
																	},
																	Tokens: tk[2:3],
																},
																Tokens: tk[2:3],
															},
															Tokens: tk[2:3],
														},
														Tokens: tk[2:3],
													},
													Tokens: tk[2:3],
												},
												Tokens: tk[2:3],
											},
											Tokens: tk[2:3],
										},
										Tokens: tk[2:3],
									},
									Tokens: tk[2:3],
								},
								Tokens: tk[2:3],
							},
							Tokens: tk[2:3],
						},
						Tokens: tk[2:3],
					},
					Tokens: tk[2:3],
				},
				Tokens: tk[:3],
			}
		}},
		{"a ? b", func(t *test, tk Tokens) { // 5
			t.Output = QueryExpression{
				AssignmentExpression: &AssignmentExpression{
					FormulaeExpression: FormulaeExpression{
						OrExpression: &OrExpression{
							AndExpression: AndExpression{
								NotExpression: NotExpression{
									RelationalExpression: RelationalExpression{
										AdditionExpression: AdditionExpression{
											MultiplicationExpression: MultiplicationExpression{
												PipeOrSpecialExpression: PipeOrSpecialExpression{
													SequenceExpression: SequenceExpression{
														UnaryExpression: UnaryExpression{
															ExponentiationExpression: ExponentiationExpression{
																SubsetExpression: SubsetExpression{
																	ScopeExpression: ScopeExpression{
																		IndexOrCallExpression: IndexOrCallExpression{
																			SimpleExpression: &SimpleExpression{
																				Identifier: &tk[0],
																				Tokens:     tk[:1],
																			},
																			Tokens: tk[:1],
																		},
																		Tokens: tk[:1],
																	},
																	Tokens: tk[:1],
																},
																Tokens: tk[:1],
															},
															Tokens: tk[:1],
														},
														Tokens: tk[:1],
													},
													Tokens: tk[:1],
												},
												Tokens: tk[:1],
											},
											Tokens: tk[:1],
										},
										Tokens: tk[:1],
									},
									Tokens: tk[:1],
								},
								Tokens: tk[:1],
							},
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				QueryExpression: &QueryExpression{
					AssignmentExpression: &AssignmentExpression{
						FormulaeExpression: FormulaeExpression{
							OrExpression: &OrExpression{
								AndExpression: AndExpression{
									NotExpression: NotExpression{
										RelationalExpression: RelationalExpression{
											AdditionExpression: AdditionExpression{
												MultiplicationExpression: MultiplicationExpression{
													PipeOrSpecialExpression: PipeOrSpecialExpression{
														SequenceExpression: SequenceExpression{
															UnaryExpression: UnaryExpression{
																ExponentiationExpression: ExponentiationExpression{
																	SubsetExpression: SubsetExpression{
																		ScopeExpression: ScopeExpression{
																			IndexOrCallExpression: IndexOrCallExpression{
																				SimpleExpression: &SimpleExpression{
																					Identifier: &tk[4],
																					Tokens:     tk[4:5],
																				},
																				Tokens: tk[4:5],
																			},
																			Tokens: tk[4:5],
																		},
																		Tokens: tk[4:5],
																	},
																	Tokens: tk[4:5],
																},
																Tokens: tk[4:5],
															},
															Tokens: tk[4:5],
														},
														Tokens: tk[4:5],
													},
													Tokens: tk[4:5],
												},
												Tokens: tk[4:5],
											},
											Tokens: tk[4:5],
										},
										Tokens: tk[4:5],
									},
									Tokens: tk[4:5],
								},
								Tokens: tk[4:5],
							},
							Tokens: tk[4:5],
						},
						Tokens: tk[4:5],
					},
					Tokens: tk[4:5],
				},
				Tokens: tk[:5],
			}
		}},
		{"in", func(t *test, tk Tokens) { // 6
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: Error{
							Err: Error{
								Err: Error{
									Err: Error{
										Err: Error{
											Err: Error{
												Err: Error{
													Err: Error{
														Err: Error{
															Err: Error{
																Err: Error{
																	Err: Error{
																		Err: Error{
																			Err: Error{
																				Err:     ErrInvalidSimpleExpression,
																				Parsing: "SimpleExpression",
																				Token:   tk[0],
																			},
																			Parsing: "IndexOrCallExpression",
																			Token:   tk[0],
																		},
																		Parsing: "ScopeExpression",
																		Token:   tk[0],
																	},
																	Parsing: "SubsetExpression",
																	Token:   tk[0],
																},
																Parsing: "ExponentiationExpression",
																Token:   tk[0],
															},
															Parsing: "UnaryExpression",
															Token:   tk[0],
														},
														Parsing: "SequenceExpression",
														Token:   tk[0],
													},
													Parsing: "PipeOrSpecialExpression",
													Token:   tk[0],
												},
												Parsing: "MultiplicationExpression",
												Token:   tk[0],
											},
											Parsing: "AdditionExpression",
											Token:   tk[0],
										},
										Parsing: "RelationalExpression",
										Token:   tk[0],
									},
									Parsing: "NotExpression",
									Token:   tk[0],
								},
								Parsing: "AndExpression",
								Token:   tk[0],
							},
							Parsing: "OrExpression",
							Token:   tk[0],
						},
						Parsing: "FormulaeExpression",
						Token:   tk[0],
					},
					Parsing: "AssignmentExpression",
					Token:   tk[0],
				},
				Parsing: "QueryExpression",
				Token:   tk[0],
			}
		}},
		{"?in", func(t *test, tk Tokens) { // 7
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: Error{
							Err: Error{
								Err: Error{
									Err: Error{
										Err: Error{
											Err: Error{
												Err: Error{
													Err: Error{
														Err: Error{
															Err: Error{
																Err: Error{
																	Err: Error{
																		Err: Error{
																			Err: Error{
																				Err:     ErrInvalidSimpleExpression,
																				Parsing: "SimpleExpression",
																				Token:   tk[1],
																			},
																			Parsing: "IndexOrCallExpression",
																			Token:   tk[1],
																		},
																		Parsing: "ScopeExpression",
																		Token:   tk[1],
																	},
																	Parsing: "SubsetExpression",
																	Token:   tk[1],
																},
																Parsing: "ExponentiationExpression",
																Token:   tk[1],
															},
															Parsing: "UnaryExpression",
															Token:   tk[1],
														},
														Parsing: "SequenceExpression",
														Token:   tk[1],
													},
													Parsing: "PipeOrSpecialExpression",
													Token:   tk[1],
												},
												Parsing: "MultiplicationExpression",
												Token:   tk[1],
											},
											Parsing: "AdditionExpression",
											Token:   tk[1],
										},
										Parsing: "RelationalExpression",
										Token:   tk[1],
									},
									Parsing: "NotExpression",
									Token:   tk[1],
								},
								Parsing: "AndExpression",
								Token:   tk[1],
							},
							Parsing: "OrExpression",
							Token:   tk[1],
						},
						Parsing: "FormulaeExpression",
						Token:   tk[1],
					},
					Parsing: "AssignmentExpression",
					Token:   tk[1],
				},
				Parsing: "QueryExpression",
				Token:   tk[1],
			}
		}},
		{"a?in", func(t *test, tk Tokens) { // 8
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: Error{
							Err: Error{
								Err: Error{
									Err: Error{
										Err: Error{
											Err: Error{
												Err: Error{
													Err: Error{
														Err: Error{
															Err: Error{
																Err: Error{
																	Err: Error{
																		Err: Error{
																			Err: Error{
																				Err: Error{
																					Err:     ErrInvalidSimpleExpression,
																					Parsing: "SimpleExpression",
																					Token:   tk[2],
																				},
																				Parsing: "IndexOrCallExpression",
																				Token:   tk[2],
																			},
																			Parsing: "ScopeExpression",
																			Token:   tk[2],
																		},
																		Parsing: "SubsetExpression",
																		Token:   tk[2],
																	},
																	Parsing: "ExponentiationExpression",
																	Token:   tk[2],
																},
																Parsing: "UnaryExpression",
																Token:   tk[2],
															},
															Parsing: "SequenceExpression",
															Token:   tk[2],
														},
														Parsing: "PipeOrSpecialExpression",
														Token:   tk[2],
													},
													Parsing: "MultiplicationExpression",
													Token:   tk[2],
												},
												Parsing: "AdditionExpression",
												Token:   tk[2],
											},
											Parsing: "RelationalExpression",
											Token:   tk[2],
										},
										Parsing: "NotExpression",
										Token:   tk[2],
									},
									Parsing: "AndExpression",
									Token:   tk[2],
								},
								Parsing: "OrExpression",
								Token:   tk[2],
							},
							Parsing: "FormulaeExpression",
							Token:   tk[2],
						},
						Parsing: "AssignmentExpression",
						Token:   tk[2],
					},
					Parsing: "QueryExpression",
					Token:   tk[2],
				},
				Parsing: "QueryExpression",
				Token:   tk[2],
			}
		}},
	}, func(t *test) (Type, error) {
		var qe QueryExpression

		err := qe.parse(&t.Tokens)

		return qe, err
	})
}

func TestAssignmentExpression(t *testing.T) {
	doTests(t, []sourceFn{
		{"a", func(t *test, tk Tokens) { // 1
			t.Output = AssignmentExpression{
				FormulaeExpression: FormulaeExpression{
					OrExpression: &OrExpression{
						AndExpression: AndExpression{
							NotExpression: NotExpression{
								RelationalExpression: RelationalExpression{
									AdditionExpression: AdditionExpression{
										MultiplicationExpression: MultiplicationExpression{
											PipeOrSpecialExpression: PipeOrSpecialExpression{
												SequenceExpression: SequenceExpression{
													UnaryExpression: UnaryExpression{
														ExponentiationExpression: ExponentiationExpression{
															SubsetExpression: SubsetExpression{
																ScopeExpression: ScopeExpression{
																	IndexOrCallExpression: IndexOrCallExpression{
																		SimpleExpression: &SimpleExpression{
																			Identifier: &tk[0],
																			Tokens:     tk[:1],
																		},
																		Tokens: tk[:1],
																	},
																	Tokens: tk[:1],
																},
																Tokens: tk[:1],
															},
															Tokens: tk[:1],
														},
														Tokens: tk[:1],
													},
													Tokens: tk[:1],
												},
												Tokens: tk[:1],
											},
											Tokens: tk[:1],
										},
										Tokens: tk[:1],
									},
									Tokens: tk[:1],
								},
								Tokens: tk[:1],
							},
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				Tokens: tk[:1],
			}
		}},
		{"a=b", func(t *test, tk Tokens) { // 2
			t.Output = AssignmentExpression{
				FormulaeExpression: FormulaeExpression{
					OrExpression: &OrExpression{
						AndExpression: AndExpression{
							NotExpression: NotExpression{
								RelationalExpression: RelationalExpression{
									AdditionExpression: AdditionExpression{
										MultiplicationExpression: MultiplicationExpression{
											PipeOrSpecialExpression: PipeOrSpecialExpression{
												SequenceExpression: SequenceExpression{
													UnaryExpression: UnaryExpression{
														ExponentiationExpression: ExponentiationExpression{
															SubsetExpression: SubsetExpression{
																ScopeExpression: ScopeExpression{
																	IndexOrCallExpression: IndexOrCallExpression{
																		SimpleExpression: &SimpleExpression{
																			Identifier: &tk[0],
																			Tokens:     tk[:1],
																		},
																		Tokens: tk[:1],
																	},
																	Tokens: tk[:1],
																},
																Tokens: tk[:1],
															},
															Tokens: tk[:1],
														},
														Tokens: tk[:1],
													},
													Tokens: tk[:1],
												},
												Tokens: tk[:1],
											},
											Tokens: tk[:1],
										},
										Tokens: tk[:1],
									},
									Tokens: tk[:1],
								},
								Tokens: tk[:1],
							},
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				AssignmentType: AssignmentEquals,
				AssignmentExpression: &AssignmentExpression{
					FormulaeExpression: FormulaeExpression{
						OrExpression: &OrExpression{
							AndExpression: AndExpression{
								NotExpression: NotExpression{
									RelationalExpression: RelationalExpression{
										AdditionExpression: AdditionExpression{
											MultiplicationExpression: MultiplicationExpression{
												PipeOrSpecialExpression: PipeOrSpecialExpression{
													SequenceExpression: SequenceExpression{
														UnaryExpression: UnaryExpression{
															ExponentiationExpression: ExponentiationExpression{
																SubsetExpression: SubsetExpression{
																	ScopeExpression: ScopeExpression{
																		IndexOrCallExpression: IndexOrCallExpression{
																			SimpleExpression: &SimpleExpression{
																				Identifier: &tk[2],
																				Tokens:     tk[2:3],
																			},
																			Tokens: tk[2:3],
																		},
																		Tokens: tk[2:3],
																	},
																	Tokens: tk[2:3],
																},
																Tokens: tk[2:3],
															},
															Tokens: tk[2:3],
														},
														Tokens: tk[2:3],
													},
													Tokens: tk[2:3],
												},
												Tokens: tk[2:3],
											},
											Tokens: tk[2:3],
										},
										Tokens: tk[2:3],
									},
									Tokens: tk[2:3],
								},
								Tokens: tk[2:3],
							},
							Tokens: tk[2:3],
						},
						Tokens: tk[2:3],
					},
					Tokens: tk[2:3],
				},
				Tokens: tk[:3],
			}
		}},
		{"a = b", func(t *test, tk Tokens) { // 3
			t.Output = AssignmentExpression{
				FormulaeExpression: FormulaeExpression{
					OrExpression: &OrExpression{
						AndExpression: AndExpression{
							NotExpression: NotExpression{
								RelationalExpression: RelationalExpression{
									AdditionExpression: AdditionExpression{
										MultiplicationExpression: MultiplicationExpression{
											PipeOrSpecialExpression: PipeOrSpecialExpression{
												SequenceExpression: SequenceExpression{
													UnaryExpression: UnaryExpression{
														ExponentiationExpression: ExponentiationExpression{
															SubsetExpression: SubsetExpression{
																ScopeExpression: ScopeExpression{
																	IndexOrCallExpression: IndexOrCallExpression{
																		SimpleExpression: &SimpleExpression{
																			Identifier: &tk[0],
																			Tokens:     tk[:1],
																		},
																		Tokens: tk[:1],
																	},
																	Tokens: tk[:1],
																},
																Tokens: tk[:1],
															},
															Tokens: tk[:1],
														},
														Tokens: tk[:1],
													},
													Tokens: tk[:1],
												},
												Tokens: tk[:1],
											},
											Tokens: tk[:1],
										},
										Tokens: tk[:1],
									},
									Tokens: tk[:1],
								},
								Tokens: tk[:1],
							},
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				AssignmentType: AssignmentEquals,
				AssignmentExpression: &AssignmentExpression{
					FormulaeExpression: FormulaeExpression{
						OrExpression: &OrExpression{
							AndExpression: AndExpression{
								NotExpression: NotExpression{
									RelationalExpression: RelationalExpression{
										AdditionExpression: AdditionExpression{
											MultiplicationExpression: MultiplicationExpression{
												PipeOrSpecialExpression: PipeOrSpecialExpression{
													SequenceExpression: SequenceExpression{
														UnaryExpression: UnaryExpression{
															ExponentiationExpression: ExponentiationExpression{
																SubsetExpression: SubsetExpression{
																	ScopeExpression: ScopeExpression{
																		IndexOrCallExpression: IndexOrCallExpression{
																			SimpleExpression: &SimpleExpression{
																				Identifier: &tk[4],
																				Tokens:     tk[4:5],
																			},
																			Tokens: tk[4:5],
																		},
																		Tokens: tk[4:5],
																	},
																	Tokens: tk[4:5],
																},
																Tokens: tk[4:5],
															},
															Tokens: tk[4:5],
														},
														Tokens: tk[4:5],
													},
													Tokens: tk[4:5],
												},
												Tokens: tk[4:5],
											},
											Tokens: tk[4:5],
										},
										Tokens: tk[4:5],
									},
									Tokens: tk[4:5],
								},
								Tokens: tk[4:5],
							},
							Tokens: tk[4:5],
						},
						Tokens: tk[4:5],
					},
					Tokens: tk[4:5],
				},
				Tokens: tk[:5],
			}
		}},
		{"a<-b", func(t *test, tk Tokens) { // 4
			t.Output = AssignmentExpression{
				FormulaeExpression: FormulaeExpression{
					OrExpression: &OrExpression{
						AndExpression: AndExpression{
							NotExpression: NotExpression{
								RelationalExpression: RelationalExpression{
									AdditionExpression: AdditionExpression{
										MultiplicationExpression: MultiplicationExpression{
											PipeOrSpecialExpression: PipeOrSpecialExpression{
												SequenceExpression: SequenceExpression{
													UnaryExpression: UnaryExpression{
														ExponentiationExpression: ExponentiationExpression{
															SubsetExpression: SubsetExpression{
																ScopeExpression: ScopeExpression{
																	IndexOrCallExpression: IndexOrCallExpression{
																		SimpleExpression: &SimpleExpression{
																			Identifier: &tk[0],
																			Tokens:     tk[:1],
																		},
																		Tokens: tk[:1],
																	},
																	Tokens: tk[:1],
																},
																Tokens: tk[:1],
															},
															Tokens: tk[:1],
														},
														Tokens: tk[:1],
													},
													Tokens: tk[:1],
												},
												Tokens: tk[:1],
											},
											Tokens: tk[:1],
										},
										Tokens: tk[:1],
									},
									Tokens: tk[:1],
								},
								Tokens: tk[:1],
							},
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				AssignmentType: AssignmentLeftAssign,
				AssignmentExpression: &AssignmentExpression{
					FormulaeExpression: FormulaeExpression{
						OrExpression: &OrExpression{
							AndExpression: AndExpression{
								NotExpression: NotExpression{
									RelationalExpression: RelationalExpression{
										AdditionExpression: AdditionExpression{
											MultiplicationExpression: MultiplicationExpression{
												PipeOrSpecialExpression: PipeOrSpecialExpression{
													SequenceExpression: SequenceExpression{
														UnaryExpression: UnaryExpression{
															ExponentiationExpression: ExponentiationExpression{
																SubsetExpression: SubsetExpression{
																	ScopeExpression: ScopeExpression{
																		IndexOrCallExpression: IndexOrCallExpression{
																			SimpleExpression: &SimpleExpression{
																				Identifier: &tk[2],
																				Tokens:     tk[2:3],
																			},
																			Tokens: tk[2:3],
																		},
																		Tokens: tk[2:3],
																	},
																	Tokens: tk[2:3],
																},
																Tokens: tk[2:3],
															},
															Tokens: tk[2:3],
														},
														Tokens: tk[2:3],
													},
													Tokens: tk[2:3],
												},
												Tokens: tk[2:3],
											},
											Tokens: tk[2:3],
										},
										Tokens: tk[2:3],
									},
									Tokens: tk[2:3],
								},
								Tokens: tk[2:3],
							},
							Tokens: tk[2:3],
						},
						Tokens: tk[2:3],
					},
					Tokens: tk[2:3],
				},
				Tokens: tk[:3],
			}
		}},
		{"a <- b", func(t *test, tk Tokens) { // 5
			t.Output = AssignmentExpression{
				FormulaeExpression: FormulaeExpression{
					OrExpression: &OrExpression{
						AndExpression: AndExpression{
							NotExpression: NotExpression{
								RelationalExpression: RelationalExpression{
									AdditionExpression: AdditionExpression{
										MultiplicationExpression: MultiplicationExpression{
											PipeOrSpecialExpression: PipeOrSpecialExpression{
												SequenceExpression: SequenceExpression{
													UnaryExpression: UnaryExpression{
														ExponentiationExpression: ExponentiationExpression{
															SubsetExpression: SubsetExpression{
																ScopeExpression: ScopeExpression{
																	IndexOrCallExpression: IndexOrCallExpression{
																		SimpleExpression: &SimpleExpression{
																			Identifier: &tk[0],
																			Tokens:     tk[:1],
																		},
																		Tokens: tk[:1],
																	},
																	Tokens: tk[:1],
																},
																Tokens: tk[:1],
															},
															Tokens: tk[:1],
														},
														Tokens: tk[:1],
													},
													Tokens: tk[:1],
												},
												Tokens: tk[:1],
											},
											Tokens: tk[:1],
										},
										Tokens: tk[:1],
									},
									Tokens: tk[:1],
								},
								Tokens: tk[:1],
							},
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				AssignmentType: AssignmentLeftAssign,
				AssignmentExpression: &AssignmentExpression{
					FormulaeExpression: FormulaeExpression{
						OrExpression: &OrExpression{
							AndExpression: AndExpression{
								NotExpression: NotExpression{
									RelationalExpression: RelationalExpression{
										AdditionExpression: AdditionExpression{
											MultiplicationExpression: MultiplicationExpression{
												PipeOrSpecialExpression: PipeOrSpecialExpression{
													SequenceExpression: SequenceExpression{
														UnaryExpression: UnaryExpression{
															ExponentiationExpression: ExponentiationExpression{
																SubsetExpression: SubsetExpression{
																	ScopeExpression: ScopeExpression{
																		IndexOrCallExpression: IndexOrCallExpression{
																			SimpleExpression: &SimpleExpression{
																				Identifier: &tk[4],
																				Tokens:     tk[4:5],
																			},
																			Tokens: tk[4:5],
																		},
																		Tokens: tk[4:5],
																	},
																	Tokens: tk[4:5],
																},
																Tokens: tk[4:5],
															},
															Tokens: tk[4:5],
														},
														Tokens: tk[4:5],
													},
													Tokens: tk[4:5],
												},
												Tokens: tk[4:5],
											},
											Tokens: tk[4:5],
										},
										Tokens: tk[4:5],
									},
									Tokens: tk[4:5],
								},
								Tokens: tk[4:5],
							},
							Tokens: tk[4:5],
						},
						Tokens: tk[4:5],
					},
					Tokens: tk[4:5],
				},
				Tokens: tk[:5],
			}
		}},
		{"a->b", func(t *test, tk Tokens) { // 6
			t.Output = AssignmentExpression{
				FormulaeExpression: FormulaeExpression{
					OrExpression: &OrExpression{
						AndExpression: AndExpression{
							NotExpression: NotExpression{
								RelationalExpression: RelationalExpression{
									AdditionExpression: AdditionExpression{
										MultiplicationExpression: MultiplicationExpression{
											PipeOrSpecialExpression: PipeOrSpecialExpression{
												SequenceExpression: SequenceExpression{
													UnaryExpression: UnaryExpression{
														ExponentiationExpression: ExponentiationExpression{
															SubsetExpression: SubsetExpression{
																ScopeExpression: ScopeExpression{
																	IndexOrCallExpression: IndexOrCallExpression{
																		SimpleExpression: &SimpleExpression{
																			Identifier: &tk[0],
																			Tokens:     tk[:1],
																		},
																		Tokens: tk[:1],
																	},
																	Tokens: tk[:1],
																},
																Tokens: tk[:1],
															},
															Tokens: tk[:1],
														},
														Tokens: tk[:1],
													},
													Tokens: tk[:1],
												},
												Tokens: tk[:1],
											},
											Tokens: tk[:1],
										},
										Tokens: tk[:1],
									},
									Tokens: tk[:1],
								},
								Tokens: tk[:1],
							},
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				AssignmentType: AssignmentRightAssign,
				AssignmentExpression: &AssignmentExpression{
					FormulaeExpression: FormulaeExpression{
						OrExpression: &OrExpression{
							AndExpression: AndExpression{
								NotExpression: NotExpression{
									RelationalExpression: RelationalExpression{
										AdditionExpression: AdditionExpression{
											MultiplicationExpression: MultiplicationExpression{
												PipeOrSpecialExpression: PipeOrSpecialExpression{
													SequenceExpression: SequenceExpression{
														UnaryExpression: UnaryExpression{
															ExponentiationExpression: ExponentiationExpression{
																SubsetExpression: SubsetExpression{
																	ScopeExpression: ScopeExpression{
																		IndexOrCallExpression: IndexOrCallExpression{
																			SimpleExpression: &SimpleExpression{
																				Identifier: &tk[2],
																				Tokens:     tk[2:3],
																			},
																			Tokens: tk[2:3],
																		},
																		Tokens: tk[2:3],
																	},
																	Tokens: tk[2:3],
																},
																Tokens: tk[2:3],
															},
															Tokens: tk[2:3],
														},
														Tokens: tk[2:3],
													},
													Tokens: tk[2:3],
												},
												Tokens: tk[2:3],
											},
											Tokens: tk[2:3],
										},
										Tokens: tk[2:3],
									},
									Tokens: tk[2:3],
								},
								Tokens: tk[2:3],
							},
							Tokens: tk[2:3],
						},
						Tokens: tk[2:3],
					},
					Tokens: tk[2:3],
				},
				Tokens: tk[:3],
			}
		}},
		{"a -> b", func(t *test, tk Tokens) { // 7
			t.Output = AssignmentExpression{
				FormulaeExpression: FormulaeExpression{
					OrExpression: &OrExpression{
						AndExpression: AndExpression{
							NotExpression: NotExpression{
								RelationalExpression: RelationalExpression{
									AdditionExpression: AdditionExpression{
										MultiplicationExpression: MultiplicationExpression{
											PipeOrSpecialExpression: PipeOrSpecialExpression{
												SequenceExpression: SequenceExpression{
													UnaryExpression: UnaryExpression{
														ExponentiationExpression: ExponentiationExpression{
															SubsetExpression: SubsetExpression{
																ScopeExpression: ScopeExpression{
																	IndexOrCallExpression: IndexOrCallExpression{
																		SimpleExpression: &SimpleExpression{
																			Identifier: &tk[0],
																			Tokens:     tk[:1],
																		},
																		Tokens: tk[:1],
																	},
																	Tokens: tk[:1],
																},
																Tokens: tk[:1],
															},
															Tokens: tk[:1],
														},
														Tokens: tk[:1],
													},
													Tokens: tk[:1],
												},
												Tokens: tk[:1],
											},
											Tokens: tk[:1],
										},
										Tokens: tk[:1],
									},
									Tokens: tk[:1],
								},
								Tokens: tk[:1],
							},
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				AssignmentType: AssignmentRightAssign,
				AssignmentExpression: &AssignmentExpression{
					FormulaeExpression: FormulaeExpression{
						OrExpression: &OrExpression{
							AndExpression: AndExpression{
								NotExpression: NotExpression{
									RelationalExpression: RelationalExpression{
										AdditionExpression: AdditionExpression{
											MultiplicationExpression: MultiplicationExpression{
												PipeOrSpecialExpression: PipeOrSpecialExpression{
													SequenceExpression: SequenceExpression{
														UnaryExpression: UnaryExpression{
															ExponentiationExpression: ExponentiationExpression{
																SubsetExpression: SubsetExpression{
																	ScopeExpression: ScopeExpression{
																		IndexOrCallExpression: IndexOrCallExpression{
																			SimpleExpression: &SimpleExpression{
																				Identifier: &tk[4],
																				Tokens:     tk[4:5],
																			},
																			Tokens: tk[4:5],
																		},
																		Tokens: tk[4:5],
																	},
																	Tokens: tk[4:5],
																},
																Tokens: tk[4:5],
															},
															Tokens: tk[4:5],
														},
														Tokens: tk[4:5],
													},
													Tokens: tk[4:5],
												},
												Tokens: tk[4:5],
											},
											Tokens: tk[4:5],
										},
										Tokens: tk[4:5],
									},
									Tokens: tk[4:5],
								},
								Tokens: tk[4:5],
							},
							Tokens: tk[4:5],
						},
						Tokens: tk[4:5],
					},
					Tokens: tk[4:5],
				},
				Tokens: tk[:5],
			}
		}},
		{"a<<-b", func(t *test, tk Tokens) { // 8
			t.Output = AssignmentExpression{
				FormulaeExpression: FormulaeExpression{
					OrExpression: &OrExpression{
						AndExpression: AndExpression{
							NotExpression: NotExpression{
								RelationalExpression: RelationalExpression{
									AdditionExpression: AdditionExpression{
										MultiplicationExpression: MultiplicationExpression{
											PipeOrSpecialExpression: PipeOrSpecialExpression{
												SequenceExpression: SequenceExpression{
													UnaryExpression: UnaryExpression{
														ExponentiationExpression: ExponentiationExpression{
															SubsetExpression: SubsetExpression{
																ScopeExpression: ScopeExpression{
																	IndexOrCallExpression: IndexOrCallExpression{
																		SimpleExpression: &SimpleExpression{
																			Identifier: &tk[0],
																			Tokens:     tk[:1],
																		},
																		Tokens: tk[:1],
																	},
																	Tokens: tk[:1],
																},
																Tokens: tk[:1],
															},
															Tokens: tk[:1],
														},
														Tokens: tk[:1],
													},
													Tokens: tk[:1],
												},
												Tokens: tk[:1],
											},
											Tokens: tk[:1],
										},
										Tokens: tk[:1],
									},
									Tokens: tk[:1],
								},
								Tokens: tk[:1],
							},
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				AssignmentType: AssignmentLeftParentAssign,
				AssignmentExpression: &AssignmentExpression{
					FormulaeExpression: FormulaeExpression{
						OrExpression: &OrExpression{
							AndExpression: AndExpression{
								NotExpression: NotExpression{
									RelationalExpression: RelationalExpression{
										AdditionExpression: AdditionExpression{
											MultiplicationExpression: MultiplicationExpression{
												PipeOrSpecialExpression: PipeOrSpecialExpression{
													SequenceExpression: SequenceExpression{
														UnaryExpression: UnaryExpression{
															ExponentiationExpression: ExponentiationExpression{
																SubsetExpression: SubsetExpression{
																	ScopeExpression: ScopeExpression{
																		IndexOrCallExpression: IndexOrCallExpression{
																			SimpleExpression: &SimpleExpression{
																				Identifier: &tk[2],
																				Tokens:     tk[2:3],
																			},
																			Tokens: tk[2:3],
																		},
																		Tokens: tk[2:3],
																	},
																	Tokens: tk[2:3],
																},
																Tokens: tk[2:3],
															},
															Tokens: tk[2:3],
														},
														Tokens: tk[2:3],
													},
													Tokens: tk[2:3],
												},
												Tokens: tk[2:3],
											},
											Tokens: tk[2:3],
										},
										Tokens: tk[2:3],
									},
									Tokens: tk[2:3],
								},
								Tokens: tk[2:3],
							},
							Tokens: tk[2:3],
						},
						Tokens: tk[2:3],
					},
					Tokens: tk[2:3],
				},
				Tokens: tk[:3],
			}
		}},
		{"a <<- b", func(t *test, tk Tokens) { // 9
			t.Output = AssignmentExpression{
				FormulaeExpression: FormulaeExpression{
					OrExpression: &OrExpression{
						AndExpression: AndExpression{
							NotExpression: NotExpression{
								RelationalExpression: RelationalExpression{
									AdditionExpression: AdditionExpression{
										MultiplicationExpression: MultiplicationExpression{
											PipeOrSpecialExpression: PipeOrSpecialExpression{
												SequenceExpression: SequenceExpression{
													UnaryExpression: UnaryExpression{
														ExponentiationExpression: ExponentiationExpression{
															SubsetExpression: SubsetExpression{
																ScopeExpression: ScopeExpression{
																	IndexOrCallExpression: IndexOrCallExpression{
																		SimpleExpression: &SimpleExpression{
																			Identifier: &tk[0],
																			Tokens:     tk[:1],
																		},
																		Tokens: tk[:1],
																	},
																	Tokens: tk[:1],
																},
																Tokens: tk[:1],
															},
															Tokens: tk[:1],
														},
														Tokens: tk[:1],
													},
													Tokens: tk[:1],
												},
												Tokens: tk[:1],
											},
											Tokens: tk[:1],
										},
										Tokens: tk[:1],
									},
									Tokens: tk[:1],
								},
								Tokens: tk[:1],
							},
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				AssignmentType: AssignmentLeftParentAssign,
				AssignmentExpression: &AssignmentExpression{
					FormulaeExpression: FormulaeExpression{
						OrExpression: &OrExpression{
							AndExpression: AndExpression{
								NotExpression: NotExpression{
									RelationalExpression: RelationalExpression{
										AdditionExpression: AdditionExpression{
											MultiplicationExpression: MultiplicationExpression{
												PipeOrSpecialExpression: PipeOrSpecialExpression{
													SequenceExpression: SequenceExpression{
														UnaryExpression: UnaryExpression{
															ExponentiationExpression: ExponentiationExpression{
																SubsetExpression: SubsetExpression{
																	ScopeExpression: ScopeExpression{
																		IndexOrCallExpression: IndexOrCallExpression{
																			SimpleExpression: &SimpleExpression{
																				Identifier: &tk[4],
																				Tokens:     tk[4:5],
																			},
																			Tokens: tk[4:5],
																		},
																		Tokens: tk[4:5],
																	},
																	Tokens: tk[4:5],
																},
																Tokens: tk[4:5],
															},
															Tokens: tk[4:5],
														},
														Tokens: tk[4:5],
													},
													Tokens: tk[4:5],
												},
												Tokens: tk[4:5],
											},
											Tokens: tk[4:5],
										},
										Tokens: tk[4:5],
									},
									Tokens: tk[4:5],
								},
								Tokens: tk[4:5],
							},
							Tokens: tk[4:5],
						},
						Tokens: tk[4:5],
					},
					Tokens: tk[4:5],
				},
				Tokens: tk[:5],
			}
		}},
		{"a->>b", func(t *test, tk Tokens) { // 10
			t.Output = AssignmentExpression{
				FormulaeExpression: FormulaeExpression{
					OrExpression: &OrExpression{
						AndExpression: AndExpression{
							NotExpression: NotExpression{
								RelationalExpression: RelationalExpression{
									AdditionExpression: AdditionExpression{
										MultiplicationExpression: MultiplicationExpression{
											PipeOrSpecialExpression: PipeOrSpecialExpression{
												SequenceExpression: SequenceExpression{
													UnaryExpression: UnaryExpression{
														ExponentiationExpression: ExponentiationExpression{
															SubsetExpression: SubsetExpression{
																ScopeExpression: ScopeExpression{
																	IndexOrCallExpression: IndexOrCallExpression{
																		SimpleExpression: &SimpleExpression{
																			Identifier: &tk[0],
																			Tokens:     tk[:1],
																		},
																		Tokens: tk[:1],
																	},
																	Tokens: tk[:1],
																},
																Tokens: tk[:1],
															},
															Tokens: tk[:1],
														},
														Tokens: tk[:1],
													},
													Tokens: tk[:1],
												},
												Tokens: tk[:1],
											},
											Tokens: tk[:1],
										},
										Tokens: tk[:1],
									},
									Tokens: tk[:1],
								},
								Tokens: tk[:1],
							},
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				AssignmentType: AssignmentRightParentAssign,
				AssignmentExpression: &AssignmentExpression{
					FormulaeExpression: FormulaeExpression{
						OrExpression: &OrExpression{
							AndExpression: AndExpression{
								NotExpression: NotExpression{
									RelationalExpression: RelationalExpression{
										AdditionExpression: AdditionExpression{
											MultiplicationExpression: MultiplicationExpression{
												PipeOrSpecialExpression: PipeOrSpecialExpression{
													SequenceExpression: SequenceExpression{
														UnaryExpression: UnaryExpression{
															ExponentiationExpression: ExponentiationExpression{
																SubsetExpression: SubsetExpression{
																	ScopeExpression: ScopeExpression{
																		IndexOrCallExpression: IndexOrCallExpression{
																			SimpleExpression: &SimpleExpression{
																				Identifier: &tk[2],
																				Tokens:     tk[2:3],
																			},
																			Tokens: tk[2:3],
																		},
																		Tokens: tk[2:3],
																	},
																	Tokens: tk[2:3],
																},
																Tokens: tk[2:3],
															},
															Tokens: tk[2:3],
														},
														Tokens: tk[2:3],
													},
													Tokens: tk[2:3],
												},
												Tokens: tk[2:3],
											},
											Tokens: tk[2:3],
										},
										Tokens: tk[2:3],
									},
									Tokens: tk[2:3],
								},
								Tokens: tk[2:3],
							},
							Tokens: tk[2:3],
						},
						Tokens: tk[2:3],
					},
					Tokens: tk[2:3],
				},
				Tokens: tk[:3],
			}
		}},
		{"a ->> b", func(t *test, tk Tokens) { // 11
			t.Output = AssignmentExpression{
				FormulaeExpression: FormulaeExpression{
					OrExpression: &OrExpression{
						AndExpression: AndExpression{
							NotExpression: NotExpression{
								RelationalExpression: RelationalExpression{
									AdditionExpression: AdditionExpression{
										MultiplicationExpression: MultiplicationExpression{
											PipeOrSpecialExpression: PipeOrSpecialExpression{
												SequenceExpression: SequenceExpression{
													UnaryExpression: UnaryExpression{
														ExponentiationExpression: ExponentiationExpression{
															SubsetExpression: SubsetExpression{
																ScopeExpression: ScopeExpression{
																	IndexOrCallExpression: IndexOrCallExpression{
																		SimpleExpression: &SimpleExpression{
																			Identifier: &tk[0],
																			Tokens:     tk[:1],
																		},
																		Tokens: tk[:1],
																	},
																	Tokens: tk[:1],
																},
																Tokens: tk[:1],
															},
															Tokens: tk[:1],
														},
														Tokens: tk[:1],
													},
													Tokens: tk[:1],
												},
												Tokens: tk[:1],
											},
											Tokens: tk[:1],
										},
										Tokens: tk[:1],
									},
									Tokens: tk[:1],
								},
								Tokens: tk[:1],
							},
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				AssignmentType: AssignmentRightParentAssign,
				AssignmentExpression: &AssignmentExpression{
					FormulaeExpression: FormulaeExpression{
						OrExpression: &OrExpression{
							AndExpression: AndExpression{
								NotExpression: NotExpression{
									RelationalExpression: RelationalExpression{
										AdditionExpression: AdditionExpression{
											MultiplicationExpression: MultiplicationExpression{
												PipeOrSpecialExpression: PipeOrSpecialExpression{
													SequenceExpression: SequenceExpression{
														UnaryExpression: UnaryExpression{
															ExponentiationExpression: ExponentiationExpression{
																SubsetExpression: SubsetExpression{
																	ScopeExpression: ScopeExpression{
																		IndexOrCallExpression: IndexOrCallExpression{
																			SimpleExpression: &SimpleExpression{
																				Identifier: &tk[4],
																				Tokens:     tk[4:5],
																			},
																			Tokens: tk[4:5],
																		},
																		Tokens: tk[4:5],
																	},
																	Tokens: tk[4:5],
																},
																Tokens: tk[4:5],
															},
															Tokens: tk[4:5],
														},
														Tokens: tk[4:5],
													},
													Tokens: tk[4:5],
												},
												Tokens: tk[4:5],
											},
											Tokens: tk[4:5],
										},
										Tokens: tk[4:5],
									},
									Tokens: tk[4:5],
								},
								Tokens: tk[4:5],
							},
							Tokens: tk[4:5],
						},
						Tokens: tk[4:5],
					},
					Tokens: tk[4:5],
				},
				Tokens: tk[:5],
			}
		}},
		{"in", func(t *test, tk Tokens) { // 12
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: Error{
							Err: Error{
								Err: Error{
									Err: Error{
										Err: Error{
											Err: Error{
												Err: Error{
													Err: Error{
														Err: Error{
															Err: Error{
																Err: Error{
																	Err: Error{
																		Err: Error{
																			Err:     ErrInvalidSimpleExpression,
																			Parsing: "SimpleExpression",
																			Token:   tk[0],
																		},
																		Parsing: "IndexOrCallExpression",
																		Token:   tk[0],
																	},
																	Parsing: "ScopeExpression",
																	Token:   tk[0],
																},
																Parsing: "SubsetExpression",
																Token:   tk[0],
															},
															Parsing: "ExponentiationExpression",
															Token:   tk[0],
														},
														Parsing: "UnaryExpression",
														Token:   tk[0],
													},
													Parsing: "SequenceExpression",
													Token:   tk[0],
												},
												Parsing: "PipeOrSpecialExpression",
												Token:   tk[0],
											},
											Parsing: "MultiplicationExpression",
											Token:   tk[0],
										},
										Parsing: "AdditionExpression",
										Token:   tk[0],
									},
									Parsing: "RelationalExpression",
									Token:   tk[0],
								},
								Parsing: "NotExpression",
								Token:   tk[0],
							},
							Parsing: "AndExpression",
							Token:   tk[0],
						},
						Parsing: "OrExpression",
						Token:   tk[0],
					},
					Parsing: "FormulaeExpression",
					Token:   tk[0],
				},
				Parsing: "AssignmentExpression",
				Token:   tk[0],
			}
		}},
		{"a=in", func(t *test, tk Tokens) { // 13
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: Error{
							Err: Error{
								Err: Error{
									Err: Error{
										Err: Error{
											Err: Error{
												Err: Error{
													Err: Error{
														Err: Error{
															Err: Error{
																Err: Error{
																	Err: Error{
																		Err: Error{
																			Err: Error{
																				Err:     ErrInvalidSimpleExpression,
																				Parsing: "SimpleExpression",
																				Token:   tk[2],
																			},
																			Parsing: "IndexOrCallExpression",
																			Token:   tk[2],
																		},
																		Parsing: "ScopeExpression",
																		Token:   tk[2],
																	},
																	Parsing: "SubsetExpression",
																	Token:   tk[2],
																},
																Parsing: "ExponentiationExpression",
																Token:   tk[2],
															},
															Parsing: "UnaryExpression",
															Token:   tk[2],
														},
														Parsing: "SequenceExpression",
														Token:   tk[2],
													},
													Parsing: "PipeOrSpecialExpression",
													Token:   tk[2],
												},
												Parsing: "MultiplicationExpression",
												Token:   tk[2],
											},
											Parsing: "AdditionExpression",
											Token:   tk[2],
										},
										Parsing: "RelationalExpression",
										Token:   tk[2],
									},
									Parsing: "NotExpression",
									Token:   tk[2],
								},
								Parsing: "AndExpression",
								Token:   tk[2],
							},
							Parsing: "OrExpression",
							Token:   tk[2],
						},
						Parsing: "FormulaeExpression",
						Token:   tk[2],
					},
					Parsing: "AssignmentExpression",
					Token:   tk[2],
				},
				Parsing: "AssignmentExpression",
				Token:   tk[2],
			}
		}},
	}, func(t *test) (Type, error) {
		var ae AssignmentExpression

		err := ae.parse(&t.Tokens)

		return ae, err
	})
}

func TestFormulaeExpression(t *testing.T) {
	doTests(t, []sourceFn{
		{"a", func(t *test, tk Tokens) { // 1
			t.Output = FormulaeExpression{
				OrExpression: &OrExpression{
					AndExpression: AndExpression{
						NotExpression: NotExpression{
							RelationalExpression: RelationalExpression{
								AdditionExpression: AdditionExpression{
									MultiplicationExpression: MultiplicationExpression{
										PipeOrSpecialExpression: PipeOrSpecialExpression{
											SequenceExpression: SequenceExpression{
												UnaryExpression: UnaryExpression{
													ExponentiationExpression: ExponentiationExpression{
														SubsetExpression: SubsetExpression{
															ScopeExpression: ScopeExpression{
																IndexOrCallExpression: IndexOrCallExpression{
																	SimpleExpression: &SimpleExpression{
																		Identifier: &tk[0],
																		Tokens:     tk[:1],
																	},
																	Tokens: tk[:1],
																},
																Tokens: tk[:1],
															},
															Tokens: tk[:1],
														},
														Tokens: tk[:1],
													},
													Tokens: tk[:1],
												},
												Tokens: tk[:1],
											},
											Tokens: tk[:1],
										},
										Tokens: tk[:1],
									},
									Tokens: tk[:1],
								},
								Tokens: tk[:1],
							},
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				Tokens: tk[:1],
			}
		}},
		{"~a", func(t *test, tk Tokens) { // 2
			t.Output = FormulaeExpression{
				FormulaeExpression: &FormulaeExpression{
					OrExpression: &OrExpression{
						AndExpression: AndExpression{
							NotExpression: NotExpression{
								RelationalExpression: RelationalExpression{
									AdditionExpression: AdditionExpression{
										MultiplicationExpression: MultiplicationExpression{
											PipeOrSpecialExpression: PipeOrSpecialExpression{
												SequenceExpression: SequenceExpression{
													UnaryExpression: UnaryExpression{
														ExponentiationExpression: ExponentiationExpression{
															SubsetExpression: SubsetExpression{
																ScopeExpression: ScopeExpression{
																	IndexOrCallExpression: IndexOrCallExpression{
																		SimpleExpression: &SimpleExpression{
																			Identifier: &tk[1],
																			Tokens:     tk[1:2],
																		},
																		Tokens: tk[1:2],
																	},
																	Tokens: tk[1:2],
																},
																Tokens: tk[1:2],
															},
															Tokens: tk[1:2],
														},
														Tokens: tk[1:2],
													},
													Tokens: tk[1:2],
												},
												Tokens: tk[1:2],
											},
											Tokens: tk[1:2],
										},
										Tokens: tk[1:2],
									},
									Tokens: tk[1:2],
								},
								Tokens: tk[1:2],
							},
							Tokens: tk[1:2],
						},
						Tokens: tk[1:2],
					},
					Tokens: tk[1:2],
				},
				Tokens: tk[:2],
			}
		}},
		{"~ a", func(t *test, tk Tokens) { // 3
			t.Output = FormulaeExpression{
				FormulaeExpression: &FormulaeExpression{
					OrExpression: &OrExpression{
						AndExpression: AndExpression{
							NotExpression: NotExpression{
								RelationalExpression: RelationalExpression{
									AdditionExpression: AdditionExpression{
										MultiplicationExpression: MultiplicationExpression{
											PipeOrSpecialExpression: PipeOrSpecialExpression{
												SequenceExpression: SequenceExpression{
													UnaryExpression: UnaryExpression{
														ExponentiationExpression: ExponentiationExpression{
															SubsetExpression: SubsetExpression{
																ScopeExpression: ScopeExpression{
																	IndexOrCallExpression: IndexOrCallExpression{
																		SimpleExpression: &SimpleExpression{
																			Identifier: &tk[2],
																			Tokens:     tk[2:3],
																		},
																		Tokens: tk[2:3],
																	},
																	Tokens: tk[2:3],
																},
																Tokens: tk[2:3],
															},
															Tokens: tk[2:3],
														},
														Tokens: tk[2:3],
													},
													Tokens: tk[2:3],
												},
												Tokens: tk[2:3],
											},
											Tokens: tk[2:3],
										},
										Tokens: tk[2:3],
									},
									Tokens: tk[2:3],
								},
								Tokens: tk[2:3],
							},
							Tokens: tk[2:3],
						},
						Tokens: tk[2:3],
					},
					Tokens: tk[2:3],
				},
				Tokens: tk[:3],
			}
		}},
		{"a~b", func(t *test, tk Tokens) { // 4
			t.Output = FormulaeExpression{
				OrExpression: &OrExpression{
					AndExpression: AndExpression{
						NotExpression: NotExpression{
							RelationalExpression: RelationalExpression{
								AdditionExpression: AdditionExpression{
									MultiplicationExpression: MultiplicationExpression{
										PipeOrSpecialExpression: PipeOrSpecialExpression{
											SequenceExpression: SequenceExpression{
												UnaryExpression: UnaryExpression{
													ExponentiationExpression: ExponentiationExpression{
														SubsetExpression: SubsetExpression{
															ScopeExpression: ScopeExpression{
																IndexOrCallExpression: IndexOrCallExpression{
																	SimpleExpression: &SimpleExpression{
																		Identifier: &tk[0],
																		Tokens:     tk[:1],
																	},
																	Tokens: tk[:1],
																},
																Tokens: tk[:1],
															},
															Tokens: tk[:1],
														},
														Tokens: tk[:1],
													},
													Tokens: tk[:1],
												},
												Tokens: tk[:1],
											},
											Tokens: tk[:1],
										},
										Tokens: tk[:1],
									},
									Tokens: tk[:1],
								},
								Tokens: tk[:1],
							},
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				FormulaeExpression: &FormulaeExpression{
					OrExpression: &OrExpression{
						AndExpression: AndExpression{
							NotExpression: NotExpression{
								RelationalExpression: RelationalExpression{
									AdditionExpression: AdditionExpression{
										MultiplicationExpression: MultiplicationExpression{
											PipeOrSpecialExpression: PipeOrSpecialExpression{
												SequenceExpression: SequenceExpression{
													UnaryExpression: UnaryExpression{
														ExponentiationExpression: ExponentiationExpression{
															SubsetExpression: SubsetExpression{
																ScopeExpression: ScopeExpression{
																	IndexOrCallExpression: IndexOrCallExpression{
																		SimpleExpression: &SimpleExpression{
																			Identifier: &tk[2],
																			Tokens:     tk[2:3],
																		},
																		Tokens: tk[2:3],
																	},
																	Tokens: tk[2:3],
																},
																Tokens: tk[2:3],
															},
															Tokens: tk[2:3],
														},
														Tokens: tk[2:3],
													},
													Tokens: tk[2:3],
												},
												Tokens: tk[2:3],
											},
											Tokens: tk[2:3],
										},
										Tokens: tk[2:3],
									},
									Tokens: tk[2:3],
								},
								Tokens: tk[2:3],
							},
							Tokens: tk[2:3],
						},
						Tokens: tk[2:3],
					},
					Tokens: tk[2:3],
				},
				Tokens: tk[:3],
			}
		}},
		{"in", func(t *test, tk Tokens) { // 5
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: Error{
							Err: Error{
								Err: Error{
									Err: Error{
										Err: Error{
											Err: Error{
												Err: Error{
													Err: Error{
														Err: Error{
															Err: Error{
																Err: Error{
																	Err: Error{
																		Err:     ErrInvalidSimpleExpression,
																		Parsing: "SimpleExpression",
																		Token:   tk[0],
																	},
																	Parsing: "IndexOrCallExpression",
																	Token:   tk[0],
																},
																Parsing: "ScopeExpression",
																Token:   tk[0],
															},
															Parsing: "SubsetExpression",
															Token:   tk[0],
														},
														Parsing: "ExponentiationExpression",
														Token:   tk[0],
													},
													Parsing: "UnaryExpression",
													Token:   tk[0],
												},
												Parsing: "SequenceExpression",
												Token:   tk[0],
											},
											Parsing: "PipeOrSpecialExpression",
											Token:   tk[0],
										},
										Parsing: "MultiplicationExpression",
										Token:   tk[0],
									},
									Parsing: "AdditionExpression",
									Token:   tk[0],
								},
								Parsing: "RelationalExpression",
								Token:   tk[0],
							},
							Parsing: "NotExpression",
							Token:   tk[0],
						},
						Parsing: "AndExpression",
						Token:   tk[0],
					},
					Parsing: "OrExpression",
					Token:   tk[0],
				},
				Parsing: "FormulaeExpression",
				Token:   tk[0],
			}
		}},
		{"~in", func(t *test, tk Tokens) { // 6
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: Error{
							Err: Error{
								Err: Error{
									Err: Error{
										Err: Error{
											Err: Error{
												Err: Error{
													Err: Error{
														Err: Error{
															Err: Error{
																Err: Error{
																	Err: Error{
																		Err:     ErrInvalidSimpleExpression,
																		Parsing: "SimpleExpression",
																		Token:   tk[1],
																	},
																	Parsing: "IndexOrCallExpression",
																	Token:   tk[1],
																},
																Parsing: "ScopeExpression",
																Token:   tk[1],
															},
															Parsing: "SubsetExpression",
															Token:   tk[1],
														},
														Parsing: "ExponentiationExpression",
														Token:   tk[1],
													},
													Parsing: "UnaryExpression",
													Token:   tk[1],
												},
												Parsing: "SequenceExpression",
												Token:   tk[1],
											},
											Parsing: "PipeOrSpecialExpression",
											Token:   tk[1],
										},
										Parsing: "MultiplicationExpression",
										Token:   tk[1],
									},
									Parsing: "AdditionExpression",
									Token:   tk[1],
								},
								Parsing: "RelationalExpression",
								Token:   tk[1],
							},
							Parsing: "NotExpression",
							Token:   tk[1],
						},
						Parsing: "AndExpression",
						Token:   tk[1],
					},
					Parsing: "OrExpression",
					Token:   tk[1],
				},
				Parsing: "FormulaeExpression",
				Token:   tk[1],
			}
		}},
	}, func(t *test) (Type, error) {
		var fe FormulaeExpression

		err := fe.parse(&t.Tokens)

		return fe, err
	})
}

func TestOrExpression(t *testing.T) {
	doTests(t, []sourceFn{
		{"a", func(t *test, tk Tokens) { // 1
			t.Output = OrExpression{
				AndExpression: AndExpression{
					NotExpression: NotExpression{
						RelationalExpression: RelationalExpression{
							AdditionExpression: AdditionExpression{
								MultiplicationExpression: MultiplicationExpression{
									PipeOrSpecialExpression: PipeOrSpecialExpression{
										SequenceExpression: SequenceExpression{
											UnaryExpression: UnaryExpression{
												ExponentiationExpression: ExponentiationExpression{
													SubsetExpression: SubsetExpression{
														ScopeExpression: ScopeExpression{
															IndexOrCallExpression: IndexOrCallExpression{
																SimpleExpression: &SimpleExpression{
																	Identifier: &tk[0],
																	Tokens:     tk[:1],
																},
																Tokens: tk[:1],
															},
															Tokens: tk[:1],
														},
														Tokens: tk[:1],
													},
													Tokens: tk[:1],
												},
												Tokens: tk[:1],
											},
											Tokens: tk[:1],
										},
										Tokens: tk[:1],
									},
									Tokens: tk[:1],
								},
								Tokens: tk[:1],
							},
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				Tokens: tk[:1],
			}
		}},
		{"a|b", func(t *test, tk Tokens) { // 2
			t.Output = OrExpression{
				AndExpression: AndExpression{
					NotExpression: NotExpression{
						RelationalExpression: RelationalExpression{
							AdditionExpression: AdditionExpression{
								MultiplicationExpression: MultiplicationExpression{
									PipeOrSpecialExpression: PipeOrSpecialExpression{
										SequenceExpression: SequenceExpression{
											UnaryExpression: UnaryExpression{
												ExponentiationExpression: ExponentiationExpression{
													SubsetExpression: SubsetExpression{
														ScopeExpression: ScopeExpression{
															IndexOrCallExpression: IndexOrCallExpression{
																SimpleExpression: &SimpleExpression{
																	Identifier: &tk[0],
																	Tokens:     tk[:1],
																},
																Tokens: tk[:1],
															},
															Tokens: tk[:1],
														},
														Tokens: tk[:1],
													},
													Tokens: tk[:1],
												},
												Tokens: tk[:1],
											},
											Tokens: tk[:1],
										},
										Tokens: tk[:1],
									},
									Tokens: tk[:1],
								},
								Tokens: tk[:1],
							},
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				OrType: OrVectorized,
				OrExpression: &OrExpression{
					AndExpression: AndExpression{
						NotExpression: NotExpression{
							RelationalExpression: RelationalExpression{
								AdditionExpression: AdditionExpression{
									MultiplicationExpression: MultiplicationExpression{
										PipeOrSpecialExpression: PipeOrSpecialExpression{
											SequenceExpression: SequenceExpression{
												UnaryExpression: UnaryExpression{
													ExponentiationExpression: ExponentiationExpression{
														SubsetExpression: SubsetExpression{
															ScopeExpression: ScopeExpression{
																IndexOrCallExpression: IndexOrCallExpression{
																	SimpleExpression: &SimpleExpression{
																		Identifier: &tk[2],
																		Tokens:     tk[2:3],
																	},
																	Tokens: tk[2:3],
																},
																Tokens: tk[2:3],
															},
															Tokens: tk[2:3],
														},
														Tokens: tk[2:3],
													},
													Tokens: tk[2:3],
												},
												Tokens: tk[2:3],
											},
											Tokens: tk[2:3],
										},
										Tokens: tk[2:3],
									},
									Tokens: tk[2:3],
								},
								Tokens: tk[2:3],
							},
							Tokens: tk[2:3],
						},
						Tokens: tk[2:3],
					},
					Tokens: tk[2:3],
				},
				Tokens: tk[:3],
			}
		}},
		{"a | b", func(t *test, tk Tokens) { // 3
			t.Output = OrExpression{
				AndExpression: AndExpression{
					NotExpression: NotExpression{
						RelationalExpression: RelationalExpression{
							AdditionExpression: AdditionExpression{
								MultiplicationExpression: MultiplicationExpression{
									PipeOrSpecialExpression: PipeOrSpecialExpression{
										SequenceExpression: SequenceExpression{
											UnaryExpression: UnaryExpression{
												ExponentiationExpression: ExponentiationExpression{
													SubsetExpression: SubsetExpression{
														ScopeExpression: ScopeExpression{
															IndexOrCallExpression: IndexOrCallExpression{
																SimpleExpression: &SimpleExpression{
																	Identifier: &tk[0],
																	Tokens:     tk[:1],
																},
																Tokens: tk[:1],
															},
															Tokens: tk[:1],
														},
														Tokens: tk[:1],
													},
													Tokens: tk[:1],
												},
												Tokens: tk[:1],
											},
											Tokens: tk[:1],
										},
										Tokens: tk[:1],
									},
									Tokens: tk[:1],
								},
								Tokens: tk[:1],
							},
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				OrType: OrVectorized,
				OrExpression: &OrExpression{
					AndExpression: AndExpression{
						NotExpression: NotExpression{
							RelationalExpression: RelationalExpression{
								AdditionExpression: AdditionExpression{
									MultiplicationExpression: MultiplicationExpression{
										PipeOrSpecialExpression: PipeOrSpecialExpression{
											SequenceExpression: SequenceExpression{
												UnaryExpression: UnaryExpression{
													ExponentiationExpression: ExponentiationExpression{
														SubsetExpression: SubsetExpression{
															ScopeExpression: ScopeExpression{
																IndexOrCallExpression: IndexOrCallExpression{
																	SimpleExpression: &SimpleExpression{
																		Identifier: &tk[4],
																		Tokens:     tk[4:5],
																	},
																	Tokens: tk[4:5],
																},
																Tokens: tk[4:5],
															},
															Tokens: tk[4:5],
														},
														Tokens: tk[4:5],
													},
													Tokens: tk[4:5],
												},
												Tokens: tk[4:5],
											},
											Tokens: tk[4:5],
										},
										Tokens: tk[4:5],
									},
									Tokens: tk[4:5],
								},
								Tokens: tk[4:5],
							},
							Tokens: tk[4:5],
						},
						Tokens: tk[4:5],
					},
					Tokens: tk[4:5],
				},
				Tokens: tk[:5],
			}
		}},
		{"a||b", func(t *test, tk Tokens) { // 4
			t.Output = OrExpression{
				AndExpression: AndExpression{
					NotExpression: NotExpression{
						RelationalExpression: RelationalExpression{
							AdditionExpression: AdditionExpression{
								MultiplicationExpression: MultiplicationExpression{
									PipeOrSpecialExpression: PipeOrSpecialExpression{
										SequenceExpression: SequenceExpression{
											UnaryExpression: UnaryExpression{
												ExponentiationExpression: ExponentiationExpression{
													SubsetExpression: SubsetExpression{
														ScopeExpression: ScopeExpression{
															IndexOrCallExpression: IndexOrCallExpression{
																SimpleExpression: &SimpleExpression{
																	Identifier: &tk[0],
																	Tokens:     tk[:1],
																},
																Tokens: tk[:1],
															},
															Tokens: tk[:1],
														},
														Tokens: tk[:1],
													},
													Tokens: tk[:1],
												},
												Tokens: tk[:1],
											},
											Tokens: tk[:1],
										},
										Tokens: tk[:1],
									},
									Tokens: tk[:1],
								},
								Tokens: tk[:1],
							},
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				OrType: OrNotVectorized,
				OrExpression: &OrExpression{
					AndExpression: AndExpression{
						NotExpression: NotExpression{
							RelationalExpression: RelationalExpression{
								AdditionExpression: AdditionExpression{
									MultiplicationExpression: MultiplicationExpression{
										PipeOrSpecialExpression: PipeOrSpecialExpression{
											SequenceExpression: SequenceExpression{
												UnaryExpression: UnaryExpression{
													ExponentiationExpression: ExponentiationExpression{
														SubsetExpression: SubsetExpression{
															ScopeExpression: ScopeExpression{
																IndexOrCallExpression: IndexOrCallExpression{
																	SimpleExpression: &SimpleExpression{
																		Identifier: &tk[2],
																		Tokens:     tk[2:3],
																	},
																	Tokens: tk[2:3],
																},
																Tokens: tk[2:3],
															},
															Tokens: tk[2:3],
														},
														Tokens: tk[2:3],
													},
													Tokens: tk[2:3],
												},
												Tokens: tk[2:3],
											},
											Tokens: tk[2:3],
										},
										Tokens: tk[2:3],
									},
									Tokens: tk[2:3],
								},
								Tokens: tk[2:3],
							},
							Tokens: tk[2:3],
						},
						Tokens: tk[2:3],
					},
					Tokens: tk[2:3],
				},
				Tokens: tk[:3],
			}
		}},
		{"a || b", func(t *test, tk Tokens) { // 5
			t.Output = OrExpression{
				AndExpression: AndExpression{
					NotExpression: NotExpression{
						RelationalExpression: RelationalExpression{
							AdditionExpression: AdditionExpression{
								MultiplicationExpression: MultiplicationExpression{
									PipeOrSpecialExpression: PipeOrSpecialExpression{
										SequenceExpression: SequenceExpression{
											UnaryExpression: UnaryExpression{
												ExponentiationExpression: ExponentiationExpression{
													SubsetExpression: SubsetExpression{
														ScopeExpression: ScopeExpression{
															IndexOrCallExpression: IndexOrCallExpression{
																SimpleExpression: &SimpleExpression{
																	Identifier: &tk[0],
																	Tokens:     tk[:1],
																},
																Tokens: tk[:1],
															},
															Tokens: tk[:1],
														},
														Tokens: tk[:1],
													},
													Tokens: tk[:1],
												},
												Tokens: tk[:1],
											},
											Tokens: tk[:1],
										},
										Tokens: tk[:1],
									},
									Tokens: tk[:1],
								},
								Tokens: tk[:1],
							},
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				OrType: OrNotVectorized,
				OrExpression: &OrExpression{
					AndExpression: AndExpression{
						NotExpression: NotExpression{
							RelationalExpression: RelationalExpression{
								AdditionExpression: AdditionExpression{
									MultiplicationExpression: MultiplicationExpression{
										PipeOrSpecialExpression: PipeOrSpecialExpression{
											SequenceExpression: SequenceExpression{
												UnaryExpression: UnaryExpression{
													ExponentiationExpression: ExponentiationExpression{
														SubsetExpression: SubsetExpression{
															ScopeExpression: ScopeExpression{
																IndexOrCallExpression: IndexOrCallExpression{
																	SimpleExpression: &SimpleExpression{
																		Identifier: &tk[4],
																		Tokens:     tk[4:5],
																	},
																	Tokens: tk[4:5],
																},
																Tokens: tk[4:5],
															},
															Tokens: tk[4:5],
														},
														Tokens: tk[4:5],
													},
													Tokens: tk[4:5],
												},
												Tokens: tk[4:5],
											},
											Tokens: tk[4:5],
										},
										Tokens: tk[4:5],
									},
									Tokens: tk[4:5],
								},
								Tokens: tk[4:5],
							},
							Tokens: tk[4:5],
						},
						Tokens: tk[4:5],
					},
					Tokens: tk[4:5],
				},
				Tokens: tk[:5],
			}
		}},
		{"in", func(t *test, tk Tokens) { // 6
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: Error{
							Err: Error{
								Err: Error{
									Err: Error{
										Err: Error{
											Err: Error{
												Err: Error{
													Err: Error{
														Err: Error{
															Err: Error{
																Err: Error{
																	Err:     ErrInvalidSimpleExpression,
																	Parsing: "SimpleExpression",
																	Token:   tk[0],
																},
																Parsing: "IndexOrCallExpression",
																Token:   tk[0],
															},
															Parsing: "ScopeExpression",
															Token:   tk[0],
														},
														Parsing: "SubsetExpression",
														Token:   tk[0],
													},
													Parsing: "ExponentiationExpression",
													Token:   tk[0],
												},
												Parsing: "UnaryExpression",
												Token:   tk[0],
											},
											Parsing: "SequenceExpression",
											Token:   tk[0],
										},
										Parsing: "PipeOrSpecialExpression",
										Token:   tk[0],
									},
									Parsing: "MultiplicationExpression",
									Token:   tk[0],
								},
								Parsing: "AdditionExpression",
								Token:   tk[0],
							},
							Parsing: "RelationalExpression",
							Token:   tk[0],
						},
						Parsing: "NotExpression",
						Token:   tk[0],
					},
					Parsing: "AndExpression",
					Token:   tk[0],
				},
				Parsing: "OrExpression",
				Token:   tk[0],
			}
		}},
		{"a|in", func(t *test, tk Tokens) { // 7
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: Error{
							Err: Error{
								Err: Error{
									Err: Error{
										Err: Error{
											Err: Error{
												Err: Error{
													Err: Error{
														Err: Error{
															Err: Error{
																Err: Error{
																	Err: Error{
																		Err:     ErrInvalidSimpleExpression,
																		Parsing: "SimpleExpression",
																		Token:   tk[2],
																	},
																	Parsing: "IndexOrCallExpression",
																	Token:   tk[2],
																},
																Parsing: "ScopeExpression",
																Token:   tk[2],
															},
															Parsing: "SubsetExpression",
															Token:   tk[2],
														},
														Parsing: "ExponentiationExpression",
														Token:   tk[2],
													},
													Parsing: "UnaryExpression",
													Token:   tk[2],
												},
												Parsing: "SequenceExpression",
												Token:   tk[2],
											},
											Parsing: "PipeOrSpecialExpression",
											Token:   tk[2],
										},
										Parsing: "MultiplicationExpression",
										Token:   tk[2],
									},
									Parsing: "AdditionExpression",
									Token:   tk[2],
								},
								Parsing: "RelationalExpression",
								Token:   tk[2],
							},
							Parsing: "NotExpression",
							Token:   tk[2],
						},
						Parsing: "AndExpression",
						Token:   tk[2],
					},
					Parsing: "OrExpression",
					Token:   tk[2],
				},
				Parsing: "OrExpression",
				Token:   tk[2],
			}
		}},
	}, func(t *test) (Type, error) {
		var oe OrExpression

		err := oe.parse(&t.Tokens)

		return oe, err
	})
}

func TestAndExpression(t *testing.T) {
	doTests(t, []sourceFn{
		{"a", func(t *test, tk Tokens) { // 1
			t.Output = AndExpression{
				NotExpression: NotExpression{
					RelationalExpression: RelationalExpression{
						AdditionExpression: AdditionExpression{
							MultiplicationExpression: MultiplicationExpression{
								PipeOrSpecialExpression: PipeOrSpecialExpression{
									SequenceExpression: SequenceExpression{
										UnaryExpression: UnaryExpression{
											ExponentiationExpression: ExponentiationExpression{
												SubsetExpression: SubsetExpression{
													ScopeExpression: ScopeExpression{
														IndexOrCallExpression: IndexOrCallExpression{
															SimpleExpression: &SimpleExpression{
																Identifier: &tk[0],
																Tokens:     tk[:1],
															},
															Tokens: tk[:1],
														},
														Tokens: tk[:1],
													},
													Tokens: tk[:1],
												},
												Tokens: tk[:1],
											},
											Tokens: tk[:1],
										},
										Tokens: tk[:1],
									},
									Tokens: tk[:1],
								},
								Tokens: tk[:1],
							},
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				Tokens: tk[:1],
			}
		}},
		{"a&b", func(t *test, tk Tokens) { // 2
			t.Output = AndExpression{
				NotExpression: NotExpression{
					RelationalExpression: RelationalExpression{
						AdditionExpression: AdditionExpression{
							MultiplicationExpression: MultiplicationExpression{
								PipeOrSpecialExpression: PipeOrSpecialExpression{
									SequenceExpression: SequenceExpression{
										UnaryExpression: UnaryExpression{
											ExponentiationExpression: ExponentiationExpression{
												SubsetExpression: SubsetExpression{
													ScopeExpression: ScopeExpression{
														IndexOrCallExpression: IndexOrCallExpression{
															SimpleExpression: &SimpleExpression{
																Identifier: &tk[0],
																Tokens:     tk[:1],
															},
															Tokens: tk[:1],
														},
														Tokens: tk[:1],
													},
													Tokens: tk[:1],
												},
												Tokens: tk[:1],
											},
											Tokens: tk[:1],
										},
										Tokens: tk[:1],
									},
									Tokens: tk[:1],
								},
								Tokens: tk[:1],
							},
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				AndType: AndVectorized,
				AndExpression: &AndExpression{
					NotExpression: NotExpression{
						RelationalExpression: RelationalExpression{
							AdditionExpression: AdditionExpression{
								MultiplicationExpression: MultiplicationExpression{
									PipeOrSpecialExpression: PipeOrSpecialExpression{
										SequenceExpression: SequenceExpression{
											UnaryExpression: UnaryExpression{
												ExponentiationExpression: ExponentiationExpression{
													SubsetExpression: SubsetExpression{
														ScopeExpression: ScopeExpression{
															IndexOrCallExpression: IndexOrCallExpression{
																SimpleExpression: &SimpleExpression{
																	Identifier: &tk[2],
																	Tokens:     tk[2:3],
																},
																Tokens: tk[2:3],
															},
															Tokens: tk[2:3],
														},
														Tokens: tk[2:3],
													},
													Tokens: tk[2:3],
												},
												Tokens: tk[2:3],
											},
											Tokens: tk[2:3],
										},
										Tokens: tk[2:3],
									},
									Tokens: tk[2:3],
								},
								Tokens: tk[2:3],
							},
							Tokens: tk[2:3],
						},
						Tokens: tk[2:3],
					},
					Tokens: tk[2:3],
				},
				Tokens: tk[:3],
			}
		}},
		{"a & b", func(t *test, tk Tokens) { // 3
			t.Output = AndExpression{
				NotExpression: NotExpression{
					RelationalExpression: RelationalExpression{
						AdditionExpression: AdditionExpression{
							MultiplicationExpression: MultiplicationExpression{
								PipeOrSpecialExpression: PipeOrSpecialExpression{
									SequenceExpression: SequenceExpression{
										UnaryExpression: UnaryExpression{
											ExponentiationExpression: ExponentiationExpression{
												SubsetExpression: SubsetExpression{
													ScopeExpression: ScopeExpression{
														IndexOrCallExpression: IndexOrCallExpression{
															SimpleExpression: &SimpleExpression{
																Identifier: &tk[0],
																Tokens:     tk[:1],
															},
															Tokens: tk[:1],
														},
														Tokens: tk[:1],
													},
													Tokens: tk[:1],
												},
												Tokens: tk[:1],
											},
											Tokens: tk[:1],
										},
										Tokens: tk[:1],
									},
									Tokens: tk[:1],
								},
								Tokens: tk[:1],
							},
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				AndType: AndVectorized,
				AndExpression: &AndExpression{
					NotExpression: NotExpression{
						RelationalExpression: RelationalExpression{
							AdditionExpression: AdditionExpression{
								MultiplicationExpression: MultiplicationExpression{
									PipeOrSpecialExpression: PipeOrSpecialExpression{
										SequenceExpression: SequenceExpression{
											UnaryExpression: UnaryExpression{
												ExponentiationExpression: ExponentiationExpression{
													SubsetExpression: SubsetExpression{
														ScopeExpression: ScopeExpression{
															IndexOrCallExpression: IndexOrCallExpression{
																SimpleExpression: &SimpleExpression{
																	Identifier: &tk[4],
																	Tokens:     tk[4:5],
																},
																Tokens: tk[4:5],
															},
															Tokens: tk[4:5],
														},
														Tokens: tk[4:5],
													},
													Tokens: tk[4:5],
												},
												Tokens: tk[4:5],
											},
											Tokens: tk[4:5],
										},
										Tokens: tk[4:5],
									},
									Tokens: tk[4:5],
								},
								Tokens: tk[4:5],
							},
							Tokens: tk[4:5],
						},
						Tokens: tk[4:5],
					},
					Tokens: tk[4:5],
				},
				Tokens: tk[:5],
			}
		}},
		{"a&&b", func(t *test, tk Tokens) { // 4
			t.Output = AndExpression{
				NotExpression: NotExpression{
					RelationalExpression: RelationalExpression{
						AdditionExpression: AdditionExpression{
							MultiplicationExpression: MultiplicationExpression{
								PipeOrSpecialExpression: PipeOrSpecialExpression{
									SequenceExpression: SequenceExpression{
										UnaryExpression: UnaryExpression{
											ExponentiationExpression: ExponentiationExpression{
												SubsetExpression: SubsetExpression{
													ScopeExpression: ScopeExpression{
														IndexOrCallExpression: IndexOrCallExpression{
															SimpleExpression: &SimpleExpression{
																Identifier: &tk[0],
																Tokens:     tk[:1],
															},
															Tokens: tk[:1],
														},
														Tokens: tk[:1],
													},
													Tokens: tk[:1],
												},
												Tokens: tk[:1],
											},
											Tokens: tk[:1],
										},
										Tokens: tk[:1],
									},
									Tokens: tk[:1],
								},
								Tokens: tk[:1],
							},
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				AndType: AndNotVectorized,
				AndExpression: &AndExpression{
					NotExpression: NotExpression{
						RelationalExpression: RelationalExpression{
							AdditionExpression: AdditionExpression{
								MultiplicationExpression: MultiplicationExpression{
									PipeOrSpecialExpression: PipeOrSpecialExpression{
										SequenceExpression: SequenceExpression{
											UnaryExpression: UnaryExpression{
												ExponentiationExpression: ExponentiationExpression{
													SubsetExpression: SubsetExpression{
														ScopeExpression: ScopeExpression{
															IndexOrCallExpression: IndexOrCallExpression{
																SimpleExpression: &SimpleExpression{
																	Identifier: &tk[2],
																	Tokens:     tk[2:3],
																},
																Tokens: tk[2:3],
															},
															Tokens: tk[2:3],
														},
														Tokens: tk[2:3],
													},
													Tokens: tk[2:3],
												},
												Tokens: tk[2:3],
											},
											Tokens: tk[2:3],
										},
										Tokens: tk[2:3],
									},
									Tokens: tk[2:3],
								},
								Tokens: tk[2:3],
							},
							Tokens: tk[2:3],
						},
						Tokens: tk[2:3],
					},
					Tokens: tk[2:3],
				},
				Tokens: tk[:3],
			}
		}},
		{"a && b", func(t *test, tk Tokens) { // 5
			t.Output = AndExpression{
				NotExpression: NotExpression{
					RelationalExpression: RelationalExpression{
						AdditionExpression: AdditionExpression{
							MultiplicationExpression: MultiplicationExpression{
								PipeOrSpecialExpression: PipeOrSpecialExpression{
									SequenceExpression: SequenceExpression{
										UnaryExpression: UnaryExpression{
											ExponentiationExpression: ExponentiationExpression{
												SubsetExpression: SubsetExpression{
													ScopeExpression: ScopeExpression{
														IndexOrCallExpression: IndexOrCallExpression{
															SimpleExpression: &SimpleExpression{
																Identifier: &tk[0],
																Tokens:     tk[:1],
															},
															Tokens: tk[:1],
														},
														Tokens: tk[:1],
													},
													Tokens: tk[:1],
												},
												Tokens: tk[:1],
											},
											Tokens: tk[:1],
										},
										Tokens: tk[:1],
									},
									Tokens: tk[:1],
								},
								Tokens: tk[:1],
							},
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				AndType: AndNotVectorized,
				AndExpression: &AndExpression{
					NotExpression: NotExpression{
						RelationalExpression: RelationalExpression{
							AdditionExpression: AdditionExpression{
								MultiplicationExpression: MultiplicationExpression{
									PipeOrSpecialExpression: PipeOrSpecialExpression{
										SequenceExpression: SequenceExpression{
											UnaryExpression: UnaryExpression{
												ExponentiationExpression: ExponentiationExpression{
													SubsetExpression: SubsetExpression{
														ScopeExpression: ScopeExpression{
															IndexOrCallExpression: IndexOrCallExpression{
																SimpleExpression: &SimpleExpression{
																	Identifier: &tk[4],
																	Tokens:     tk[4:5],
																},
																Tokens: tk[4:5],
															},
															Tokens: tk[4:5],
														},
														Tokens: tk[4:5],
													},
													Tokens: tk[4:5],
												},
												Tokens: tk[4:5],
											},
											Tokens: tk[4:5],
										},
										Tokens: tk[4:5],
									},
									Tokens: tk[4:5],
								},
								Tokens: tk[4:5],
							},
							Tokens: tk[4:5],
						},
						Tokens: tk[4:5],
					},
					Tokens: tk[4:5],
				},
				Tokens: tk[:5],
			}
		}},
		{"in", func(t *test, tk Tokens) { // 6
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: Error{
							Err: Error{
								Err: Error{
									Err: Error{
										Err: Error{
											Err: Error{
												Err: Error{
													Err: Error{
														Err: Error{
															Err: Error{
																Err:     ErrInvalidSimpleExpression,
																Parsing: "SimpleExpression",
																Token:   tk[0],
															},
															Parsing: "IndexOrCallExpression",
															Token:   tk[0],
														},
														Parsing: "ScopeExpression",
														Token:   tk[0],
													},
													Parsing: "SubsetExpression",
													Token:   tk[0],
												},
												Parsing: "ExponentiationExpression",
												Token:   tk[0],
											},
											Parsing: "UnaryExpression",
											Token:   tk[0],
										},
										Parsing: "SequenceExpression",
										Token:   tk[0],
									},
									Parsing: "PipeOrSpecialExpression",
									Token:   tk[0],
								},
								Parsing: "MultiplicationExpression",
								Token:   tk[0],
							},
							Parsing: "AdditionExpression",
							Token:   tk[0],
						},
						Parsing: "RelationalExpression",
						Token:   tk[0],
					},
					Parsing: "NotExpression",
					Token:   tk[0],
				},
				Parsing: "AndExpression",
				Token:   tk[0],
			}
		}},
		{"a&in", func(t *test, tk Tokens) { // 7
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: Error{
							Err: Error{
								Err: Error{
									Err: Error{
										Err: Error{
											Err: Error{
												Err: Error{
													Err: Error{
														Err: Error{
															Err: Error{
																Err: Error{
																	Err:     ErrInvalidSimpleExpression,
																	Parsing: "SimpleExpression",
																	Token:   tk[2],
																},
																Parsing: "IndexOrCallExpression",
																Token:   tk[2],
															},
															Parsing: "ScopeExpression",
															Token:   tk[2],
														},
														Parsing: "SubsetExpression",
														Token:   tk[2],
													},
													Parsing: "ExponentiationExpression",
													Token:   tk[2],
												},
												Parsing: "UnaryExpression",
												Token:   tk[2],
											},
											Parsing: "SequenceExpression",
											Token:   tk[2],
										},
										Parsing: "PipeOrSpecialExpression",
										Token:   tk[2],
									},
									Parsing: "MultiplicationExpression",
									Token:   tk[2],
								},
								Parsing: "AdditionExpression",
								Token:   tk[2],
							},
							Parsing: "RelationalExpression",
							Token:   tk[2],
						},
						Parsing: "NotExpression",
						Token:   tk[2],
					},
					Parsing: "AndExpression",
					Token:   tk[2],
				},
				Parsing: "AndExpression",
				Token:   tk[2],
			}
		}},
	}, func(t *test) (Type, error) {
		var ae AndExpression

		err := ae.parse(&t.Tokens)

		return ae, err
	})
}

func TestNotExpression(t *testing.T) {
	doTests(t, []sourceFn{
		{"a", func(t *test, tk Tokens) { // 1
			t.Output = NotExpression{
				RelationalExpression: RelationalExpression{
					AdditionExpression: AdditionExpression{
						MultiplicationExpression: MultiplicationExpression{
							PipeOrSpecialExpression: PipeOrSpecialExpression{
								SequenceExpression: SequenceExpression{
									UnaryExpression: UnaryExpression{
										ExponentiationExpression: ExponentiationExpression{
											SubsetExpression: SubsetExpression{
												ScopeExpression: ScopeExpression{
													IndexOrCallExpression: IndexOrCallExpression{
														SimpleExpression: &SimpleExpression{
															Identifier: &tk[0],
															Tokens:     tk[:1],
														},
														Tokens: tk[:1],
													},
													Tokens: tk[:1],
												},
												Tokens: tk[:1],
											},
											Tokens: tk[:1],
										},
										Tokens: tk[:1],
									},
									Tokens: tk[:1],
								},
								Tokens: tk[:1],
							},
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				Tokens: tk[:1],
			}
		}},
		{"!a", func(t *test, tk Tokens) { // 2
			t.Output = NotExpression{
				Nots: 1,
				RelationalExpression: RelationalExpression{
					AdditionExpression: AdditionExpression{
						MultiplicationExpression: MultiplicationExpression{
							PipeOrSpecialExpression: PipeOrSpecialExpression{
								SequenceExpression: SequenceExpression{
									UnaryExpression: UnaryExpression{
										ExponentiationExpression: ExponentiationExpression{
											SubsetExpression: SubsetExpression{
												ScopeExpression: ScopeExpression{
													IndexOrCallExpression: IndexOrCallExpression{
														SimpleExpression: &SimpleExpression{
															Identifier: &tk[1],
															Tokens:     tk[1:2],
														},
														Tokens: tk[1:2],
													},
													Tokens: tk[1:2],
												},
												Tokens: tk[1:2],
											},
											Tokens: tk[1:2],
										},
										Tokens: tk[1:2],
									},
									Tokens: tk[1:2],
								},
								Tokens: tk[1:2],
							},
							Tokens: tk[1:2],
						},
						Tokens: tk[1:2],
					},
					Tokens: tk[1:2],
				},
				Tokens: tk[:2],
			}
		}},
		{"! a", func(t *test, tk Tokens) { // 3
			t.Output = NotExpression{
				Nots: 1,
				RelationalExpression: RelationalExpression{
					AdditionExpression: AdditionExpression{
						MultiplicationExpression: MultiplicationExpression{
							PipeOrSpecialExpression: PipeOrSpecialExpression{
								SequenceExpression: SequenceExpression{
									UnaryExpression: UnaryExpression{
										ExponentiationExpression: ExponentiationExpression{
											SubsetExpression: SubsetExpression{
												ScopeExpression: ScopeExpression{
													IndexOrCallExpression: IndexOrCallExpression{
														SimpleExpression: &SimpleExpression{
															Identifier: &tk[2],
															Tokens:     tk[2:3],
														},
														Tokens: tk[2:3],
													},
													Tokens: tk[2:3],
												},
												Tokens: tk[2:3],
											},
											Tokens: tk[2:3],
										},
										Tokens: tk[2:3],
									},
									Tokens: tk[2:3],
								},
								Tokens: tk[2:3],
							},
							Tokens: tk[2:3],
						},
						Tokens: tk[2:3],
					},
					Tokens: tk[2:3],
				},
				Tokens: tk[:3],
			}
		}},
		{"! !! a", func(t *test, tk Tokens) { // 4
			t.Output = NotExpression{
				Nots: 3,
				RelationalExpression: RelationalExpression{
					AdditionExpression: AdditionExpression{
						MultiplicationExpression: MultiplicationExpression{
							PipeOrSpecialExpression: PipeOrSpecialExpression{
								SequenceExpression: SequenceExpression{
									UnaryExpression: UnaryExpression{
										ExponentiationExpression: ExponentiationExpression{
											SubsetExpression: SubsetExpression{
												ScopeExpression: ScopeExpression{
													IndexOrCallExpression: IndexOrCallExpression{
														SimpleExpression: &SimpleExpression{
															Identifier: &tk[5],
															Tokens:     tk[5:6],
														},
														Tokens: tk[5:6],
													},
													Tokens: tk[5:6],
												},
												Tokens: tk[5:6],
											},
											Tokens: tk[5:6],
										},
										Tokens: tk[5:6],
									},
									Tokens: tk[5:6],
								},
								Tokens: tk[5:6],
							},
							Tokens: tk[5:6],
						},
						Tokens: tk[5:6],
					},
					Tokens: tk[5:6],
				},
				Tokens: tk[:6],
			}
		}},
		{"in", func(t *test, tk Tokens) { // 5
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: Error{
							Err: Error{
								Err: Error{
									Err: Error{
										Err: Error{
											Err: Error{
												Err: Error{
													Err: Error{
														Err: Error{
															Err:     ErrInvalidSimpleExpression,
															Parsing: "SimpleExpression",
															Token:   tk[0],
														},
														Parsing: "IndexOrCallExpression",
														Token:   tk[0],
													},
													Parsing: "ScopeExpression",
													Token:   tk[0],
												},
												Parsing: "SubsetExpression",
												Token:   tk[0],
											},
											Parsing: "ExponentiationExpression",
											Token:   tk[0],
										},
										Parsing: "UnaryExpression",
										Token:   tk[0],
									},
									Parsing: "SequenceExpression",
									Token:   tk[0],
								},
								Parsing: "PipeOrSpecialExpression",
								Token:   tk[0],
							},
							Parsing: "MultiplicationExpression",
							Token:   tk[0],
						},
						Parsing: "AdditionExpression",
						Token:   tk[0],
					},
					Parsing: "RelationalExpression",
					Token:   tk[0],
				},
				Parsing: "NotExpression",
				Token:   tk[0],
			}
		}},
	}, func(t *test) (Type, error) {
		var ne NotExpression

		err := ne.parse(&t.Tokens)

		return ne, err
	})
}

func TestRelationalExpression(t *testing.T) {
	doTests(t, []sourceFn{
		{"a", func(t *test, tk Tokens) { // 1
			t.Output = RelationalExpression{
				AdditionExpression: AdditionExpression{
					MultiplicationExpression: MultiplicationExpression{
						PipeOrSpecialExpression: PipeOrSpecialExpression{
							SequenceExpression: SequenceExpression{
								UnaryExpression: UnaryExpression{
									ExponentiationExpression: ExponentiationExpression{
										SubsetExpression: SubsetExpression{
											ScopeExpression: ScopeExpression{
												IndexOrCallExpression: IndexOrCallExpression{
													SimpleExpression: &SimpleExpression{
														Identifier: &tk[0],
														Tokens:     tk[:1],
													},
													Tokens: tk[:1],
												},
												Tokens: tk[:1],
											},
											Tokens: tk[:1],
										},
										Tokens: tk[:1],
									},
									Tokens: tk[:1],
								},
								Tokens: tk[:1],
							},
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				Tokens: tk[:1],
			}
		}},
		{"a>b", func(t *test, tk Tokens) { // 2
			t.Output = RelationalExpression{
				AdditionExpression: AdditionExpression{
					MultiplicationExpression: MultiplicationExpression{
						PipeOrSpecialExpression: PipeOrSpecialExpression{
							SequenceExpression: SequenceExpression{
								UnaryExpression: UnaryExpression{
									ExponentiationExpression: ExponentiationExpression{
										SubsetExpression: SubsetExpression{
											ScopeExpression: ScopeExpression{
												IndexOrCallExpression: IndexOrCallExpression{
													SimpleExpression: &SimpleExpression{
														Identifier: &tk[0],
														Tokens:     tk[:1],
													},
													Tokens: tk[:1],
												},
												Tokens: tk[:1],
											},
											Tokens: tk[:1],
										},
										Tokens: tk[:1],
									},
									Tokens: tk[:1],
								},
								Tokens: tk[:1],
							},
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				RelationalOperator: RelationalGreaterThan,
				RelationalExpression: &RelationalExpression{
					AdditionExpression: AdditionExpression{
						MultiplicationExpression: MultiplicationExpression{
							PipeOrSpecialExpression: PipeOrSpecialExpression{
								SequenceExpression: SequenceExpression{
									UnaryExpression: UnaryExpression{
										ExponentiationExpression: ExponentiationExpression{
											SubsetExpression: SubsetExpression{
												ScopeExpression: ScopeExpression{
													IndexOrCallExpression: IndexOrCallExpression{
														SimpleExpression: &SimpleExpression{
															Identifier: &tk[2],
															Tokens:     tk[2:3],
														},
														Tokens: tk[2:3],
													},
													Tokens: tk[2:3],
												},
												Tokens: tk[2:3],
											},
											Tokens: tk[2:3],
										},
										Tokens: tk[2:3],
									},
									Tokens: tk[2:3],
								},
								Tokens: tk[2:3],
							},
							Tokens: tk[2:3],
						},
						Tokens: tk[2:3],
					},
					Tokens: tk[2:3],
				},
				Tokens: tk[:3],
			}
		}},
		{"a < b", func(t *test, tk Tokens) { // 3
			t.Output = RelationalExpression{
				AdditionExpression: AdditionExpression{
					MultiplicationExpression: MultiplicationExpression{
						PipeOrSpecialExpression: PipeOrSpecialExpression{
							SequenceExpression: SequenceExpression{
								UnaryExpression: UnaryExpression{
									ExponentiationExpression: ExponentiationExpression{
										SubsetExpression: SubsetExpression{
											ScopeExpression: ScopeExpression{
												IndexOrCallExpression: IndexOrCallExpression{
													SimpleExpression: &SimpleExpression{
														Identifier: &tk[0],
														Tokens:     tk[:1],
													},
													Tokens: tk[:1],
												},
												Tokens: tk[:1],
											},
											Tokens: tk[:1],
										},
										Tokens: tk[:1],
									},
									Tokens: tk[:1],
								},
								Tokens: tk[:1],
							},
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				RelationalOperator: RelationalLessThan,
				RelationalExpression: &RelationalExpression{
					AdditionExpression: AdditionExpression{
						MultiplicationExpression: MultiplicationExpression{
							PipeOrSpecialExpression: PipeOrSpecialExpression{
								SequenceExpression: SequenceExpression{
									UnaryExpression: UnaryExpression{
										ExponentiationExpression: ExponentiationExpression{
											SubsetExpression: SubsetExpression{
												ScopeExpression: ScopeExpression{
													IndexOrCallExpression: IndexOrCallExpression{
														SimpleExpression: &SimpleExpression{
															Identifier: &tk[4],
															Tokens:     tk[4:5],
														},
														Tokens: tk[4:5],
													},
													Tokens: tk[4:5],
												},
												Tokens: tk[4:5],
											},
											Tokens: tk[4:5],
										},
										Tokens: tk[4:5],
									},
									Tokens: tk[4:5],
								},
								Tokens: tk[4:5],
							},
							Tokens: tk[4:5],
						},
						Tokens: tk[4:5],
					},
					Tokens: tk[4:5],
				},
				Tokens: tk[:5],
			}
		}},
		{"a >= b<=c ==d!= e", func(t *test, tk Tokens) { // 4
			t.Output = RelationalExpression{
				AdditionExpression: AdditionExpression{
					MultiplicationExpression: MultiplicationExpression{
						PipeOrSpecialExpression: PipeOrSpecialExpression{
							SequenceExpression: SequenceExpression{
								UnaryExpression: UnaryExpression{
									ExponentiationExpression: ExponentiationExpression{
										SubsetExpression: SubsetExpression{
											ScopeExpression: ScopeExpression{
												IndexOrCallExpression: IndexOrCallExpression{
													SimpleExpression: &SimpleExpression{
														Identifier: &tk[0],
														Tokens:     tk[:1],
													},
													Tokens: tk[:1],
												},
												Tokens: tk[:1],
											},
											Tokens: tk[:1],
										},
										Tokens: tk[:1],
									},
									Tokens: tk[:1],
								},
								Tokens: tk[:1],
							},
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				RelationalOperator: RelationalGreaterThanOrEqual,
				RelationalExpression: &RelationalExpression{
					AdditionExpression: AdditionExpression{
						MultiplicationExpression: MultiplicationExpression{
							PipeOrSpecialExpression: PipeOrSpecialExpression{
								SequenceExpression: SequenceExpression{
									UnaryExpression: UnaryExpression{
										ExponentiationExpression: ExponentiationExpression{
											SubsetExpression: SubsetExpression{
												ScopeExpression: ScopeExpression{
													IndexOrCallExpression: IndexOrCallExpression{
														SimpleExpression: &SimpleExpression{
															Identifier: &tk[4],
															Tokens:     tk[4:5],
														},
														Tokens: tk[4:5],
													},
													Tokens: tk[4:5],
												},
												Tokens: tk[4:5],
											},
											Tokens: tk[4:5],
										},
										Tokens: tk[4:5],
									},
									Tokens: tk[4:5],
								},
								Tokens: tk[4:5],
							},
							Tokens: tk[4:5],
						},
						Tokens: tk[4:5],
					},
					RelationalOperator: RelationalLessThanOrEqual,
					RelationalExpression: &RelationalExpression{
						AdditionExpression: AdditionExpression{
							MultiplicationExpression: MultiplicationExpression{
								PipeOrSpecialExpression: PipeOrSpecialExpression{
									SequenceExpression: SequenceExpression{
										UnaryExpression: UnaryExpression{
											ExponentiationExpression: ExponentiationExpression{
												SubsetExpression: SubsetExpression{
													ScopeExpression: ScopeExpression{
														IndexOrCallExpression: IndexOrCallExpression{
															SimpleExpression: &SimpleExpression{
																Identifier: &tk[6],
																Tokens:     tk[6:7],
															},
															Tokens: tk[6:7],
														},
														Tokens: tk[6:7],
													},
													Tokens: tk[6:7],
												},
												Tokens: tk[6:7],
											},
											Tokens: tk[6:7],
										},
										Tokens: tk[6:7],
									},
									Tokens: tk[6:7],
								},
								Tokens: tk[6:7],
							},
							Tokens: tk[6:7],
						},
						RelationalOperator: RelationalEqual,
						RelationalExpression: &RelationalExpression{
							AdditionExpression: AdditionExpression{
								MultiplicationExpression: MultiplicationExpression{
									PipeOrSpecialExpression: PipeOrSpecialExpression{
										SequenceExpression: SequenceExpression{
											UnaryExpression: UnaryExpression{
												ExponentiationExpression: ExponentiationExpression{
													SubsetExpression: SubsetExpression{
														ScopeExpression: ScopeExpression{
															IndexOrCallExpression: IndexOrCallExpression{
																SimpleExpression: &SimpleExpression{
																	Identifier: &tk[9],
																	Tokens:     tk[9:10],
																},
																Tokens: tk[9:10],
															},
															Tokens: tk[9:10],
														},
														Tokens: tk[9:10],
													},
													Tokens: tk[9:10],
												},
												Tokens: tk[9:10],
											},
											Tokens: tk[9:10],
										},
										Tokens: tk[9:10],
									},
									Tokens: tk[9:10],
								},
								Tokens: tk[9:10],
							},
							RelationalOperator: RelationalNotEqual,
							RelationalExpression: &RelationalExpression{
								AdditionExpression: AdditionExpression{
									MultiplicationExpression: MultiplicationExpression{
										PipeOrSpecialExpression: PipeOrSpecialExpression{
											SequenceExpression: SequenceExpression{
												UnaryExpression: UnaryExpression{
													ExponentiationExpression: ExponentiationExpression{
														SubsetExpression: SubsetExpression{
															ScopeExpression: ScopeExpression{
																IndexOrCallExpression: IndexOrCallExpression{
																	SimpleExpression: &SimpleExpression{
																		Identifier: &tk[12],
																		Tokens:     tk[12:13],
																	},
																	Tokens: tk[12:13],
																},
																Tokens: tk[12:13],
															},
															Tokens: tk[12:13],
														},
														Tokens: tk[12:13],
													},
													Tokens: tk[12:13],
												},
												Tokens: tk[12:13],
											},
											Tokens: tk[12:13],
										},
										Tokens: tk[12:13],
									},
									Tokens: tk[12:13],
								},
								Tokens: tk[12:13],
							},
							Tokens: tk[9:13],
						},
						Tokens: tk[6:13],
					},
					Tokens: tk[4:13],
				},
				Tokens: tk[:13],
			}
		}},
		{"in", func(t *test, tk Tokens) { // 5
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: Error{
							Err: Error{
								Err: Error{
									Err: Error{
										Err: Error{
											Err: Error{
												Err: Error{
													Err: Error{
														Err:     ErrInvalidSimpleExpression,
														Parsing: "SimpleExpression",
														Token:   tk[0],
													},
													Parsing: "IndexOrCallExpression",
													Token:   tk[0],
												},
												Parsing: "ScopeExpression",
												Token:   tk[0],
											},
											Parsing: "SubsetExpression",
											Token:   tk[0],
										},
										Parsing: "ExponentiationExpression",
										Token:   tk[0],
									},
									Parsing: "UnaryExpression",
									Token:   tk[0],
								},
								Parsing: "SequenceExpression",
								Token:   tk[0],
							},
							Parsing: "PipeOrSpecialExpression",
							Token:   tk[0],
						},
						Parsing: "MultiplicationExpression",
						Token:   tk[0],
					},
					Parsing: "AdditionExpression",
					Token:   tk[0],
				},
				Parsing: "RelationalExpression",
				Token:   tk[0],
			}
		}},
		{"a>in", func(t *test, tk Tokens) { // 6
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: Error{
							Err: Error{
								Err: Error{
									Err: Error{
										Err: Error{
											Err: Error{
												Err: Error{
													Err: Error{
														Err: Error{
															Err:     ErrInvalidSimpleExpression,
															Parsing: "SimpleExpression",
															Token:   tk[2],
														},
														Parsing: "IndexOrCallExpression",
														Token:   tk[2],
													},
													Parsing: "ScopeExpression",
													Token:   tk[2],
												},
												Parsing: "SubsetExpression",
												Token:   tk[2],
											},
											Parsing: "ExponentiationExpression",
											Token:   tk[2],
										},
										Parsing: "UnaryExpression",
										Token:   tk[2],
									},
									Parsing: "SequenceExpression",
									Token:   tk[2],
								},
								Parsing: "PipeOrSpecialExpression",
								Token:   tk[2],
							},
							Parsing: "MultiplicationExpression",
							Token:   tk[2],
						},
						Parsing: "AdditionExpression",
						Token:   tk[2],
					},
					Parsing: "RelationalExpression",
					Token:   tk[2],
				},
				Parsing: "RelationalExpression",
				Token:   tk[2],
			}
		}},
	}, func(t *test) (Type, error) {
		var re RelationalExpression

		err := re.parse(&t.Tokens)

		return re, err
	})
}

func TestAdditionExpression(t *testing.T) {
	doTests(t, []sourceFn{
		{"a", func(t *test, tk Tokens) { // 1
			t.Output = AdditionExpression{
				MultiplicationExpression: MultiplicationExpression{
					PipeOrSpecialExpression: PipeOrSpecialExpression{
						SequenceExpression: SequenceExpression{
							UnaryExpression: UnaryExpression{
								ExponentiationExpression: ExponentiationExpression{
									SubsetExpression: SubsetExpression{
										ScopeExpression: ScopeExpression{
											IndexOrCallExpression: IndexOrCallExpression{
												SimpleExpression: &SimpleExpression{
													Identifier: &tk[0],
													Tokens:     tk[:1],
												},
												Tokens: tk[:1],
											},
											Tokens: tk[:1],
										},
										Tokens: tk[:1],
									},
									Tokens: tk[:1],
								},
								Tokens: tk[:1],
							},
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				Tokens: tk[:1],
			}
		}},
		{"a+b", func(t *test, tk Tokens) { // 2
			t.Output = AdditionExpression{
				MultiplicationExpression: MultiplicationExpression{
					PipeOrSpecialExpression: PipeOrSpecialExpression{
						SequenceExpression: SequenceExpression{
							UnaryExpression: UnaryExpression{
								ExponentiationExpression: ExponentiationExpression{
									SubsetExpression: SubsetExpression{
										ScopeExpression: ScopeExpression{
											IndexOrCallExpression: IndexOrCallExpression{
												SimpleExpression: &SimpleExpression{
													Identifier: &tk[0],
													Tokens:     tk[:1],
												},
												Tokens: tk[:1],
											},
											Tokens: tk[:1],
										},
										Tokens: tk[:1],
									},
									Tokens: tk[:1],
								},
								Tokens: tk[:1],
							},
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				AdditionType: AdditionAdd,
				AdditionExpression: &AdditionExpression{
					MultiplicationExpression: MultiplicationExpression{
						PipeOrSpecialExpression: PipeOrSpecialExpression{
							SequenceExpression: SequenceExpression{
								UnaryExpression: UnaryExpression{
									ExponentiationExpression: ExponentiationExpression{
										SubsetExpression: SubsetExpression{
											ScopeExpression: ScopeExpression{
												IndexOrCallExpression: IndexOrCallExpression{
													SimpleExpression: &SimpleExpression{
														Identifier: &tk[2],
														Tokens:     tk[2:3],
													},
													Tokens: tk[2:3],
												},
												Tokens: tk[2:3],
											},
											Tokens: tk[2:3],
										},
										Tokens: tk[2:3],
									},
									Tokens: tk[2:3],
								},
								Tokens: tk[2:3],
							},
							Tokens: tk[2:3],
						},
						Tokens: tk[2:3],
					},
					Tokens: tk[2:3],
				},
				Tokens: tk[:3],
			}
		}},
		{"a-b", func(t *test, tk Tokens) { // 3
			t.Output = AdditionExpression{
				MultiplicationExpression: MultiplicationExpression{
					PipeOrSpecialExpression: PipeOrSpecialExpression{
						SequenceExpression: SequenceExpression{
							UnaryExpression: UnaryExpression{
								ExponentiationExpression: ExponentiationExpression{
									SubsetExpression: SubsetExpression{
										ScopeExpression: ScopeExpression{
											IndexOrCallExpression: IndexOrCallExpression{
												SimpleExpression: &SimpleExpression{
													Identifier: &tk[0],
													Tokens:     tk[:1],
												},
												Tokens: tk[:1],
											},
											Tokens: tk[:1],
										},
										Tokens: tk[:1],
									},
									Tokens: tk[:1],
								},
								Tokens: tk[:1],
							},
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				AdditionType: AdditionSubtract,
				AdditionExpression: &AdditionExpression{
					MultiplicationExpression: MultiplicationExpression{
						PipeOrSpecialExpression: PipeOrSpecialExpression{
							SequenceExpression: SequenceExpression{
								UnaryExpression: UnaryExpression{
									ExponentiationExpression: ExponentiationExpression{
										SubsetExpression: SubsetExpression{
											ScopeExpression: ScopeExpression{
												IndexOrCallExpression: IndexOrCallExpression{
													SimpleExpression: &SimpleExpression{
														Identifier: &tk[2],
														Tokens:     tk[2:3],
													},
													Tokens: tk[2:3],
												},
												Tokens: tk[2:3],
											},
											Tokens: tk[2:3],
										},
										Tokens: tk[2:3],
									},
									Tokens: tk[2:3],
								},
								Tokens: tk[2:3],
							},
							Tokens: tk[2:3],
						},
						Tokens: tk[2:3],
					},
					Tokens: tk[2:3],
				},
				Tokens: tk[:3],
			}
		}},
		{"a + b", func(t *test, tk Tokens) { // 4
			t.Output = AdditionExpression{
				MultiplicationExpression: MultiplicationExpression{
					PipeOrSpecialExpression: PipeOrSpecialExpression{
						SequenceExpression: SequenceExpression{
							UnaryExpression: UnaryExpression{
								ExponentiationExpression: ExponentiationExpression{
									SubsetExpression: SubsetExpression{
										ScopeExpression: ScopeExpression{
											IndexOrCallExpression: IndexOrCallExpression{
												SimpleExpression: &SimpleExpression{
													Identifier: &tk[0],
													Tokens:     tk[:1],
												},
												Tokens: tk[:1],
											},
											Tokens: tk[:1],
										},
										Tokens: tk[:1],
									},
									Tokens: tk[:1],
								},
								Tokens: tk[:1],
							},
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				AdditionType: AdditionAdd,
				AdditionExpression: &AdditionExpression{
					MultiplicationExpression: MultiplicationExpression{
						PipeOrSpecialExpression: PipeOrSpecialExpression{
							SequenceExpression: SequenceExpression{
								UnaryExpression: UnaryExpression{
									ExponentiationExpression: ExponentiationExpression{
										SubsetExpression: SubsetExpression{
											ScopeExpression: ScopeExpression{
												IndexOrCallExpression: IndexOrCallExpression{
													SimpleExpression: &SimpleExpression{
														Identifier: &tk[4],
														Tokens:     tk[4:5],
													},
													Tokens: tk[4:5],
												},
												Tokens: tk[4:5],
											},
											Tokens: tk[4:5],
										},
										Tokens: tk[4:5],
									},
									Tokens: tk[4:5],
								},
								Tokens: tk[4:5],
							},
							Tokens: tk[4:5],
						},
						Tokens: tk[4:5],
					},
					Tokens: tk[4:5],
				},
				Tokens: tk[:5],
			}
		}},
		{"a - b", func(t *test, tk Tokens) { // 5
			t.Output = AdditionExpression{
				MultiplicationExpression: MultiplicationExpression{
					PipeOrSpecialExpression: PipeOrSpecialExpression{
						SequenceExpression: SequenceExpression{
							UnaryExpression: UnaryExpression{
								ExponentiationExpression: ExponentiationExpression{
									SubsetExpression: SubsetExpression{
										ScopeExpression: ScopeExpression{
											IndexOrCallExpression: IndexOrCallExpression{
												SimpleExpression: &SimpleExpression{
													Identifier: &tk[0],
													Tokens:     tk[:1],
												},
												Tokens: tk[:1],
											},
											Tokens: tk[:1],
										},
										Tokens: tk[:1],
									},
									Tokens: tk[:1],
								},
								Tokens: tk[:1],
							},
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				AdditionType: AdditionSubtract,
				AdditionExpression: &AdditionExpression{
					MultiplicationExpression: MultiplicationExpression{
						PipeOrSpecialExpression: PipeOrSpecialExpression{
							SequenceExpression: SequenceExpression{
								UnaryExpression: UnaryExpression{
									ExponentiationExpression: ExponentiationExpression{
										SubsetExpression: SubsetExpression{
											ScopeExpression: ScopeExpression{
												IndexOrCallExpression: IndexOrCallExpression{
													SimpleExpression: &SimpleExpression{
														Identifier: &tk[4],
														Tokens:     tk[4:5],
													},
													Tokens: tk[4:5],
												},
												Tokens: tk[4:5],
											},
											Tokens: tk[4:5],
										},
										Tokens: tk[4:5],
									},
									Tokens: tk[4:5],
								},
								Tokens: tk[4:5],
							},
							Tokens: tk[4:5],
						},
						Tokens: tk[4:5],
					},
					Tokens: tk[4:5],
				},
				Tokens: tk[:5],
			}
		}},
		{"in", func(t *test, tk Tokens) { // 6
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: Error{
							Err: Error{
								Err: Error{
									Err: Error{
										Err: Error{
											Err: Error{
												Err: Error{
													Err:     ErrInvalidSimpleExpression,
													Parsing: "SimpleExpression",
													Token:   tk[0],
												},
												Parsing: "IndexOrCallExpression",
												Token:   tk[0],
											},
											Parsing: "ScopeExpression",
											Token:   tk[0],
										},
										Parsing: "SubsetExpression",
										Token:   tk[0],
									},
									Parsing: "ExponentiationExpression",
									Token:   tk[0],
								},
								Parsing: "UnaryExpression",
								Token:   tk[0],
							},
							Parsing: "SequenceExpression",
							Token:   tk[0],
						},
						Parsing: "PipeOrSpecialExpression",
						Token:   tk[0],
					},
					Parsing: "MultiplicationExpression",
					Token:   tk[0],
				},
				Parsing: "AdditionExpression",
				Token:   tk[0],
			}
		}},
		{"a+in", func(t *test, tk Tokens) { // 7
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: Error{
							Err: Error{
								Err: Error{
									Err: Error{
										Err: Error{
											Err: Error{
												Err: Error{
													Err: Error{
														Err:     ErrInvalidSimpleExpression,
														Parsing: "SimpleExpression",
														Token:   tk[2],
													},
													Parsing: "IndexOrCallExpression",
													Token:   tk[2],
												},
												Parsing: "ScopeExpression",
												Token:   tk[2],
											},
											Parsing: "SubsetExpression",
											Token:   tk[2],
										},
										Parsing: "ExponentiationExpression",
										Token:   tk[2],
									},
									Parsing: "UnaryExpression",
									Token:   tk[2],
								},
								Parsing: "SequenceExpression",
								Token:   tk[2],
							},
							Parsing: "PipeOrSpecialExpression",
							Token:   tk[2],
						},
						Parsing: "MultiplicationExpression",
						Token:   tk[2],
					},
					Parsing: "AdditionExpression",
					Token:   tk[2],
				},
				Parsing: "AdditionExpression",
				Token:   tk[2],
			}
		}},
	}, func(t *test) (Type, error) {
		var ae AdditionExpression

		err := ae.parse(&t.Tokens)

		return ae, err
	})
}

func TestMultiplicationExpression(t *testing.T) {
	doTests(t, []sourceFn{
		{"a", func(t *test, tk Tokens) { // 1
			t.Output = MultiplicationExpression{
				PipeOrSpecialExpression: PipeOrSpecialExpression{
					SequenceExpression: SequenceExpression{
						UnaryExpression: UnaryExpression{
							ExponentiationExpression: ExponentiationExpression{
								SubsetExpression: SubsetExpression{
									ScopeExpression: ScopeExpression{
										IndexOrCallExpression: IndexOrCallExpression{
											SimpleExpression: &SimpleExpression{
												Identifier: &tk[0],
												Tokens:     tk[:1],
											},
											Tokens: tk[:1],
										},
										Tokens: tk[:1],
									},
									Tokens: tk[:1],
								},
								Tokens: tk[:1],
							},
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				Tokens: tk[:1],
			}
		}},
		{"a*b", func(t *test, tk Tokens) { // 2
			t.Output = MultiplicationExpression{
				PipeOrSpecialExpression: PipeOrSpecialExpression{
					SequenceExpression: SequenceExpression{
						UnaryExpression: UnaryExpression{
							ExponentiationExpression: ExponentiationExpression{
								SubsetExpression: SubsetExpression{
									ScopeExpression: ScopeExpression{
										IndexOrCallExpression: IndexOrCallExpression{
											SimpleExpression: &SimpleExpression{
												Identifier: &tk[0],
												Tokens:     tk[:1],
											},
											Tokens: tk[:1],
										},
										Tokens: tk[:1],
									},
									Tokens: tk[:1],
								},
								Tokens: tk[:1],
							},
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				MultiplicationType: MultiplicationMultiply,
				MultiplicationExpression: &MultiplicationExpression{
					PipeOrSpecialExpression: PipeOrSpecialExpression{
						SequenceExpression: SequenceExpression{
							UnaryExpression: UnaryExpression{
								ExponentiationExpression: ExponentiationExpression{
									SubsetExpression: SubsetExpression{
										ScopeExpression: ScopeExpression{
											IndexOrCallExpression: IndexOrCallExpression{
												SimpleExpression: &SimpleExpression{
													Identifier: &tk[2],
													Tokens:     tk[2:3],
												},
												Tokens: tk[2:3],
											},
											Tokens: tk[2:3],
										},
										Tokens: tk[2:3],
									},
									Tokens: tk[2:3],
								},
								Tokens: tk[2:3],
							},
							Tokens: tk[2:3],
						},
						Tokens: tk[2:3],
					},
					Tokens: tk[2:3],
				},
				Tokens: tk[:3],
			}
		}},
		{"a/b", func(t *test, tk Tokens) { // 3
			t.Output = MultiplicationExpression{
				PipeOrSpecialExpression: PipeOrSpecialExpression{
					SequenceExpression: SequenceExpression{
						UnaryExpression: UnaryExpression{
							ExponentiationExpression: ExponentiationExpression{
								SubsetExpression: SubsetExpression{
									ScopeExpression: ScopeExpression{
										IndexOrCallExpression: IndexOrCallExpression{
											SimpleExpression: &SimpleExpression{
												Identifier: &tk[0],
												Tokens:     tk[:1],
											},
											Tokens: tk[:1],
										},
										Tokens: tk[:1],
									},
									Tokens: tk[:1],
								},
								Tokens: tk[:1],
							},
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				MultiplicationType: MultiplicationDivide,
				MultiplicationExpression: &MultiplicationExpression{
					PipeOrSpecialExpression: PipeOrSpecialExpression{
						SequenceExpression: SequenceExpression{
							UnaryExpression: UnaryExpression{
								ExponentiationExpression: ExponentiationExpression{
									SubsetExpression: SubsetExpression{
										ScopeExpression: ScopeExpression{
											IndexOrCallExpression: IndexOrCallExpression{
												SimpleExpression: &SimpleExpression{
													Identifier: &tk[2],
													Tokens:     tk[2:3],
												},
												Tokens: tk[2:3],
											},
											Tokens: tk[2:3],
										},
										Tokens: tk[2:3],
									},
									Tokens: tk[2:3],
								},
								Tokens: tk[2:3],
							},
							Tokens: tk[2:3],
						},
						Tokens: tk[2:3],
					},
					Tokens: tk[2:3],
				},
				Tokens: tk[:3],
			}
		}},
		{"a * b", func(t *test, tk Tokens) { // 4
			t.Output = MultiplicationExpression{
				PipeOrSpecialExpression: PipeOrSpecialExpression{
					SequenceExpression: SequenceExpression{
						UnaryExpression: UnaryExpression{
							ExponentiationExpression: ExponentiationExpression{
								SubsetExpression: SubsetExpression{
									ScopeExpression: ScopeExpression{
										IndexOrCallExpression: IndexOrCallExpression{
											SimpleExpression: &SimpleExpression{
												Identifier: &tk[0],
												Tokens:     tk[:1],
											},
											Tokens: tk[:1],
										},
										Tokens: tk[:1],
									},
									Tokens: tk[:1],
								},
								Tokens: tk[:1],
							},
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				MultiplicationType: MultiplicationMultiply,
				MultiplicationExpression: &MultiplicationExpression{
					PipeOrSpecialExpression: PipeOrSpecialExpression{
						SequenceExpression: SequenceExpression{
							UnaryExpression: UnaryExpression{
								ExponentiationExpression: ExponentiationExpression{
									SubsetExpression: SubsetExpression{
										ScopeExpression: ScopeExpression{
											IndexOrCallExpression: IndexOrCallExpression{
												SimpleExpression: &SimpleExpression{
													Identifier: &tk[4],
													Tokens:     tk[4:5],
												},
												Tokens: tk[4:5],
											},
											Tokens: tk[4:5],
										},
										Tokens: tk[4:5],
									},
									Tokens: tk[4:5],
								},
								Tokens: tk[4:5],
							},
							Tokens: tk[4:5],
						},
						Tokens: tk[4:5],
					},
					Tokens: tk[4:5],
				},
				Tokens: tk[:5],
			}
		}},
		{"a / b", func(t *test, tk Tokens) { // 5
			t.Output = MultiplicationExpression{
				PipeOrSpecialExpression: PipeOrSpecialExpression{
					SequenceExpression: SequenceExpression{
						UnaryExpression: UnaryExpression{
							ExponentiationExpression: ExponentiationExpression{
								SubsetExpression: SubsetExpression{
									ScopeExpression: ScopeExpression{
										IndexOrCallExpression: IndexOrCallExpression{
											SimpleExpression: &SimpleExpression{
												Identifier: &tk[0],
												Tokens:     tk[:1],
											},
											Tokens: tk[:1],
										},
										Tokens: tk[:1],
									},
									Tokens: tk[:1],
								},
								Tokens: tk[:1],
							},
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				MultiplicationType: MultiplicationDivide,
				MultiplicationExpression: &MultiplicationExpression{
					PipeOrSpecialExpression: PipeOrSpecialExpression{
						SequenceExpression: SequenceExpression{
							UnaryExpression: UnaryExpression{
								ExponentiationExpression: ExponentiationExpression{
									SubsetExpression: SubsetExpression{
										ScopeExpression: ScopeExpression{
											IndexOrCallExpression: IndexOrCallExpression{
												SimpleExpression: &SimpleExpression{
													Identifier: &tk[4],
													Tokens:     tk[4:5],
												},
												Tokens: tk[4:5],
											},
											Tokens: tk[4:5],
										},
										Tokens: tk[4:5],
									},
									Tokens: tk[4:5],
								},
								Tokens: tk[4:5],
							},
							Tokens: tk[4:5],
						},
						Tokens: tk[4:5],
					},
					Tokens: tk[4:5],
				},
				Tokens: tk[:5],
			}
		}},
		{"in", func(t *test, tk Tokens) { // 6
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: Error{
							Err: Error{
								Err: Error{
									Err: Error{
										Err: Error{
											Err: Error{
												Err:     ErrInvalidSimpleExpression,
												Parsing: "SimpleExpression",
												Token:   tk[0],
											},
											Parsing: "IndexOrCallExpression",
											Token:   tk[0],
										},
										Parsing: "ScopeExpression",
										Token:   tk[0],
									},
									Parsing: "SubsetExpression",
									Token:   tk[0],
								},
								Parsing: "ExponentiationExpression",
								Token:   tk[0],
							},
							Parsing: "UnaryExpression",
							Token:   tk[0],
						},
						Parsing: "SequenceExpression",
						Token:   tk[0],
					},
					Parsing: "PipeOrSpecialExpression",
					Token:   tk[0],
				},
				Parsing: "MultiplicationExpression",
				Token:   tk[0],
			}
		}},
		{"a*in", func(t *test, tk Tokens) { // 7
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: Error{
							Err: Error{
								Err: Error{
									Err: Error{
										Err: Error{
											Err: Error{
												Err: Error{
													Err:     ErrInvalidSimpleExpression,
													Parsing: "SimpleExpression",
													Token:   tk[2],
												},
												Parsing: "IndexOrCallExpression",
												Token:   tk[2],
											},
											Parsing: "ScopeExpression",
											Token:   tk[2],
										},
										Parsing: "SubsetExpression",
										Token:   tk[2],
									},
									Parsing: "ExponentiationExpression",
									Token:   tk[2],
								},
								Parsing: "UnaryExpression",
								Token:   tk[2],
							},
							Parsing: "SequenceExpression",
							Token:   tk[2],
						},
						Parsing: "PipeOrSpecialExpression",
						Token:   tk[2],
					},
					Parsing: "MultiplicationExpression",
					Token:   tk[2],
				},
				Parsing: "MultiplicationExpression",
				Token:   tk[2],
			}
		}},
	}, func(t *test) (Type, error) {
		var me MultiplicationExpression

		err := me.parse(&t.Tokens)

		return me, err
	})
}

func TestPipeOrSpecialExpression(t *testing.T) {
	doTests(t, []sourceFn{
		{"a", func(t *test, tk Tokens) { // 1
			t.Output = PipeOrSpecialExpression{
				SequenceExpression: SequenceExpression{
					UnaryExpression: UnaryExpression{
						ExponentiationExpression: ExponentiationExpression{
							SubsetExpression: SubsetExpression{
								ScopeExpression: ScopeExpression{
									IndexOrCallExpression: IndexOrCallExpression{
										SimpleExpression: &SimpleExpression{
											Identifier: &tk[0],
											Tokens:     tk[:1],
										},
										Tokens: tk[:1],
									},
									Tokens: tk[:1],
								},
								Tokens: tk[:1],
							},
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				Tokens: tk[:1],
			}
		}},
		{"a|>b", func(t *test, tk Tokens) { // 2
			t.Output = PipeOrSpecialExpression{
				SequenceExpression: SequenceExpression{
					UnaryExpression: UnaryExpression{
						ExponentiationExpression: ExponentiationExpression{
							SubsetExpression: SubsetExpression{
								ScopeExpression: ScopeExpression{
									IndexOrCallExpression: IndexOrCallExpression{
										SimpleExpression: &SimpleExpression{
											Identifier: &tk[0],
											Tokens:     tk[:1],
										},
										Tokens: tk[:1],
									},
									Tokens: tk[:1],
								},
								Tokens: tk[:1],
							},
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				Operator: &tk[1],
				PipeOrSpecialExpression: &PipeOrSpecialExpression{
					SequenceExpression: SequenceExpression{
						UnaryExpression: UnaryExpression{
							ExponentiationExpression: ExponentiationExpression{
								SubsetExpression: SubsetExpression{
									ScopeExpression: ScopeExpression{
										IndexOrCallExpression: IndexOrCallExpression{
											SimpleExpression: &SimpleExpression{
												Identifier: &tk[2],
												Tokens:     tk[2:3],
											},
											Tokens: tk[2:3],
										},
										Tokens: tk[2:3],
									},
									Tokens: tk[2:3],
								},
								Tokens: tk[2:3],
							},
							Tokens: tk[2:3],
						},
						Tokens: tk[2:3],
					},
					Tokens: tk[2:3],
				},
				Tokens: tk[:3],
			}
		}},
		{"a |> b", func(t *test, tk Tokens) { // 3
			t.Output = PipeOrSpecialExpression{
				SequenceExpression: SequenceExpression{
					UnaryExpression: UnaryExpression{
						ExponentiationExpression: ExponentiationExpression{
							SubsetExpression: SubsetExpression{
								ScopeExpression: ScopeExpression{
									IndexOrCallExpression: IndexOrCallExpression{
										SimpleExpression: &SimpleExpression{
											Identifier: &tk[0],
											Tokens:     tk[:1],
										},
										Tokens: tk[:1],
									},
									Tokens: tk[:1],
								},
								Tokens: tk[:1],
							},
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				Operator: &tk[2],
				PipeOrSpecialExpression: &PipeOrSpecialExpression{
					SequenceExpression: SequenceExpression{
						UnaryExpression: UnaryExpression{
							ExponentiationExpression: ExponentiationExpression{
								SubsetExpression: SubsetExpression{
									ScopeExpression: ScopeExpression{
										IndexOrCallExpression: IndexOrCallExpression{
											SimpleExpression: &SimpleExpression{
												Identifier: &tk[4],
												Tokens:     tk[4:5],
											},
											Tokens: tk[4:5],
										},
										Tokens: tk[4:5],
									},
									Tokens: tk[4:5],
								},
								Tokens: tk[4:5],
							},
							Tokens: tk[4:5],
						},
						Tokens: tk[4:5],
					},
					Tokens: tk[4:5],
				},
				Tokens: tk[:5],
			}
		}},
		{"a%%b", func(t *test, tk Tokens) { // 4
			t.Output = PipeOrSpecialExpression{
				SequenceExpression: SequenceExpression{
					UnaryExpression: UnaryExpression{
						ExponentiationExpression: ExponentiationExpression{
							SubsetExpression: SubsetExpression{
								ScopeExpression: ScopeExpression{
									IndexOrCallExpression: IndexOrCallExpression{
										SimpleExpression: &SimpleExpression{
											Identifier: &tk[0],
											Tokens:     tk[:1],
										},
										Tokens: tk[:1],
									},
									Tokens: tk[:1],
								},
								Tokens: tk[:1],
							},
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				Operator: &tk[1],
				PipeOrSpecialExpression: &PipeOrSpecialExpression{
					SequenceExpression: SequenceExpression{
						UnaryExpression: UnaryExpression{
							ExponentiationExpression: ExponentiationExpression{
								SubsetExpression: SubsetExpression{
									ScopeExpression: ScopeExpression{
										IndexOrCallExpression: IndexOrCallExpression{
											SimpleExpression: &SimpleExpression{
												Identifier: &tk[2],
												Tokens:     tk[2:3],
											},
											Tokens: tk[2:3],
										},
										Tokens: tk[2:3],
									},
									Tokens: tk[2:3],
								},
								Tokens: tk[2:3],
							},
							Tokens: tk[2:3],
						},
						Tokens: tk[2:3],
					},
					Tokens: tk[2:3],
				},
				Tokens: tk[:3],
			}
		}},
		{"a %abc% b", func(t *test, tk Tokens) { // 5
			t.Output = PipeOrSpecialExpression{
				SequenceExpression: SequenceExpression{
					UnaryExpression: UnaryExpression{
						ExponentiationExpression: ExponentiationExpression{
							SubsetExpression: SubsetExpression{
								ScopeExpression: ScopeExpression{
									IndexOrCallExpression: IndexOrCallExpression{
										SimpleExpression: &SimpleExpression{
											Identifier: &tk[0],
											Tokens:     tk[:1],
										},
										Tokens: tk[:1],
									},
									Tokens: tk[:1],
								},
								Tokens: tk[:1],
							},
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				Operator: &tk[2],
				PipeOrSpecialExpression: &PipeOrSpecialExpression{
					SequenceExpression: SequenceExpression{
						UnaryExpression: UnaryExpression{
							ExponentiationExpression: ExponentiationExpression{
								SubsetExpression: SubsetExpression{
									ScopeExpression: ScopeExpression{
										IndexOrCallExpression: IndexOrCallExpression{
											SimpleExpression: &SimpleExpression{
												Identifier: &tk[4],
												Tokens:     tk[4:5],
											},
											Tokens: tk[4:5],
										},
										Tokens: tk[4:5],
									},
									Tokens: tk[4:5],
								},
								Tokens: tk[4:5],
							},
							Tokens: tk[4:5],
						},
						Tokens: tk[4:5],
					},
					Tokens: tk[4:5],
				},
				Tokens: tk[:5],
			}
		}},
		{"in", func(t *test, tk Tokens) { // 6
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: Error{
							Err: Error{
								Err: Error{
									Err: Error{
										Err: Error{
											Err:     ErrInvalidSimpleExpression,
											Parsing: "SimpleExpression",
											Token:   tk[0],
										},
										Parsing: "IndexOrCallExpression",
										Token:   tk[0],
									},
									Parsing: "ScopeExpression",
									Token:   tk[0],
								},
								Parsing: "SubsetExpression",
								Token:   tk[0],
							},
							Parsing: "ExponentiationExpression",
							Token:   tk[0],
						},
						Parsing: "UnaryExpression",
						Token:   tk[0],
					},
					Parsing: "SequenceExpression",
					Token:   tk[0],
				},
				Parsing: "PipeOrSpecialExpression",
				Token:   tk[0],
			}
		}},
		{"a|>in", func(t *test, tk Tokens) { // 7
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: Error{
							Err: Error{
								Err: Error{
									Err: Error{
										Err: Error{
											Err: Error{
												Err:     ErrInvalidSimpleExpression,
												Parsing: "SimpleExpression",
												Token:   tk[2],
											},
											Parsing: "IndexOrCallExpression",
											Token:   tk[2],
										},
										Parsing: "ScopeExpression",
										Token:   tk[2],
									},
									Parsing: "SubsetExpression",
									Token:   tk[2],
								},
								Parsing: "ExponentiationExpression",
								Token:   tk[2],
							},
							Parsing: "UnaryExpression",
							Token:   tk[2],
						},
						Parsing: "SequenceExpression",
						Token:   tk[2],
					},
					Parsing: "PipeOrSpecialExpression",
					Token:   tk[2],
				},
				Parsing: "PipeOrSpecialExpression",
				Token:   tk[2],
			}
		}},
	}, func(t *test) (Type, error) {
		var pe PipeOrSpecialExpression

		err := pe.parse(&t.Tokens)

		return pe, err
	})
}

func TestSequenceExpression(t *testing.T) {
	doTests(t, []sourceFn{
		{"a", func(t *test, tk Tokens) { // 1
			t.Output = SequenceExpression{
				UnaryExpression: UnaryExpression{
					ExponentiationExpression: ExponentiationExpression{
						SubsetExpression: SubsetExpression{
							ScopeExpression: ScopeExpression{
								IndexOrCallExpression: IndexOrCallExpression{
									SimpleExpression: &SimpleExpression{
										Identifier: &tk[0],
										Tokens:     tk[:1],
									},
									Tokens: tk[:1],
								},
								Tokens: tk[:1],
							},
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				Tokens: tk[:1],
			}
		}},
		{"a:b", func(t *test, tk Tokens) { // 2
			t.Output = SequenceExpression{
				UnaryExpression: UnaryExpression{
					ExponentiationExpression: ExponentiationExpression{
						SubsetExpression: SubsetExpression{
							ScopeExpression: ScopeExpression{
								IndexOrCallExpression: IndexOrCallExpression{
									SimpleExpression: &SimpleExpression{
										Identifier: &tk[0],
										Tokens:     tk[:1],
									},
									Tokens: tk[:1],
								},
								Tokens: tk[:1],
							},
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				SequenceExpression: &SequenceExpression{
					UnaryExpression: UnaryExpression{
						ExponentiationExpression: ExponentiationExpression{
							SubsetExpression: SubsetExpression{
								ScopeExpression: ScopeExpression{
									IndexOrCallExpression: IndexOrCallExpression{
										SimpleExpression: &SimpleExpression{
											Identifier: &tk[2],
											Tokens:     tk[2:3],
										},
										Tokens: tk[2:3],
									},
									Tokens: tk[2:3],
								},
								Tokens: tk[2:3],
							},
							Tokens: tk[2:3],
						},
						Tokens: tk[2:3],
					},
					Tokens: tk[2:3],
				},
				Tokens: tk[:3],
			}
		}},
		{"a : b", func(t *test, tk Tokens) { // 3
			t.Output = SequenceExpression{
				UnaryExpression: UnaryExpression{
					ExponentiationExpression: ExponentiationExpression{
						SubsetExpression: SubsetExpression{
							ScopeExpression: ScopeExpression{
								IndexOrCallExpression: IndexOrCallExpression{
									SimpleExpression: &SimpleExpression{
										Identifier: &tk[0],
										Tokens:     tk[:1],
									},
									Tokens: tk[:1],
								},
								Tokens: tk[:1],
							},
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				SequenceExpression: &SequenceExpression{
					UnaryExpression: UnaryExpression{
						ExponentiationExpression: ExponentiationExpression{
							SubsetExpression: SubsetExpression{
								ScopeExpression: ScopeExpression{
									IndexOrCallExpression: IndexOrCallExpression{
										SimpleExpression: &SimpleExpression{
											Identifier: &tk[4],
											Tokens:     tk[4:5],
										},
										Tokens: tk[4:5],
									},
									Tokens: tk[4:5],
								},
								Tokens: tk[4:5],
							},
							Tokens: tk[4:5],
						},
						Tokens: tk[4:5],
					},
					Tokens: tk[4:5],
				},
				Tokens: tk[:5],
			}
		}},
		{"in", func(t *test, tk Tokens) { // 4
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: Error{
							Err: Error{
								Err: Error{
									Err: Error{
										Err:     ErrInvalidSimpleExpression,
										Parsing: "SimpleExpression",
										Token:   tk[0],
									},
									Parsing: "IndexOrCallExpression",
									Token:   tk[0],
								},
								Parsing: "ScopeExpression",
								Token:   tk[0],
							},
							Parsing: "SubsetExpression",
							Token:   tk[0],
						},
						Parsing: "ExponentiationExpression",
						Token:   tk[0],
					},
					Parsing: "UnaryExpression",
					Token:   tk[0],
				},
				Parsing: "SequenceExpression",
				Token:   tk[0],
			}
		}},
		{"a:in", func(t *test, tk Tokens) { // 5
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: Error{
							Err: Error{
								Err: Error{
									Err: Error{
										Err: Error{
											Err:     ErrInvalidSimpleExpression,
											Parsing: "SimpleExpression",
											Token:   tk[2],
										},
										Parsing: "IndexOrCallExpression",
										Token:   tk[2],
									},
									Parsing: "ScopeExpression",
									Token:   tk[2],
								},
								Parsing: "SubsetExpression",
								Token:   tk[2],
							},
							Parsing: "ExponentiationExpression",
							Token:   tk[2],
						},
						Parsing: "UnaryExpression",
						Token:   tk[2],
					},
					Parsing: "SequenceExpression",
					Token:   tk[2],
				},
				Parsing: "SequenceExpression",
				Token:   tk[2],
			}
		}},
	}, func(t *test) (Type, error) {
		var se SequenceExpression

		err := se.parse(&t.Tokens)

		return se, err
	})
}

func TestUnaryExpression(t *testing.T) {
	doTests(t, []sourceFn{
		{"a", func(t *test, tk Tokens) { // 1
			t.Output = UnaryExpression{
				ExponentiationExpression: ExponentiationExpression{
					SubsetExpression: SubsetExpression{
						ScopeExpression: ScopeExpression{
							IndexOrCallExpression: IndexOrCallExpression{
								SimpleExpression: &SimpleExpression{
									Identifier: &tk[0],
									Tokens:     tk[:1],
								},
								Tokens: tk[:1],
							},
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				Tokens: tk[:1],
			}
		}},
		{"+a", func(t *test, tk Tokens) { // 2
			t.Output = UnaryExpression{
				UnaryType: []UnaryType{
					UnaryAdd,
				},
				ExponentiationExpression: ExponentiationExpression{
					SubsetExpression: SubsetExpression{
						ScopeExpression: ScopeExpression{
							IndexOrCallExpression: IndexOrCallExpression{
								SimpleExpression: &SimpleExpression{
									Identifier: &tk[1],
									Tokens:     tk[1:2],
								},
								Tokens: tk[1:2],
							},
							Tokens: tk[1:2],
						},
						Tokens: tk[1:2],
					},
					Tokens: tk[1:2],
				},
				Tokens: tk[:2],
			}
		}},
		{"-a", func(t *test, tk Tokens) { // 3
			t.Output = UnaryExpression{
				UnaryType: []UnaryType{
					UnaryMinus,
				},
				ExponentiationExpression: ExponentiationExpression{
					SubsetExpression: SubsetExpression{
						ScopeExpression: ScopeExpression{
							IndexOrCallExpression: IndexOrCallExpression{
								SimpleExpression: &SimpleExpression{
									Identifier: &tk[1],
									Tokens:     tk[1:2],
								},
								Tokens: tk[1:2],
							},
							Tokens: tk[1:2],
						},
						Tokens: tk[1:2],
					},
					Tokens: tk[1:2],
				},
				Tokens: tk[:2],
			}
		}},
		{"+ a", func(t *test, tk Tokens) { // 4
			t.Output = UnaryExpression{
				UnaryType: []UnaryType{
					UnaryAdd,
				},
				ExponentiationExpression: ExponentiationExpression{
					SubsetExpression: SubsetExpression{
						ScopeExpression: ScopeExpression{
							IndexOrCallExpression: IndexOrCallExpression{
								SimpleExpression: &SimpleExpression{
									Identifier: &tk[2],
									Tokens:     tk[2:3],
								},
								Tokens: tk[2:3],
							},
							Tokens: tk[2:3],
						},
						Tokens: tk[2:3],
					},
					Tokens: tk[2:3],
				},
				Tokens: tk[:3],
			}
		}},
		{"+- + a", func(t *test, tk Tokens) { // 5
			t.Output = UnaryExpression{
				UnaryType: []UnaryType{
					UnaryAdd,
					UnaryMinus,
					UnaryAdd,
				},
				ExponentiationExpression: ExponentiationExpression{
					SubsetExpression: SubsetExpression{
						ScopeExpression: ScopeExpression{
							IndexOrCallExpression: IndexOrCallExpression{
								SimpleExpression: &SimpleExpression{
									Identifier: &tk[5],
									Tokens:     tk[5:6],
								},
								Tokens: tk[5:6],
							},
							Tokens: tk[5:6],
						},
						Tokens: tk[5:6],
					},
					Tokens: tk[5:6],
				},
				Tokens: tk[:6],
			}
		}},
		{"in", func(t *test, tk Tokens) { // 6
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: Error{
							Err: Error{
								Err: Error{
									Err:     ErrInvalidSimpleExpression,
									Parsing: "SimpleExpression",
									Token:   tk[0],
								},
								Parsing: "IndexOrCallExpression",
								Token:   tk[0],
							},
							Parsing: "ScopeExpression",
							Token:   tk[0],
						},
						Parsing: "SubsetExpression",
						Token:   tk[0],
					},
					Parsing: "ExponentiationExpression",
					Token:   tk[0],
				},
				Parsing: "UnaryExpression",
				Token:   tk[0],
			}
		}},
	}, func(t *test) (Type, error) {
		var ue UnaryExpression

		err := ue.parse(&t.Tokens)

		return ue, err
	})
}

func TestExponentiationExpression(t *testing.T) {
	doTests(t, []sourceFn{
		{"a", func(t *test, tk Tokens) { // 1
			t.Output = ExponentiationExpression{
				SubsetExpression: SubsetExpression{
					ScopeExpression: ScopeExpression{
						IndexOrCallExpression: IndexOrCallExpression{
							SimpleExpression: &SimpleExpression{
								Identifier: &tk[0],
								Tokens:     tk[:1],
							},
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				Tokens: tk[:1],
			}
		}},
		{"a^b", func(t *test, tk Tokens) { // 2
			t.Output = ExponentiationExpression{
				SubsetExpression: SubsetExpression{
					ScopeExpression: ScopeExpression{
						IndexOrCallExpression: IndexOrCallExpression{
							SimpleExpression: &SimpleExpression{
								Identifier: &tk[0],
								Tokens:     tk[:1],
							},
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				ExponentiationExpression: &ExponentiationExpression{
					SubsetExpression: SubsetExpression{
						ScopeExpression: ScopeExpression{
							IndexOrCallExpression: IndexOrCallExpression{
								SimpleExpression: &SimpleExpression{
									Identifier: &tk[2],
									Tokens:     tk[2:3],
								},
								Tokens: tk[2:3],
							},
							Tokens: tk[2:3],
						},
						Tokens: tk[2:3],
					},
					Tokens: tk[2:3],
				},
				Tokens: tk[:3],
			}
		}},
		{"a ^ b", func(t *test, tk Tokens) { // 3
			t.Output = ExponentiationExpression{
				SubsetExpression: SubsetExpression{
					ScopeExpression: ScopeExpression{
						IndexOrCallExpression: IndexOrCallExpression{
							SimpleExpression: &SimpleExpression{
								Identifier: &tk[0],
								Tokens:     tk[:1],
							},
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				ExponentiationExpression: &ExponentiationExpression{
					SubsetExpression: SubsetExpression{
						ScopeExpression: ScopeExpression{
							IndexOrCallExpression: IndexOrCallExpression{
								SimpleExpression: &SimpleExpression{
									Identifier: &tk[4],
									Tokens:     tk[4:5],
								},
								Tokens: tk[4:5],
							},
							Tokens: tk[4:5],
						},
						Tokens: tk[4:5],
					},
					Tokens: tk[4:5],
				},
				Tokens: tk[:5],
			}
		}},
		{"in", func(t *test, tk Tokens) { // 4
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: Error{
							Err: Error{
								Err:     ErrInvalidSimpleExpression,
								Parsing: "SimpleExpression",
								Token:   tk[0],
							},
							Parsing: "IndexOrCallExpression",
							Token:   tk[0],
						},
						Parsing: "ScopeExpression",
						Token:   tk[0],
					},
					Parsing: "SubsetExpression",
					Token:   tk[0],
				},
				Parsing: "ExponentiationExpression",
				Token:   tk[0],
			}
		}},
		{"a^in", func(t *test, tk Tokens) { // 5
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: Error{
							Err: Error{
								Err: Error{
									Err:     ErrInvalidSimpleExpression,
									Parsing: "SimpleExpression",
									Token:   tk[2],
								},
								Parsing: "IndexOrCallExpression",
								Token:   tk[2],
							},
							Parsing: "ScopeExpression",
							Token:   tk[2],
						},
						Parsing: "SubsetExpression",
						Token:   tk[2],
					},
					Parsing: "ExponentiationExpression",
					Token:   tk[2],
				},
				Parsing: "ExponentiationExpression",
				Token:   tk[2],
			}
		}},
	}, func(t *test) (Type, error) {
		var ee ExponentiationExpression

		err := ee.parse(&t.Tokens)

		return ee, err
	})
}

func TestSubsetExpression(t *testing.T) {
	doTests(t, []sourceFn{
		{"a", func(t *test, tk Tokens) { // 1
			t.Output = SubsetExpression{
				ScopeExpression: ScopeExpression{
					IndexOrCallExpression: IndexOrCallExpression{
						SimpleExpression: &SimpleExpression{
							Identifier: &tk[0],
							Tokens:     tk[:1],
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				Tokens: tk[:1],
			}
		}},
		{"a$b", func(t *test, tk Tokens) { // 2
			t.Output = SubsetExpression{
				SubsetType: SubsetList,
				ScopeExpression: ScopeExpression{
					IndexOrCallExpression: IndexOrCallExpression{
						SimpleExpression: &SimpleExpression{
							Identifier: &tk[0],
							Tokens:     tk[:1],
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				SubsetExpression: &SubsetExpression{
					ScopeExpression: ScopeExpression{
						IndexOrCallExpression: IndexOrCallExpression{
							SimpleExpression: &SimpleExpression{
								Identifier: &tk[2],
								Tokens:     tk[2:3],
							},
							Tokens: tk[2:3],
						},
						Tokens: tk[2:3],
					},
					Tokens: tk[2:3],
				},
				Tokens: tk[:3],
			}
		}},
		{"a $ b", func(t *test, tk Tokens) { // 3
			t.Output = SubsetExpression{
				SubsetType: SubsetList,
				ScopeExpression: ScopeExpression{
					IndexOrCallExpression: IndexOrCallExpression{
						SimpleExpression: &SimpleExpression{
							Identifier: &tk[0],
							Tokens:     tk[:1],
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				SubsetExpression: &SubsetExpression{
					ScopeExpression: ScopeExpression{
						IndexOrCallExpression: IndexOrCallExpression{
							SimpleExpression: &SimpleExpression{
								Identifier: &tk[4],
								Tokens:     tk[4:5],
							},
							Tokens: tk[4:5],
						},
						Tokens: tk[4:5],
					},
					Tokens: tk[4:5],
				},
				Tokens: tk[:5],
			}
		}},
		{"a@b", func(t *test, tk Tokens) { // 4
			t.Output = SubsetExpression{
				SubsetType: SubsetStructure,
				ScopeExpression: ScopeExpression{
					IndexOrCallExpression: IndexOrCallExpression{
						SimpleExpression: &SimpleExpression{
							Identifier: &tk[0],
							Tokens:     tk[:1],
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				SubsetExpression: &SubsetExpression{
					ScopeExpression: ScopeExpression{
						IndexOrCallExpression: IndexOrCallExpression{
							SimpleExpression: &SimpleExpression{
								Identifier: &tk[2],
								Tokens:     tk[2:3],
							},
							Tokens: tk[2:3],
						},
						Tokens: tk[2:3],
					},
					Tokens: tk[2:3],
				},
				Tokens: tk[:3],
			}
		}},
		{"a @ b", func(t *test, tk Tokens) { // 5
			t.Output = SubsetExpression{
				SubsetType: SubsetStructure,
				ScopeExpression: ScopeExpression{
					IndexOrCallExpression: IndexOrCallExpression{
						SimpleExpression: &SimpleExpression{
							Identifier: &tk[0],
							Tokens:     tk[:1],
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				SubsetExpression: &SubsetExpression{
					ScopeExpression: ScopeExpression{
						IndexOrCallExpression: IndexOrCallExpression{
							SimpleExpression: &SimpleExpression{
								Identifier: &tk[4],
								Tokens:     tk[4:5],
							},
							Tokens: tk[4:5],
						},
						Tokens: tk[4:5],
					},
					Tokens: tk[4:5],
				},
				Tokens: tk[:5],
			}
		}},
		{"in", func(t *test, tk Tokens) { // 6
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: Error{
							Err:     ErrInvalidSimpleExpression,
							Parsing: "SimpleExpression",
							Token:   tk[0],
						},
						Parsing: "IndexOrCallExpression",
						Token:   tk[0],
					},
					Parsing: "ScopeExpression",
					Token:   tk[0],
				},
				Parsing: "SubsetExpression",
				Token:   tk[0],
			}
		}},
		{"a$in", func(t *test, tk Tokens) { // 7
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: Error{
							Err: Error{
								Err:     ErrInvalidSimpleExpression,
								Parsing: "SimpleExpression",
								Token:   tk[2],
							},
							Parsing: "IndexOrCallExpression",
							Token:   tk[2],
						},
						Parsing: "ScopeExpression",
						Token:   tk[2],
					},
					Parsing: "SubsetExpression",
					Token:   tk[2],
				},
				Parsing: "SubsetExpression",
				Token:   tk[2],
			}
		}},
		{"a@in", func(t *test, tk Tokens) { // 8
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: Error{
							Err: Error{
								Err:     ErrInvalidSimpleExpression,
								Parsing: "SimpleExpression",
								Token:   tk[2],
							},
							Parsing: "IndexOrCallExpression",
							Token:   tk[2],
						},
						Parsing: "ScopeExpression",
						Token:   tk[2],
					},
					Parsing: "SubsetExpression",
					Token:   tk[2],
				},
				Parsing: "SubsetExpression",
				Token:   tk[2],
			}
		}},
	}, func(t *test) (Type, error) {
		var se SubsetExpression

		err := se.parse(&t.Tokens)

		return se, err
	})
}

func TestScopeExpression(t *testing.T) {
	doTests(t, []sourceFn{
		{"a", func(t *test, tk Tokens) { // 1
			t.Output = ScopeExpression{
				IndexOrCallExpression: IndexOrCallExpression{
					SimpleExpression: &SimpleExpression{
						Identifier: &tk[0],
						Tokens:     tk[:1],
					},
					Tokens: tk[:1],
				},
				Tokens: tk[:1],
			}
		}},
		{"a::b", func(t *test, tk Tokens) { // 2
			t.Output = ScopeExpression{
				IndexOrCallExpression: IndexOrCallExpression{
					SimpleExpression: &SimpleExpression{
						Identifier: &tk[0],
						Tokens:     tk[:1],
					},
					Tokens: tk[:1],
				},
				ScopeExpression: &ScopeExpression{
					IndexOrCallExpression: IndexOrCallExpression{
						SimpleExpression: &SimpleExpression{
							Identifier: &tk[2],
							Tokens:     tk[2:3],
						},
						Tokens: tk[2:3],
					},
					Tokens: tk[2:3],
				},
				Tokens: tk[:3],
			}
		}},
		{"a :: b", func(t *test, tk Tokens) { // 3
			t.Output = ScopeExpression{
				IndexOrCallExpression: IndexOrCallExpression{
					SimpleExpression: &SimpleExpression{
						Identifier: &tk[0],
						Tokens:     tk[:1],
					},
					Tokens: tk[:1],
				},
				ScopeExpression: &ScopeExpression{
					IndexOrCallExpression: IndexOrCallExpression{
						SimpleExpression: &SimpleExpression{
							Identifier: &tk[4],
							Tokens:     tk[4:5],
						},
						Tokens: tk[4:5],
					},
					Tokens: tk[4:5],
				},
				Tokens: tk[:5],
			}
		}},
		{"a::b::c", func(t *test, tk Tokens) { // 4
			t.Output = ScopeExpression{
				IndexOrCallExpression: IndexOrCallExpression{
					SimpleExpression: &SimpleExpression{
						Identifier: &tk[0],
						Tokens:     tk[:1],
					},
					Tokens: tk[:1],
				},
				ScopeExpression: &ScopeExpression{
					IndexOrCallExpression: IndexOrCallExpression{
						SimpleExpression: &SimpleExpression{
							Identifier: &tk[2],
							Tokens:     tk[2:3],
						},
						Tokens: tk[2:3],
					},
					ScopeExpression: &ScopeExpression{
						IndexOrCallExpression: IndexOrCallExpression{
							SimpleExpression: &SimpleExpression{
								Identifier: &tk[4],
								Tokens:     tk[4:5],
							},
							Tokens: tk[4:5],
						},
						Tokens: tk[4:5],
					},
					Tokens: tk[2:5],
				},
				Tokens: tk[:5],
			}
		}},
		{"in", func(t *test, tk Tokens) { // 5
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err:     ErrInvalidSimpleExpression,
						Parsing: "SimpleExpression",
						Token:   tk[0],
					},
					Parsing: "IndexOrCallExpression",
					Token:   tk[0],
				},
				Parsing: "ScopeExpression",
				Token:   tk[0],
			}
		}},
		{"a::in", func(t *test, tk Tokens) { // 6
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: Error{
							Err:     ErrInvalidSimpleExpression,
							Parsing: "SimpleExpression",
							Token:   tk[2],
						},
						Parsing: "IndexOrCallExpression",
						Token:   tk[2],
					},
					Parsing: "ScopeExpression",
					Token:   tk[2],
				},
				Parsing: "ScopeExpression",
				Token:   tk[2],
			}
		}},
	}, func(t *test) (Type, error) {
		var se ScopeExpression

		err := se.parse(&t.Tokens)

		return se, err
	})
}

func TestIndexOrCallExpression(t *testing.T) {
	doTests(t, []sourceFn{
		{"a[b]", func(t *test, tk Tokens) { // 1
			t.Output = IndexOrCallExpression{
				IndexOrCallExpression: &IndexOrCallExpression{
					SimpleExpression: &SimpleExpression{
						Identifier: &tk[0],
						Tokens:     tk[:1],
					},
					Tokens: tk[:1],
				},
				Index: &Index{
					Args: []QueryExpression{
						*WrapQuery(&SimpleExpression{
							Identifier: &tk[2],
							Tokens:     tk[2:3],
						}),
					},
					Tokens: tk[1:4],
				},
				Tokens: tk[:4],
			}
		}},
		{"a [b]", func(t *test, tk Tokens) { // 2
			t.Output = IndexOrCallExpression{
				IndexOrCallExpression: &IndexOrCallExpression{
					SimpleExpression: &SimpleExpression{
						Identifier: &tk[0],
						Tokens:     tk[:1],
					},
					Tokens: tk[:1],
				},
				Index: &Index{
					Args: []QueryExpression{
						*WrapQuery(&SimpleExpression{
							Identifier: &tk[3],
							Tokens:     tk[3:4],
						}),
					},
					Tokens: tk[2:5],
				},
				Tokens: tk[:5],
			}
		}},
		{"a[[b]]", func(t *test, tk Tokens) { // 3
			t.Output = IndexOrCallExpression{
				IndexOrCallExpression: &IndexOrCallExpression{
					SimpleExpression: &SimpleExpression{
						Identifier: &tk[0],
						Tokens:     tk[:1],
					},
					Tokens: tk[:1],
				},
				Index: &Index{
					Double: true,
					Args: []QueryExpression{
						*WrapQuery(&SimpleExpression{
							Identifier: &tk[2],
							Tokens:     tk[2:3],
						}),
					},
					Tokens: tk[1:4],
				},
				Tokens: tk[:4],
			}
		}},
		{"a [[b]]", func(t *test, tk Tokens) { // 4
			t.Output = IndexOrCallExpression{
				IndexOrCallExpression: &IndexOrCallExpression{
					SimpleExpression: &SimpleExpression{
						Identifier: &tk[0],
						Tokens:     tk[:1],
					},
					Tokens: tk[:1],
				},
				Index: &Index{
					Double: true,
					Args: []QueryExpression{
						*WrapQuery(&SimpleExpression{
							Identifier: &tk[3],
							Tokens:     tk[3:4],
						}),
					},
					Tokens: tk[2:5],
				},
				Tokens: tk[:5],
			}
		}},
		{"a(b)", func(t *test, tk Tokens) { // 5
			t.Output = IndexOrCallExpression{
				IndexOrCallExpression: &IndexOrCallExpression{
					SimpleExpression: &SimpleExpression{
						Identifier: &tk[0],
						Tokens:     tk[:1],
					},
					Tokens: tk[:1],
				},
				Call: &Call{
					Args: []Arg{
						{
							QueryExpression: WrapQuery(&SimpleExpression{
								Identifier: &tk[2],
								Tokens:     tk[2:3],
							}),
							Tokens: tk[2:3],
						},
					},
					Tokens: tk[1:4],
				},
				Tokens: tk[:4],
			}
		}},
		{"a (b)", func(t *test, tk Tokens) { // 6
			t.Output = IndexOrCallExpression{
				IndexOrCallExpression: &IndexOrCallExpression{
					SimpleExpression: &SimpleExpression{
						Identifier: &tk[0],
						Tokens:     tk[:1],
					},
					Tokens: tk[:1],
				},
				Call: &Call{
					Args: []Arg{
						{
							QueryExpression: WrapQuery(&SimpleExpression{
								Identifier: &tk[3],
								Tokens:     tk[3:4],
							}),
							Tokens: tk[3:4],
						},
					},
					Tokens: tk[2:5],
				},
				Tokens: tk[:5],
			}
		}},
		{"in", func(t *test, tk Tokens) { // 7
			t.Err = Error{
				Err: Error{
					Err:     ErrInvalidSimpleExpression,
					Parsing: "SimpleExpression",
					Token:   tk[0],
				},
				Parsing: "IndexOrCallExpression",
				Token:   tk[0],
			}
		}},
		{"a[in]", func(t *test, tk Tokens) { // 8
			t.Err = Error{
				Err: Error{
					Err: wrapQueryExpressionError(Error{
						Err:     ErrInvalidSimpleExpression,
						Parsing: "SimpleExpression",
						Token:   tk[2],
					}),
					Parsing: "Index",
					Token:   tk[2],
				},
				Parsing: "IndexOrCallExpression",
				Token:   tk[1],
			}
		}},
		{"a(in)", func(t *test, tk Tokens) { // 9
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: wrapQueryExpressionError(Error{
							Err:     ErrInvalidSimpleExpression,
							Parsing: "SimpleExpression",
							Token:   tk[2],
						}),
						Parsing: "Arg",
						Token:   tk[2],
					},
					Parsing: "Call",
					Token:   tk[2],
				},
				Parsing: "IndexOrCallExpression",
				Token:   tk[1],
			}
		}},
	}, func(t *test) (Type, error) {
		var ice IndexOrCallExpression

		err := ice.parse(&t.Tokens)

		return ice, err
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
				ParenthesizedExpression: &ParenthesizedExpression{
					Expression: Expression{
						QueryExpression: WrapQuery(&SimpleExpression{
							Identifier: &tk[1],
							Tokens:     tk[1:2],
						}),
						Tokens: tk[1:2],
					},
					Tokens: tk[:3],
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
		{"(a b)", func(t *test, tk Tokens) { // 15
			t.Err = Error{
				Err: Error{
					Err:     ErrMissingClosingParen,
					Parsing: "ParenthesizedExpression",
					Token:   tk[3],
				},
				Parsing: "SimpleExpression",
				Token:   tk[0],
			}
		}},
		{"(in)", func(t *test, tk Tokens) { // 16
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: wrapQueryExpressionError(Error{
							Err:     ErrInvalidSimpleExpression,
							Parsing: "SimpleExpression",
							Token:   tk[1],
						}),
						Parsing: "Expression",
						Token:   tk[1],
					},
					Parsing: "ParenthesizedExpression",
					Token:   tk[1],
				},
				Parsing: "SimpleExpression",
				Token:   tk[0],
			}
		}},
		{"{in}", func(t *test, tk Tokens) { // 17
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: wrapQueryExpressionError(Error{
							Err:     ErrInvalidSimpleExpression,
							Parsing: "SimpleExpression",
							Token:   tk[1],
						}),
						Parsing: "Expression",
						Token:   tk[1],
					},
					Parsing: "CompoundExpression",
					Token:   tk[1],
				},
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

func TestIndex(t *testing.T) {
	doTests(t, []sourceFn{
		{"[]", func(t *test, tk Tokens) { // 1
			t.Output = Index{
				Tokens: tk[:2],
			}
		}},
		{"[a]", func(t *test, tk Tokens) { // 2
			t.Output = Index{
				Args: []QueryExpression{
					*WrapQuery(&SimpleExpression{
						Identifier: &tk[1],
						Tokens:     tk[1:2],
					}),
				},
				Tokens: tk[:3],
			}
		}},
		{"[ a ]", func(t *test, tk Tokens) { // 3
			t.Output = Index{
				Args: []QueryExpression{
					*WrapQuery(&SimpleExpression{
						Identifier: &tk[2],
						Tokens:     tk[2:3],
					}),
				},
				Tokens: tk[:5],
			}
		}},
		{"[a,b]", func(t *test, tk Tokens) { // 4
			t.Output = Index{
				Args: []QueryExpression{
					*WrapQuery(&SimpleExpression{
						Identifier: &tk[1],
						Tokens:     tk[1:2],
					}),
					*WrapQuery(&SimpleExpression{
						Identifier: &tk[3],
						Tokens:     tk[3:4],
					}),
				},
				Tokens: tk[:5],
			}
		}},
		{"[a , b]", func(t *test, tk Tokens) { // 5
			t.Output = Index{
				Args: []QueryExpression{
					*WrapQuery(&SimpleExpression{
						Identifier: &tk[1],
						Tokens:     tk[1:2],
					}),
					*WrapQuery(&SimpleExpression{
						Identifier: &tk[5],
						Tokens:     tk[5:6],
					}),
				},
				Tokens: tk[:7],
			}
		}},
		{"[[a]]", func(t *test, tk Tokens) { // 6
			t.Output = Index{
				Double: true,
				Args: []QueryExpression{
					*WrapQuery(&SimpleExpression{
						Identifier: &tk[1],
						Tokens:     tk[1:2],
					}),
				},
				Tokens: tk[:3],
			}
		}},
		{"[[ a ]]", func(t *test, tk Tokens) { // 7
			t.Output = Index{
				Double: true,
				Args: []QueryExpression{
					*WrapQuery(&SimpleExpression{
						Identifier: &tk[2],
						Tokens:     tk[2:3],
					}),
				},
				Tokens: tk[:5],
			}
		}},
		{"[in]", func(t *test, tk Tokens) { // 8
			t.Err = Error{
				Err: wrapQueryExpressionError(Error{
					Err:     ErrInvalidSimpleExpression,
					Parsing: "SimpleExpression",
					Token:   tk[1],
				}),
				Parsing: "Index",
				Token:   tk[1],
			}
		}},
		{"[a b]", func(t *test, tk Tokens) { // 9
			t.Err = Error{
				Err:     ErrMissingComma,
				Parsing: "Index",
				Token:   tk[3],
			}
		}},
		{"[[in]]", func(t *test, tk Tokens) { // 10
			t.Err = Error{
				Err: wrapQueryExpressionError(Error{
					Err:     ErrInvalidSimpleExpression,
					Parsing: "SimpleExpression",
					Token:   tk[1],
				}),
				Parsing: "Index",
				Token:   tk[1],
			}
		}},
		{"[[a b]]", func(t *test, tk Tokens) { // 11
			t.Err = Error{
				Err:     ErrMissingClosingDoubleBracket,
				Parsing: "Index",
				Token:   tk[3],
			}
		}},
	}, func(t *test) (Type, error) {
		var i Index

		err := i.parse(&t.Tokens)

		return i, err
	})
}

func TestCall(t *testing.T) {
	doTests(t, []sourceFn{
		{"()", func(t *test, tk Tokens) { // 1
			t.Output = Call{
				Tokens: tk[:2],
			}
		}},
		{"(a)", func(t *test, tk Tokens) { // 2
			t.Output = Call{
				Args: []Arg{
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
		{"( a )", func(t *test, tk Tokens) { // 3
			t.Output = Call{
				Args: []Arg{
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
		{"(a,b)", func(t *test, tk Tokens) { // 4
			t.Output = Call{
				Args: []Arg{
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
		{"(a , b)", func(t *test, tk Tokens) { // 5
			t.Output = Call{
				Args: []Arg{
					{
						QueryExpression: WrapQuery(&SimpleExpression{
							Identifier: &tk[1],
							Tokens:     tk[1:2],
						}),
						Tokens: tk[1:2],
					},
					{
						QueryExpression: WrapQuery(&SimpleExpression{
							Identifier: &tk[5],
							Tokens:     tk[5:6],
						}),
						Tokens: tk[5:6],
					},
				},
				Tokens: tk[:7],
			}
		}},
		{"(,)", func(t *test, tk Tokens) { // 6
			t.Output = Call{
				Args:   []Arg{},
				Tokens: tk[:3],
			}
		}},
		{"(a,)", func(t *test, tk Tokens) { // 7
			t.Output = Call{
				Args: []Arg{
					{
						QueryExpression: WrapQuery(&SimpleExpression{
							Identifier: &tk[1],
							Tokens:     tk[1:2],
						}),
						Tokens: tk[1:2],
					},
				},
				Tokens: tk[:4],
			}
		}},
		{"(,a)", func(t *test, tk Tokens) { // 8
			t.Output = Call{
				Args: []Arg{
					{
						Tokens: tk[1:1],
					},
					{
						QueryExpression: WrapQuery(&SimpleExpression{
							Identifier: &tk[2],
							Tokens:     tk[2:3],
						}),
						Tokens: tk[2:3],
					},
				},
				Tokens: tk[:4],
			}
		}},
		{"(,,)", func(t *test, tk Tokens) { // 9
			t.Output = Call{
				Args:   []Arg{},
				Tokens: tk[:4],
			}
		}},
		{"(a,,)", func(t *test, tk Tokens) { // 10
			t.Output = Call{
				Args: []Arg{
					{
						QueryExpression: WrapQuery(&SimpleExpression{
							Identifier: &tk[1],
							Tokens:     tk[1:2],
						}),
						Tokens: tk[1:2],
					},
				},
				Tokens: tk[:5],
			}
		}},
		{"(,,a)", func(t *test, tk Tokens) { // 11
			t.Output = Call{
				Args: []Arg{
					{
						Tokens: tk[1:1],
					},
					{
						Tokens: tk[2:2],
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
		{"(,a,)", func(t *test, tk Tokens) { // 12
			t.Output = Call{
				Args: []Arg{
					{
						Tokens: tk[1:1],
					},
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
		{"(a,,b)", func(t *test, tk Tokens) { // 13
			t.Output = Call{
				Args: []Arg{
					{
						QueryExpression: WrapQuery(&SimpleExpression{
							Identifier: &tk[1],
							Tokens:     tk[1:2],
						}),
						Tokens: tk[1:2],
					},
					{
						Tokens: tk[3:3],
					},
					{
						QueryExpression: WrapQuery(&SimpleExpression{
							Identifier: &tk[4],
							Tokens:     tk[4:5],
						}),
						Tokens: tk[4:5],
					},
				},
				Tokens: tk[:6],
			}
		}},
		{"(in)", func(t *test, tk Tokens) { // 14
			t.Err = Error{
				Err: Error{
					Err: wrapQueryExpressionError(Error{
						Err:     ErrInvalidSimpleExpression,
						Parsing: "SimpleExpression",
						Token:   tk[1],
					}),
					Parsing: "Arg",
					Token:   tk[1],
				},
				Parsing: "Call",
				Token:   tk[1],
			}
		}},
		{"(a b)", func(t *test, tk Tokens) { // 15
			t.Err = Error{
				Err:     ErrMissingComma,
				Parsing: "Call",
				Token:   tk[3],
			}
		}},
	}, func(t *test) (Type, error) {
		var c Call

		err := c.parse(&t.Tokens)

		return c, err
	})
}

func TestArg(t *testing.T) {
	doTests(t, []sourceFn{
		{"a", func(t *test, tk Tokens) { // 1
			t.Output = Arg{
				QueryExpression: WrapQuery(&SimpleExpression{
					Identifier: &tk[0],
					Tokens:     tk[:1],
				}),
				Tokens: tk[:1],
			}
		}},
		{"...", func(t *test, tk Tokens) { // 2
			t.Output = Arg{
				Ellipsis: &tk[0],
				Tokens:   tk[:1],
			}
		}},
		{"", func(t *test, tk Tokens) { // 3
			t.Output = Arg{
				Tokens: tk[:0],
			}
		}},
		{"in", func(t *test, tk Tokens) { // 4
			t.Err = Error{
				Err: wrapQueryExpressionError(Error{
					Err:     ErrInvalidSimpleExpression,
					Parsing: "SimpleExpression",
					Token:   tk[0],
				}),
				Parsing: "Arg",
				Token:   tk[0],
			}
		}},
	}, func(t *test) (Type, error) {
		var a Arg

		err := a.parse(&t.Tokens)

		return a, err
	})
}
