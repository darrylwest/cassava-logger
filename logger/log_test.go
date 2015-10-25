package logger

import (
	"os"
	"testing"
)

func TestStdStreamLog(t *testing.T) {
	Info("TestStdStreamLog: basic stream tests for type and instance methods")

	h, _ := NewStreamHandler(os.Stdout)
	s := NewLogger(h)
	s.Info("instance message")
	s.Warn("instance warn statement")
	s.Error("instance error statement")
	s.Fatal("instance fatal statement")

	s.Close()

	SetLevel(TraceLevel)
	Trace("type trace statement")
	Debug("type debug statement")
	Info("type info statement")
	Warn("type warn statement")
	Error("type error statement")
	Fatal("type fatal statement")

}

func TestRotatingDayHandler(t *testing.T) {
	Info("TestRotatingDayHanlder: test the file name and rolling based on time")
	path := "./daytest"
	os.RemoveAll(path)

	os.Mkdir(path, 0777)
	filename := path + "/test"

	h, err := NewRotatingDayHandler(filename)
	if err != nil {
		t.Fatal(err)
	}

	Info("log basename: %s", h.baseName)
	if h.baseName != "./daytest/test" {
		t.Fatalf("file basename is incorrect")
	}

	Info("interval: %d", h.interval)
	if h.interval != 86400 {
		t.Fatal("intervali is incorrect")
	}

	realFilename := h.CreateFilename()

	Info("real filename: %s", realFilename)
	finfo, _ := os.Stat(realFilename)
	Info("file: %v", finfo.Name())

	h.Write([]byte("my first test\n"))
	h.Write([]byte("my second test\n"))
	h.Close()

	os.RemoveAll(path)
}

func TestRotatingFileLog(t *testing.T) {
	Info("TestRotatingFileLog: test the file names and rolling logic based on size")
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
