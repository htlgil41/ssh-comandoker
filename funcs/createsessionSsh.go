package funcs

import (
	"fmt"
	"ssh-comandoker/config/ssgclient"

	"golang.org/x/crypto/ssh"
)

func CreateSessionSsh(ip string) {

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

	_, errIn := sessionSsh.StdinPipe()
	if errIn != nil {

		fmt.Println("Error al establcer el pipe en la entrada de datos")
		return
	}
	_, errOut := sessionSsh.StdoutPipe()
	if errOut != nil {

		fmt.Println("Error al establcer el pipe en la salida de datos")
		return
	}

	sessionSsh.Start("Comando")
	go ReadWriteInServer()

	sessionSsh.Wait()
	fmt.Println("Terminado para ", ip)
}
