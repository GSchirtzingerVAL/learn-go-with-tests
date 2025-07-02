package di

import (
	"fmt"
	"io"
	"log"
)

func Greet(writer io.Writer, name string) int {
	bytes, err := fmt.Fprintf(writer, "Hello, %s", name)
	if err != nil {
		log.Println(err)
	}
	return bytes
}
