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




type rofiMenu struct {
	cmd  string
	args []string
}

func (r rofiMenu) Run(input string) (string, error) {
	return util.RunCmdWithInput(input, r.cmd, r.args)
}

type RofiBuilder struct {
	Prompt     string
	AutoSelect bool
	ThemeStr   string
	//	//'s' selected string
	//	//'i' index (0 - (N-1))
	//	//'d' index (1 - N)
	//	//'q' quote string
	//	//'p' Selected string stripped from Pango markup (Needs to be a valid string)
	//	//'f' filter string (user action)
	//	//'F' quoted filter string (user action)
	Format string
	Mode   string
	DMenu  bool
}


func (r RofiBuilder) Build() Menu {
	argSlice := make([]string, 0)
	
	argSlice = util.AppendIf(argSlice, rofiMode, r.Mode)
	argSlice = util.AppendIf(argSlice, rofiDmenu, r.DMenu)
	
	argSlice = util.AppendIf(argSlice, rofiPrompt, r.Prompt)
	argSlice = util.AppendIf(argSlice, rofiAutoSelect, r.AutoSelect)
	
	argSlice = util.AppendIf(argSlice, rofiThemeStr, r.ThemeStr)
	argSlice = util.AppendIf(argSlice, rofiFormat, r.Format)
	
	return rofiMenu{
		cmd:  rofiCmd,
		args: argSlice,
	}
	
}
