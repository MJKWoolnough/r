package r

import (
	"fmt"
	"strings"
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
			"function(a,b) b",
			"function(a,b)b\n",
			"function(a, b) b\n",
		},
		{ // 15
			"function(a = b) c",
			"function(a=b)c\n",
			"function(a = b) c\n",
		},
		{ // 16
			"a=b",
			"a=b\n",
			"a = b\n",
		},
		{ // 17
			"a<-b",
			"a<-b\n",
			"a <- b\n",
		},
		{ // 18
			"a<<-b",
			"a<<-b\n",
			"a <<- b\n",
		},
		{ // 19
			"a->b",
			"a->b\n",
			"a -> b\n",
		},
		{ // 20
			"a->>b",
			"a->>b\n",
			"a ->> b\n",
		},
		{ // 21
			"{}",
			"{}\n",
			"{}\n",
		},
		{ // 22
			"{a}",
			"{\n\ta\n}\n",
			"{\n\ta\n}\n",
		},
		{ // 23
			"{a;b}",
			"{\n\ta\n\tb\n}\n",
			"{\n\ta\n\tb\n}\n",
		},
		{ // 24
			"a^b",
			"a^b\n",
			"a^b\n",
		},
		{ // 25
			"a;b",
			"a\nb\n",
			"a\nb\n",
		},
		{ // 26
			"a\nb",
			"a\nb\n",
			"a\nb\n",
		},
		{ // 27
			"for(a in b) c",
			"for(a in b)c\n",
			"for (a in b) c\n",
		},
		{ // 28
			"~a",
			"~a\n",
			"~ a\n",
		},
		{ // 29
			"a~b",
			"a~b\n",
			"a ~ b\n",
		},
		{ // 30
			"if(a)b",
			"if(a)b\n",
			"if (a) b\n",
		},
		{ // 31
			"if(a)b else c",
			"if(a)b else c\n",
			"if (a) b else c\n",
		},
		{ // 32
			"a[0]",
			"a[0]\n",
			"a[0]\n",
		},
		{ // 33
			"a[b,c]",
			"a[b,c]\n",
			"a[b, c]\n",
		},
		{ // 34
			"a[[0]]",
			"a[[0]]\n",
			"a[[0]]\n",
		},
		{ // 35
			"a*b",
			"a*b\n",
			"a * b\n",
		},
		{ // 36
			"a/b",
			"a/b\n",
			"a / b\n",
		},
		{ // 37
			"a * b/c",
			"a*b/c\n",
			"a * b / c\n",
		},
		{ // 38
			"!a",
			"!a\n",
			"!a\n",
		},
		{ // 39
			"!!!!a",
			"!!!!a\n",
			"!!!!a\n",
		},
		{ // 40
			"a|b",
			"a|b\n",
			"a | b\n",
		},
		{ // 41
			"a||b",
			"a||b\n",
			"a || b\n",
		},
		{ // 42
			"a | b||c",
			"a|b||c\n",
			"a | b || c\n",
		},
		{ // 43
			"a|>b",
			"a|>b\n",
			"a |> b\n",
		},
		{ // 44
			"a%%b",
			"a%%b\n",
			"a %% b\n",
		},
		{ // 45
			"a%b%c",
			"a%b%c\n",
			"a %b% c\n",
		},
		{ // 46
			"a?b",
			"a?b\n",
			"a ? b\n",
		},
		{ // 47
			"a??b",
			"a??b\n",
			"a ?? b\n",
		},
		{ // 48
			"?a",
			"?a\n",
			"? a\n",
		},
		{ // 49
			"??a",
			"??a\n",
			"?? a\n",
		},
		{ // 50
			"???a",
			"???a\n",
			"??? a\n",
		},
		{ // 51
			"a>b",
			"a>b\n",
			"a > b\n",
		},
		{ // 52
			"a>=b",
			"a>=b\n",
			"a >= b\n",
		},
		{ // 53
			"a<b",
			"a<b\n",
			"a < b\n",
		},
		{ // 54
			"a<=b",
			"a<=b\n",
			"a <= b\n",
		},
		{ // 55
			"a==b",
			"a==b\n",
			"a == b\n",
		},
		{ // 56
			"a!=b",
			"a!=b\n",
			"a != b\n",
		},
		{ // 57
			"repeat a",
			"repeat a\n",
			"repeat a\n",
		},
		{ // 58
			"a::b",
			"a::b\n",
			"a::b\n",
		},
		{ // 59
			"a:b",
			"a:b\n",
			"a:b\n",
		},
		{ // 60
			"...",
			"...\n",
			"...\n",
		},
		{ // 61
			"..1",
			"..1\n",
			"..1\n",
		},
		{ // 62
			"(a)",
			"(a)\n",
			"(a)\n",
		},
		{ // 63
			"(a + b) * c",
			"(a+b)*c\n",
			"(a + b) * c\n",
		},
		{ // 64
			"a$b",
			"a$b\n",
			"a$b\n",
		},
		{ // 65
			"a@b",
			"a@b\n",
			"a@b\n",
		},
		{ // 66
			"+a",
			"+a\n",
			"+a\n",
		},
		{ // 67
			"-a",
			"-a\n",
			"-a\n",
		},
		{ // 68
			"while(a)b",
			"while(a)b\n",
			"while (a) b\n",
		},
		{ // 69
			"# abc\na",
			"a\n",
			"# abc\na\n",
		},
		{ // 70
			"# abc\n# def\na",
			"a\n",
			"# abc\n# def\na\n",
		},
		{ // 71
			"# abc\n# def\na\n#ghi\nb",
			"a\nb\n",
			"# abc\n# def\na\n#ghi\nb\n",
		},
		{ // 72
			"# abc\na #ghi",
			"a\n",
			"# abc\na #ghi\n\n",
		},
		{ // 73
			"# abc\n\n\n# def\na #ghi",
			"a\n",
			"# abc\n\n# def\na #ghi\n\n",
		},
		{ // 74
			"# abc\n\n\n# def\na #ghi\n\n#123",
			"a\n",
			"# abc\n\n# def\na #ghi\n\n#123\n",
		},
		{ // 75
			"{\n# abc\n\t# def\na #ghi\n\n#123\n}",
			"{\n\ta\n}\n",
			"{\n\t# abc\n\t# def\n\ta #ghi\n\t\n#123\n}\n",
		},
		{ // 76
			"(a\n#abc\n)",
			"(a)\n",
			"(a\n#abc\n)\n",
		},
		{ // 77
			"a(b#a comment\n)",
			"a(b)\n",
			"a(b #a comment\n)\n",
		},
		{ // 78
			"a(#abc\nb#def\n#ghi\n\n#jkl\n)",
			"a(b)\n",
			"a(#abc\nb #def\n#ghi\n\n#jkl\n)\n",
		},
		{ // 78
			"a(#abc\nb#def\n,#ghi\nc\n#jkl\n)",
			"a(b,c)\n",
			"a(#abc\nb #def\n, #ghi\nc #jkl\n)\n",
		},
	} {
		for m, input := range test {
			tk := parser.NewStringTokeniser(input)

			if f, err := Parse(&tk); err != nil {
				t.Errorf("test %d.%d: unexpected error: %s", n+1, m+1, err)
			} else if simple := fmt.Sprintf("%s", f); simple != test[1] {
				t.Errorf("test %d.%d.1: expecting output %q, got %q", n+1, m+1, test[1], simple)
			} else if verbose := fmt.Sprintf("%+s", f); verbose != test[2] && (m != 1 || !strings.ContainsRune(test[0], '#')) {
				t.Errorf("test %d.%d.2: expecting output %q, got %q", n+1, m+1, test[2], verbose)
			}
		}
	}
}
