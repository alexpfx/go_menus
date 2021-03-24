package main

import (
	"fmt"
	"github.com/alexpfx/go_menus/menu"
)

func main() {
	testFzfBuilder()
	testFzfBuilder2()
}
func testFzfBuilder() {
	const input1 = "1 janeiro sexta true\n2 fevereiro quarta true\n2 março terça false"

	b := menu.MenuBuilder{
		Prompt:     "Selecione",
		AutoSelect: false,
		WithNth:    "1,2,3",
		Delimiter:  "",
	}

	m := b.Build()

	run, err := m.Run(input1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(run)

}

func testFzfBuilder2() {
	const input2 = "1;gremio fbpa\n2;palmeiras fc\n3;figueira fc"

	b := menu.MenuBuilder{
		Prompt:     "Time\n",
		AutoSelect: true,
		WithNth:    "2",
		Delimiter:  ";",
	}

	m := b.Build()

	run, err := m.Run(input2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(run)

}
