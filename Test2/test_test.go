package Test2

import "testing"
import "bytes"
import "os"

func TestHello(t *testing.T) {
	// Capture stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	Hello()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	expected := "你好!\n"
	if output != expected {
		t.Errorf("Hello() output = %q, want %q", output, expected)
	}
}

func TestKain(t *testing.T) {
	// Capture stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	Kain()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	expected := "好男人就是我\n"
	if output != expected {
		t.Errorf("Kain() output = %q, want %q", output, expected)
	}
}
