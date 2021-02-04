package util

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"reflect"
)

func RunCmdWithReader(finput func(closer io.WriteCloser), cmd string, args []string) (string, error) {
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

func RunCmdWithInput(input string, cmd string, args []string) (string, error) {
	finput := func(in io.WriteCloser) {
		fmt.Fprintln(in, input)
	}
	return RunCmdWithReader(finput, cmd, args)

}

func AppendIf(res []string, argName string, pValue interface{}) []string {
	if pValue == nil {
		return res
	}
	vs := reflect.ValueOf(pValue)

	switch vs.Kind() {
	case reflect.String:
		if vs.String() != ""{
			res = appendArgName(res, argName)
			res = append(res, fmt.Sprintf("%s", vs.String()))
			return res
		}

	case reflect.Bool:
		if vs.Bool() {
			res = appendArgName(res, argName)
			return res
		}
	case reflect.Slice:
		for i := 0; i < vs.Len(); i++ {
			vIndex := vs.Index(i)

			if i == 0 {
				res = AppendIf(res, argName, vIndex.Interface())
			} else {
				res = AppendIf(res, "", vIndex.Interface())
			}
		}
	case reflect.Int:
		log.Fatal(fmt.Errorf("cannot be int, use string instead"))
	}

	return res
}

func appendArgName(slice []string, argName string) []string {
	if argName == "" {
		return slice
	}
	slice = append(slice, argName)
	return slice
}
