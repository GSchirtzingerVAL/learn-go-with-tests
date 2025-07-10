package poker_test

import (
	"io"
	"testing"

	poker "gerrod.com/http-server"
)

func TestTape_Write(t *testing.T) {
	file, clean := createTempFile(t, "12345")
	defer clean()

	tape := poker.NewTape(file)

	_, _ = tape.Write([]byte("abc"))

	_, _ = file.Seek(0, io.SeekStart)
	newFileContents, _ := io.ReadAll(file)

	got := string(newFileContents)
	want := "abc"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
