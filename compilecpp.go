package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {

	//fmt.Println(cmd)
	//c = cmd.Str
	cmd := exec.Command("cmd", "/C", "C:\\MinGW\\bin\\g++ -o  D:\\Testprograms\\C++\\main.o -c D:\\Testprograms\\C++\\main.cpp")
	out, err := cmd.CombinedOutput()
	if err != nil || out != nil {
		if out != nil && len(out) > 0 {
			errmsg := string(out)
			errmsg = strings.Replace(errmsg, "D:\\Testprograms\\C++\\main.cpp:", "", -1)
			fmt.Println(errmsg)
		} else if err != nil {
			fmt.Printf("%s", err.Error())
		} else {
			goto Co
		}
		return
	}
Co:
	cmd = exec.Command("cmd", "/C", "C:\\MinGW\\bin\\g++ -o D:\\Testprograms\\C++\\main.exe   D:\\Testprograms\\C++\\main.o")
	out, err = cmd.CombinedOutput()
	if err != nil || out != nil {
		if out != nil && len(out) > 0 {
			fmt.Println(string(out))
		} else if err != nil {
			fmt.Printf("%s", err.Error())
		} else {
			goto Ex
		}
	}

Ex:
	cmd = exec.Command("cmd", "/C", "D:\\Testprograms\\C++\\main.exe 6,5")
	out, err = cmd.CombinedOutput()
	if err != nil || out != nil {
		if out != nil && len(out) > 0 {
			fmt.Println(string(out))
		} else if err != nil {
			fmt.Printf(err.Error())
		}

	}

}
