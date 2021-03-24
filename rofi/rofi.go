package menu

import (
	"github.com/alexpfx/go_menus/internal/util"
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


func NewRofiInputBuilder(prompt string) RofiBuilder {
	return RofiBuilder{
		dMenu:    true,
		themeStr: "listview { enabled: false;}",
		format:   "f",
		prompt:   prompt,
	}
}


type rofiMenu struct {
	cmd  string
	args []string
}

func (r rofiMenu) Run(input string) (string, error) {
	//return util.RunCmdWithInput(input, r.cmd, r.args)
	return util.RunCmdWithNoInput(r.cmd, r.args)
}

type RofiBuilder struct {
	prompt     string
	AutoSelect bool
	themeStr   string
	//	//'s' selected string
	//	//'i' index (0 - (N-1))
	//	//'d' index (1 - N)
	//	//'q' quote string
	//	//'p' Selected string stripped from Pango markup (Needs to be a valid string)
	//	//'f' filter string (user action)
	//	//'F' quoted filter string (user action)
	format string
	Mode   string
	dMenu  bool
}


func (r RofiBuilder) Build() Menu {
	argSlice := make([]string, 0)
	
	argSlice = util.AppendIf(argSlice, rofiMode, r.Mode)
	argSlice = util.AppendIf(argSlice, rofiDmenu, r.dMenu)
	
	argSlice = util.AppendIf(argSlice, rofiPrompt, r.prompt)
	argSlice = util.AppendIf(argSlice, rofiAutoSelect, r.AutoSelect)
	
	argSlice = util.AppendIf(argSlice, rofiThemeStr, r.themeStr)
	argSlice = util.AppendIf(argSlice, rofiFormat, r.format)
	
	return rofiMenu{
		cmd:  rofiCmd,
		args: argSlice,
	}
	
}
