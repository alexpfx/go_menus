package rofi

import (
	"fmt"
	"github.com/alexpfx/go_menus/internal/util"
	"github.com/alexpfx/go_menus/menu"
)

const (
	rofiCmd        = "rofi"
	rofiPrompt     = "-p"
	rofiAutoSelect = "-auto-select"
	rofiThemeStr   = "-theme-str"
	rofiFormat     = "-format"
	rofiMode       = "-show"
	rofiDmenu      = "-dmenu"
)

func NewInput(prompt string) menu.Menu {
	b := builder{
		dMenu:    true,
		themeStr: "listview { enabled: false;}",
		format:   "f",
		prompt:   prompt,
	}
	args := b.buildArgs()
	return rofiMenu{
		args: args,
	}
}

func NewDMenu(prompt string) menu.Menu {
	args := builder{
		prompt:     prompt,
		autoSelect: false,
		format:     "i;s",
		mode:       "",
		dMenu:      true,
	}.buildArgs()
	return rofiDMenu{
		args: args,
	}
}

type rofiMenu struct {
	args []string
}
type rofiDMenu struct {
	args []string
}

func (r rofiMenu) Run(interface{}) (string, error) {
	//return util.RunCmdWithInput(input, r.cmd, r.args)
	return util.RunCmdWithNoInput(rofiCmd, r.args)
}

func (r rofiDMenu) Run(input interface{}) (string, error) {
	fmt.Println("input ", input)
	return util.RunCmdWithInput(input.(string), rofiCmd, r.args)
}

type builder struct {
	prompt     string
	autoSelect bool
	themeStr   string
	//	//'s' selected string
	//	//'i' index (0 - (N-1))
	//	//'d' index (1 - N)
	//	//'q' quote string
	//	//'p' Selected string stripped from Pango markup (Needs to be a valid string)
	//	//'f' filter string (user action)
	//	//'F' quoted filter string (user action)
	format string
	mode   string
	dMenu  bool
}

func (r builder) buildArgs() []string {
	argSlice := make([]string, 0)

	argSlice = util.AppendIf(argSlice, rofiMode, r.mode)
	argSlice = util.AppendIf(argSlice, rofiDmenu, r.dMenu)

	argSlice = util.AppendIf(argSlice, rofiPrompt, r.prompt)
	argSlice = util.AppendIf(argSlice, rofiAutoSelect, r.autoSelect)

	argSlice = util.AppendIf(argSlice, rofiThemeStr, r.themeStr)
	argSlice = util.AppendIf(argSlice, rofiFormat, r.format)

	return argSlice

}
