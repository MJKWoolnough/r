package r

// File automatically generated with format.sh.

import "fmt"

// Type is an interface satisfied by all R structural types.
type Type interface {
	fmt.Formatter
	rType()
}

func (Tokens) rType() {}

func (AdditionExpression) rType() {}

func (AndExpression) rType() {}

func (Arg) rType() {}

func (ArgList) rType() {}

func (Argument) rType() {}

func (AssignmentExpression) rType() {}

func (Call) rType() {}

func (CompoundExpression) rType() {}

func (ExponentiationExpression) rType() {}

func (Expression) rType() {}

func (FlowControl) rType() {}

func (ForControl) rType() {}

func (FormulaeExpression) rType() {}

func (FunctionDefinition) rType() {}

func (IfControl) rType() {}

func (Index) rType() {}

func (IndexOrCallExpression) rType() {}

func (MultiplicationExpression) rType() {}

func (NotExpression) rType() {}

func (OrExpression) rType() {}

func (PipeOrSpecialExpression) rType() {}

func (QueryExpression) rType() {}

func (RelationalExpression) rType() {}

func (RepeatControl) rType() {}

func (ScopeExpression) rType() {}

func (SequenceExpression) rType() {}

func (SimpleExpression) rType() {}

func (SubsetExpression) rType() {}

func (UnaryExpression) rType() {}

func (WhileControl) rType() {}
