package main

import (
	"fmt"
	"os"
)

func main() {
	file, error := os.Open("example.txt")

	if error != nil {
		fmt.Println(error)
		panic(error)
	}

	fileinfno, fileError := file.Stat()
	if fileError != nil {
		panic(fileError)
	}

	fmt.Println("File Size", fileinfno.Size())
	fmt.Println("File Name", fileinfno.Name())

	fmt.Println("File Creatin time", fileinfno.ModTime())
	fmt.Println("File is directory or not", fileinfno.IsDir())
	fmt.Println("File type", fileinfno.Mode())

	// This is one way to read the file manually without using os.ReadFile function
	scratchFile := make([]byte, fileinfno.Size())
	readfile, e := file.Read(scratchFile)
	if e != nil {
		panic(e)
	}
	defer file.Close()

	var fileText string
	for i := 0; i < len(scratchFile); i++ {
		fmt.Println(readfile)
		fileText += string(scratchFile[i])
	}

	fmt.Println(fileText)

	// This is second way to read the file
	newFile, errorNew := os.ReadFile("example.txt")
	if errorNew != nil {
		panic(errorNew)
	}

	fmt.Println(string(newFile))

	data, error := os.Create("example2.txt")
	if error != nil {
		panic(error)
	}

	defer data.Close()
	data.WriteString("Hello World")

	dfe := os.Remove("example2.txt")
	if dfe != nil {
		panic(dfe)
	}

}
