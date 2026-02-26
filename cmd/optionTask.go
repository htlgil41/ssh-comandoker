package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func OptionTask() int64 {

	var opt int64
	scannerOpt := bufio.NewScanner(os.Stdin)

	fmt.Println("Seleccione una de las opciones para proceder con la transaccion.")
	fmt.Printf("1.Recontruir el VisorWeb\n2. Reiniciar el VisorWeb\n3. Agregar compartida\n0. Salir\n")

	for scannerOpt.Scan() {

		valueOpt, errCon := strconv.ParseInt(scannerOpt.Text(), 10, 64)
		if errCon != nil {

			fmt.Println("No se ha podido leer una de las opciones ya que no es valido lo que ha ingresado")

			fmt.Println("Seleccione una de las opciones para proceder con la transaccion.")
			fmt.Printf("1.Recontruir el VisorWeb\n2. Reiniciar el VisorWeb\n3. Agregar compartida\n")
			continue
		}

		opt = valueOpt
		break
	}

	return opt
}
