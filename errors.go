package r

import "errors"

// Errors
var (
	ErrInvalidCharacter            = errors.New("invalid character")
	ErrInvalidNumber               = errors.New("invalid number")
	ErrInvalidOperator             = errors.New("invalid operator")
	ErrInvalidSimpleExpression     = errors.New("invalid simple expression")
	ErrInvalidString               = errors.New("invalid string")
	ErrMissingClosingDoubleBracket = errors.New("missing closing double-bracket")
	ErrMissingClosingParen         = errors.New("missing closing paren")
	ErrMissingComma                = errors.New("missing comma")
	ErrMissingIdentifier           = errors.New("missing identifier")
	ErrMissingIn                   = errors.New("missing in keyword")
	ErrMissingOpeningParen         = errors.New("missing opening paren")
	ErrMissingStatementTerminator  = errors.New("missing statement terminator")
	ErrMissingTerminator           = errors.New("missing terminator")
)
