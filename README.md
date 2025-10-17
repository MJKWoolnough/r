# r

[![CI](https://github.com/MJKWoolnough/r/actions/workflows/go-checks.yml/badge.svg)](https://github.com/MJKWoolnough/r/actions)
[![Go Reference](https://pkg.go.dev/badge/vimagination.zapto.org/r.svg)](https://pkg.go.dev/vimagination.zapto.org/r)
[![Go Report Card](https://goreportcard.com/badge/vimagination.zapto.org/r)](https://goreportcard.com/report/vimagination.zapto.org/r)

--
    import "vimagination.zapto.org/r"

Package r implements an R tokeniser and parser.

## Highlights

 - Parse R into AST.
 - Modify parsed code.
 - Consistant R formatting.

## Usage

```go
package main

import (
	"fmt"

	"vimagination.zapto.org/parser"
	"vimagination.zapto.org/r"
)

func main() {
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
```

## Documentation

Full API docs can be found at:

https://pkg.go.dev/vimagination.zapto.org/r
