# walk

[![CI](https://github.com/MJKWoolnough/r/actions/workflows/go-checks.yml/badge.svg)](https://github.com/MJKWoolnough/r/actions)
[![Go Reference](https://pkg.go.dev/badge/vimagination.zapto.org/r.svg)](https://pkg.go.dev/vimagination.zapto.org/r/walk)
[![Go Report Card](https://goreportcard.com/badge/vimagination.zapto.org/r)](https://goreportcard.com/report/vimagination.zapto.org/r)

--
    import "vimagination.zapto.org/r/walk"

Package walk provides an R type walker.

## Highlights

 - Simple interface to allow control over walking through parsed R.
 - Allows modification to the tree as it's being walked.

## Usage

```go
package main

import (
	"fmt"

	"vimagination.zapto.org/parser"
	"vimagination.zapto.org/r"
	"vimagination.zapto.org/r/walk"
)

func main() {
	src := `a <- "b" - "c"`
	tk := parser.NewStringTokeniser(src)

	m, _ := r.Parse(&tk)

	var walkFn walk.Handler

	walkFn = walk.HandlerFunc(func(t r.Type) error {
		switch t := t.(type) {
		case *r.AdditionExpression:
			if t.AdditionType == r.AdditionSubtract {
				t.AdditionType = r.AdditionAdd
			}
		case *r.SimpleExpression:
			if t.Constant == nil {
				break
			}

			switch t.Constant.Data {
			case `"b"`:
				t.Constant.Data = `"Hello"`
			case `"c"`:
				t.Constant.Data = `", World"`
			}
		}

		return walk.Walk(t, walkFn)
	})

	walk.Walk(m, walkFn)

	fmt.Printf("%+s", m)

	// Output:
	// a <- "Hello" + ", World"
}
```

## Documentation

Full API docs can be found at:

https://pkg.go.dev/vimagination.zapto.org/r/walk
