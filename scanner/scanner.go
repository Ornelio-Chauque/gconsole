package scanner

import (
	"os"
)

// Get the user input from console and retun as s string
func Get() string{
	s :=os.Stdin
	buff := make([]byte, 128)
	_, err:= s.Read(buff)
	defer s.Close()
	if err!=nil {
		panic(" something unexpected happens, can't read from stdio")
	}
	return string(buff);
}