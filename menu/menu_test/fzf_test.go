package menu_test

import (
	"fmt"
	"github.com/alexpfx/go_menus/menu"
	"testing"
)

//FIXME deixar teste estático e não interativo
func TestNewBuilder(t *testing.T) {
	
	rmenu := menu.NewFzfBuilder().Build()
	
	out, err := rmenu.Run("abc\nabcd\nxpto\ntesxpto")
	if err != nil {
		fmt.Println(err)
	}
	t.Log(out)
	
}

//FIXME deixar teste estático e não interativo
func TestNewBuilderWithArgs(t *testing.T) {
	rmenu := menu.NewFzfBuilder().
		Prompt("Mes:\n").AutoSelect(true).(menu.FzfBuilder).WithNth("1,2,3", "").Build()
	
	out, err := rmenu.Run("1 janeiro sexta true\n2 fevereiro quarta true\n2 março terça false")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("mes: ", out)
	
}

func TestNewBuilderWithIndex(t *testing.T) {
	mmenu := menu.NewFzfBuilder().Prompt("selecione").
		AutoSelect(true).
		Prompt("selecione o time").(menu.FzfBuilder).
		WithNth("2", ";").Build()
	
	out, err := mmenu.Run("1;gremio fbpa\n2;palmeiras fc\n3;figueira fc")
	
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("time: ", out)
	
}
