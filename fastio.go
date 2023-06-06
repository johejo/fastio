// Original license for golang/go
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fastio

import (
	"bytes"
	"io"
	"os"
	"reflect"
	"strings"
)

// ReadAll wraps io.ReadAll with optimization to some io.Reader implementations.
func ReadAll(r io.Reader) ([]byte, error) {
	switch ir := r.(type) {
	case *bytes.Buffer:
		return ir.Bytes(), nil
	case *bytes.Reader:
		b := make([]byte, ir.Len())
		_, err := io.ReadFull(ir, b)
		return b, err
	case *strings.Reader:
		b := make([]byte, ir.Len())
		_, err := io.ReadFull(ir, b)
		return b, err
	case *os.File:
		return readAllFile(ir)
	}
	if rr, ok := unwrapNopCloser(r); ok {
		return ReadAll(rr)
	}

	return io.ReadAll(r)
}

// Copied from os.ReadFile in os/file.go

func readAllFile(f *os.File) ([]byte, error) {
	var size int
	if info, err := f.Stat(); err == nil {
		size64 := info.Size()
		if int64(int(size64)) == size64 {
			size = int(size64)
		}
	}
	size++ // one byte for final read at EOF

	// If a file claims a small size, read at least 512 bytes.
	// In particular, files in Linux's /proc claim size 0 but
	// then do not work right if read in small pieces,
	// so an initial read of 1 byte would not work correctly.
	if size < 512 {
		size = 512
	}

	data := make([]byte, 0, size)
	for {
		if len(data) >= cap(data) {
			d := append(data[:cap(data)], 0)
			data = d[:len(data)]
		}
		n, err := f.Read(data[len(data):cap(data)])
		data = data[:len(data)+n]
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			return data, err
		}
	}
}

// Copied from net/http/transfer.go

var nopCloserType = reflect.TypeOf(io.NopCloser(nil))
var nopCloserWriterToType = reflect.TypeOf(io.NopCloser(struct {
	io.Reader
	io.WriterTo
}{}))

// unwrapNopCloser return the underlying reader and true if r is a NopCloser
// else it return false.
func unwrapNopCloser(r io.Reader) (underlyingReader io.Reader, isNopCloser bool) {
	switch reflect.TypeOf(r) {
	case nopCloserType, nopCloserWriterToType:
		return reflect.ValueOf(r).Field(0).Interface().(io.Reader), true
	default:
		return nil, false
	}
}
