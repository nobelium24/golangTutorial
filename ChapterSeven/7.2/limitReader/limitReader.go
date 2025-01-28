package limitreader

import "io"

type LimitedReader struct {
	r io.Reader
	n int64
}

// func (lm *LimitedReader) Read(b []byte) (int, error) {
// 	nr, err := lm.r.Read(b)
// 	if err != nil {
// 		return 0, err
// 	}
// 	if lm.n <= int64(nr) {
// 		return 0, io.EOF
// 	}
// 	lm.n += int64(nr)
// 	return nr, nil
//

func (lm *LimitedReader) Read(b []byte) (int, error) {
	if lm.n <= 0 {
		return 0, io.EOF
	}

	if int64(len(b)) > lm.n {
		b = b[lm.n:]
	}
	nr, err := lm.r.Read(b)
	lm.n -= int64(nr)
	return nr, err
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &LimitedReader{n: n, r: r}
}
