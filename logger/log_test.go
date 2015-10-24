package logger

import (
	"os"
	"testing"
)

func TestStdStreamLog(t *testing.T) {
	h, _ := NewStreamHandler(os.Stdout)
	s := NewLogger(h)
	s.Info("hello world")

	s.Close()

	Info("hello world")
}

func TestRotatingFileLog(t *testing.T) {
	path := "./testFolder"
	os.RemoveAll(path)

	os.Mkdir(path, 0777)
	fileName := path + "/test"

	h, err := NewRotatingFileHandler(fileName, 10, 2)
	if err != nil {
		t.Fatal(err)
	}

	buf := make([]byte, 10)

	h.Write(buf)

	h.Write(buf)

	if _, err := os.Stat(fileName + ".1"); err != nil {
		t.Fatal(err)
	}

	if _, err := os.Stat(fileName + ".2"); err == nil {
		t.Fatal(err)
	}

	h.Write(buf)
	if _, err := os.Stat(fileName + ".2"); err != nil {
		t.Fatal(err)
	}

	h.Close()

	os.RemoveAll(path)
}
