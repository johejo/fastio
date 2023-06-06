# fastio

Optimize io package for some io.Reader implementations

## fastio.ReadAll

Optimized for in-memory io.Reader (*bytes.Buffer, *bytes.Reader, *strings.Reader) and *os.File, and also applies optimizations by unwrapping it if it is wrapped with io.NopCloser.

# License

BSD-3-Clause

Much was copied from Go's standard library. That part follows the original license.
