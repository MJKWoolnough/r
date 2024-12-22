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
		{"(a b)", func(t *test, tk Tokens) { // 15
			t.Err = Error{
				Err:     ErrMissingClosingParen,
				Parsing: "SimpleExpression",
				Token:   tk[3],
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
		{"(in)", func(t *test, tk Tokens) { // 6
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
		{"(a b)", func(t *test, tk Tokens) { // 7
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
		{"in", func(t *test, tk Tokens) { // 3
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
