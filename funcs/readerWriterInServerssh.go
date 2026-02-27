package funcs

import (
	"fmt"
	"io"
	"strings"
)

func ReadWriterInServerSsh(ip string, in io.WriteCloser, out io.Reader) {
	for {
		bufferResponse := make([]byte, 1024)
		n, err := out.Read(bufferResponse)
		if n > 0 {

			text := string(bufferResponse[:n])
			if strings.Contains(text, "password for") {
				in.Write([]byte("PASSWORD FOR EXECUTE SUDO COMAND\n"))
				fmt.Println(text)
			}
		}

		if err != nil && err == io.EOF {

			fmt.Printf("Se ha producido un erro inesperado al tratar de ejecutar el comando en %s \n", ip)
			break
		}
	}
}
