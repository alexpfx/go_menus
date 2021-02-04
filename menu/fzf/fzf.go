package fzf

import (
	"fmt"
	"github.com/alexpfx/go_menus/internal/util"
	"github.com/alexpfx/go_menus/menu"
	"io"
	"log"
	"strconv"
	"strings"
)

const (
	cmd        = "fzf"
	prompt     = "--prompt"
	autoSelect = "-1"
	withPrefix = "--with-nth"
)

func NewBuilder() menu.Builder {
	return &fzfMenuBuilder{}
}



type Builder interface {

	menu.Builder

	Prefix(position int, sep string, prefixFunction func(input string, index int) string) menu.Builder
}

func (f fzfMenu) Run(input string) (string, error) {
	if f.prefixFunction != nil {
		splited := strings.Split(input, f.sep)

		finput := func(in io.WriteCloser) {
			for index, line := range splited {
				pre := f.prefixFunction(line, index)
				fmt.Println("pre: ", pre)
				_, err := fmt.Fprintf(in, "%v\n", pre)
				if err != nil {
					log.Fatal(err)
				}
			}
		}
		return util.RunCmdWithReader(finput, f.cmd, f.args)
	} else {
		return util.RunCmdWithInput(input, f.cmd, f.args)
	}

}

type fzfMenu struct {
	cmd            string
	args           []string
	prefixPosition int
	sep            string
	prefixFunction func(input string, index int) string
}

type fzfMenuBuilder struct {
	prompt         string
	autoSelect     bool
	sep            string
	prefixPosition string
	prefixFunction func(input string, index int) string
}

func (f *fzfMenuBuilder) Prefix(position int, sep string, prefixFunction func(input string, index int) string) menu.Builder {
	f.prefixPosition = strconv.Itoa(position)
	f.prefixFunction = prefixFunction
	f.sep = sep
	return f
}


func (f *fzfMenuBuilder) Prompt(s string) menu.Builder {
	f.prompt = s
	return f
}

func (f *fzfMenuBuilder) AutoSelect(b bool) menu.Builder{
	f.autoSelect = b
	return f
}

func (f *fzfMenuBuilder) Build() menu.Menu {
	argSlice := make([]string, 0)

	argSlice = util.AppendIf(argSlice, prompt, f.prompt)
	argSlice = util.AppendIf(argSlice, autoSelect, f.autoSelect)
	argSlice = util.AppendIf(argSlice, withPrefix, f.prefixPosition)

	return fzfMenu{
		cmd:            cmd,
		args:           argSlice,
		sep:            f.sep,
		prefixFunction: f.prefixFunction,
	}
}
