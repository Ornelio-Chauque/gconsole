package scanner

import (
	"os"
	"strings"
)
const (
	BUFFER_SIZE = 256// set a fixed buffer size, to allow better use of memory
)
// Get the user input from console and retun as s string
func Get() string {
	s := os.Stdin
	buff := make([]byte, BUFFER_SIZE)
	defer s.Close()
    var err error;
	var inputResult strings.Builder

	for n:=0; err == nil;{
		
		n,err = s.Read(buff)
		if err != nil {
			panic(" something unexpected happens, can't read from stdio")
			break
		}
		inputResult.Write(buff[:n])
		if n < BUFFER_SIZE{
			break
		}


	}
	
	return inputResult.String()
}
