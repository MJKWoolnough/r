package walk_test

import (
	"fmt"

	"vimagination.zapto.org/parser"
	"vimagination.zapto.org/r"
	"vimagination.zapto.org/r/walk"
)

func Example() {
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
