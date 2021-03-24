package main

import (
	"fmt"
	"github.com/alexpfx/go_menus/menu"
	"log"
)

func main() {
	const input1 = "1 janeiro sexta true\n2 fevereiro quarta true\n2 março terça false"

	b := menu.RofiBuilder{
		Prompt:     "data",
		AutoSelect: true,
		ThemeStr:   "",
		Format:     "i",
		Mode:       "",
		DMenu:      true,
	}

	m := b.Build()

	run, err := m.Run(input1)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(run)

}
