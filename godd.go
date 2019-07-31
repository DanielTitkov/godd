package main

import (
	"flag"
	"fmt"
	"io"
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

func copyFile(inpPath, outpPath string, bufSize int) error {
	// Open input file
	inp, err := os.Open(inpPath)
	if err != nil {
		return err
	}
	defer inp.Close()

	inpLength := getFileLength(inp)
	var doneLength int64

	// Check if output file exists
	_, err = os.Stat(outpPath)
	if err == nil {
		log.Fatalf("Output file already exists: %s", outpPath)
	}

	// Create output file
	outp, err := os.Create(outpPath)
	if err != nil {
		return err
	}
	defer outp.Close()

	// Move bytes between files using buffer
	buf := make([]byte, bufSize)
	for {
		n, err := inp.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}

		if n == 0 {
			break
		}

		if _, err := outp.Write(buf[:n]); err != nil {
			return err
		}

		if doneLength = doneLength + int64(n); (doneLength % 100) == 0 {
			donePercent := int(float64(doneLength) / float64(inpLength) * 100)
			fmt.Printf("\rCopying.. %v%%", donePercent)
		}
	}
	fmt.Printf("\rAll done %v%%\n", int(float64(doneLength)/float64(inpLength)*100))
	return err
}

var inpPath string
var outpPath string
var bufSize int

func init() {
	flag.StringVar(&inpPath, "i", "", "Input file path (required)")
	flag.StringVar(&outpPath, "o", "", "Output file path (required)")
	flag.IntVar(&bufSize, "b", 1024, "Buffer size")
}

func main() {
	flag.Parse()
	if len(inpPath) < 1 || len(outpPath) < 1 {
		flag.Usage()
		os.Exit(1)
	}

	log.Printf("Started copying from %s to %s", inpPath, outpPath)
	err := copyFile(inpPath, outpPath, bufSize)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Finished")
}
