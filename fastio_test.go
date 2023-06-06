package fastio_test

import (
	"bytes"
	"io"
	"reflect"
	"strings"
	"testing"

	"github.com/johejo/fastio"
)

func TestReadAll(t *testing.T) {
	s := strings.Repeat("0123456789", 8192)
	want, err := io.ReadAll(strings.NewReader(s))
	if err != nil {
		t.Fatal(err)
	}
	got, err := fastio.ReadAll(strings.NewReader(s))
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want=%v, got=%v", want, got)
	}
}

func Benchmark_ioReadAll_stringsReader(b *testing.B) {
	s := strings.Repeat("0123456789", 8192)
	sr := strings.NewReader(s)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := io.ReadAll(sr)
		if err != nil {
			b.Fatal(err)
		}
		sr.Reset(s)
	}
}

func Benchmark_ioReadAll_ioNopCloser_stringsReader(b *testing.B) {
	s := strings.Repeat("0123456789", 8192)
	readers := make([]io.Reader, 0, b.N)
	for i := 0; i < b.N; i++ {
		readers = append(readers, io.NopCloser(strings.NewReader(s)))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := io.ReadAll(readers[i])
		if err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_fastioReadAll_ioNopCloser_stringsReader(b *testing.B) {
	s := strings.Repeat("0123456789", 8192)
	readers := make([]io.Reader, 0, b.N)
	for i := 0; i < b.N; i++ {
		readers = append(readers, io.NopCloser(strings.NewReader(s)))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := fastio.ReadAll(readers[i])
		if err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_fastioReadAll_stringsReader(b *testing.B) {
	s := strings.Repeat("0123456789", 8192)
	sr := strings.NewReader(s)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := fastio.ReadAll(sr)
		if err != nil {
			b.Fatal(err)
		}
		sr.Reset(s)
	}
}

func Benchmark_ioReadAll_bytesReader(b *testing.B) {
	bs := bytes.Repeat([]byte("0123456789"), 8192)
	br := bytes.NewReader(bs)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := io.ReadAll(br)
		if err != nil {
			b.Fatal(err)
		}
		br.Reset(bs)
	}
}

func Benchmark_fastioReadAll_bytesReader(b *testing.B) {
	bs := bytes.Repeat([]byte("0123456789"), 8192)
	br := bytes.NewReader(bs)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := fastio.ReadAll(br)
		if err != nil {
			b.Fatal(err)
		}
		br.Reset(bs)
	}
}

func Benchmark_ioReadAll_bytesBuffer(b *testing.B) {
	bs := bytes.Repeat([]byte("0123456789"), 8192)
	buffers := make([]*bytes.Buffer, 0, b.N)
	for i := 0; i < b.N; i++ {
		buffers = append(buffers, bytes.NewBuffer(bs))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := io.ReadAll(buffers[i])
		if err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_fastioReadAll_bytesBuffer(b *testing.B) {
	bs := bytes.Repeat([]byte("0123456789"), 8192)
	buffers := make([]*bytes.Buffer, 0, b.N)
	for i := 0; i < b.N; i++ {
		buffers = append(buffers, bytes.NewBuffer(bs))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := fastio.ReadAll(buffers[i])
		if err != nil {
			b.Fatal(err)
		}
	}
}
