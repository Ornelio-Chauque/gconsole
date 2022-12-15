package printer

import (
	"fmt"
	"os"
)

// Will print the content of string to the stdout
func Print(s string) (string, error) {
	_, err := writeStdout([]byte(s))
	if err != nil {
		return "", fmt.Errorf("on printer.Print: %w", err)
	}
	return s, nil
}

// will print the content of the string to stdou and then LF
func Println(s string) (string, error) {
	newString := s + "\n"
	_, err := writeStdout([]byte(newString))
	if err != nil {
		return "", fmt.Errorf("on printer.Println: %w", err)
	}

	return s, nil
}

func Printf(format string, args ...string) {

}

// write the content to the specified file name
func WriteTo(filename string, data []byte, mode os.FileMode) error {

	f, err := processFile(filename, mode)
	if err != nil{
		return err
	}
	defer f.Close()
	
	_, err = write(f, data)
	if err != nil {
		return err
	}
	return nil
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
func processFile(filename string, mode os.FileMode) (*os.File, error) {

	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, mode)
	if err != nil {
		return nil, fmt.Errorf("in printer.processprocessFile: %w", err)
	}
	return f, nil
}
