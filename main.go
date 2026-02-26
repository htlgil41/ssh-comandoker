package main

import (
	"fmt"
	"ssh-comandoker/cmd"
	"ssh-comandoker/funcs"
)

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

				funcs.CreateSessionSsh("")
				break haveOption
			}
		default:
			{
				fmt.Println("Debe ingresar una de las opciones marcadas")
			}
		}
	}
}
