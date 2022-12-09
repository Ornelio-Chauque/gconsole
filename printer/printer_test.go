package printer

import (
	"os"
	"testing"
)

func TestPrintStringToStdout(t *testing.T) {
	s := "Hello Word"
	result := Print(s)

	if result != s {
		t.Errorf("expect %s, and found %s", s, result)
	}

}

/* func TestEndWithEOFWhenPrintln(t *testing.T) {
	s := "\\n"
	result := Println("Hello word")
	fmt.Println([]byte(result))

	if !strings.HasSuffix(result, s) {
		t.Errorf("Expect %s, but not found %s", s, result[len(result):])
	}

} */

func TestWhenWhriteToFileThenWriteContent(t *testing.T) {
	input := struct {
		filename string
		content  []byte
		mode     os.FileMode
	}{
		filename: "text.txt",
		content:  []byte("hello world"),
		mode:     777,
	}

	WriteTo(input.filename, input.content, input.mode)

	f, err := os.Open(input.filename)
	defer f.Close()
	if err != nil {
		t.Error("File was not created")
	}

	buff := make([]byte, len((input.content)))

	if _, err2 := f.Read(buff); err2 != nil {
		t.Error(" Not able to read the file")
	}

	if exp, found := string(input.content), string(buff); exp != found{
		t.Errorf("Expect %s, found %s ", exp, found)
	}
}

func TestWhenCallWriteThenReturnNumByteWritten(t *testing.T){
	data:= struct{
		content []byte
		expectedBytes int
	}{[]byte("hello"), 5}

	foundBytes, err:= Write(data.content);
	
	if err != nil{
		t.Error(" Problem when writing")
	}

	if data.expectedBytes != foundBytes {
		t.Errorf("Expect %d, found %d", data.expectedBytes, foundBytes)
	}

}

func Test_writeToStdout(t *testing.T){
	_, err := writeStdout([]byte("hello"))
	if err != nil{
		t.Errorf("expect %v, found %s", nil, err.Error())
	}
}

func Test_write(t *testing.T){

	data:= struct{
		content string
		size int
	}{"Hello", 5}

	totalWrittenBytes, err := write(os.Stdout, []byte(data.content))
	
	if err != nil{
		t.Error(" error while writteing")
	}

	if data.size != totalWrittenBytes{
		t.Errorf("Expected %d written bytes, found %d", data.size, totalWrittenBytes)
	}


}

