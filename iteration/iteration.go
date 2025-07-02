package iteration

import (
	"strings"
)

func Repeat(character string, repeat int) string {
	var repeated strings.Builder
	for i := 0; i < repeat; i++ {
		repeated.WriteString(character)
	}
	return repeated.String()
}
