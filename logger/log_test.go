package logger

import (
	"os"
	"testing"
	"time"
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
		t.Fatal("interval is incorrect")
	}

	Info("rollover at: %d", h.rolloverAt)
	if time.Now().Unix()+(24*3600) <= h.rolloverAt {
		t.Fatal("rollover time is incorrect")
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

func TestMiddlewareSkip(t *testing.T) {
	Info("TestMiddlewareSkip: test middleware logic to insure skip of /ping from health check")

	h, _ := NewStreamHandler(os.Stdout)
	l := NewLogger(h)
	m := NewMiddlewareLogger(l)

	path := "/ping"
	agent := "ELB-HealthChecker/1.0"

	skip := m.Skip(path, agent)

	Info("path: %s, agent: %s, skip: %v", path, agent, skip)

	if skip != true {
		t.Fatal("should skip ping healthcheck")
	}

	path = "/index.html"
	skip = m.Skip(path, agent)
	Info("path: %s, agent: %s, skip: %v", path, agent, skip)

	if skip == true {
		t.Fatal("should not skip index healthcheck")
	}

	path = "/ping"
	agent = "curl"
	skip = m.Skip(path, agent)
	Info("path: %s, agent: %s, skip: %v", path, agent, skip)

	if skip == true {
		t.Fatal("should not skip ping from curl")
	}

}
