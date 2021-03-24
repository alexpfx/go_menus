package rofi

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os/exec"
)

const rofi = "rofi"

func NewDMenuBuilder() DMenuBuilder {
	return &dmenuBuilder{}
}

func NewMenuBuilder() MenuBuilder {
	return &menuBuilder{}
}

func NewMessage(message string) Menu {
	return NewMenuBuilder().ErrorMessage(message).Build()
}

func NewUserInputDMenu(prompt string) DMenu {
	b := NewDMenuBuilder().
		Format("f").
		Prompt(prompt).
		ThemeStr("listview { enabled: false;}").
		Build()

	return b

}

type MenuBuilder interface {
	ErrorMessage(message string) MenuBuilder
	Build() Menu
}

type DMenuBuilder interface {
	Prompt(string) DMenuBuilder
	MultiSelect(bool) DMenuBuilder
	AutoSelect(bool) DMenuBuilder
	//'s' selected string
	//'i' index (0 - (N-1))
	//'d' index (1 - N)
	//'q' quote string
	//'p' Selected string stripped from Pango markup (Needs to be a valid string)
	//'f' filter string (user action)
	//'F' quoted filter string (user action)
	Format(string) DMenuBuilder

	/*
	 -matching method

	       Specify  the matching algorithm used.  Current the fol‐
	       lowing methods are supported.

	       • normal: match the int string

	       • regex: match a regex input

	       • glob: match a glob pattern

	       • fuzzy: do a fuzzy match

	       Default: normal
	*/
	Matching(string) DMenuBuilder
	Sep(string) DMenuBuilder
	ThemeStr(string) DMenuBuilder
	Build() DMenu
}

type Menu interface {
	Exec(input string) (string, error)
}

type menuBuilder struct {
	message string
}

func (mb *menuBuilder) ErrorMessage(message string) MenuBuilder {
	mb.message = message
	return mb
}

func (mb *menuBuilder) Build() Menu {
	argSlice := make([]string, 0)
	argSlice = appendIf(argSlice, "-e", mb.message)

	return menu{
		cmd:  rofi,
		args: argSlice,
	}

}

type menu struct {
	cmd  string
	args []string
}

func (m menu) Exec(input string) (string, error) {
	rofi := exec.Command(m.cmd, m.args...)

	var stdout, stderr bytes.Buffer

	rofi.Stdout = &stdout
	rofi.Stderr = &stderr

	err := rofi.Run()
	check(err)

	return string(stdout.Bytes()), fmt.Errorf(string(stderr.Bytes()))
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type dmenuBuilder struct {
	autoSelect  bool
	prompt      string
	multiSelect bool
	format      string
	matching    string
	sep         string
	themeStr    string
}

func (dmb *dmenuBuilder) ThemeStr(s string) DMenuBuilder {
	dmb.themeStr = s
	return dmb
}

func (dmb *dmenuBuilder) Matching(s string) DMenuBuilder {
	dmb.matching = s
	return dmb
}

func (dmb *dmenuBuilder) Prompt(s string) DMenuBuilder {
	dmb.prompt = s
	return dmb
}

func (dmb *dmenuBuilder) AutoSelect(b bool) DMenuBuilder {
	dmb.autoSelect = b
	return dmb
}

func (dmb *dmenuBuilder) MultiSelect(b bool) DMenuBuilder {
	dmb.multiSelect = b
	return dmb
}

func (dmb *dmenuBuilder) Format(s string) DMenuBuilder {
	dmb.format = s
	return dmb
}

func (dmb *dmenuBuilder) Sep(s string) DMenuBuilder {
	dmb.sep = s
	return dmb
}

func (dmb *dmenuBuilder) Build() DMenu {
	argSlice := []string{"-dmenu"}
	argSlice = appendIf(argSlice, "-p", dmb.prompt)
	argSlice = appendIf(argSlice, "-format", dmb.format)
	argSlice = appendIf(argSlice, "-multi-select", dmb.multiSelect)
	argSlice = appendIf(argSlice, "-sep", fmt.Sprintf("%s", dmb.sep))
	argSlice = appendIf(argSlice, "-auto-select", dmb.autoSelect)
	argSlice = appendIf(argSlice, "-matching", dmb.matching)
	argSlice = appendIf(argSlice, "-theme-str", dmb.themeStr)

	fmt.Printf("action args: \n %v\n", argSlice)
	return dmenu{
		cmd:  rofi,
		args: argSlice,
	}
}

func appendIf(slice []string, argName string, argValue interface{}) []string {

	switch v := argValue.(type) {
	case bool:
		if v {
			slice = append(slice, argName)
		}
	case string:
		if v != "" {
			slice = append(slice, argName)
			slice = append(slice, fmt.Sprintf("%s", v))
		}
	case int:
		if v != 0 {
			slice = append(slice, argName)
			slice = append(slice, fmt.Sprintf("%d", v))
		}
	}

	return slice

}

type DMenu interface {
	Exec(input string) (string, error)
}

type dmenu struct {
	cmd  string
	args []string
}

func (d dmenu) Exec(input string) (string, error) {
	rofi := exec.Command(d.cmd, d.args...)

	stdin, err := rofi.StdinPipe()
	check(err)

	go func() {
		defer stdin.Close()
		_, _ = io.WriteString(stdin, input)
	}()

	var stdout, stderr bytes.Buffer
	rofi.Stdout = &stdout
	rofi.Stderr = &stderr

	err = rofi.Run()
	check(err)

	return string(stdout.Bytes()), fmt.Errorf(string(stderr.Bytes()))
}
