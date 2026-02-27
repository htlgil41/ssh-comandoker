package main

import (
	"fmt"
	"ssh-comandoker/cmd"
	"ssh-comandoker/funcs"
)

var ips []string = []string{
	":22",
	":22",
	":22",
	":22",
	":22",
	":22",
	":22",
	":22",
	":22",
	":22",
	":22",
	":22",
	":22",
	":22",
	":22",
}

func main() {

haveOption:
	for {

		optionsToDo := cmd.OptionTask()
		if optionsToDo == 0 {

			fmt.Println("Chaoo!!!")
			break haveOption
		}

		switch optionsToDo {

		case 1:
			{

				for _, v := range ips {

					go funcs.RestartAllServicesVisor(v)
				}
				break haveOption
			}
		default:
			{
				fmt.Println("Debe ingresar una de las opciones marcadas")
			}
		}
	}

	select {}
}
