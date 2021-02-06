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
	mmenu := NewBuilder().
		Prompt("Mes:\n").(menu2.Builder).AutoSelect(true).(Builder).WithNth("1,2,3", "").Build()
	
	out, err := mmenu.Run("1 janeiro sexta true\n2 fevereiro quarta true\n2 março terça false")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("mes: ", out)
	
}

func TestNewBuilderWithIndex(t *testing.T) {
	mmenu := NewBuilder().Prompt("selecione").
		AutoSelect(true).
		Prompt("selecione o time").(Builder).
		WithNth("2", ";").Build()
	
	out, err := mmenu.Run("1;gremio fbpa\n2;palmeiras fc\n3;figueira fc")
	
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("time: ", out)
	
}
