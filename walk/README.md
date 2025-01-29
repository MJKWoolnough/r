# walk
--
    import "vimagination.zapto.org/r/walk"

Package walk provides a R type walker.

## Usage

#### func  Walk

```go
func Walk(t r.Type, fn Handler) error
```
Walk calls the Handle function on the given interface for each non-nil,
non-Token field of the given R type.

#### type Handler

```go
type Handler interface {
	Handle(r.Type) error
}
```

Handler is used to process R types.

#### type HandlerFunc

```go
type HandlerFunc func(r.Type) error
```

HandlerFunc wraps a func to implement Handler interface.

#### func (HandlerFunc) Handle

```go
func (h HandlerFunc) Handle(t r.Type) error
```
Handle implements the Handler interface.
