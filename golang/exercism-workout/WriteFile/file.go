package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
)

// Creating a file and writting it
func CreateFile(filename string) {
	file, error := os.Create(filename)
	if error != nil {
		log.Fatal(error)
		return
	}
	defer file.Close()

	_, writeErr := file.WriteString("Hey new file")
	if writeErr != nil {
		log.Fatal(writeErr)
		return
	}
}

// Creating function to write byetesize data into a file
func WriteByte(filename string, data []byte) {
	file, error := os.Create(filename)
	if error != nil {
		log.Fatal(error)
		return
	}
	defer file.Close()
	size, writeError := file.Write(data)
	if writeError != nil {
		log.Fatal(writeError)
		return
	}
	fmt.Printf("Wrote %d bytes to file", size)
}

// Creating a file and writing multiple line of strings example like logs
func WriteByLines(filename string, lines []string) {
	file, error := os.Create(filename)
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

// Create a fuction that writes to the existing file without overwrite it.
func Append(filename string, lines []string) {

	// os.O_APEEND = append data to the file when writing
	// OS.O_WRONLY =  open the file write-only
	// fs.ModeApeend = append-only
	file, error := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, fs.ModeAppend)
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

func main() {
	CreateFile("test.txt")
	WriteByte("bytes.txt", []byte("Hello, I am a robot"))
	Append("log.txt", []string{"India", "Canada", "Japan"})
}
