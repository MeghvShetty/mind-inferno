package filescanner

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// The package get the information about the file.
// return is written in a file for the scan engine to idenity the file type and path to file.

// list all the files in the dir
func PathInfo() {
	rootDir, erro := filepath.Abs("../Go_lang")
	if erro != nil {
		log.Fatal(erro)
	}

	AbsFilePath := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		if info.IsDir() == false {
			fmt.Println(path)
		}
		return nil

	})
	if AbsFilePath != nil {
		fmt.Println(AbsFilePath)
	}

}
