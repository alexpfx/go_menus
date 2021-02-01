package fzf

import (
	"github.com/alexpfx/go_menus/internal/util"
	"github.com/alexpfx/go_menus/menu"
)

const (
	cmd        = "fzf"
	prompt     = "--prompt"
	autoSelect = "-1"
	sep        = "-d"
)

func NewBuilder() menu.Builder {
	return &fzfMenuBuilder{}
}

func (f fzfMenu) Run(input string) (string, error) {

	return util.RunCmdWithInput(input, f.cmd, f.args)
}

type fzfMenu struct {
	cmd  string
	args []string
}

type fzfMenuBuilder struct {
	prompt     string
	autoSelect bool
	sep        string
}

func (f *fzfMenuBuilder) Prompt(s string) menu.Builder {
	f.prompt = s
	return f
}

func (f *fzfMenuBuilder) AutoSelect(b bool) menu.Builder {
	f.autoSelect = b
	return f
}

func (f *fzfMenuBuilder) Sep(s string) menu.Builder {
	f.sep = s
	return f
}

func (f *fzfMenuBuilder) Build() menu.Menu {
	argSlice := make([]string, 0)

	argSlice = util.AppendIf(argSlice, prompt, f.prompt)
	argSlice = util.AppendIf(argSlice, autoSelect, f.autoSelect)
	argSlice = util.AppendIf(argSlice, sep, f.sep)

	return fzfMenu{
		cmd:  cmd,
		args: argSlice,
	}
}
