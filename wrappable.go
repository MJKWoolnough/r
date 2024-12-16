package r

type QueryWrappable interface {
	Type
	queryWrappable()
}

func (QueryExpression) queryWrappable() {}

func (AssignmentExpression) queryWrappable() {}

func (FormulaeExpression) queryWrappable() {}

func (OrExpression) queryWrappable() {}

func (AndExpression) queryWrappable() {}

func (NotExpression) queryWrappable() {}

func (RelationalExpression) queryWrappable() {}

func (AdditionExpression) queryWrappable() {}

func (MultiplicationExpression) queryWrappable() {}

func (PipeOrSpecialExpression) queryWrappable() {}

func (SequenceExpression) queryWrappable() {}

func (UnaryExpression) queryWrappable() {}

func (ExponentiationExpression) queryWrappable() {}

func (SubsetExpression) queryWrappable() {}

func (ScopeExpression) queryWrappable() {}

func (IndexOrCallExpression) queryWrappable() {}

func (SimpleExpression) queryWrappable() {}

func (CompoundExpression) queryWrappable() {}
