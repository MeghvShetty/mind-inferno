package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// Read stats (file info name , create date, IsDir)
func ReadStats(filename string) {
	data, error := os.Open(filename)
	if error != nil {
		log.Fatal(error)
	}
	defer data.Close()
	stats, error := data.Stat()
	if error != nil {
		log.Fatal(error)
	}
	fmt.Printf("File Name: %s\n", stats.Name())
	fmt.Printf("File Modified: %s\n", stats.ModTime().Format("Monday, Januray 01, 15:04:05"))

}

// Basic way to read the whole file for the contect - This will not work if you have very large file like log files.
func ReadWholeFile(filename string) {
	data, error := ioutil.ReadFile(filename)
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println(string(data))
}

// Line By Line
func ReadByLine(filename string) {
	data, error := os.Open(filename)
	if error != nil {
		log.Fatal(error)
	}
	defer data.Close()

	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

// Word by word
func ReadByWord(filename string) {
	data, error := os.Open(filename)
	if error != nil {
		log.Fatal(error)
	}
	defer data.Close()

	scanner := bufio.NewScanner(data)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

// Read the files in Bytes
func ReadByBytes(filename string, size uint8) {
	data, error := os.Open(filename)
	if error != nil {
		log.Fatal(error)
	}
	defer data.Close()

	buf := make([]byte, size)
	for {
		totalRead, err := data.Read(buf)
		if err != nil {
			if err != io.EOF { //End of File
				log.Fatal(err)
			}
			break
		}
		fmt.Println(string(buf[:totalRead]))
	}

}

// Reading from a config file
func ReadConfig(filename string) {
	data, error := os.Open(filename)
	if error != nil {
		log.Fatal(error)
	}
	defer data.Close()

	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		raw := strings.Split(scanner.Text(), "=") // [Key, value]
		fmt.Println(raw[0])
		fmt.Println(raw[1])
	}
}

func main() {
	ReadByBytes("/home/Exercism/Go_lang/readFile/file.dat", 8)
	ReadConfig("/home/Exercism/Go_lang/readFile/test.conf")

}
