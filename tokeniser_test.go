package r

import (
	"testing"

	"vimagination.zapto.org/parser"
)

func TestTokeniser(t *testing.T) {
	for n, test := range [...]struct {
		Input  string
		Output []parser.Token
	}{
		{ // 1
			" \t\n\v\f\u00a0\ufeff",
			[]parser.Token{
				{Type: TokenWhitespace, Data: " \t"},
				{Type: TokenLineTerminator, Data: "\n"},
				{Type: TokenWhitespace, Data: "\v\f\u00a0\ufeff"},
				{Type: parser.TokenDone, Data: ""},
			},
		},
		{ // 2
			"\n\r \u2028\u2029",
			[]parser.Token{
				{Type: TokenLineTerminator, Data: "\n\r"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenLineTerminator, Data: "\u2028\u2029"},
				{Type: parser.TokenDone, Data: ""},
			},
		},
		{ // 3
			"#A comment\n # Another comment",
			[]parser.Token{
				{Type: TokenComment, Data: "#A comment"},
				{Type: TokenLineTerminator, Data: "\n"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenComment, Data: "# Another comment"},
				{Type: parser.TokenDone, Data: ""},
			},
		},
		{ // 4
			"\"abc\"'def'\"a\\n\\t\\\"\"",
			[]parser.Token{
				{Type: TokenStringLiteral, Data: "\"abc\""},
				{Type: TokenStringLiteral, Data: "'def'"},
				{Type: TokenStringLiteral, Data: "\"a\\n\\t\\\"\""},
				{Type: parser.TokenDone, Data: ""},
			},
		},
		{ // 5
			"0 1 23 0x1a2f30 .123 0.456 123.456 9.8e+7 7E-6 0x123.FEDpFF 0xFFP+FF",
			[]parser.Token{
				{Type: TokenNumericLiteral, Data: "0"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenNumericLiteral, Data: "1"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenNumericLiteral, Data: "23"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenNumericLiteral, Data: "0x1a2f30"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenNumericLiteral, Data: ".123"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenNumericLiteral, Data: "0.456"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenNumericLiteral, Data: "123.456"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenNumericLiteral, Data: "9.8e+7"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenNumericLiteral, Data: "7E-6"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenNumericLiteral, Data: "0x123.FEDpFF"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenNumericLiteral, Data: "0xFFP+FF"},
				{Type: parser.TokenDone, Data: ""},
			},
		},
		{ // 6
			"1L 99.88L 1e1L 1.2E-23L 0x123L 0xFEEDp101L",
			[]parser.Token{
				{Type: TokenIntegerLiteral, Data: "1L"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenIntegerLiteral, Data: "99.88L"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenIntegerLiteral, Data: "1e1L"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenIntegerLiteral, Data: "1.2E-23L"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenIntegerLiteral, Data: "0x123L"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenIntegerLiteral, Data: "0xFEEDp101L"},
				{Type: parser.TokenDone, Data: ""},
			},
		},
	} {
		p := parser.NewStringTokeniser(test.Input)

		SetTokeniser(&p)

		for m, tkn := range test.Output {
			tk, _ := p.GetToken()
			if tk.Type != tkn.Type {
				if tk.Type == parser.TokenError {
					t.Errorf("test %d.%d: unexpected error: %s", n+1, m+1, tk.Data)
				} else {
					t.Errorf("test %d.%d: Incorrect type, expecting %d, got %d", n+1, m+1, tkn.Type, tk.Type)
				}

				break
			} else if tk.Data != tkn.Data {
				t.Errorf("test %d.%d: Incorrect data, expecting %q, got %q", n+1, m+1, tkn.Data, tk.Data)

				break
			}
		}
	}
}
