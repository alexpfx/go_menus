package util

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

func RunCmdWithInput(input string, cmd string, args []string) (string, error) {
	finput := func(in io.WriteCloser) {
		fmt.Fprintln(in, input)
	}

	menu := exec.Command(cmd, args...)
	menu.Stderr = os.Stderr
	in, _ := menu.StdinPipe()
	go func() {
		finput(in)
		in.Close()
	}()
	result, _ := menu.Output()
	return string(result), nil
}

func AppendIf(slice []string, argName string, argValue interface{}) []string {

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
