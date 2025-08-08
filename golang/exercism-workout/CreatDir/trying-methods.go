package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	"example.com/Go_lang/CreatDir/mdir"
)

var projectName string = "testing"

func main() {
	mdir.CreateDirStr(projectName)
	var files []string
	ext := []string{".conf", ".dat"}
	root, erro := filepath.Abs("../")
	if erro != nil {
		log.Fatal(erro)
	}

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

		if err != nil {

			fmt.Println(err)
			return nil
		}
		for _, value := range ext {
			if !info.IsDir() && filepath.Ext(path) == value {
				files = append(files, path)

			}
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
	Append(files, "General")
	for _, file := range files {
		fmt.Println(file)

	}
	//////////////////////////////// Clear zone

}

func Append(lines []string, ReportPattern string) {
	TempDir := fmt.Sprintf("/home/Exercism/Go_lang/CreatDir/%s/Reports", projectName)
	pattern := fmt.Sprintf("%s-%s", projectName, ReportPattern)
	Newfile, error := os.CreateTemp(TempDir, pattern)
	if error != nil {
		log.Fatal(error)
		return
	}
	defer Newfile.Close()
	// os.O_APEEND = append data to the file when writing
	// OS.O_WRONLY =  open the file write-only
	// fs.ModeApeend = append-only
	file, error := os.OpenFile(Newfile.Name(), os.O_APPEND|os.O_WRONLY, fs.ModeAppend)
	if error != nil {
		log.Fatal(error)
		return
	}
	defer file.Close()
	// loop to ignore the index of the slice
	for _, val := range lines {
		_, writeError := fmt.Fprintln(file, val)
		if writeError != nil {
			log.Fatal(writeError)
			return
		}
	}
}

// func CreateDirStr(FileName string) {
// 	root, erro := filepath.Abs("./")
// 	if erro != nil {
// 		log.Fatal(erro)
// 	}
// 	dirCreate := fmt.Sprintf("%s/%s/Reports", root, FileName)
// 	fmt.Println(dirCreate)
// 	if err := os.MkdirAll(dirCreate, os.ModePerm); err != nil {
// 		log.Fatal(err)
// 	}
// }
