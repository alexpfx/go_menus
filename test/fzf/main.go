package main

import (
	"fmt"
	"github.com/alexpfx/go_menus/fzf"
)

func main() {
	testFzfBuilder()
	testFzfBuilder2()
}
func testFzfBuilder() {
	const input1 = "1 janeiro sexta true\n2 fevereiro quarta true\n2 março terça false"

	m := fzf.New("Selecione", false, "", "")

	run, err := m.Run(input1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(run)

}

func testFzfBuilder2() {
	const input2 = "1;gremio fbpa\n2;palmeiras fc\n3;figueira fc"

	m := fzf.NewIndexed("Selecione")

	run, err := m.Run(input2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(run)

}
