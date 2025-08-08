package mdir

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var ProjectName string = "testing"

func CreateDirStr(FileName string) {
	root, erro := filepath.Abs("./")
	if erro != nil {
		log.Fatal(erro)
	}
	dirCreate := fmt.Sprintf("%s/%s/Reports", root, FileName)
	fmt.Println(dirCreate)
	if err := os.MkdirAll(dirCreate, os.ModePerm); err != nil {
		log.Fatal(err)
	}
}
