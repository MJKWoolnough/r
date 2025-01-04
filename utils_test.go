package r

import (
	"reflect"
	"testing"

	"vimagination.zapto.org/parser"
)

func TestWrapQuery(t *testing.T) {
	tks := Tokens{
		{
			Token: parser.Token{
				Type: TokenGrouping,
				Data: "{",
			},
		},
		{
			Token: parser.Token{
				Type: TokenIdentifier,
				Data: "a",
			},
		},
		{
			Token: parser.Token{
				Type: TokenGrouping,
				Data: "}",
			},
		},
	}
	compound := &CompoundExpression{
		Expressions: []Expression{
			{
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
																					Identifier: &tks[1],
																					Tokens:     tks[1:2],
																				},
																				Tokens: tks[1:2],
																			},
																			Tokens: tks[1:2],
																		},
																		Tokens: tks[1:2],
																	},
																	Tokens: tks[1:2],
																},
																Tokens: tks[1:2],
															},
															Tokens: tks[1:2],
														},
														Tokens: tks[1:2],
													},
													Tokens: tks[1:2],
												},
												Tokens: tks[1:2],
											},
											Tokens: tks[1:2],
										},
										Tokens: tks[1:2],
									},
									Tokens: tks[1:2],
								},
								Tokens: tks[1:2],
							},
							Tokens: tks[1:2],
						},
						Tokens: tks[1:2],
					},
					Tokens: tks[1:2],
				},
				Tokens: tks[1:2],
			},
		},
		Tokens: tks,
	}
	expectedOutput := QueryExpression{
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
																		CompoundExpression: compound,
																		Tokens:             tks,
																	},
																	Tokens: tks,
																},
																Tokens: tks,
															},
															Tokens: tks,
														},
														Tokens: tks,
													},
													Tokens: tks,
												},
												Tokens: tks,
											},
											Tokens: tks,
										},
										Tokens: tks,
									},
									Tokens: tks,
								},
								Tokens: tks,
							},
							Tokens: tks,
						},
						Tokens: tks,
					},
					Tokens: tks,
				},
				Tokens: tks,
			},
			Tokens: tks,
		},
		Tokens: tks,
	}

	for n, test := range [...]QueryWrappable{
		compound,  // 1
		*compound, // 2
		&SimpleExpression{ // 3
			CompoundExpression: compound,
			Tokens:             tks,
		},
		SimpleExpression{ // 4
			CompoundExpression: compound,
			Tokens:             tks,
		},
		&IndexOrCallExpression{ // 5
			SimpleExpression: &SimpleExpression{
				CompoundExpression: compound,
				Tokens:             tks,
			},
			Tokens: tks,
		},
		IndexOrCallExpression{ // 6
			SimpleExpression: &SimpleExpression{
				CompoundExpression: compound,
				Tokens:             tks,
			},
			Tokens: tks,
		},
		&ScopeExpression{ // 7
			IndexOrCallExpression: IndexOrCallExpression{
				SimpleExpression: &SimpleExpression{
					CompoundExpression: compound,
					Tokens:             tks,
				},
				Tokens: tks,
			},
			Tokens: tks,
		},
		ScopeExpression{ // 8
			IndexOrCallExpression: IndexOrCallExpression{
				SimpleExpression: &SimpleExpression{
					CompoundExpression: compound,
					Tokens:             tks,
				},
				Tokens: tks,
			},
			Tokens: tks,
		},
		&SubsetExpression{ // 9
			ScopeExpression: ScopeExpression{
				IndexOrCallExpression: IndexOrCallExpression{
					SimpleExpression: &SimpleExpression{
						CompoundExpression: compound,
						Tokens:             tks,
					},
					Tokens: tks,
				},
				Tokens: tks,
			},
			Tokens: tks,
		},
		SubsetExpression{ // 10
			ScopeExpression: ScopeExpression{
				IndexOrCallExpression: IndexOrCallExpression{
					SimpleExpression: &SimpleExpression{
						CompoundExpression: compound,
						Tokens:             tks,
					},
					Tokens: tks,
				},
				Tokens: tks,
			},
			Tokens: tks,
		},
		&ExponentiationExpression{ // 11
			SubsetExpression: SubsetExpression{
				ScopeExpression: ScopeExpression{
					IndexOrCallExpression: IndexOrCallExpression{
						SimpleExpression: &SimpleExpression{
							CompoundExpression: compound,
							Tokens:             tks,
						},
						Tokens: tks,
					},
					Tokens: tks,
				},
				Tokens: tks,
			},
			Tokens: tks,
		},
		ExponentiationExpression{ // 12
			SubsetExpression: SubsetExpression{
				ScopeExpression: ScopeExpression{
					IndexOrCallExpression: IndexOrCallExpression{
						SimpleExpression: &SimpleExpression{
							CompoundExpression: compound,
							Tokens:             tks,
						},
						Tokens: tks,
					},
					Tokens: tks,
				},
				Tokens: tks,
			},
			Tokens: tks,
		},
		&UnaryExpression{ // 13
			ExponentiationExpression: ExponentiationExpression{
				SubsetExpression: SubsetExpression{
					ScopeExpression: ScopeExpression{
						IndexOrCallExpression: IndexOrCallExpression{
							SimpleExpression: &SimpleExpression{
								CompoundExpression: compound,
								Tokens:             tks,
							},
							Tokens: tks,
						},
						Tokens: tks,
					},
					Tokens: tks,
				},
				Tokens: tks,
			},
			Tokens: tks,
		},
		UnaryExpression{ // 14
			ExponentiationExpression: ExponentiationExpression{
				SubsetExpression: SubsetExpression{
					ScopeExpression: ScopeExpression{
						IndexOrCallExpression: IndexOrCallExpression{
							SimpleExpression: &SimpleExpression{
								CompoundExpression: compound,
								Tokens:             tks,
							},
							Tokens: tks,
						},
						Tokens: tks,
					},
					Tokens: tks,
				},
				Tokens: tks,
			},
			Tokens: tks,
		},
		&SequenceExpression{ // 15
			UnaryExpression: UnaryExpression{
				ExponentiationExpression: ExponentiationExpression{
					SubsetExpression: SubsetExpression{
						ScopeExpression: ScopeExpression{
							IndexOrCallExpression: IndexOrCallExpression{
								SimpleExpression: &SimpleExpression{
									CompoundExpression: compound,
									Tokens:             tks,
								},
								Tokens: tks,
							},
							Tokens: tks,
						},
						Tokens: tks,
					},
					Tokens: tks,
				},
				Tokens: tks,
			},
			Tokens: tks,
		},
		SequenceExpression{ // 16
			UnaryExpression: UnaryExpression{
				ExponentiationExpression: ExponentiationExpression{
					SubsetExpression: SubsetExpression{
						ScopeExpression: ScopeExpression{
							IndexOrCallExpression: IndexOrCallExpression{
								SimpleExpression: &SimpleExpression{
									CompoundExpression: compound,
									Tokens:             tks,
								},
								Tokens: tks,
							},
							Tokens: tks,
						},
						Tokens: tks,
					},
					Tokens: tks,
				},
				Tokens: tks,
			},
			Tokens: tks,
		},
		&PipeOrSpecialExpression{ // 17
			SequenceExpression: SequenceExpression{
				UnaryExpression: UnaryExpression{
					ExponentiationExpression: ExponentiationExpression{
						SubsetExpression: SubsetExpression{
							ScopeExpression: ScopeExpression{
								IndexOrCallExpression: IndexOrCallExpression{
									SimpleExpression: &SimpleExpression{
										CompoundExpression: compound,
										Tokens:             tks,
									},
									Tokens: tks,
								},
								Tokens: tks,
							},
							Tokens: tks,
						},
						Tokens: tks,
					},
					Tokens: tks,
				},
				Tokens: tks,
			},
			Tokens: tks,
		},
		PipeOrSpecialExpression{ // 18
			SequenceExpression: SequenceExpression{
				UnaryExpression: UnaryExpression{
					ExponentiationExpression: ExponentiationExpression{
						SubsetExpression: SubsetExpression{
							ScopeExpression: ScopeExpression{
								IndexOrCallExpression: IndexOrCallExpression{
									SimpleExpression: &SimpleExpression{
										CompoundExpression: compound,
										Tokens:             tks,
									},
									Tokens: tks,
								},
								Tokens: tks,
							},
							Tokens: tks,
						},
						Tokens: tks,
					},
					Tokens: tks,
				},
				Tokens: tks,
			},
			Tokens: tks,
		},
		&MultiplicationExpression{ // 19
			PipeOrSpecialExpression: PipeOrSpecialExpression{
				SequenceExpression: SequenceExpression{
					UnaryExpression: UnaryExpression{
						ExponentiationExpression: ExponentiationExpression{
							SubsetExpression: SubsetExpression{
								ScopeExpression: ScopeExpression{
									IndexOrCallExpression: IndexOrCallExpression{
										SimpleExpression: &SimpleExpression{
											CompoundExpression: compound,
											Tokens:             tks,
										},
										Tokens: tks,
									},
									Tokens: tks,
								},
								Tokens: tks,
							},
							Tokens: tks,
						},
						Tokens: tks,
					},
					Tokens: tks,
				},
				Tokens: tks,
			},
			Tokens: tks,
		},
		MultiplicationExpression{ // 20
			PipeOrSpecialExpression: PipeOrSpecialExpression{
				SequenceExpression: SequenceExpression{
					UnaryExpression: UnaryExpression{
						ExponentiationExpression: ExponentiationExpression{
							SubsetExpression: SubsetExpression{
								ScopeExpression: ScopeExpression{
									IndexOrCallExpression: IndexOrCallExpression{
										SimpleExpression: &SimpleExpression{
											CompoundExpression: compound,
											Tokens:             tks,
										},
										Tokens: tks,
									},
									Tokens: tks,
								},
								Tokens: tks,
							},
							Tokens: tks,
						},
						Tokens: tks,
					},
					Tokens: tks,
				},
				Tokens: tks,
			},
			Tokens: tks,
		},
		&AdditionExpression{ // 21
			MultiplicationExpression: MultiplicationExpression{
				PipeOrSpecialExpression: PipeOrSpecialExpression{
					SequenceExpression: SequenceExpression{
						UnaryExpression: UnaryExpression{
							ExponentiationExpression: ExponentiationExpression{
								SubsetExpression: SubsetExpression{
									ScopeExpression: ScopeExpression{
										IndexOrCallExpression: IndexOrCallExpression{
											SimpleExpression: &SimpleExpression{
												CompoundExpression: compound,
												Tokens:             tks,
											},
											Tokens: tks,
										},
										Tokens: tks,
									},
									Tokens: tks,
								},
								Tokens: tks,
							},
							Tokens: tks,
						},
						Tokens: tks,
					},
					Tokens: tks,
				},
				Tokens: tks,
			},
			Tokens: tks,
		},
		AdditionExpression{ // 22
			MultiplicationExpression: MultiplicationExpression{
				PipeOrSpecialExpression: PipeOrSpecialExpression{
					SequenceExpression: SequenceExpression{
						UnaryExpression: UnaryExpression{
							ExponentiationExpression: ExponentiationExpression{
								SubsetExpression: SubsetExpression{
									ScopeExpression: ScopeExpression{
										IndexOrCallExpression: IndexOrCallExpression{
											SimpleExpression: &SimpleExpression{
												CompoundExpression: compound,
												Tokens:             tks,
											},
											Tokens: tks,
										},
										Tokens: tks,
									},
									Tokens: tks,
								},
								Tokens: tks,
							},
							Tokens: tks,
						},
						Tokens: tks,
					},
					Tokens: tks,
				},
				Tokens: tks,
			},
			Tokens: tks,
		},
		&RelationalExpression{ // 23
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
													CompoundExpression: compound,
													Tokens:             tks,
												},
												Tokens: tks,
											},
											Tokens: tks,
										},
										Tokens: tks,
									},
									Tokens: tks,
								},
								Tokens: tks,
							},
							Tokens: tks,
						},
						Tokens: tks,
					},
					Tokens: tks,
				},
				Tokens: tks,
			},
			Tokens: tks,
		},
		RelationalExpression{ // 24
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
													CompoundExpression: compound,
													Tokens:             tks,
												},
												Tokens: tks,
											},
											Tokens: tks,
										},
										Tokens: tks,
									},
									Tokens: tks,
								},
								Tokens: tks,
							},
							Tokens: tks,
						},
						Tokens: tks,
					},
					Tokens: tks,
				},
				Tokens: tks,
			},
			Tokens: tks,
		},
		&NotExpression{ // 25
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
														CompoundExpression: compound,
														Tokens:             tks,
													},
													Tokens: tks,
												},
												Tokens: tks,
											},
											Tokens: tks,
										},
										Tokens: tks,
									},
									Tokens: tks,
								},
								Tokens: tks,
							},
							Tokens: tks,
						},
						Tokens: tks,
					},
					Tokens: tks,
				},
				Tokens: tks,
			},
			Tokens: tks,
		},
		NotExpression{ // 26
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
														CompoundExpression: compound,
														Tokens:             tks,
													},
													Tokens: tks,
												},
												Tokens: tks,
											},
											Tokens: tks,
										},
										Tokens: tks,
									},
									Tokens: tks,
								},
								Tokens: tks,
							},
							Tokens: tks,
						},
						Tokens: tks,
					},
					Tokens: tks,
				},
				Tokens: tks,
			},
			Tokens: tks,
		},
		&AndExpression{ // 27
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
															CompoundExpression: compound,
															Tokens:             tks,
														},
														Tokens: tks,
													},
													Tokens: tks,
												},
												Tokens: tks,
											},
											Tokens: tks,
										},
										Tokens: tks,
									},
									Tokens: tks,
								},
								Tokens: tks,
							},
							Tokens: tks,
						},
						Tokens: tks,
					},
					Tokens: tks,
				},
				Tokens: tks,
			},
			Tokens: tks,
		},
		AndExpression{ // 28
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
															CompoundExpression: compound,
															Tokens:             tks,
														},
														Tokens: tks,
													},
													Tokens: tks,
												},
												Tokens: tks,
											},
											Tokens: tks,
										},
										Tokens: tks,
									},
									Tokens: tks,
								},
								Tokens: tks,
							},
							Tokens: tks,
						},
						Tokens: tks,
					},
					Tokens: tks,
				},
				Tokens: tks,
			},
			Tokens: tks,
		},
		&OrExpression{ // 29
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
																CompoundExpression: compound,
																Tokens:             tks,
															},
															Tokens: tks,
														},
														Tokens: tks,
													},
													Tokens: tks,
												},
												Tokens: tks,
											},
											Tokens: tks,
										},
										Tokens: tks,
									},
									Tokens: tks,
								},
								Tokens: tks,
							},
							Tokens: tks,
						},
						Tokens: tks,
					},
					Tokens: tks,
				},
				Tokens: tks,
			},
			Tokens: tks,
		},
		OrExpression{ // 30
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
																CompoundExpression: compound,
																Tokens:             tks,
															},
															Tokens: tks,
														},
														Tokens: tks,
													},
													Tokens: tks,
												},
												Tokens: tks,
											},
											Tokens: tks,
										},
										Tokens: tks,
									},
									Tokens: tks,
								},
								Tokens: tks,
							},
							Tokens: tks,
						},
						Tokens: tks,
					},
					Tokens: tks,
				},
				Tokens: tks,
			},
			Tokens: tks,
		},
		&FormulaeExpression{ // 31
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
																	CompoundExpression: compound,
																	Tokens:             tks,
																},
																Tokens: tks,
															},
															Tokens: tks,
														},
														Tokens: tks,
													},
													Tokens: tks,
												},
												Tokens: tks,
											},
											Tokens: tks,
										},
										Tokens: tks,
									},
									Tokens: tks,
								},
								Tokens: tks,
							},
							Tokens: tks,
						},
						Tokens: tks,
					},
					Tokens: tks,
				},
				Tokens: tks,
			},
			Tokens: tks,
		},
		FormulaeExpression{ // 32
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
																	CompoundExpression: compound,
																	Tokens:             tks,
																},
																Tokens: tks,
															},
															Tokens: tks,
														},
														Tokens: tks,
													},
													Tokens: tks,
												},
												Tokens: tks,
											},
											Tokens: tks,
										},
										Tokens: tks,
									},
									Tokens: tks,
								},
								Tokens: tks,
							},
							Tokens: tks,
						},
						Tokens: tks,
					},
					Tokens: tks,
				},
				Tokens: tks,
			},
			Tokens: tks,
		},
		&AssignmentExpression{ // 33
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
																		CompoundExpression: compound,
																		Tokens:             tks,
																	},
																	Tokens: tks,
																},
																Tokens: tks,
															},
															Tokens: tks,
														},
														Tokens: tks,
													},
													Tokens: tks,
												},
												Tokens: tks,
											},
											Tokens: tks,
										},
										Tokens: tks,
									},
									Tokens: tks,
								},
								Tokens: tks,
							},
							Tokens: tks,
						},
						Tokens: tks,
					},
					Tokens: tks,
				},
				Tokens: tks,
			},
			Tokens: tks,
		},
		AssignmentExpression{ // 34
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
																		CompoundExpression: compound,
																		Tokens:             tks,
																	},
																	Tokens: tks,
																},
																Tokens: tks,
															},
															Tokens: tks,
														},
														Tokens: tks,
													},
													Tokens: tks,
												},
												Tokens: tks,
											},
											Tokens: tks,
										},
										Tokens: tks,
									},
									Tokens: tks,
								},
								Tokens: tks,
							},
							Tokens: tks,
						},
						Tokens: tks,
					},
					Tokens: tks,
				},
				Tokens: tks,
			},
			Tokens: tks,
		},
		&QueryExpression{ // 35
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
																			CompoundExpression: compound,
																			Tokens:             tks,
																		},
																		Tokens: tks,
																	},
																	Tokens: tks,
																},
																Tokens: tks,
															},
															Tokens: tks,
														},
														Tokens: tks,
													},
													Tokens: tks,
												},
												Tokens: tks,
											},
											Tokens: tks,
										},
										Tokens: tks,
									},
									Tokens: tks,
								},
								Tokens: tks,
							},
							Tokens: tks,
						},
						Tokens: tks,
					},
					Tokens: tks,
				},
				Tokens: tks,
			},
			Tokens: tks,
		},
		QueryExpression{ // 36
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
																			CompoundExpression: compound,
																			Tokens:             tks,
																		},
																		Tokens: tks,
																	},
																	Tokens: tks,
																},
																Tokens: tks,
															},
															Tokens: tks,
														},
														Tokens: tks,
													},
													Tokens: tks,
												},
												Tokens: tks,
											},
											Tokens: tks,
										},
										Tokens: tks,
									},
									Tokens: tks,
								},
								Tokens: tks,
							},
							Tokens: tks,
						},
						Tokens: tks,
					},
					Tokens: tks,
				},
				Tokens: tks,
			},
			Tokens: tks,
		},
	} {
		if output := WrapQuery(test); !reflect.DeepEqual(output, &expectedOutput) {
			t.Errorf("test %d: expecting\n%v\n...got...\n%v", n+1, expectedOutput, output)
		}
	}
}

func TestUnwrapQuery(t *testing.T) {
	tks := Tokens{
		{
			Token: parser.Token{
				Type: TokenIdentifier,
				Data: "a",
			},
		},
		{
			Token: parser.Token{
				Type: TokenIdentifier,
				Data: "b",
			},
		},
	}
	identA := &tks[0]
	identB := &tks[1]

	for n, test := range [...]QueryWrappable{
		&CompoundExpression{ // 1
			Expressions: []Expression{
				{
					QueryExpression: WrapQuery(&SimpleExpression{
						Identifier: identA,
						Tokens:     tks[:1],
					}),
					Tokens: tks[:1],
				},
			},
			Tokens: tks[:1],
		},
		&SimpleExpression{ // 2
			Identifier: identA,
			Tokens:     tks[:1],
		},
		&IndexOrCallExpression{ // 3
			IndexOrCallExpression: &IndexOrCallExpression{
				SimpleExpression: &SimpleExpression{
					Identifier: identA,
					Tokens:     tks[:1],
				},
			},
			Call: &Call{
				Args: []Arg{
					{
						QueryExpression: WrapQuery(&SimpleExpression{
							Identifier: identB,
							Tokens:     tks[1:2],
						}),
						Tokens: tks[1:2],
					},
				},
				Tokens: tks[1:2],
			},
			Tokens: tks[:2],
		},
		&IndexOrCallExpression{ // 4
			IndexOrCallExpression: &IndexOrCallExpression{
				SimpleExpression: &SimpleExpression{
					Identifier: identA,
					Tokens:     tks[:1],
				},
			},
			Index: &Index{
				Args: []QueryExpression{
					*WrapQuery(&SimpleExpression{
						Identifier: identB,
						Tokens:     tks[1:2],
					}),
				},
				Tokens: tks[1:2],
			},
			Tokens: tks[:2],
		},
		&ScopeExpression{ // 5
			IndexOrCallExpression: IndexOrCallExpression{
				SimpleExpression: &SimpleExpression{
					Identifier: identA,
					Tokens:     tks[:1],
				},
				Tokens: tks[:1],
			},
			ScopeExpression: &ScopeExpression{
				IndexOrCallExpression: IndexOrCallExpression{
					SimpleExpression: &SimpleExpression{
						Identifier: identB,
						Tokens:     tks[1:2],
					},
					Tokens: tks[1:2],
				},
				Tokens: tks[1:2],
			},
			Tokens: tks[:2],
		},
		&SubsetExpression{ // 6
			ScopeExpression: ScopeExpression{
				IndexOrCallExpression: IndexOrCallExpression{
					SimpleExpression: &SimpleExpression{
						Identifier: identA,
						Tokens:     tks[:1],
					},
					Tokens: tks[:1],
				},
				Tokens: tks[:1],
			},
			SubsetExpression: &SubsetExpression{
				ScopeExpression: ScopeExpression{
					IndexOrCallExpression: IndexOrCallExpression{
						SimpleExpression: &SimpleExpression{
							Identifier: identB,
							Tokens:     tks[1:2],
						},
						Tokens: tks[1:2],
					},
					Tokens: tks[1:2],
				},
				Tokens: tks[1:2],
			},
			Tokens: tks[:2],
		},
		&ExponentiationExpression{ // 7
			SubsetExpression: SubsetExpression{
				ScopeExpression: ScopeExpression{
					IndexOrCallExpression: IndexOrCallExpression{
						SimpleExpression: &SimpleExpression{
							Identifier: identA,
							Tokens:     tks[:1],
						},
						Tokens: tks[:1],
					},
					Tokens: tks[:1],
				},
				Tokens: tks[:1],
			},
			ExponentiationExpression: &ExponentiationExpression{
				SubsetExpression: SubsetExpression{
					ScopeExpression: ScopeExpression{
						IndexOrCallExpression: IndexOrCallExpression{
							SimpleExpression: &SimpleExpression{
								Identifier: identB,
								Tokens:     tks[1:2],
							},
							Tokens: tks[1:2],
						},
						Tokens: tks[1:2],
					},
					Tokens: tks[1:2],
				},
				Tokens: tks[1:2],
			},
			Tokens: tks[:2],
		},
		&UnaryExpression{ // 8
			UnaryType: []UnaryType{UnaryAdd},
			ExponentiationExpression: ExponentiationExpression{
				SubsetExpression: SubsetExpression{
					ScopeExpression: ScopeExpression{
						IndexOrCallExpression: IndexOrCallExpression{
							SimpleExpression: &SimpleExpression{
								Identifier: identA,
								Tokens:     tks[:1],
							},
							Tokens: tks[:1],
						},
						Tokens: tks[:1],
					},
					Tokens: tks[:1],
				},
				Tokens: tks[:1],
			},
			Tokens: tks[:1],
		},
		&SequenceExpression{ // 9
			UnaryExpression: UnaryExpression{
				ExponentiationExpression: ExponentiationExpression{
					SubsetExpression: SubsetExpression{
						ScopeExpression: ScopeExpression{
							IndexOrCallExpression: IndexOrCallExpression{
								SimpleExpression: &SimpleExpression{
									Identifier: identA,
									Tokens:     tks[:1],
								},
								Tokens: tks[:1],
							},
							Tokens: tks[:1],
						},
						Tokens: tks[:1],
					},
					Tokens: tks[:1],
				},
				Tokens: tks[:1],
			},
			SequenceExpression: &SequenceExpression{
				UnaryExpression: UnaryExpression{
					ExponentiationExpression: ExponentiationExpression{
						SubsetExpression: SubsetExpression{
							ScopeExpression: ScopeExpression{
								IndexOrCallExpression: IndexOrCallExpression{
									SimpleExpression: &SimpleExpression{
										Identifier: identB,
										Tokens:     tks[1:2],
									},
									Tokens: tks[1:2],
								},
								Tokens: tks[1:2],
							},
							Tokens: tks[1:2],
						},
						Tokens: tks[1:2],
					},
					Tokens: tks[1:2],
				},
				Tokens: tks[:2],
			},
		},
		&PipeOrSpecialExpression{ // 10
			SequenceExpression: SequenceExpression{
				UnaryExpression: UnaryExpression{
					ExponentiationExpression: ExponentiationExpression{
						SubsetExpression: SubsetExpression{
							ScopeExpression: ScopeExpression{
								IndexOrCallExpression: IndexOrCallExpression{
									SimpleExpression: &SimpleExpression{
										Identifier: identA,
										Tokens:     tks[:1],
									},
									Tokens: tks[:1],
								},
								Tokens: tks[:1],
							},
							Tokens: tks[:1],
						},
						Tokens: tks[:1],
					},
					Tokens: tks[:1],
				},
				Tokens: tks[:1],
			},
			PipeOrSpecialExpression: &PipeOrSpecialExpression{
				SequenceExpression: SequenceExpression{
					UnaryExpression: UnaryExpression{
						ExponentiationExpression: ExponentiationExpression{
							SubsetExpression: SubsetExpression{
								ScopeExpression: ScopeExpression{
									IndexOrCallExpression: IndexOrCallExpression{
										SimpleExpression: &SimpleExpression{
											Identifier: identB,
											Tokens:     tks[1:2],
										},
										Tokens: tks[1:2],
									},
									Tokens: tks[1:2],
								},
								Tokens: tks[1:2],
							},
							Tokens: tks[1:2],
						},
						Tokens: tks[1:2],
					},
					Tokens: tks[1:2],
				},
				Tokens: tks[1:2],
			},
			Tokens: tks[:2],
		},
		&MultiplicationExpression{ // 11
			PipeOrSpecialExpression: PipeOrSpecialExpression{
				SequenceExpression: SequenceExpression{
					UnaryExpression: UnaryExpression{
						ExponentiationExpression: ExponentiationExpression{
							SubsetExpression: SubsetExpression{
								ScopeExpression: ScopeExpression{
									IndexOrCallExpression: IndexOrCallExpression{
										SimpleExpression: &SimpleExpression{
											Identifier: identA,
											Tokens:     tks[:1],
										},
										Tokens: tks[:1],
									},
									Tokens: tks[:1],
								},
								Tokens: tks[:1],
							},
							Tokens: tks[:1],
						},
						Tokens: tks[:1],
					},
					Tokens: tks[:1],
				},
				Tokens: tks[:1],
			},
			MultiplicationExpression: &MultiplicationExpression{
				PipeOrSpecialExpression: PipeOrSpecialExpression{
					SequenceExpression: SequenceExpression{
						UnaryExpression: UnaryExpression{
							ExponentiationExpression: ExponentiationExpression{
								SubsetExpression: SubsetExpression{
									ScopeExpression: ScopeExpression{
										IndexOrCallExpression: IndexOrCallExpression{
											SimpleExpression: &SimpleExpression{
												Identifier: identB,
												Tokens:     tks[1:2],
											},
											Tokens: tks[1:2],
										},
										Tokens: tks[1:2],
									},
									Tokens: tks[1:2],
								},
								Tokens: tks[1:2],
							},
							Tokens: tks[1:2],
						},
						Tokens: tks[1:2],
					},
					Tokens: tks[1:2],
				},
				Tokens: tks[1:2],
			},
			Tokens: tks[:2],
		},
		&AdditionExpression{ // 12
			MultiplicationExpression: MultiplicationExpression{
				PipeOrSpecialExpression: PipeOrSpecialExpression{
					SequenceExpression: SequenceExpression{
						UnaryExpression: UnaryExpression{
							ExponentiationExpression: ExponentiationExpression{
								SubsetExpression: SubsetExpression{
									ScopeExpression: ScopeExpression{
										IndexOrCallExpression: IndexOrCallExpression{
											SimpleExpression: &SimpleExpression{
												Identifier: identA,
												Tokens:     tks[:1],
											},
											Tokens: tks[:1],
										},
										Tokens: tks[:1],
									},
									Tokens: tks[:1],
								},
								Tokens: tks[:1],
							},
							Tokens: tks[:1],
						},
						Tokens: tks[:1],
					},
					Tokens: tks[:1],
				},
				Tokens: tks[:1],
			},
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
													Identifier: identB,
													Tokens:     tks[1:2],
												},
												Tokens: tks[1:2],
											},
											Tokens: tks[1:2],
										},
										Tokens: tks[1:2],
									},
									Tokens: tks[1:2],
								},
								Tokens: tks[1:2],
							},
							Tokens: tks[1:2],
						},
						Tokens: tks[1:2],
					},
					Tokens: tks[1:2],
				},
				Tokens: tks[1:2],
			},
			Tokens: tks[:2],
		},
		&RelationalExpression{ // 13
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
													Identifier: identA,
													Tokens:     tks[:1],
												},
												Tokens: tks[:1],
											},
											Tokens: tks[:1],
										},
										Tokens: tks[:1],
									},
									Tokens: tks[:1],
								},
								Tokens: tks[:1],
							},
							Tokens: tks[:1],
						},
						Tokens: tks[:1],
					},
					Tokens: tks[:1],
				},
				Tokens: tks[:1],
			},
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
														Identifier: identB,
														Tokens:     tks[1:2],
													},
													Tokens: tks[1:2],
												},
												Tokens: tks[1:2],
											},
											Tokens: tks[1:2],
										},
										Tokens: tks[1:2],
									},
									Tokens: tks[1:2],
								},
								Tokens: tks[1:2],
							},
							Tokens: tks[1:2],
						},
						Tokens: tks[1:2],
					},
					Tokens: tks[1:2],
				},
				Tokens: tks[1:2],
			},
			Tokens: tks[:2],
		},
		&NotExpression{ // 14
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
														Identifier: identA,
														Tokens:     tks[:1],
													},
													Tokens: tks[:1],
												},
												Tokens: tks[:1],
											},
											Tokens: tks[:1],
										},
										Tokens: tks[:1],
									},
									Tokens: tks[:1],
								},
								Tokens: tks[:1],
							},
							Tokens: tks[:1],
						},
						Tokens: tks[:1],
					},
					Tokens: tks[:1],
				},
				Tokens: tks[:1],
			},
			Tokens: tks[:1],
		},
		&AndExpression{ // 15
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
															Identifier: identA,
															Tokens:     tks[:1],
														},
														Tokens: tks[:1],
													},
													Tokens: tks[:1],
												},
												Tokens: tks[:1],
											},
											Tokens: tks[:1],
										},
										Tokens: tks[:1],
									},
									Tokens: tks[:1],
								},
								Tokens: tks[:1],
							},
							Tokens: tks[:1],
						},
						Tokens: tks[:1],
					},
					Tokens: tks[:1],
				},
				Tokens: tks[:1],
			},
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
																Identifier: identB,
																Tokens:     tks[1:2],
															},
															Tokens: tks[1:2],
														},
														Tokens: tks[1:2],
													},
													Tokens: tks[1:2],
												},
												Tokens: tks[1:2],
											},
											Tokens: tks[1:2],
										},
										Tokens: tks[1:2],
									},
									Tokens: tks[1:2],
								},
								Tokens: tks[1:2],
							},
							Tokens: tks[1:2],
						},
						Tokens: tks[1:2],
					},
					Tokens: tks[1:2],
				},
				Tokens: tks[1:2],
			},
			Tokens: tks[:2],
		},
		&OrExpression{ // 16
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
																Identifier: identA,
																Tokens:     tks[:1],
															},
															Tokens: tks[:1],
														},
														Tokens: tks[:1],
													},
													Tokens: tks[:1],
												},
												Tokens: tks[:1],
											},
											Tokens: tks[:1],
										},
										Tokens: tks[:1],
									},
									Tokens: tks[:1],
								},
								Tokens: tks[:1],
							},
							Tokens: tks[:1],
						},
						Tokens: tks[:1],
					},
					Tokens: tks[:1],
				},
				Tokens: tks[:1],
			},
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
																	Identifier: identB,
																	Tokens:     tks[1:2],
																},
																Tokens: tks[1:2],
															},
															Tokens: tks[1:2],
														},
														Tokens: tks[1:2],
													},
													Tokens: tks[1:2],
												},
												Tokens: tks[1:2],
											},
											Tokens: tks[1:2],
										},
										Tokens: tks[1:2],
									},
									Tokens: tks[1:2],
								},
								Tokens: tks[1:2],
							},
							Tokens: tks[1:2],
						},
						Tokens: tks[1:2],
					},
					Tokens: tks[1:2],
				},
				Tokens: tks[1:2],
			},
			Tokens: tks[:2],
		},
		&FormulaeExpression{ // 17
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
																		Identifier: identA,
																		Tokens:     tks[:1],
																	},
																	Tokens: tks[:1],
																},
																Tokens: tks[:1],
															},
															Tokens: tks[:1],
														},
														Tokens: tks[:1],
													},
													Tokens: tks[:1],
												},
												Tokens: tks[:1],
											},
											Tokens: tks[:1],
										},
										Tokens: tks[:1],
									},
									Tokens: tks[:1],
								},
								Tokens: tks[:1],
							},
							Tokens: tks[:1],
						},
						Tokens: tks[:1],
					},
					Tokens: tks[:1],
				},
				Tokens: tks[:1],
			},
			Tokens: tks[:1],
		},
		&AssignmentExpression{ // 18
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
																		Identifier: identA,
																		Tokens:     tks[:1],
																	},
																	Tokens: tks[:1],
																},
																Tokens: tks[:1],
															},
															Tokens: tks[:1],
														},
														Tokens: tks[:1],
													},
													Tokens: tks[:1],
												},
												Tokens: tks[:1],
											},
											Tokens: tks[:1],
										},
										Tokens: tks[:1],
									},
									Tokens: tks[:1],
								},
								Tokens: tks[:1],
							},
							Tokens: tks[:1],
						},
						Tokens: tks[:1],
					},
					Tokens: tks[:1],
				},
				Tokens: tks[:1],
			},
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
																			Identifier: identB,
																			Tokens:     tks[1:2],
																		},
																		Tokens: tks[1:2],
																	},
																	Tokens: tks[1:2],
																},
																Tokens: tks[1:2],
															},
															Tokens: tks[1:2],
														},
														Tokens: tks[1:2],
													},
													Tokens: tks[1:2],
												},
												Tokens: tks[1:2],
											},
											Tokens: tks[1:2],
										},
										Tokens: tks[1:2],
									},
									Tokens: tks[1:2],
								},
								Tokens: tks[1:2],
							},
							Tokens: tks[1:2],
						},
						Tokens: tks[1:2],
					},
					Tokens: tks[1:2],
				},
			},
			Tokens: tks[:2],
		},
		&QueryExpression{ // 19
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
																				Identifier: identA,
																				Tokens:     tks[:1],
																			},
																			Tokens: tks[:1],
																		},
																		Tokens: tks[:1],
																	},
																	Tokens: tks[:1],
																},
																Tokens: tks[:1],
															},
															Tokens: tks[:1],
														},
														Tokens: tks[:1],
													},
													Tokens: tks[:1],
												},
												Tokens: tks[:1],
											},
											Tokens: tks[:1],
										},
										Tokens: tks[:1],
									},
									Tokens: tks[:1],
								},
								Tokens: tks[:1],
							},
							Tokens: tks[:1],
						},
						Tokens: tks[:1],
					},
					Tokens: tks[:1],
				},
				Tokens: tks[:1],
			},
			Tokens: tks[:1],
		},
	} {
		if output := UnwrapQuery(WrapQuery(test)); !reflect.DeepEqual(output, test) {
			t.Errorf("test %d: expecting\n%v\n...got...\n%v", n+1, test, output)
		}
	}
}

func TestUnwrapQueryExtra(t *testing.T) {
	if res := UnwrapQuery(nil); res != nil {
		t.Errorf("test 1: expecting nil, got %v", res)
	}

	if res := UnwrapQuery(WrapQuery(&FormulaeExpression{})); res != nil {
		t.Errorf("test 2: expecting nil, got %v", res)
	}

	if res := UnwrapQuery(WrapQuery(&IndexOrCallExpression{})); res != nil {
		t.Errorf("test 3: expecting nil, got %v", res)
	}
}
