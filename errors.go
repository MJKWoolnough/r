package r

import "errors"

// Errors
var (
	ErrInvalidSimpleExpression     = errors.New("invalid simple expression")
	ErrMissingClosingDoubleBracket = errors.New("missing closing double-bracket")
	ErrMissingClosingParen         = errors.New("missing closing paren")
	ErrMissingComma                = errors.New("missing comma")
	ErrMissingIdentifier           = errors.New("missing identifier")
	ErrMissingIn                   = errors.New("missing in keyword")
	ErrMissingOpeningParen         = errors.New("missing opening paren")
	ErrMissingStatementTerminator  = errors.New("missing statement terminator")
	ErrMissingTerminator           = errors.New("missing terminator")
)
