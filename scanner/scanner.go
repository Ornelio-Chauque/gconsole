package scanner

import (
	"fmt"
	"os"
	"strings"
)
const (
	BUFFER_SIZE = 256// set a fixed buffer size, to allow better use of memory
)
// Get the user input from console and retun as s string
func Get() (string, error) {
	s := os.Stdin
	buff := make([]byte, BUFFER_SIZE)
	defer s.Close()
    var err error;
	var inputResultBuilder strings.Builder

	for n:=0; err == nil;{
		
		n,err = s.Read(buff)
		if err != nil {
			return "", fmt.Errorf("on Get: %w", err)
		}
		inputResultBuilder.Write(buff[:n])
		
		//When the number of byte written to buff (c) is less than the overall buff capacity (BUFFER_SIZE),
		// means no more bytes reamins to be read, so break the loop
		if n < BUFFER_SIZE{
			break
		}


	}
	
	return inputResultBuilder.String(), nil
}

//GetBytes return the  slice of byte representation of content read from Stdin
func GetBytes() ([]byte, error) {
	s := os.Stdin
	buff := make([]byte, BUFFER_SIZE)
	defer s.Close()
    var err error;
	inputResult := make([]byte, BUFFER_SIZE)

	for n:=0; err == nil;{
		
		n,err = s.Read(buff)
		if err != nil {
			return nil, fmt.Errorf("on GetBytes: %w", err)
		}
	
		inputResult = append(inputResult, buff...)
		//When the number of byte written to buff (c) is less than the overall buff capacity (BUFFER_SIZE),
		// means no more bytes reamins to be read, so break the loop
		if n < BUFFER_SIZE{
			break
		}


	}
	
	return inputResult, nil
}

