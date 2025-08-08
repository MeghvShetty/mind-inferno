package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func main() {

	var files []string
	ext := []string{".conf", ".dat"}
	root, erro := filepath.Abs("../Go_lang")
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
	Append(files)
	for _, file := range files {
		fmt.Println(file)

	}
}

func Append(lines []string) {

	Newfile, error := os.Create("testdata.txt")
	if error != nil {
		log.Fatal(error)
		return
	}
	defer Newfile.Close()
	// os.O_APEEND = append data to the file when writing
	// OS.O_WRONLY =  open the file write-only
	// fs.ModeApeend = append-only
	file, error := os.OpenFile("testdata.txt", os.O_APPEND|os.O_WRONLY, fs.ModeAppend)
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
