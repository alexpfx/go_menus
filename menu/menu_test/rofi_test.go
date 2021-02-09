package menu_test

import (
	"github.com/alexpfx/go_menus/menu"
	"log"
	"testing"
)

const (
	StrSilabas = "ba be bi bo bu ta te ti to tu sa se si so su ra ra ri ro ru pa pe pi po pu da de di do du la le li lo lu lha lhe a e i o u nha nhe na ne ni no nu"
)

//FIXME deixar teste estático e não interativo
func Test_NewBuilder(t *testing.T) {
	rmenu := menu.NewRofiDMenuBuilder().Build()
	
	out, err := rmenu.Run("abc\nabcd\nxpto\ntesxpto")
	if err != nil {
		log.Fatal(err)
	}
	t.Log(out)
}

//FIXME deixar teste estático e não interativo
func Test_NewBuilderInput(t *testing.T) {
	rmenu := menu.NewRofiInputBuilder("informe o texto de entrada").Build()
	
	out, err := rmenu.Run("")
	if err != nil {
		log.Fatal(err)
	}
	t.Log(out)
	
}
