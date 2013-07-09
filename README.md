# lw
--
    import "github.com/MJKWoolnough/io-lw"

Limited Writer acts similarly to io.LimitedReader in so far as it wraps a writer
and only allows a certain number of bytes to be writter.

## Usage

#### func  LimitWriter

```go
func LimitWriter(w io.Writer, n int64) io.Writer
```

#### type LimitedWriter

```go
type LimitedWriter struct {
	W io.Writer
	N int64
}
```


#### func (*LimitedWriter) Write

```go
func (l *LimitedWriter) Write(p []byte) (n int, err error)
```

#### type NoSpace

```go
type NoSpace struct{}
```


#### func (*NoSpace) Error

```go
func (n *NoSpace) Error() string
```
