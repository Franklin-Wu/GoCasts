package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	argc := len(os.Args)
	if argc != 2 {
		fmt.Println("usage: <go execution command> <filename>")
	} else {
		filename := os.Args[1]
		fmt.Println("processing", filename, "...")
		file, err := os.Open(filename)
		if err != nil {
			fmt.Println(err)
		} else {
			written, err := io.Copy(os.Stdout, file)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("... wrote", written, "bytes")
			}
		}
	}
}
