package printer

import (
	"fmt"
	"strings"
	"testing"
)

func TestPrintStringToStdout(t *testing.T){
	s:= "Hello Word"
	result:= Print(s);

	if result!=s {
		t.Errorf("expect %s, and found %s", s, result)
	}

}

func TestEndWithEOFWhenPrintln(t *testing.T){
	s:= "\\n"
	result:= Println("Hello word");
	fmt.Println([]byte(result))
	
	if !strings.HasSuffix(result, s){
		t.Errorf("Expect %s, but not found %s", s, result[len(result):])
	}

}
