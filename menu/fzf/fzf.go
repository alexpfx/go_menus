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

func NewBuilder() menu.Builder {
	return &fzfMenuBuilder{}
}

type Builder interface {
	menu.Builder
	WithNth(nth, sep string) menu.Builder
}

type fzfMenu struct {
	cmd  string
	args []string
}

func (f fzfMenu) Run(input string) (string, error) {
	return util.RunCmdWithInput(input, f.cmd, f.args)
}

type fzfMenuBuilder struct {
	prompt       string
	autoSelect   bool
	withNth      string
	nthDelimiter string
}

func (f *fzfMenuBuilder) WithNth(nth string, sep string) menu.Builder {
	f.withNth = nth
	f.nthDelimiter = sep
	return f
}

func (f *fzfMenuBuilder) Prompt(s string) menu.Builder {
	f.prompt = s
	return f
}

func (f *fzfMenuBuilder) AutoSelect(b bool) menu.Builder {
	f.autoSelect = b
	return f
}

func (f *fzfMenuBuilder) Build() menu.Menu {
	argSlice := make([]string, 0)
	
	argSlice = util.AppendIf(argSlice, prompt, f.prompt)
	argSlice = util.AppendIf(argSlice, autoSelect, f.autoSelect)
	argSlice = util.AppendIf(argSlice, withNth, f.withNth)
	if f.withNth != "" {
		argSlice = util.AppendIf(argSlice, delimiter, f.nthDelimiter)
	}
	
	return fzfMenu{
		cmd:  cmd,
		args: argSlice,
	}
}
