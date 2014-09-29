package deletefolder

import "os"

func RemoveAllFiles(dirName string) error {

	err := os.RemoveAll(dirName)
	if err != nil {
		return err
	}
	return nil

}
