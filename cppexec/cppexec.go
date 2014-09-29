package cppexec

import (
	et "github.com/senthilkgithub/learngolang/exectypes"
	"os/exec"
	"strings"
)

func CPlusPlus_Executer(req et.ExecRequest, CompilerResponse chan<- et.ExecResponse) {
	exeresult := et.ExecResponse{}
	exeresult.Language = req.Language
	resp, err := compileAndExecuteCPlusPlus(req.MainPath, req.CommandLineArgs)
	if err != nil {
		exeresult.Response = err.Error()
	} else {
		exeresult.Response = resp
	}
	CompilerResponse <- exeresult
}

// The process of compilation and execution is taken care by function CompileAndExecuteC(filepath)(output,error)
func compileAndExecuteCPlusPlus(mainFilePath string, commandLineArgs string) (output string, err error) {
	// if commandLineArgs !=""  {
	// 	commandLineArgs = ""
	// }
	outputpath := strings.Replace(mainFilePath, ".cpp", ".o", 1)
	exepath := strings.Replace(mainFilePath, ".cpp", ".exe", 1)
	cmd := exec.Command("cmd", "/C", "C:\\MinGW\\bin\\g++ -o "+outputpath+" -c "+mainFilePath)
	out, err := cmd.CombinedOutput()
	if err != nil || out != nil {
		if out != nil && len(out) > 0 {
			errmsg := string(out)
			errmsg = strings.Replace(errmsg, mainFilePath, "", -1)
			return errmsg, nil
		} else if err != nil {
			return "", err
		} else {
			goto Co
		}
	}
Co:
	cmd = exec.Command("cmd", "/C", "C:\\MinGW\\bin\\g++ -o "+exepath+" "+outputpath)
	out, err = cmd.CombinedOutput()
	if err != nil || out != nil {
		if out != nil && len(out) > 0 {
			return string(out), nil
		} else if err != nil {
			return "", err
		} else {
			goto Ex
		}
	}

Ex:
	cmd = exec.Command("cmd", "/C", exepath+commandLineArgs)
	out, err = cmd.CombinedOutput()
	if err != nil || out != nil {
		if out != nil && len(out) > 0 {
			return string(out), nil
		} else if err != nil {
			return "", err
		}
	}
	return "", nil

}
