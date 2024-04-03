# go-iopsep

Package **iopsep** provides a limited-reader that reads a single paragraph from an io.Reader, for the Go programming language.

A paragraph is terminated by one of these:

* `"\n\n"`
* `"\r\r"`
* `"\n\r\n\r"`
* `"\n\r\u0085"`
* `"\r\n\r\n"`
* `"\r\n\u0085"`
* `"\u0085\u0085"`
* `"\u0085\n\r"`
* `"\u0085\r\n"`
* `"\u2029"`
* `io.EOF`

(This package takes care of the complexity with the paragraph terminator characters.)

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-iopsep

[![GoDoc](https://godoc.org/github.com/reiver/go-iopsep?status.svg)](https://godoc.org/github.com/reiver/go-iopsep)

## Usage

```golang
var readcloser io.ReadCloser = iopsep.NewParagraphReadCloser(reader)
```

## Import

To import package **iopsep** use `import` code like the follownig:
```
import "github.com/reiver/go-iopsep"
```

## Installation

To install package **iopsep** do the following:
```
GOPROXY=direct go get https://github.com/reiver/go-iopsep
```

## Author

Package **iopsep** was written by [Charles Iliya Krempeaux](http://reiver.link)
