package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {

	//fmt.Println(cmd)
	//c = cmd.Str
	cmd := exec.Command("cmd", "/C", "C:\\MinGW\\bin\\gcc -o  D:\\Testprograms\\C\\main.o -c D:\\Testprograms\\C\\main.c")
	out, err := cmd.CombinedOutput()
	if err != nil || out != nil {
		if out != nil && len(out) > 0 {
			errmsg := string(out)
			errmsg = strings.Replace(errmsg, "D:\\Testprograms\\C\\main.c:", "", -1)
			fmt.Println(errmsg)
		} else if err != nil {
			fmt.Printf("%s", err.Error())
		} else {
			goto Co
		}
		return
	}
Co:
	cmd = exec.Command("cmd", "/C", "C:\\MinGW\\bin\\gcc -o D:\\Testprograms\\C\\main.exe   D:\\Testprograms\\C\\main.o")
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
	cmd = exec.Command("cmd", "/C", "D:\\Testprograms\\C\\main.exe 6,5")
	out, err = cmd.CombinedOutput()
	if err != nil || out != nil {
		if out != nil && len(out) > 0 {
			fmt.Println(string(out))
		} else if err != nil {
			fmt.Printf(err.Error())
		}

	}

	// ccmd := exec.Command("D:\\Testprograms\\C\\compilec.bat")
	// out, err := ccmd.CombinedOutput()
	// if err != nil || out != nil {
	// 	fmt.Println(err)
	// 	fmt.Printf("%s", err.Error(), string(out))
	// } else {
	// 	time.Sleep(1000)
	// 	cmd := exec.Command("D:\\Testprograms\\C\\executec.bat")

	// 	output, err := cmd.CombinedOutput()
	// 	//	exec.Command("c:\\del.bat").Run()
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		fmt.Printf("%s", err.Error())
	// 	}
	// 	fmt.Printf("%s", output)
	// }
}
