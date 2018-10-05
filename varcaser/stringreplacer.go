package varcaser

import (
	"io"
)

// replacer is the interface that a replacement algorithm needs to implement. Provided for compatibility with the
// strings.replacer interface. Copied from strings.replacer: https://golang.org/src/strings/replace.go
// It is the good way to standardize all string-replacers. The interface allows to use varcaser in any place, where
// strings.replacer is expected.
type replacer interface {
	Replace(s string) string
	WriteString(w io.Writer, s string) (n int, err error)
}
