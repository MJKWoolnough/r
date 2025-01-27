package walk

import "vimagination.zapto.org/r"

// Handler is used to process R types.
type Handler interface {
	Handle(r.Type) error
}

// HandlerFunc wraps a func to implement Handler interface.
type HandlerFunc func(r.Type) error

// Handle implements the Handler interface.
func (h HandlerFunc) Handle(t r.Type) error {
	return h(t)
}

// Walk calls the Handle function on the given interface for each non-nil, non-Token field of the given R type.
func Walk(t r.Type, fn Handler) error {
	switch t := t.(type) {
	case r.AdditionExpression:
		return walkAdditionExpression(&t, fn)
	case *r.AdditionExpression:
		return walkAdditionExpression(t, fn)
	case r.AndExpression:
		return walkAndExpression(&t, fn)
	case *r.AndExpression:
		return walkAndExpression(t, fn)
	case r.Arg:
		return walkArg(&t, fn)
	case *r.Arg:
		return walkArg(t, fn)
	case r.ArgList:
	case *r.ArgList:
	case r.Argument:
	case *r.Argument:
	case r.AssignmentExpression:
	case *r.AssignmentExpression:
	case r.Call:
	case *r.Call:
	case r.CompoundExpression:
	case *r.CompoundExpression:
	case r.ExponentiationExpression:
	case *r.ExponentiationExpression:
	case r.Expression:
	case *r.Expression:
	case r.File:
	case *r.File:
	case r.FlowControl:
	case *r.FlowControl:
	case r.ForControl:
	case *r.ForControl:
	case r.FormulaeExpression:
	case *r.FormulaeExpression:
	case r.FunctionDefinition:
	case *r.FunctionDefinition:
	case r.IfControl:
	case *r.IfControl:
	case r.Index:
	case *r.Index:
	case r.IndexExpression:
	case *r.IndexExpression:
	case r.IndexOrCallExpression:
	case *r.IndexOrCallExpression:
	case r.MultiplicationExpression:
	case *r.MultiplicationExpression:
	case r.NotExpression:
	case *r.NotExpression:
	case r.OrExpression:
	case *r.OrExpression:
	case r.ParenthesizedExpression:
	case *r.ParenthesizedExpression:
	case r.PipeOrSpecialExpression:
	case *r.PipeOrSpecialExpression:
	case r.QueryExpression:
	case *r.QueryExpression:
	}

	return nil
}

func walkAdditionExpression(t *r.AdditionExpression, fn Handler) error {
	if err := fn.Handle(&t.MultiplicationExpression); err != nil {
		return err
	}

	if t.AdditionExpression != nil {
		return fn.Handle(t.AdditionExpression)
	}

	return nil
}

func walkAndExpression(t *r.AndExpression, fn Handler) error {
	if err := fn.Handle(&t.NotExpression); err != nil {
		return err
	}

	if t.AndExpression != nil {
		return fn.Handle(t.AndExpression)
	}

	return nil
}

func walkArg(t *r.Arg, fn Handler) error {
	if t.QueryExpression != nil {
		return fn.Handle(t.QueryExpression)
	}

	return nil
}

func walkArgList(t *r.ArgList, fn Handler) error { return nil }

func walkArgument(t *r.Argument, fn Handler) error { return nil }

func walkAssignmentExpression(t *r.AssignmentExpression, fn Handler) error { return nil }

func walkCall(t *r.Call, fn Handler) error { return nil }

func walkCompoundExpression(t *r.CompoundExpression, fn Handler) error { return nil }

func walkExponentiationExpression(t *r.ExponentiationExpression, fn Handler) error { return nil }

func walkExpression(t *r.Expression, fn Handler) error { return nil }

func walkFile(t *r.File, fn Handler) error { return nil }

func walkFlowControl(t *r.FlowControl, fn Handler) error { return nil }

func walkForControl(t *r.ForControl, fn Handler) error { return nil }

func walkFormulaeExpression(t *r.FormulaeExpression, fn Handler) error { return nil }

func walkFunctionDefinition(t *r.FunctionDefinition, fn Handler) error { return nil }

func walkIfControl(t *r.IfControl, fn Handler) error { return nil }

func walkIndex(t *r.Index, fn Handler) error { return nil }

func walkIndexExpression(t *r.IndexExpression, fn Handler) error { return nil }

func walkIndexOrCallExpression(t *r.IndexOrCallExpression, fn Handler) error { return nil }

func walkMultiplicationExpression(t *r.MultiplicationExpression, fn Handler) error { return nil }

func walkNotExpression(t *r.NotExpression, fn Handler) error { return nil }

func walkOrExpression(t *r.OrExpression, fn Handler) error { return nil }

func walkParenthesizedExpression(t *r.ParenthesizedExpression, fn Handler) error { return nil }

func walkPipeOrSpecialExpression(t *r.PipeOrSpecialExpression, fn Handler) error { return nil }

func walkQueryExpression(t *r.QueryExpression, fn Handler) error { return nil }
