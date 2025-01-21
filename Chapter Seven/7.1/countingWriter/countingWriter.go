package countingwriter

import "io"

type NewWriter struct {
	writer io.Writer
	count  int64
}

func (nw *NewWriter) Write(p []byte) (int, error) {
	n, err := nw.writer.Write(p)
	nw.count += int64(n)
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	nw := &NewWriter{writer: w}
	return nw, &nw.count
}
