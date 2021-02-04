package fzf

import (
	"fmt"
	menu2 "github.com/alexpfx/go_menus/menu"
	"testing"
)

func TestNewBuilder(t *testing.T) {

	menu := NewBuilder().Build()

	out, err := menu.Run("abc\nabcd\nxpto\ntesxpto")
	if err != nil {
		fmt.Println(err)
	}
	t.Log(out)

}

func TestNewBuilderWithArgs(t *testing.T) {
	mmenu := NewBuilder().Prompt("Mes:\n").(menu2.Builder).AutoSelect(true).Build()

	out, err := mmenu.Run("janeiro\nfevereiro\nmarço")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("mes: ", out)

}

func TestNewBuilderWithIndex(t *testing.T) {

	mmenu := NewBuilder().Prompt("selecione").
		AutoSelect(true).
		Prompt("selecione o time").(Builder).
		Prefix(2, "\n", func(input string, index int) string {
			return fmt.Sprintf("%d %s", index, input)
		}).Build()

	out, err := mmenu.Run("gremio\npalmeiras\nfigueira")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("time: ", out)

}
