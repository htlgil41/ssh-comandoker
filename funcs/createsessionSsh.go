package funcs

import (
	"fmt"
	"io"
	"ssh-comandoker/config/ssgclient"
	"strings"

	"golang.org/x/crypto/ssh"
)

func RestartAllServicesVisor(ip string) {

	cliente, errCliente := ssgclient.CreateClientSsh(ip)
	if errCliente != nil {

		fmt.Println(errCliente.Error())
		return
	}
	defer cliente.Close()

	sessionSsh, errSessionSsh := cliente.NewSession()
	if errSessionSsh != nil {

		fmt.Println("Ha ocurddio un erro al crear la sesison en ", ip)
		return
	}
	defer sessionSsh.Close()
	sessionSsh.RequestPty("xterm", 40, 80, ssh.TerminalModes{})

	in, errIn := sessionSsh.StdinPipe()
	if errIn != nil {

		fmt.Println("Error al establcer el pipe en la entrada de datos")
		return
	}
	out, errOut := sessionSsh.StdoutPipe()
	if errOut != nil {

		fmt.Println("Error al establcer el pipe en la salida de datos")
		return
	}

	fmt.Println("Todo listo para ejcutar el comando")
	erroExecute := sessionSsh.Start("sudo docker compose -p visor -f docker-compose.visor.yml restart")
	if erroExecute != nil {

		fmt.Println("Error al ejecutar el comando deseado")
		return
	}
	go func() {
		for {
			bufferResponse := make([]byte, 1024)
			n, err := out.Read(bufferResponse)
			if n > 0 {

				text := string(bufferResponse[:n])
				if strings.Contains(text, "password for") {
					in.Write([]byte("PASSWORD FOR EXECUTE SUDO COMAND\n"))
					fmt.Println(text)
				} else {

					fmt.Println(string(bufferResponse[:n]))
				}
			}

			if err != nil && err == io.EOF {

				fmt.Println("Se ha producido un erro inesperado al tratar de ejecutar el comando en ", ip)
				break
			}
		}
	}()

	sessionSsh.Wait()
	fmt.Println("Terminado para ", ip)
}
