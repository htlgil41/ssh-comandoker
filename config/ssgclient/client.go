package ssgclient

import (
	"fmt"
	"time"

	"golang.org/x/crypto/ssh"
)

var configClientSsh *ssh.ClientConfig = &ssh.ClientConfig{
	User: "aaron",
	Auth: []ssh.AuthMethod{
		ssh.Password("PASSWORD"),
	},
	HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	Timeout:         time.Second * 5,
}

func CreateClientSsh(ip string) (*ssh.Client, error) {

	cliente, errCreateClient := ssh.Dial("tcp", ip, configClientSsh)
	if errCreateClient != nil {

		return nil, fmt.Errorf("Error al tratar de crear la conexion con %q en la ip %q", errCreateClient.Error(), ip)
	}

	return cliente, nil
}
