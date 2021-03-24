package menu

import (
	"github.com/alexpfx/go_menus/internal/util"
)

const (
	cmd        = "fzf"
	prompt     = "--Prompt"
	autoSelect = "-1"
	withNth = "--with-nth"
	delimiter = "-d"
)

type MenuBuilder struct {
	Prompt     string
	AutoSelect bool
	//WithNth limita a saída. formato: "1,2,.."
	WithNth string
	//Delimiter define o separador para Nth e WithNth
	Delimiter string
}

func (f fzfmenu) Run(input string) (string, error) {
	return util.RunCmdWithInput(input, f.cmd, f.args)
}

type fzfmenu struct {
	cmd  string
	args []string
}

func (f *MenuBuilder) Build() Menu {
	argSlice := make([]string, 0)

	argSlice = util.AppendIf(argSlice, prompt, f.Prompt)
	argSlice = util.AppendIf(argSlice, autoSelect, f.AutoSelect)
	argSlice = util.AppendIf(argSlice, withNth, f.WithNth)
	argSlice = util.AppendIf(argSlice, delimiter, f.Delimiter)

	return fzfmenu{
		cmd:  cmd,
		args: argSlice,
	}
}
