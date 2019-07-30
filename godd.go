package main

import (
	"fmt"
	"log"
	"os"
)

func getFileLength(file *os.File) int64 {
	stat, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	return stat.Size()
}

func main() {
	fmt.Println("foo")
	file, err := os.Open("foo.txt")
	if err != nil {
		log.Fatal(err)
	}
	fl := getFileLength(file)
	fmt.Println(fl)
}
