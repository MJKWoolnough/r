package r

// WrapQuery takes one of many types and wraps it in a *QueryExpression
//
// The accepted types/pointers are as follows:
//
//	QueryExpression
//	*QueryExpression
//	AssignmentExpression
//	*AssignmentExpression
//	FormulaeExpression
//	*FormulaeExpression
//	OrExpression
//	*OrExpression
//	AndExpression
//	*AndExpression
//	NotExpression
//	*NotExpression
//	RelationalExpression
//	*RelationalExpression
//	AdditionExpression
//	*AdditionExpression
//	MultiplicationExpression
//	*MultiplicationExpression
//	PipeOrSpecialExpression
//	*PipeOrSpecialExpression
//	SequenceExpression
//	*SequenceExpression
//	UnaryExpression
//	*UnaryExpression
//	ExponentiationExpression
//	*ExponentiationExpression
//	SubsetExpression
//	*SubsetExpression
//	ScopeExpression
//	*ScopeExpression
//	IndexOrCallExpression
//	*IndexOrCallExpression
//	SimpleExpression
//	*SimpleExpression
//	CompoundExpression
//	*CompoundExpression
func WrapQuery(p QueryWrappable) *QueryExpression {
	if q, ok := p.(*QueryExpression); ok {
		return q
	}

	if q, ok := p.(QueryExpression); ok {
		return &q
	}

	q := new(QueryExpression)

	switch p := p.(type) {
	case *AssignmentExpression:
		q.AssignmentExpression = p

		goto AssignmentExpression
	case AssignmentExpression:
		q.AssignmentExpression = &p

		goto AssignmentExpression
	default:
		q.AssignmentExpression = new(AssignmentExpression)

		switch p := p.(type) {
		case *FormulaeExpression:
			q.AssignmentExpression.FormulaeExpression = *p

			goto FormulaeExpression
		case FormulaeExpression:
			q.AssignmentExpression.FormulaeExpression = p

			goto FormulaeExpression
		case *OrExpression:
			q.AssignmentExpression.FormulaeExpression.OrExpression = p

			goto OrExpression
		case OrExpression:
			q.AssignmentExpression.FormulaeExpression.OrExpression = &p

			goto OrExpression
		default:
			q.AssignmentExpression.FormulaeExpression.OrExpression = new(OrExpression)

			switch p := p.(type) {
			case *AndExpression:
				q.AssignmentExpression.FormulaeExpression.OrExpression.AndExpression = *p

				goto AndExpression
			case AndExpression:
				q.AssignmentExpression.FormulaeExpression.OrExpression.AndExpression = p

				goto AndExpression
			case *NotExpression:
				q.AssignmentExpression.FormulaeExpression.OrExpression.AndExpression.NotExpression = *p

				goto NotExpression
			case NotExpression:
				q.AssignmentExpression.FormulaeExpression.OrExpression.AndExpression.NotExpression = p

				goto NotExpression
			case *RelationalExpression:
				q.AssignmentExpression.FormulaeExpression.OrExpression.AndExpression.NotExpression.RelationalExpression = *p

				goto RelationalExpression
			case RelationalExpression:
				q.AssignmentExpression.FormulaeExpression.OrExpression.AndExpression.NotExpression.RelationalExpression = p

				goto RelationalExpression
			case *AdditionExpression:
				q.AssignmentExpression.FormulaeExpression.OrExpression.AndExpression.NotExpression.RelationalExpression.AdditionExpression = *p

				goto AdditionExpression
			case AdditionExpression:
				q.AssignmentExpression.FormulaeExpression.OrExpression.AndExpression.NotExpression.RelationalExpression.AdditionExpression = p

				goto AdditionExpression
			case *MultiplicationExpression:
				q.AssignmentExpression.FormulaeExpression.OrExpression.AndExpression.NotExpression.RelationalExpression.AdditionExpression.MultiplicationExpression = *p

				goto MultiplicationExpression
			case MultiplicationExpression:
				q.AssignmentExpression.FormulaeExpression.OrExpression.AndExpression.NotExpression.RelationalExpression.AdditionExpression.MultiplicationExpression = p

				goto MultiplicationExpression
			case *PipeOrSpecialExpression:
				q.AssignmentExpression.FormulaeExpression.OrExpression.AndExpression.NotExpression.RelationalExpression.AdditionExpression.MultiplicationExpression.PipeOrSpecialExpression = *p

				goto PipeOrSpecialExpression
			case PipeOrSpecialExpression:
				q.AssignmentExpression.FormulaeExpression.OrExpression.AndExpression.NotExpression.RelationalExpression.AdditionExpression.MultiplicationExpression.PipeOrSpecialExpression = p

				goto PipeOrSpecialExpression
			case *SequenceExpression:
				q.AssignmentExpression.FormulaeExpression.OrExpression.AndExpression.NotExpression.RelationalExpression.AdditionExpression.MultiplicationExpression.PipeOrSpecialExpression.SequenceExpression = *p

				goto SequenceExpression
			case SequenceExpression:
				q.AssignmentExpression.FormulaeExpression.OrExpression.AndExpression.NotExpression.RelationalExpression.AdditionExpression.MultiplicationExpression.PipeOrSpecialExpression.SequenceExpression = p

				goto SequenceExpression
			case *UnaryExpression:
				q.AssignmentExpression.FormulaeExpression.OrExpression.AndExpression.NotExpression.RelationalExpression.AdditionExpression.MultiplicationExpression.PipeOrSpecialExpression.SequenceExpression.UnaryExpression = *p

				goto UnaryExpression
			case UnaryExpression:
				q.AssignmentExpression.FormulaeExpression.OrExpression.AndExpression.NotExpression.RelationalExpression.AdditionExpression.MultiplicationExpression.PipeOrSpecialExpression.SequenceExpression.UnaryExpression = p

				goto UnaryExpression
			case *ExponentiationExpression:
				q.AssignmentExpression.FormulaeExpression.OrExpression.AndExpression.NotExpression.RelationalExpression.AdditionExpression.MultiplicationExpression.PipeOrSpecialExpression.SequenceExpression.UnaryExpression.ExponentiationExpression = *p

				goto ExponentiationExpression
			case ExponentiationExpression:
				q.AssignmentExpression.FormulaeExpression.OrExpression.AndExpression.NotExpression.RelationalExpression.AdditionExpression.MultiplicationExpression.PipeOrSpecialExpression.SequenceExpression.UnaryExpression.ExponentiationExpression = p

				goto ExponentiationExpression
			case *SubsetExpression:
				q.AssignmentExpression.FormulaeExpression.OrExpression.AndExpression.NotExpression.RelationalExpression.AdditionExpression.MultiplicationExpression.PipeOrSpecialExpression.SequenceExpression.UnaryExpression.ExponentiationExpression.SubsetExpression = *p

				goto SubsetExpression
			case SubsetExpression:
				q.AssignmentExpression.FormulaeExpression.OrExpression.AndExpression.NotExpression.RelationalExpression.AdditionExpression.MultiplicationExpression.PipeOrSpecialExpression.SequenceExpression.UnaryExpression.ExponentiationExpression.SubsetExpression = p

				goto SubsetExpression
			case *ScopeExpression:
				q.AssignmentExpression.FormulaeExpression.OrExpression.AndExpression.NotExpression.RelationalExpression.AdditionExpression.MultiplicationExpression.PipeOrSpecialExpression.SequenceExpression.UnaryExpression.ExponentiationExpression.SubsetExpression.ScopeExpression = *p

				goto ScopeExpression
			case ScopeExpression:
				q.AssignmentExpression.FormulaeExpression.OrExpression.AndExpression.NotExpression.RelationalExpression.AdditionExpression.MultiplicationExpression.PipeOrSpecialExpression.SequenceExpression.UnaryExpression.ExponentiationExpression.SubsetExpression.ScopeExpression = p

				goto ScopeExpression
			case *IndexOrCallExpression:
				q.AssignmentExpression.FormulaeExpression.OrExpression.AndExpression.NotExpression.RelationalExpression.AdditionExpression.MultiplicationExpression.PipeOrSpecialExpression.SequenceExpression.UnaryExpression.ExponentiationExpression.SubsetExpression.ScopeExpression.IndexOrCallExpression = *p

				goto IndexOrCallExpression
			case IndexOrCallExpression:
				q.AssignmentExpression.FormulaeExpression.OrExpression.AndExpression.NotExpression.RelationalExpression.AdditionExpression.MultiplicationExpression.PipeOrSpecialExpression.SequenceExpression.UnaryExpression.ExponentiationExpression.SubsetExpression.ScopeExpression.IndexOrCallExpression = p

				goto IndexOrCallExpression
			case *SimpleExpression:
				q.AssignmentExpression.FormulaeExpression.OrExpression.AndExpression.NotExpression.RelationalExpression.AdditionExpression.MultiplicationExpression.PipeOrSpecialExpression.SequenceExpression.UnaryExpression.ExponentiationExpression.SubsetExpression.ScopeExpression.IndexOrCallExpression.SimpleExpression = p

				goto SimpleExpression
			case SimpleExpression:
				q.AssignmentExpression.FormulaeExpression.OrExpression.AndExpression.NotExpression.RelationalExpression.AdditionExpression.MultiplicationExpression.PipeOrSpecialExpression.SequenceExpression.UnaryExpression.ExponentiationExpression.SubsetExpression.ScopeExpression.IndexOrCallExpression.SimpleExpression = &p

				goto SimpleExpression
			case *CompoundExpression:
				q.AssignmentExpression.FormulaeExpression.OrExpression.AndExpression.NotExpression.RelationalExpression.AdditionExpression.MultiplicationExpression.PipeOrSpecialExpression.SequenceExpression.UnaryExpression.ExponentiationExpression.SubsetExpression.ScopeExpression.IndexOrCallExpression.SimpleExpression = &SimpleExpression{
					CompoundExpression: p,
					Tokens:             p.Tokens,
				}

				goto SimpleExpression
			case CompoundExpression:
				q.AssignmentExpression.FormulaeExpression.OrExpression.AndExpression.NotExpression.RelationalExpression.AdditionExpression.MultiplicationExpression.PipeOrSpecialExpression.SequenceExpression.UnaryExpression.ExponentiationExpression.SubsetExpression.ScopeExpression.IndexOrCallExpression.SimpleExpression = &SimpleExpression{
					CompoundExpression: &p,
					Tokens:             p.Tokens,
				}

				goto SimpleExpression
			}
		}
	}

SimpleExpression:
	q.AssignmentExpression.FormulaeExpression.OrExpression.AndExpression.NotExpression.RelationalExpression.AdditionExpression.MultiplicationExpression.PipeOrSpecialExpression.SequenceExpression.UnaryExpression.ExponentiationExpression.SubsetExpression.ScopeExpression.IndexOrCallExpression.Tokens = q.AssignmentExpression.FormulaeExpression.OrExpression.AndExpression.NotExpression.RelationalExpression.AdditionExpression.MultiplicationExpression.PipeOrSpecialExpression.SequenceExpression.UnaryExpression.ExponentiationExpression.SubsetExpression.ScopeExpression.IndexOrCallExpression.SimpleExpression.Tokens
IndexOrCallExpression:
	q.AssignmentExpression.FormulaeExpression.OrExpression.AndExpression.NotExpression.RelationalExpression.AdditionExpression.MultiplicationExpression.PipeOrSpecialExpression.SequenceExpression.UnaryExpression.ExponentiationExpression.SubsetExpression.ScopeExpression.Tokens = q.AssignmentExpression.FormulaeExpression.OrExpression.AndExpression.NotExpression.RelationalExpression.AdditionExpression.MultiplicationExpression.PipeOrSpecialExpression.SequenceExpression.UnaryExpression.ExponentiationExpression.SubsetExpression.ScopeExpression.IndexOrCallExpression.Tokens
ScopeExpression:
	q.AssignmentExpression.FormulaeExpression.OrExpression.AndExpression.NotExpression.RelationalExpression.AdditionExpression.MultiplicationExpression.PipeOrSpecialExpression.SequenceExpression.UnaryExpression.ExponentiationExpression.SubsetExpression.Tokens = q.AssignmentExpression.FormulaeExpression.OrExpression.AndExpression.NotExpression.RelationalExpression.AdditionExpression.MultiplicationExpression.PipeOrSpecialExpression.SequenceExpression.UnaryExpression.ExponentiationExpression.SubsetExpression.ScopeExpression.Tokens
SubsetExpression:
	q.AssignmentExpression.FormulaeExpression.OrExpression.AndExpression.NotExpression.RelationalExpression.AdditionExpression.MultiplicationExpression.PipeOrSpecialExpression.SequenceExpression.UnaryExpression.ExponentiationExpression.Tokens = q.AssignmentExpression.FormulaeExpression.OrExpression.AndExpression.NotExpression.RelationalExpression.AdditionExpression.MultiplicationExpression.PipeOrSpecialExpression.SequenceExpression.UnaryExpression.ExponentiationExpression.SubsetExpression.Tokens
ExponentiationExpression:
	q.AssignmentExpression.FormulaeExpression.OrExpression.AndExpression.NotExpression.RelationalExpression.AdditionExpression.MultiplicationExpression.PipeOrSpecialExpression.SequenceExpression.UnaryExpression.Tokens = q.AssignmentExpression.FormulaeExpression.OrExpression.AndExpression.NotExpression.RelationalExpression.AdditionExpression.MultiplicationExpression.PipeOrSpecialExpression.SequenceExpression.UnaryExpression.ExponentiationExpression.Tokens
UnaryExpression:
	q.AssignmentExpression.FormulaeExpression.OrExpression.AndExpression.NotExpression.RelationalExpression.AdditionExpression.MultiplicationExpression.PipeOrSpecialExpression.SequenceExpression.Tokens = q.AssignmentExpression.FormulaeExpression.OrExpression.AndExpression.NotExpression.RelationalExpression.AdditionExpression.MultiplicationExpression.PipeOrSpecialExpression.SequenceExpression.UnaryExpression.Tokens
SequenceExpression:
	q.AssignmentExpression.FormulaeExpression.OrExpression.AndExpression.NotExpression.RelationalExpression.AdditionExpression.MultiplicationExpression.PipeOrSpecialExpression.Tokens = q.AssignmentExpression.FormulaeExpression.OrExpression.AndExpression.NotExpression.RelationalExpression.AdditionExpression.MultiplicationExpression.PipeOrSpecialExpression.SequenceExpression.Tokens
PipeOrSpecialExpression:
	q.AssignmentExpression.FormulaeExpression.OrExpression.AndExpression.NotExpression.RelationalExpression.AdditionExpression.MultiplicationExpression.Tokens = q.AssignmentExpression.FormulaeExpression.OrExpression.AndExpression.NotExpression.RelationalExpression.AdditionExpression.MultiplicationExpression.PipeOrSpecialExpression.Tokens
MultiplicationExpression:
	q.AssignmentExpression.FormulaeExpression.OrExpression.AndExpression.NotExpression.RelationalExpression.AdditionExpression.Tokens = q.AssignmentExpression.FormulaeExpression.OrExpression.AndExpression.NotExpression.RelationalExpression.AdditionExpression.MultiplicationExpression.Tokens
AdditionExpression:
	q.AssignmentExpression.FormulaeExpression.OrExpression.AndExpression.NotExpression.RelationalExpression.Tokens = q.AssignmentExpression.FormulaeExpression.OrExpression.AndExpression.NotExpression.RelationalExpression.AdditionExpression.Tokens
RelationalExpression:
	q.AssignmentExpression.FormulaeExpression.OrExpression.AndExpression.NotExpression.Tokens = q.AssignmentExpression.FormulaeExpression.OrExpression.AndExpression.NotExpression.RelationalExpression.Tokens
NotExpression:
	q.AssignmentExpression.FormulaeExpression.OrExpression.AndExpression.Tokens = q.AssignmentExpression.FormulaeExpression.OrExpression.AndExpression.NotExpression.Tokens
AndExpression:
	q.AssignmentExpression.FormulaeExpression.OrExpression.Tokens = q.AssignmentExpression.FormulaeExpression.OrExpression.AndExpression.Tokens
OrExpression:
	q.AssignmentExpression.FormulaeExpression.Tokens = q.AssignmentExpression.FormulaeExpression.OrExpression.Tokens
FormulaeExpression:
	q.AssignmentExpression.Tokens = q.AssignmentExpression.FormulaeExpression.Tokens
AssignmentExpression:
	q.Tokens = q.AssignmentExpression.Tokens

	return q
}

// UnwrapQuery returns the first value up the QueryExpression chain that
// contains all of the information required to rebuild the lower chain.
//
// Possible returns types are as follows:
//
//	*QueryExpression
//	*AssignmentExpression
//	*FormulaeExpression
//	*OrExpression
//	*AndExpression
//	*NotExpression
//	*RelationalExpression
//	*AdditionExpression
//	*MultiplicationExpression
//	*PipeOrSpecialExpression
//	*SequenceExpression
//	*UnaryExpression
//	*ExponentiationExpression
//	*SubsetExpression
//	*ScopeExpression
//	*IndexOrCallExpression
//	*SimpleExpression
//	*CompoundExpression
func UnwrapQuery(q *QueryExpression) QueryWrappable {
	if q == nil {
		return nil
	} else if q.QueryExpression != nil {
		return q
	} else if q.AssignmentExpression.Expression != nil {
		return q.AssignmentExpression
	} else if q.AssignmentExpression.FormulaeExpression.FormulaeExpression != nil {
		return &q.AssignmentExpression.FormulaeExpression
	} else if q.AssignmentExpression.FormulaeExpression.OrExpression == nil {
		return nil
	} else if q.AssignmentExpression.FormulaeExpression.OrExpression.OrExpression != nil {
		return q.AssignmentExpression.FormulaeExpression.OrExpression
	} else if q.AssignmentExpression.FormulaeExpression.OrExpression.AndExpression.AndExpression != nil {
		return &q.AssignmentExpression.FormulaeExpression.OrExpression.AndExpression
	} else if q.AssignmentExpression.FormulaeExpression.OrExpression.AndExpression.NotExpression.Nots != 0 {
		return &q.AssignmentExpression.FormulaeExpression.OrExpression.AndExpression.NotExpression
	} else if q.AssignmentExpression.FormulaeExpression.OrExpression.AndExpression.NotExpression.RelationalExpression.RelationalExpression != nil {
		return &q.AssignmentExpression.FormulaeExpression.OrExpression.AndExpression.NotExpression.RelationalExpression
	} else if q.AssignmentExpression.FormulaeExpression.OrExpression.AndExpression.NotExpression.RelationalExpression.AdditionExpression.AdditionExpression != nil {
		return &q.AssignmentExpression.FormulaeExpression.OrExpression.AndExpression.NotExpression.RelationalExpression.AdditionExpression
	} else if q.AssignmentExpression.FormulaeExpression.OrExpression.AndExpression.NotExpression.RelationalExpression.AdditionExpression.MultiplicationExpression.MultiplicationExpression != nil {
		return &q.AssignmentExpression.FormulaeExpression.OrExpression.AndExpression.NotExpression.RelationalExpression.AdditionExpression.MultiplicationExpression
	} else if q.AssignmentExpression.FormulaeExpression.OrExpression.AndExpression.NotExpression.RelationalExpression.AdditionExpression.MultiplicationExpression.PipeOrSpecialExpression.PipeOrSpecialExpression != nil {
		return &q.AssignmentExpression.FormulaeExpression.OrExpression.AndExpression.NotExpression.RelationalExpression.AdditionExpression.MultiplicationExpression.PipeOrSpecialExpression
	} else if q.AssignmentExpression.FormulaeExpression.OrExpression.AndExpression.NotExpression.RelationalExpression.AdditionExpression.MultiplicationExpression.PipeOrSpecialExpression.SequenceExpression.SequenceExpression != nil {
		return &q.AssignmentExpression.FormulaeExpression.OrExpression.AndExpression.NotExpression.RelationalExpression.AdditionExpression.MultiplicationExpression.PipeOrSpecialExpression.SequenceExpression
	} else if len(q.AssignmentExpression.FormulaeExpression.OrExpression.AndExpression.NotExpression.RelationalExpression.AdditionExpression.MultiplicationExpression.PipeOrSpecialExpression.SequenceExpression.UnaryExpression.UnaryType) != 0 {
		return &q.AssignmentExpression.FormulaeExpression.OrExpression.AndExpression.NotExpression.RelationalExpression.AdditionExpression.MultiplicationExpression.PipeOrSpecialExpression.SequenceExpression.UnaryExpression
	} else if q.AssignmentExpression.FormulaeExpression.OrExpression.AndExpression.NotExpression.RelationalExpression.AdditionExpression.MultiplicationExpression.PipeOrSpecialExpression.SequenceExpression.UnaryExpression.ExponentiationExpression.ExponentiationExpression != nil {
		return &q.AssignmentExpression.FormulaeExpression.OrExpression.AndExpression.NotExpression.RelationalExpression.AdditionExpression.MultiplicationExpression.PipeOrSpecialExpression.SequenceExpression.UnaryExpression.ExponentiationExpression
	} else if q.AssignmentExpression.FormulaeExpression.OrExpression.AndExpression.NotExpression.RelationalExpression.AdditionExpression.MultiplicationExpression.PipeOrSpecialExpression.SequenceExpression.UnaryExpression.ExponentiationExpression.SubsetExpression.SubsetExpression != nil {
		return &q.AssignmentExpression.FormulaeExpression.OrExpression.AndExpression.NotExpression.RelationalExpression.AdditionExpression.MultiplicationExpression.PipeOrSpecialExpression.SequenceExpression.UnaryExpression.ExponentiationExpression.SubsetExpression
	} else if q.AssignmentExpression.FormulaeExpression.OrExpression.AndExpression.NotExpression.RelationalExpression.AdditionExpression.MultiplicationExpression.PipeOrSpecialExpression.SequenceExpression.UnaryExpression.ExponentiationExpression.SubsetExpression.ScopeExpression.ScopeExpression != nil {
		return &q.AssignmentExpression.FormulaeExpression.OrExpression.AndExpression.NotExpression.RelationalExpression.AdditionExpression.MultiplicationExpression.PipeOrSpecialExpression.SequenceExpression.UnaryExpression.ExponentiationExpression.SubsetExpression.ScopeExpression
	} else if q.AssignmentExpression.FormulaeExpression.OrExpression.AndExpression.NotExpression.RelationalExpression.AdditionExpression.MultiplicationExpression.PipeOrSpecialExpression.SequenceExpression.UnaryExpression.ExponentiationExpression.SubsetExpression.ScopeExpression.IndexOrCallExpression.Call != nil || q.AssignmentExpression.FormulaeExpression.OrExpression.AndExpression.NotExpression.RelationalExpression.AdditionExpression.MultiplicationExpression.PipeOrSpecialExpression.SequenceExpression.UnaryExpression.ExponentiationExpression.SubsetExpression.ScopeExpression.IndexOrCallExpression.Index != nil || q.AssignmentExpression.FormulaeExpression.OrExpression.AndExpression.NotExpression.RelationalExpression.AdditionExpression.MultiplicationExpression.PipeOrSpecialExpression.SequenceExpression.UnaryExpression.ExponentiationExpression.SubsetExpression.ScopeExpression.IndexOrCallExpression.IndexOrCallExpression != nil {
		return &q.AssignmentExpression.FormulaeExpression.OrExpression.AndExpression.NotExpression.RelationalExpression.AdditionExpression.MultiplicationExpression.PipeOrSpecialExpression.SequenceExpression.UnaryExpression.ExponentiationExpression.SubsetExpression.ScopeExpression.IndexOrCallExpression
	} else if q.AssignmentExpression.FormulaeExpression.OrExpression.AndExpression.NotExpression.RelationalExpression.AdditionExpression.MultiplicationExpression.PipeOrSpecialExpression.SequenceExpression.UnaryExpression.ExponentiationExpression.SubsetExpression.ScopeExpression.IndexOrCallExpression.SimpleExpression == nil {
		return nil
	} else if q.AssignmentExpression.FormulaeExpression.OrExpression.AndExpression.NotExpression.RelationalExpression.AdditionExpression.MultiplicationExpression.PipeOrSpecialExpression.SequenceExpression.UnaryExpression.ExponentiationExpression.SubsetExpression.ScopeExpression.IndexOrCallExpression.SimpleExpression.CompoundExpression != nil {
		return q.AssignmentExpression.FormulaeExpression.OrExpression.AndExpression.NotExpression.RelationalExpression.AdditionExpression.MultiplicationExpression.PipeOrSpecialExpression.SequenceExpression.UnaryExpression.ExponentiationExpression.SubsetExpression.ScopeExpression.IndexOrCallExpression.SimpleExpression.CompoundExpression
	}

	return q.AssignmentExpression.FormulaeExpression.OrExpression.AndExpression.NotExpression.RelationalExpression.AdditionExpression.MultiplicationExpression.PipeOrSpecialExpression.SequenceExpression.UnaryExpression.ExponentiationExpression.SubsetExpression.ScopeExpression.IndexOrCallExpression.SimpleExpression
}
