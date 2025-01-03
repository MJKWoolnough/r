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
			";,",
			[]parser.Token{
				{Type: TokenExpressionTerminator, Data: ";"},
				{Type: TokenExpressionTerminator, Data: ","},
				{Type: parser.TokenDone, Data: ""},
			},
		},
		{ // 4
			"#A comment\n # Another comment",
			[]parser.Token{
				{Type: TokenComment, Data: "#A comment"},
				{Type: TokenLineTerminator, Data: "\n"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenComment, Data: "# Another comment"},
				{Type: parser.TokenDone, Data: ""},
			},
		},
		{ // 5
			"\"abc\"'def'\"a\\n\\t\\\"\"",
			[]parser.Token{
				{Type: TokenStringLiteral, Data: "\"abc\""},
				{Type: TokenStringLiteral, Data: "'def'"},
				{Type: TokenStringLiteral, Data: "\"a\\n\\t\\\"\""},
				{Type: parser.TokenDone, Data: ""},
			},
		},
		{ // 6
			"\"'\\\"\\n\\r\\t\\b\\a\\f\\v\\\"\"",
			[]parser.Token{
				{Type: TokenStringLiteral, Data: "\"'\\\"\\n\\r\\t\\b\\a\\f\\v\\\"\""},
				{Type: parser.TokenDone, Data: ""},
			},
		},
		{ // 7
			"\"\\132\\142=\\064\\062\"",
			[]parser.Token{
				{Type: TokenStringLiteral, Data: "\"\\132\\142=\\064\\062\""},
				{Type: parser.TokenDone, Data: ""},
			},
		},
		{ // 8
			"\"\\0a\"",
			[]parser.Token{
				{Type: parser.TokenError, Data: "invalid string"},
			},
		},
		{ // 9
			"\"\\x7A\\x42=\\x34\\x32\"",
			[]parser.Token{
				{Type: TokenStringLiteral, Data: "\"\\x7A\\x42=\\x34\\x32\""},
				{Type: parser.TokenDone, Data: ""},
			},
		},
		{ // 10
			"\"\\xz\"",
			[]parser.Token{
				{Type: parser.TokenError, Data: "invalid string"},
			},
		},
		{ // 11
			"\"\\m\"",
			[]parser.Token{
				{Type: parser.TokenError, Data: "invalid string"},
			},
		},
		{ // 12
			"\"abc",
			[]parser.Token{
				{Type: parser.TokenError, Data: "invalid string"},
			},
		},
		{ // 13
			"\"\\u2190 \\u{800}\"",
			[]parser.Token{
				{Type: TokenStringLiteral, Data: "\"\\u2190 \\u{800}\""},
				{Type: parser.TokenDone, Data: ""},
			},
		},
		{ // 14
			"\"\\u{x}\"",
			[]parser.Token{
				{Type: parser.TokenError, Data: "invalid string"},
			},
		},
		{ // 15
			"\"\\u{f\"",
			[]parser.Token{
				{Type: parser.TokenError, Data: "invalid string"},
			},
		},
		{ // 16
			"\"\\u{fffff}\"",
			[]parser.Token{
				{Type: parser.TokenError, Data: "invalid string"},
			},
		},
		{ // 17
			"\"\\ufff\"",
			[]parser.Token{
				{Type: parser.TokenError, Data: "invalid string"},
			},
		},
		{ // 18
			"\"\\U01000000 \\U{2190}\"",
			[]parser.Token{
				{Type: TokenStringLiteral, Data: "\"\\U01000000 \\U{2190}\""},
				{Type: parser.TokenDone, Data: ""},
			},
		},
		{ // 19
			"\"\\U{x}\"",
			[]parser.Token{
				{Type: parser.TokenError, Data: "invalid string"},
			},
		},
		{ // 20
			"\"\\U{f\"",
			[]parser.Token{
				{Type: parser.TokenError, Data: "invalid string"},
			},
		},
		{ // 21
			"\"\\U{fffffffff}\"",
			[]parser.Token{
				{Type: parser.TokenError, Data: "invalid string"},
			},
		},
		{ // 22
			"\"\\Ufffffff\"",
			[]parser.Token{
				{Type: parser.TokenError, Data: "invalid string"},
			},
		},
		{ // 23
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
		{ // 24
			"Inf NaN",
			[]parser.Token{
				{Type: TokenNumericLiteral, Data: "Inf"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenNumericLiteral, Data: "NaN"},
				{Type: parser.TokenDone, Data: ""},
			},
		},
		{ // 25
			"0xz",
			[]parser.Token{
				{Type: parser.TokenError, Data: "invalid number"},
			},
		},
		{ // 26
			"0x1.",
			[]parser.Token{
				{Type: parser.TokenError, Data: "invalid number"},
			},
		},
		{ // 27
			"0x1pz",
			[]parser.Token{
				{Type: parser.TokenError, Data: "invalid number"},
			},
		},
		{ // 28
			"1ea",
			[]parser.Token{
				{Type: parser.TokenError, Data: "invalid number"},
			},
		},
		{ // 29
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
		{ // 30
			"1i .2i 9.8e1i 0x123i 0x456.ffi 0xapbi",
			[]parser.Token{
				{Type: TokenComplexLiteral, Data: "1i"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenComplexLiteral, Data: ".2i"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenComplexLiteral, Data: "9.8e1i"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenComplexLiteral, Data: "0x123i"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenComplexLiteral, Data: "0x456.ffi"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenComplexLiteral, Data: "0xapbi"},
				{Type: parser.TokenDone, Data: ""},
			},
		},
		{ // 31
			"TRUE FALSE",
			[]parser.Token{
				{Type: TokenBooleanLiteral, Data: "TRUE"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenBooleanLiteral, Data: "FALSE"},
				{Type: parser.TokenDone, Data: ""},
			},
		},
		{ // 32
			"NULL",
			[]parser.Token{
				{Type: TokenNull, Data: "NULL"},
				{Type: parser.TokenDone, Data: ""},
			},
		},
		{ // 33
			"NA NA_character_ NA_integer_ NA_real_ NA_complex_",
			[]parser.Token{
				{Type: TokenNA, Data: "NA"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenNA, Data: "NA_character_"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenNA, Data: "NA_integer_"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenNA, Data: "NA_real_"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenNA, Data: "NA_complex_"},
				{Type: parser.TokenDone, Data: ""},
			},
		},
		{ // 34
			"a bc def a1 b_c abc.def",
			[]parser.Token{
				{Type: TokenIdentifier, Data: "a"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenIdentifier, Data: "bc"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenIdentifier, Data: "def"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenIdentifier, Data: "a1"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenIdentifier, Data: "b_c"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenIdentifier, Data: "abc.def"},
				{Type: parser.TokenDone, Data: ""},
			},
		},
		{ // 35
			".a",
			[]parser.Token{
				{Type: TokenIdentifier, Data: ".a"},
				{Type: parser.TokenDone, Data: ""},
			},
		},
		{ // 36
			"if else repeat while function for in next break",
			[]parser.Token{
				{Type: TokenKeyword, Data: "if"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenKeyword, Data: "else"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenKeyword, Data: "repeat"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenKeyword, Data: "while"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenKeyword, Data: "function"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenKeyword, Data: "for"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenKeyword, Data: "in"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenKeyword, Data: "next"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenKeyword, Data: "break"},
				{Type: parser.TokenDone, Data: ""},
			},
		},
		{ // 37
			"... ..1 ..2 ..99",
			[]parser.Token{
				{Type: TokenEllipsis, Data: "..."},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenEllipsis, Data: "..1"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenEllipsis, Data: "..2"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenEllipsis, Data: "..99"},
				{Type: parser.TokenDone, Data: ""},
			},
		},
		{ // 38
			"+ - * / ^ > >= < <= == != ! & && | ~ -> <- ->> <<- $ : :: || |> @ = ? ??",
			[]parser.Token{
				{Type: TokenOperator, Data: "+"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenOperator, Data: "-"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenOperator, Data: "*"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenOperator, Data: "/"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenOperator, Data: "^"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenOperator, Data: ">"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenOperator, Data: ">="},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenOperator, Data: "<"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenOperator, Data: "<="},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenOperator, Data: "=="},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenOperator, Data: "!="},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenOperator, Data: "!"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenOperator, Data: "&"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenOperator, Data: "&&"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenOperator, Data: "|"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenOperator, Data: "~"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenOperator, Data: "->"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenOperator, Data: "<-"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenOperator, Data: "->>"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenOperator, Data: "<<-"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenOperator, Data: "$"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenOperator, Data: ":"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenOperator, Data: "::"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenOperator, Data: "||"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenOperator, Data: "|>"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenOperator, Data: "@"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenOperator, Data: "="},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenOperator, Data: "?"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenOperator, Data: "?"},
				{Type: TokenOperator, Data: "?"},
				{Type: parser.TokenDone, Data: ""},
			},
		},
		{ // 39
			"%% %/% %*% %in% %o% %x% %||%",
			[]parser.Token{
				{Type: TokenSpecialOperator, Data: "%%"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenSpecialOperator, Data: "%/%"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenSpecialOperator, Data: "%*%"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenSpecialOperator, Data: "%in%"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenSpecialOperator, Data: "%o%"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenSpecialOperator, Data: "%x%"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenSpecialOperator, Data: "%||%"},
				{Type: parser.TokenDone, Data: ""},
			},
		},
		{ // 40
			"<<",
			[]parser.Token{
				{Type: parser.TokenError, Data: "invalid operator"},
			},
		},
		{ // 41
			"%\n%",
			[]parser.Token{
				{Type: parser.TokenError, Data: "invalid operator"},
			},
		},
		{ // 42
			"£",
			[]parser.Token{
				{Type: parser.TokenError, Data: "invalid character"},
			},
		},
		{ // 43
			"([{[[]]}])]",
			[]parser.Token{
				{Type: TokenGrouping, Data: "("},
				{Type: TokenGrouping, Data: "["},
				{Type: TokenGrouping, Data: "{"},
				{Type: TokenGrouping, Data: "[["},
				{Type: TokenGrouping, Data: "]]"},
				{Type: TokenGrouping, Data: "}"},
				{Type: TokenGrouping, Data: "]"},
				{Type: TokenGrouping, Data: ")"},
				{Type: parser.TokenError, Data: "invalid character"},
			},
		},
		{ // 44
			"(",
			[]parser.Token{
				{Type: TokenGrouping, Data: "("},
				{Type: parser.TokenError, Data: "unexpected EOF"},
			},
		},
		{ // 45
			"[",
			[]parser.Token{
				{Type: TokenGrouping, Data: "["},
				{Type: parser.TokenError, Data: "unexpected EOF"},
			},
		},
		{ // 46
			"{",
			[]parser.Token{
				{Type: TokenGrouping, Data: "{"},
				{Type: parser.TokenError, Data: "unexpected EOF"},
			},
		},
		{ // 47
			"{]",
			[]parser.Token{
				{Type: TokenGrouping, Data: "{"},
				{Type: parser.TokenError, Data: "invalid character"},
			},
		},
		{ // 48
			"[)",
			[]parser.Token{
				{Type: TokenGrouping, Data: "["},
				{Type: parser.TokenError, Data: "invalid character"},
			},
		},
		{ // 49
			"[[]",
			[]parser.Token{
				{Type: TokenGrouping, Data: "[["},
				{Type: parser.TokenError, Data: "invalid character"},
			},
		},
		{ // 50
			"[[}",
			[]parser.Token{
				{Type: TokenGrouping, Data: "[["},
				{Type: parser.TokenError, Data: "invalid character"},
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
