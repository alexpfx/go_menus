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

type RofiBuilder interface {
	Builder
	
	//	//'s' selected string
	//	//'i' index (0 - (N-1))
	//	//'d' index (1 - N)
	//	//'q' quote string
	//	//'p' Selected string stripped from Pango markup (Needs to be a valid string)
	//	//'f' filter string (user action)
	//	//'F' quoted filter string (user action)
	Format(string) Builder
	ThemeStr(string) Builder
	
	Mode(string) Builder
	Dmenu(bool) Builder
}

func NewRofiDMenuBuilder() Builder {
	return rofiBuilder{
		dmenu: true,
	}
}

func NewRofiInputBuilder(prompt string) Builder {
	return rofiBuilder{
		dmenu: true,
		themeStr: "listview { enabled: false;}",
		format: "f",
		prompt: prompt,
	}
}

type rofiMenu struct {
	cmd  string
	args []string
}

func (r rofiMenu) Run(_ string) (string, error) {
	return util.RunCmdWithNoImput(r.cmd, r.args)
}

type rofiBuilder struct {
	prompt     string
	autoSelect bool
	themeStr   string
	format     string
	mode       string
	dmenu bool
}

func (r rofiBuilder) Format(s string) Builder {
	r.format = s
	return r
}

func (r rofiBuilder) ThemeStr(s string) Builder {
	r.themeStr = s
	return r
}

func (r rofiBuilder) Prompt(s string) Builder {
	r.prompt = s
	return r
}

func (r rofiBuilder) AutoSelect(b bool) Builder {
	r.autoSelect = b
	return r
}

func (r rofiBuilder) Mode(s string) Builder {
	r.mode = s
	return r
}

func (r rofiBuilder) Build() Menu {
	argSlice := make([]string, 0)
	
	argSlice = util.AppendIf(argSlice, rofiMode, r.mode)
	argSlice = util.AppendIf(argSlice, rofiDmenu, r.dmenu)
	
	argSlice = util.AppendIf(argSlice, rofiPrompt, r.prompt)
	argSlice = util.AppendIf(argSlice, rofiAutoSelect, r.autoSelect)
	
	argSlice = util.AppendIf(argSlice, rofiThemeStr, r.themeStr)
	argSlice = util.AppendIf(argSlice, rofiFormat, r.format)
	
	return rofiMenu{
		cmd:  rofiCmd,
		args: argSlice,
	}
	
}
