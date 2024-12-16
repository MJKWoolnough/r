package r

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
		q.AssignmentExpression = *p

		goto AssignmentExpression
	case AssignmentExpression:
		q.AssignmentExpression = p

		goto AssignmentExpression
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
