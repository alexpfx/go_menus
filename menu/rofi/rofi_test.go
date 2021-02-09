package rofi

import (
	"log"
	"testing"
)

const (
	StrSilabas = "ba be bi bo bu ta te ti to tu sa se si so su ra ra ri ro ru pa pe pi po pu da de di do du la le li lo lu lha lhe a e i o u nha nhe na ne ni no nu"
)

func Test_NewBuilder(t *testing.T) {
	menu := NewDMenuBuilder().Build()
	
	out, err := menu.Run("abc\nabcd\nxpto\ntesxpto")
	if err != nil {
		log.Fatal(err)
	}
	t.Log(out)
}

func Test_NewBuilderInput(t *testing.T){
	menu := NewInputModeBuilder("informe o texto de entrada").Build()
	
	out, err := menu.Run("")
	if err != nil{
		log.Fatal(err)
	}
	t.Log(out)
	
}

//func Test_RofiDmenuBuilder(t *testing.T) {
//	b := NewDMenuBuilder().Format("s").MultiSelect(true).Prompt("selecione").Sep("|").Build()
//
//	fmt.Println(b.Exec(""))
//}
//
//func Test_Input(t *testing.T) {
//	b := NewUserInputDMenu("Informe a idade")
//	fmt.Println(b.Exec("s\n"))
//}
//
//func Test_generated(t *testing.T) {
//	silabas := strings.Split(StrSilabas, " ")
//	ls := len(silabas)
//
//	input := strings.Builder{}
//	for i := 0; i < 20; i++ {
//		for i := 0; i < 3; i++ {
//			rindex := rand.Intn(ls) - 1
//			s := silabas[rindex]
//			input.WriteString(s)
//		}
//		input.WriteString("\n")
//	}
//
//	dmenu := NewDMenuBuilder().Format("s").
//		MultiSelect(false).
//		AutoSelect(true).
//		Matching("regex").
//		Prompt("action").
//		Build()
//
//	out, err := dmenu.Exec(input.String())
//	if err != nil {
//		log.Println(err.Error())
//	}
//	fmt.Println(out)
//
//}
//
//func Test_RofiDmenuExec(t *testing.T) {
//	input := "[ls]\nls -[la]\nls -l [ll]\n[ca]t\nflameshot [fm]"
//	print("echo " + input)
//
//	dmenu := NewDMenuBuilder().Format("s").
//		MultiSelect(false).
//		AutoSelect(true).
//		Matching("regex").
//		Prompt("action").
//		Build()
//
//	out, err := dmenu.Exec(input)
//	if err != nil {
//		log.Println(err.Error())
//	}
//
//	fmt.Println(out)
//
//}
//
//func Test_appendIf(t *testing.T) {
//	slice := make([]string, 0)
//	slice = append(slice, "base")
//
//	slice = appendIf(slice, "-v", true)
//	slice = appendIf(slice, "-n", 22)
//	slice = appendIf(slice, "-s", "teste")
//
//	slice = appendIf(slice, "-v", false)
//	slice = appendIf(slice, "-n", 0)
//	slice = appendIf(slice, "-s", "")
//
//	assert.EqualValues(t, []string{
//		"base", "-v", "-n", "22", "-s", "teste",
//	}, slice)
//}
