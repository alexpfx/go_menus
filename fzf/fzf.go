package fzf

import (
	"github.com/alexpfx/go_menus/internal/util"
	"github.com/alexpfx/go_menus/menu"
)

const (
	cmd        = "fzf"
	prompt     = "--prompt"
	autoSelect = "-1"
	withNth    = "--with-nth"
	delimiter  = "-d"
)

func New(prompt string, autoSelect bool, withNth string, delimiter string) menu.Menu {
	b := builder{
		prompt:     prompt,
		autoSelect: autoSelect,
		withNth:    withNth,
		delimiter:  delimiter,
	}

	args := b.buildArgs()

	return fzfMenu{
		args: args,
	}

}

func NewIndexed(prompt string) menu.Menu {
	b := builder{
		prompt:     prompt,
		autoSelect: true,
		withNth:    "2",
		delimiter:  ";",
	}

	args := b.buildArgs()

	return fzfMenu{
		args: args,
	}

}

type builder struct {
	prompt     string
	autoSelect bool
	//WithNth limita a saída. formato: "1,2,.."
	withNth string
	//Delimiter define o separador para Nth e WithNth
	delimiter string
}

func (f fzfMenu) Run(input interface{}) (string, error) {
	return util.RunCmdWithInput(input.(string), cmd, f.args)
}

type fzfMenu struct {
	args []string
}

func (f builder) buildArgs() []string {
	argSlice := make([]string, 0)

	argSlice = util.AppendIf(argSlice, prompt, f.prompt)
	argSlice = util.AppendIf(argSlice, autoSelect, f.autoSelect)
	argSlice = util.AppendIf(argSlice, withNth, f.withNth)
	argSlice = util.AppendIf(argSlice, delimiter, f.delimiter)

	return argSlice
}
