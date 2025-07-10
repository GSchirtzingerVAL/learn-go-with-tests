package poker

import (
	"io"
	"os"
)

type tape struct {
	file *os.File
}

func NewTape(file *os.File) *tape {
	return &tape{file}
}

func (t *tape) Write(p []byte) (n int, err error) {
	_ = t.file.Truncate(0)
	_, _ = t.file.Seek(0, io.SeekStart)
	return t.file.Write(p)
}
