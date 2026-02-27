package funcs

import (
	"fmt"
	"io"
	"strings"
)

func ReadWriteInServer(ip string, in io.WriteCloser, out io.Reader) {
	for {
		bufferResponse := make([]byte, 1024)
		n, err := out.Read(bufferResponse)
		if n > 0 {

			text := string(bufferResponse[:0])
			if strings.Contains(text, "password for") {

				println(text)
			}
		}

		if err != nil && err != io.EOF {

			fmt.Println("Se ha producido un erro inesperado al tratar de ejecutar el comando en ", ip)
			break
		}
	}
}
