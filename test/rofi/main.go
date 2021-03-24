package main

import (
	"fmt"
	"github.com/alexpfx/go_menus/rofi"
	"log"
)

func main() {

	//teste1()



	teste2()

}

func teste2() {
	const input1 = "1 janeiro sexta true\n2 fevereiro quarta true\n2 março terça false"
	m2 := rofi.NewDMenu("Escolha")

	r, err := m2.Run(input1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(r)
}

func teste1() {
	m := rofi.NewInput("Informe a idade: ")

	run, err := m.Run(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(run)
}
