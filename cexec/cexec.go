package cexec

import (
	et "github.com/senthilkgithub/learngolang/exectypes"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func C_Executer(req et.ExecRequest, CompilerResponse chan<- et.ExecResponse) {
	exeresult := et.ExecResponse{}
	req.MainPath, exeresult.FolderPath, _ = createMainFile(req.Guid, req.SourceCode)
	resp, err := compileAndExecuteC(req.MainPath, req.CommandLineArgs)
	if err != nil {
		exeresult.Response = err.Error()
	} else {
		exeresult.Response = resp
	}
	exeresult.RespWriter = req.RespWriter
	CompilerResponse <- exeresult
}

// The process of compilation and execution is taken care by function CompileAndExecuteC(filepath)(output,error)
func compileAndExecuteC(mainFilePath string, commandLineArgs string) (output string, err error) {
	// if commandLineArgs !=""  {
	// 	commandLineArgs = ""
	// }

	outputpath := strings.Replace(mainFilePath, ".c", ".o", -1)
	exepath := strings.Replace(mainFilePath, ".c", ".exe", -1)

	cmd := exec.Command("cmd", "/C", "C:\\MinGW\\bin\\gcc -o "+outputpath+" -c "+mainFilePath)
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
	cmd = exec.Command("cmd", "/C", "C:\\MinGW\\bin\\gcc -o "+exepath+" "+outputpath)
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

//function CreateMainFile creates main.c file
func createMainFile(guid string, sourceCode string) (mainfilepath string, folderPath string, err error) {
	folderPath = "D:\\sourcecodes\\c\\" + guid
	mainfilepath = folderPath + "\\main.c"
	err = os.Mkdir(folderPath, 0644)
	err = ioutil.WriteFile(mainfilepath, []byte(sourceCode), 0644)
	if err != nil {
		return "", "", err
	}
	return mainfilepath, folderPath, nil
}
