package main

import (
	"fmt"
	"ssh-comandoker/cmd"
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

				fmt.Println("Opcion numeor 01")
				break haveOption
			}
		default:
			{
				fmt.Println("Debe ingresar una de las opciones marcadas")
			}
		}
	}
}
