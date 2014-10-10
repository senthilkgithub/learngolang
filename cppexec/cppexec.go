// Complete Package for C++ language compilation and executaion using gcc compiler
package cppexec

import (
	et "github.com/senthilkgithub/learngolang/exectypes"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

/*CPlusPlus_Executer gets source code and CommandLine Arguments and returns Final Response(Output/Error)*/
func CPlusPlus_Executer(req et.ExecRequest) et.ExecResponse { //, CompilerResponse chan<- et.ExecResponse
	exeresult := et.ExecResponse{}
	exeresult.Language = req.Language
	req.MainPath, exeresult.FolderPath, _ = createMainFile(req.Guid, req.SourceCode)
	resp, err := compileAndExecuteCPlusPlus(req.MainPath, req.CommandLineArgs)
	if err != nil {
		exeresult.Response = err.Error()
	} else {
		exeresult.Response = resp
	}
	return exeresult
	//CompilerResponse <- exeresult
}

// The process of c++ compilation and execution is taken care by this function CompileAndExecuteC(filepath)(output,error)
func compileAndExecuteCPlusPlus(mainFilePath string, commandLineArgs string) (output string, err error) {
	outputpath := strings.Replace(mainFilePath, ".cpp", ".o", 1)
	exepath := strings.Replace(mainFilePath, ".cpp", ".exe", 1)

	//Generation of output file
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
	//Generation of Exe file using output file(.o)
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
	//Stage of execution
Ex:
	cmd = exec.Command("cmd", "/C", exepath+" "+commandLineArgs)
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

//function CreateMainFile creates main.cpp file
func createMainFile(guid string, sourceCode string) (mainfilepath string, folderPath string, err error) {
	folderPath = "D:\\sourcecodes\\cpp\\" + guid
	mainfilepath = folderPath + "\\main.cpp"
	err = os.Mkdir(folderPath, 0644)
	err = ioutil.WriteFile(mainfilepath, []byte(sourceCode), 0644)
	if err != nil {
		return "", "", err
	}
	return mainfilepath, folderPath, nil
}
