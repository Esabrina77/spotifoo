package main

import (
	"fmt"
	"spotifoo/roots"
	initTemplate "spotifoo/templates"
)

func main() {
	fmt.Println("server is running...")

	fmt.Println("")
	fmt.Print("CLICK HERE to OPEN PAGE---> http://localhost:8080/jul \n")
	fmt.Println("")
	fmt.Println("TO STOP THE SERVER , PRESS  'ctrl+C' ")
	initTemplate.InitTemplate()
	roots.InitServe()

}
