package main

import (
	"fmt"
	"./structs"
)

func main () {
	opcion := 0

	input1 := ""
	input2 := ""
	input3 := 0

	cw := structs.ContenidoWeb{
		Multi: []structs.Multimedia{
		},
	}

	for opcion != 5 {
		fmt.Println("1. Agregar Imagen\n2. Agregar Audio\n3. Agregar Video\n4. Mostrar \n5. Salir")
		fmt.Scanln(&opcion)
		if opcion == 1{
			fmt.Print("Titulo: ")
			fmt.Scanln(&input1)
			fmt.Print("Formato: ")
			fmt.Scanln(&input2)
			fmt.Print("Canales: ")
			fmt.Scanln(&input3)

			i01 := structs.Imagen{input1, input2, input3}
			cw.Multi = append(cw.Multi, &i01)
		}else if opcion == 2{
			fmt.Print("Titulo: ")
			fmt.Scanln(&input1)
			fmt.Print("Formato: ")
			fmt.Scanln(&input2)
			fmt.Print("Duracion: ")
			fmt.Scanln(&input3)

			a01 := structs.Audio{input1, input2, input3}
			cw.Multi = append(cw.Multi, &a01)
		}else if opcion == 3{
			fmt.Print("Titulo: ")
			fmt.Scanln(&input1)
			fmt.Print("Formato: ")
			fmt.Scanln(&input2)
			fmt.Print("Frames: ")
			fmt.Scanln(&input3)

			v01 := structs.Video{input1, input2, input3}
			cw.Multi = append(cw.Multi, &v01)
		}else if opcion == 4{
			cw.Mostrar()
		}
	}
	
}