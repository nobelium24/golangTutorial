package reader

import (
	"io"
)

type StringReader struct {
	s   string
	pos int
}

// func (sr *StringReader) Read(p []byte) (int, error) {
// 	if len(sr.s) == len(p) {
// 		copy(p, []byte(sr.s))
// 		return len(p), nil
// 	}
// 	if sr.pos == 0 {
// 		newStr := sr.s[:len(p)]
// 		in := copy(p, []byte(newStr))
// 		sr.pos = in
// 		return in, nil
// 	} else {
// 		newStr := sr.s[sr.pos:]
// 		in := copy(p, []byte(newStr))
// 		sr.pos = sr.pos + in
// 		return in, nil
// 	}

// }

func (sr *StringReader) Read(p []byte) (int, error) {
	if sr.pos >= len(sr.s) {
		return 0, io.EOF
	}
	n := copy(p, sr.s[sr.pos:])
	sr.pos += n

	return n, nil
}

func NewReader(input string) *StringReader {
	return &StringReader{s: input, pos: 0}
}
