package r_test

import (
	"fmt"

	"vimagination.zapto.org/parser"
	"vimagination.zapto.org/r"
)

func Example() {
	src := `hello <- function(name) { message(sprintf("Hello, %s", name)) }; hello("Alice")`

	tk := parser.NewStringTokeniser(src)

	ast, err := r.Parse(&tk)
	if err != nil {
		fmt.Println(err)

		return
	}

	r.UnwrapQuery(r.UnwrapQuery(r.UnwrapQuery(r.UnwrapQuery(ast.Statements[0].QueryExpression.AssignmentExpression.Expression.FunctionDefinition.Body.QueryExpression).(*r.CompoundExpression).Expressions[0].QueryExpression).(*r.IndexOrCallExpression).Call.Args[0].QueryExpression).(*r.IndexOrCallExpression).Call.Args[0].QueryExpression).(*r.SimpleExpression).Constant.Data = `"Hi, %s"`

	fmt.Printf("%+s", ast)

	// Output:
	// hello <- function(name) {
	// 	message(sprintf("Hi, %s", name))
	// }
	// hello("Alice")
}
