package structs

import ("fmt")

type ContenidoWeb struct {
	Multi []Multimedia
}

func (cw *ContenidoWeb) Mostrar(){
	for _, f := range cw.Multi {
		f.Mostrar()
	}
}


type Multimedia interface {
	Mostrar()
}

func (i *Imagen) Mostrar() {
	fmt.Println("Titulo: ",i.Titulo)
	fmt.Println("Formato: ",i.Formato)
	fmt.Println("Canales: ",i.Canales)
}

type Imagen struct {
	Titulo string
	Formato string
	Canales int
}

func (a *Audio) Mostrar() {
	fmt.Println("Titulo: ",a.Titulo)
	fmt.Println("Formato: ",a.Formato)
	fmt.Println("Duracion: ",a.Duracion)
}

type Audio struct {
	Titulo string
	Formato string
	Duracion int
}

func (v *Video) Mostrar() {
	fmt.Println("Titulo: ",v.Titulo)
	fmt.Println("Formato: ",v.Formato)
	fmt.Println("Duracion: ",v.Frames)
}

type Video struct {
	Titulo string
	Formato string
	Frames int
}