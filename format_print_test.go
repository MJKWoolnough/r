package r

import (
	"fmt"
	"testing"

	"vimagination.zapto.org/parser"
)

func TestPrintSource(t *testing.T) {
	for n, test := range [...][3]string{
		{ // 1
			"a+b",
			"a+b\n",
			"a + b\n",
		},
		{ // 2
			"a-b",
			"a-b\n",
			"a - b\n",
		},
		{ // 3
			"a + b-c",
			"a+b-c\n",
			"a + b - c\n",
		},
		{ // 4
			"a&b",
			"a&b\n",
			"a & b\n",
		},
		{ // 5
			"a&&b",
			"a&&b\n",
			"a && b\n",
		},
		{ // 6
			"a & b&&c",
			"a&b&&c\n",
			"a & b && c\n",
		},
		{ // 7
			"a(b)",
			"a(b)\n",
			"a(b)\n",
		},
		{ // 8
			"a ( ... )",
			"a(...)\n",
			"a(...)\n",
		},
		{ // 9
			"a(..1)",
			"a(..1)\n",
			"a(..1)\n",
		},
		{ // 10
			"a(,b)",
			"a(,b)\n",
			"a(, b)\n",
		},
		{ // 11
			"a()",
			"a()\n",
			"a()\n",
		},
		{ // 12
			"a(b,c)",
			"a(b,c)\n",
			"a(b, c)\n",
		},
		{ // 13
			"function(a) b",
			"function(a)b\n",
			"function(a) b\n",
		},
		{ // 14
			"function(a = b) c",
			"function(a=b)c\n",
			"function(a = b) c\n",
		},
		{ // 15
			"a=b",
			"a=b\n",
			"a = b\n",
		},
		{ // 16
			"a<-b",
			"a<-b\n",
			"a <- b\n",
		},
		{ // 17
			"a<<-b",
			"a<<-b\n",
			"a <<- b\n",
		},
		{ // 18
			"a->b",
			"a->b\n",
			"a -> b\n",
		},
		{ // 19
			"a->>b",
			"a->>b\n",
			"a ->> b\n",
		},
		{ // 20
			"{}",
			"{}\n",
			"{}\n",
		},
		{ // 21
			"{a}",
			"{\n\ta\n}\n",
			"{\n\ta\n}\n",
		},
		{ // 22
			"{a;b}",
			"{\n\ta\n\tb\n}\n",
			"{\n\ta\n\tb\n}\n",
		},
		{ // 23
			"a^b",
			"a^b\n",
			"a^b\n",
		},
		{ // 24
			"a;b",
			"a\nb\n",
			"a\nb\n",
		},
		{ // 25
			"a\nb",
			"a\nb\n",
			"a\nb\n",
		},
	} {
		for m, input := range test {
			tk := parser.NewStringTokeniser(input)

			if f, err := Parse(&tk); err != nil {
				t.Errorf("test %d.%d: unexpected error: %s", n+1, m+1, err)
			} else if simple := fmt.Sprintf("%s", f); simple != test[1] {
				t.Errorf("test %d.%d.1: expecting output %q, got %q", n+1, m+1, test[1], simple)
			} else if verbose := fmt.Sprintf("%+s", f); verbose != test[2] {
				t.Errorf("test %d.%d.2: expecting output %q, got %q", n+1, m+1, test[2], verbose)
			}
		}
	}
}
