package fzf

import (
	"fmt"
	"testing"
)

func TestNewBuilder(t *testing.T) {

	menu := NewBuilder().AutoSelect(true).Build()

	out, err := menu.Run("abc\nabcd\nxpto")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("teste ok")
	fmt.Println(out)

}