package main

import (
	"fmt"
	"os"
)

func main() {
	var fw *FileWriter
	_, err := fw.Write([]byte("Hello World"))
	if err != nil {
		fmt.Println(err.Error())
		switch e := err.(type) {
		case *FileError:
			println("Code:", e.Code)
			println("Message:", e.Message)
		}
	} else {
		println("file successfully created and written")
	}
}

type FileWriter struct {
	FileName string
}

func New(filename string) *FileWriter {
	return &FileWriter{FileName: filename}
}

func (fw *FileWriter) Write(p []byte) (n int, err error) {
	if fw == nil {
		return 0, NewFileError(100, "nil file object")
	}
	f, err := os.OpenFile(fw.FileName, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return 0, err
	}
	// n, err = f.Write(p)
	// return n, err
	defer f.Close()
	return f.Write(p)
}

type FileError struct {
	Code    int
	Message string
}

func NewFileError(code int, message string) *FileError {
	return &FileError{code, message}
}

func (fe *FileError) Error() string {
	return fmt.Sprint("Code:", fe.Code, " Message:", fe.Message)
}
