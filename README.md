# r
--
    import "vimagination.zapto.org/r"

Package r implements an R tokeniser and parser.

## Usage

```go
const (
	TokenWhitespace parser.TokenType = iota
	TokenLineTerminator
	TokenExpressionTerminator
	TokenComment
	TokenStringLiteral
	TokenNumericLiteral
	TokenIntegerLiteral
	TokenComplexLiteral
	TokenBooleanLiteral
	TokenNull
	TokenNA
	TokenIdentifier
	TokenKeyword
	TokenEllipsis
	TokenOperator
	TokenSpecialOperator
	TokenGrouping
)
```
TokenType IDs

```go
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
```
Errors

#### func  SetTokeniser

```go
func SetTokeniser(t *parser.Tokeniser) *parser.Tokeniser
```
SetTokeniser sets the initial tokeniser state of a parser.Tokeniser.

Used if you want to manually tokeniser R source code.

#### type AdditionExpression

```go
type AdditionExpression struct {
	MultiplicationExpression MultiplicationExpression
	AdditionType             AdditionType
	AdditionExpression       *AdditionExpression
	Tokens                   Tokens
}
```

AdditionExpression represents a binary adding or subtracting of two expressions.

#### func (AdditionExpression) Format

```go
func (f AdditionExpression) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type AdditionType

```go
type AdditionType uint8
```

AdditionType determines the type of a AdditionExpression.

```go
const (
	AdditionNone AdditionType = iota
	AdditionAdd
	AdditionSubtract
)
```

#### func (AdditionType) String

```go
func (a AdditionType) String() string
```
String implements the fmt.Stringer interface.

#### type AndExpression

```go
type AndExpression struct {
	NotExpression NotExpression
	AndType       AndType
	AndExpression *AndExpression
	Tokens        Tokens
}
```

AndExpression represents one of two And expressions.

#### func (AndExpression) Format

```go
func (f AndExpression) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type AndType

```go
type AndType uint8
```

AndType defines the type of an AndExpression.

```go
const (
	AndNone AndType = iota
	AndVectorized
	AndNotVectorized
)
```

#### func (AndType) String

```go
func (a AndType) String() string
```
String implements the fmt.Stringer interface.

#### type Arg

```go
type Arg struct {
	QueryExpression *QueryExpression
	Ellipsis        *Token
	Tokens          Tokens
}
```

Arg represents a single argument passed to a function.

#### func (Arg) Format

```go
func (f Arg) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type ArgList

```go
type ArgList struct {
	Args   []Argument
	Tokens Tokens
}
```

ArgList represents a series af arguments accepted by a FunctionDefinition.

#### func (ArgList) Format

```go
func (f ArgList) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type Argument

```go
type Argument struct {
	Identifier *Token
	Default    *Expression
	Tokens     Tokens
}
```


#### func (Argument) Format

```go
func (f Argument) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type AssignmentExpression

```go
type AssignmentExpression struct {
	FormulaeExpression   FormulaeExpression
	AssignmentType       AssignmentType
	AssignmentExpression *AssignmentExpression
	Tokens               Tokens
}
```

AssignmentExpression represents a binding of an expression value.

#### func (AssignmentExpression) Format

```go
func (f AssignmentExpression) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type AssignmentType

```go
type AssignmentType uint8
```

AssignmentType defines the type of assignment in AssignmentExpression.

```go
const (
	AssignmentNone AssignmentType = iota
	AssignmentEquals
	AssignmentLeftAssign
	AssignmentRightAssign
	AssignmentLeftParentAssign
	AssignmentRightParentAssign
)
```

#### func (AssignmentType) String

```go
func (a AssignmentType) String() string
```
String implements the fmt.Stringer interface.

#### type Call

```go
type Call struct {
	Args   []Arg
	Tokens Tokens
}
```

Call represents the arguments passed to a function.

#### func (Call) Format

```go
func (f Call) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type CompoundExpression

```go
type CompoundExpression struct {
	Expressions []Expression
	Tokens      Tokens
}
```

CompoundExpression represents a series of expressions, wrapped in braces, and
seperated by semi-colons, commas, and newlines.

#### func (CompoundExpression) Format

```go
func (f CompoundExpression) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type Error

```go
type Error struct {
	Err     error
	Parsing string
	Token   Token
}
```

Error is a parsing error with trace details.

#### func (Error) Error

```go
func (e Error) Error() string
```
Error returns the error string.

#### func (Error) Unwrap

```go
func (e Error) Unwrap() error
```
Unwrap returns the wrapped error.

#### type ExponentiationExpression

```go
type ExponentiationExpression struct {
	SubsetExpression         SubsetExpression
	ExponentiationExpression *ExponentiationExpression
	Tokens                   Tokens
}
```

ExponentiationExpression represents a exponentiation operation.

#### func (ExponentiationExpression) Format

```go
func (f ExponentiationExpression) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type Expression

```go
type Expression struct {
	FlowControl        *FlowControl
	FunctionDefinition *FunctionDefinition
	QueryExpression    *QueryExpression
	Tokens             Tokens
}
```

Expression represents either a FlowControl, FunctionDefinition, or
QueryExpression.

#### func (Expression) Format

```go
func (f Expression) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type File

```go
type File struct {
	Statements []Expression
	Tokens     Tokens
}
```

File represents a parsed R file.

#### func  Parse

```go
func Parse(t Tokeniser) (*File, error)
```
Parse parses R input into AST.

#### func (File) Format

```go
func (f File) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type FlowControl

```go
type FlowControl struct {
	IfControl     *IfControl
	WhileControl  *WhileControl
	RepeatControl *RepeatControl
	ForControl    *ForControl
	Tokens        Tokens
}
```

FlowControl represents an If-, While-, Repeat-, or For-Control.

#### func (FlowControl) Format

```go
func (f FlowControl) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type ForControl

```go
type ForControl struct {
	Var    *Token
	List   FormulaeExpression
	Expr   Expression
	Tokens Tokens
}
```

ForControl represents a looping branch over an expression.

#### func (ForControl) Format

```go
func (f ForControl) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type FormulaeExpression

```go
type FormulaeExpression struct {
	OrExpression       *OrExpression
	FormulaeExpression *FormulaeExpression
	Tokens             Tokens
}
```

FormulaeExpression represents a model formula.

#### func (FormulaeExpression) Format

```go
func (f FormulaeExpression) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type FunctionDefinition

```go
type FunctionDefinition struct {
	ArgList ArgList
	Body    Expression
	Tokens  Tokens
}
```

FunctionDefinition represents a defined function.

#### func (FunctionDefinition) Format

```go
func (f FunctionDefinition) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type IfControl

```go
type IfControl struct {
	Cond   FormulaeExpression
	Expr   Expression
	Else   *Expression
	Tokens Tokens
}
```

IfControl represents a conditional branch and optional else.

#### func (IfControl) Format

```go
func (f IfControl) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type Index

```go
type Index struct {
	Double bool
	Args   []QueryExpression
	Tokens Tokens
}
```

Index represents either a single or double bracketed indexing operation.

#### func (Index) Format

```go
func (f Index) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type IndexOrCallExpression

```go
type IndexOrCallExpression struct {
	SimpleExpression      *SimpleExpression
	IndexOrCallExpression *IndexOrCallExpression
	Index                 *Index
	Call                  *Call
	Tokens                Tokens
}
```

IndexOrCallExpression represents a possible indexing or function calling
operation.

#### func (IndexOrCallExpression) Format

```go
func (f IndexOrCallExpression) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type MultiplicationExpression

```go
type MultiplicationExpression struct {
	PipeOrSpecialExpression  PipeOrSpecialExpression
	MultiplicationType       MultiplicationType
	MultiplicationExpression *MultiplicationExpression
	Tokens                   Tokens
}
```


#### func (MultiplicationExpression) Format

```go
func (f MultiplicationExpression) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type MultiplicationType

```go
type MultiplicationType uint8
```

MultiplicationType determines the type of a MultiplicationExpression.

```go
const (
	MultiplicationNone MultiplicationType = iota
	MultiplicationMultiply
	MultiplicationDivide
)
```

#### func (MultiplicationType) String

```go
func (m MultiplicationType) String() string
```
String implements the fmt.Stringer interface.

#### type NotExpression

```go
type NotExpression struct {
	Nots                 uint
	RelationalExpression RelationalExpression
	Tokens               Tokens
}
```

NotExpression represents a possibly negated expression.

#### func (NotExpression) Format

```go
func (f NotExpression) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type OrExpression

```go
type OrExpression struct {
	AndExpression AndExpression
	OrType        OrType
	OrExpression  *OrExpression
	Tokens        Tokens
}
```

OrExpression represents one of two Or expressions.

#### func (OrExpression) Format

```go
func (f OrExpression) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type OrType

```go
type OrType uint8
```

OrType defines the type of an OrExpression.

```go
const (
	OrNone OrType = iota
	OrVectorized
	OrNotVectorized
)
```

#### func (OrType) String

```go
func (o OrType) String() string
```
String implements the fmt.Stringer interface.

#### type PipeOrSpecialExpression

```go
type PipeOrSpecialExpression struct {
	SequenceExpression      SequenceExpression
	Operator                *Token
	PipeOrSpecialExpression *PipeOrSpecialExpression
	Tokens                  Tokens
}
```

PipeOrSpecialExpression represetns either a pipe (|>) or special (%%) binary
operation.

#### func (PipeOrSpecialExpression) Format

```go
func (f PipeOrSpecialExpression) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type QueryExpression

```go
type QueryExpression struct {
	AssignmentExpression *AssignmentExpression
	QueryExpression      *QueryExpression
	Tokens               Tokens
}
```

QueryExpression represents a help command.

#### func  WrapQuery

```go
func WrapQuery(p QueryWrappable) *QueryExpression
```
WrapQuery takes one of many types and wraps it in a *QueryExpression

The accepted types/pointers are as follows:

    QueryExpression
    *QueryExpression
    AssignmentExpression
    *AssignmentExpression
    FormulaeExpression
    *FormulaeExpression
    OrExpression
    *OrExpression
    AndExpression
    *AndExpression
    NotExpression
    *NotExpression
    RelationalExpression
    *RelationalExpression
    AdditionExpression
    *AdditionExpression
    MultiplicationExpression
    *MultiplicationExpression
    PipeOrSpecialExpression
    *PipeOrSpecialExpression
    SequenceExpression
    *SequenceExpression
    UnaryExpression
    *UnaryExpression
    ExponentiationExpression
    *ExponentiationExpression
    SubsetExpression
    *SubsetExpression
    ScopeExpression
    *ScopeExpression
    IndexOrCallExpression
    *IndexOrCallExpression
    SimpleExpression
    *SimpleExpression
    CompoundExpression
    *CompoundExpression

#### func (QueryExpression) Format

```go
func (f QueryExpression) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type QueryWrappable

```go
type QueryWrappable interface {
	Type
	// contains filtered or unexported methods
}
```

QueryWrappable represents the types that can be wrapped with WrapQuery and
unwrapped with UnwrapQuery.

#### func  UnwrapQuery

```go
func UnwrapQuery(q *QueryExpression) QueryWrappable
```
UnwrapQuery returns the first value up the QueryExpression chain that contains
all of the information required to rebuild the lower chain.

Possible returns types are as follows:

    *QueryExpression
    *AssignmentExpression
    *FormulaeExpression
    *OrExpression
    *AndExpression
    *NotExpression
    *RelationalExpression
    *AdditionExpression
    *MultiplicationExpression
    *PipeOrSpecialExpression
    *SequenceExpression
    *UnaryExpression
    *ExponentiationExpression
    *SubsetExpression
    *ScopeExpression
    *IndexOrCallExpression
    *SimpleExpression
    *CompoundExpression

#### type RelationalExpression

```go
type RelationalExpression struct {
	AdditionExpression   AdditionExpression
	RelationalOperator   RelationalOperator
	RelationalExpression *RelationalExpression
	Tokens               Tokens
}
```

RelationalExpression represents a logical relationship between two expressions.

#### func (RelationalExpression) Format

```go
func (f RelationalExpression) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type RelationalOperator

```go
type RelationalOperator uint8
```

RelationalOperator defines the type of relationship for a RelationalExpression.

```go
const (
	RelationalNone RelationalOperator = iota
	RelationalGreaterThan
	RelationalGreaterThanOrEqual
	RelationalLessThan
	RelationalLessThanOrEqual
	RelationalEqual
	RelationalNotEqual
)
```

#### func (RelationalOperator) String

```go
func (r RelationalOperator) String() string
```
String implements the fmt.Stringer interface.

#### type RepeatControl

```go
type RepeatControl struct {
	Expr   Expression
	Tokens Tokens
}
```

RepeatControl represents a looping branch.

#### func (RepeatControl) Format

```go
func (f RepeatControl) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type ScopeExpression

```go
type ScopeExpression struct {
	IndexOrCallExpression IndexOrCallExpression
	ScopeExpression       *ScopeExpression
	Tokens                Tokens
}
```

ScopeExpression represents a scoping operation.

#### func (ScopeExpression) Format

```go
func (f ScopeExpression) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type SequenceExpression

```go
type SequenceExpression struct {
	UnaryExpression    UnaryExpression
	SequenceExpression *SequenceExpression
	Tokens             Tokens
}
```

SequenceExpression represents a sequencing operation.

#### func (SequenceExpression) Format

```go
func (f SequenceExpression) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type SimpleExpression

```go
type SimpleExpression struct {
	Identifier              *Token
	Constant                *Token
	Ellipsis                *Token
	ParenthesizedExpression *Expression
	CompoundExpression      *CompoundExpression
	Tokens                  Tokens
}
```

SimpleExpression represents either an Identifier, a Constant, an Ellipsis, a
ParenthesizedExpression, or a CompoundExpression.

#### func (SimpleExpression) Format

```go
func (f SimpleExpression) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type SubsetExpression

```go
type SubsetExpression struct {
	ScopeExpression  ScopeExpression
	SubsetType       SubsetType
	SubsetExpression *SubsetExpression
	Tokens           Tokens
}
```

SubsetExpression represents a subsetting operation.

#### func (SubsetExpression) Format

```go
func (f SubsetExpression) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type SubsetType

```go
type SubsetType uint8
```

SubsetType determines the type of a SubsetExpression.

```go
const (
	SubsetNone SubsetType = iota
	SubsetList
	SubsetStructure
)
```

#### func (SubsetType) String

```go
func (s SubsetType) String() string
```
String implements the fmt.Stringer interface.

#### type Token

```go
type Token struct {
	parser.Token
	Pos, Line, LinePos uint64
}
```

Token represents a single parsed token with source positioning.

#### type Tokeniser

```go
type Tokeniser interface {
	TokeniserState(parser.TokenFunc)
	Iter(func(parser.Token) bool)
	GetError() error
}
```

Tokeniser is an interface representing a tokeniser.

#### type Tokens

```go
type Tokens []Token
```

Tokens is a collection of Token values.

#### type Type

```go
type Type interface {
	fmt.Formatter
	// contains filtered or unexported methods
}
```

Type is an interface satisfied by all R structural types.

#### type UnaryExpression

```go
type UnaryExpression struct {
	UnaryType                []UnaryType
	ExponentiationExpression ExponentiationExpression
	Tokens                   Tokens
}
```

UnaryExpression represents a unary addition or subtraction.

#### func (UnaryExpression) Format

```go
func (f UnaryExpression) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type UnaryType

```go
type UnaryType uint8
```

UnaryType determines the type of operation in a UnaryExpression.

```go
const (
	UnaryAdd UnaryType = iota
	UnaryMinus
)
```

#### func (UnaryType) String

```go
func (u UnaryType) String() string
```
String implements the fmt.Stringer interface.

#### type WhileControl

```go
type WhileControl struct {
	Cond   FormulaeExpression
	Expr   Expression
	Tokens Tokens
}
```

WhileControl represents a looping branch with a single condition.

#### func (WhileControl) Format

```go
func (f WhileControl) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface
