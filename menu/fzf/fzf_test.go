package fzf

import (
	"fmt"
	"testing"
)

func TestNewBuilder(t *testing.T) {

	menu := NewBuilder().Build()

	out, err := menu.Run("abc\nabcd\nxpto\ntesxpto")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("teste ok")
	fmt.Println(out)

}

func TestNewBuilderWithArgs(t *testing.T) {

	menu := NewBuilder().Prompt("select:\n").AutoSelect(true).Build()

	out, err := menu.Run("abcde\nabcde1\nxpto")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("teste ok")
	fmt.Println(out)

}