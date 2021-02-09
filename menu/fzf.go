package menu

import (
	"github.com/alexpfx/go_menus/internal/util"
)

const (
	fzfCmd        = "fzf"
	fzfPrompt     = "--prompt"
	fzfAutoselect = "-1"
	fzfWithNth    = "--with-nth"
	fzfDelimiter  = "-d"
)

func NewFzfBuilder() Builder {
	return &fzfBuilder{}
}

type FzfBuilder interface {
	Builder
	WithNth(nth, sep string) Builder
}

type fzfMenu struct {
	cmd  string
	args []string
}

func (f fzfMenu) Run(input string) (string, error) {
	return util.RunCmdWithInput(input, f.cmd, f.args)
}

type fzfBuilder struct {
	prompt       string
	autoSelect   bool
	withNth      string
	nthDelimiter string
}

func (f *fzfBuilder) WithNth(nth string, sep string) Builder {
	f.withNth = nth
	f.nthDelimiter = sep
	return f
}

func (f *fzfBuilder) Prompt(s string) Builder {
	f.prompt = s
	return f
}

func (f *fzfBuilder) AutoSelect(b bool) Builder {
	f.autoSelect = b
	return f
}

func (f *fzfBuilder) Build() Menu {
	argSlice := make([]string, 0)
	
	argSlice = util.AppendIf(argSlice, fzfPrompt, f.prompt)
	argSlice = util.AppendIf(argSlice, fzfAutoselect, f.autoSelect)
	argSlice = util.AppendIf(argSlice, fzfWithNth, f.withNth)
	if f.withNth != "" {
		argSlice = util.AppendIf(argSlice, fzfDelimiter, f.nthDelimiter)
	}
	
	return fzfMenu{
		cmd:  fzfCmd,
		args: argSlice,
	}
}
