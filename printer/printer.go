package printer

import (
	"os"
)

// Will print the content of string to the stdout
func Print(s string) string {
	_, err := writeStdout([]byte(s))
	if err != nil {
		panic(" Problem writing to stdout")
	}
	return s
}

// will print the content of the string to stdou and then LF
func Println(s string) string {
	newString := s + "\n"
	_, err := writeStdout([]byte(newString))
	if err != nil {
		panic(" Problem writing to stdout")
	}

	return s
}

func Printf(format string, args ...string) {

}

// write the content to the specified file name
func WriteTo(filename string, data []byte, mode os.FileMode) {

	f := processFile(filename, mode)
	defer f.Close()
	_, err := write(f, data)
	if err != nil {
		panic(" we find issue writing to this file")
	}
}

// Makes gconsole.printer a io.Writer
//when call writes to stdout
func Write(data []byte) (int, error) {
	n, err := writeStdout(data)
	return n, err
}

// This function write to stdout specifically
func writeStdout(data []byte) (int, error) {
	return write(os.Stdout, data)
}

// Helper fuction to write bytes to any file, including stdout
func write(f *os.File, data []byte) (int, error) {
	file := *f
	return file.Write(data)
}

// Helper function which make a syscal to open a file with read and write mode if exist
//if don't exist creates the file
func processFile(filename string, mode os.FileMode) *os.File {

	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, mode)
	if err != nil {
		panic(" expecting filename not filepath") //TODO: impprove error mesage
	}
	return f
}
