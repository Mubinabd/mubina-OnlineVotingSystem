package logger

import (
	"io"
	"log"
	"os"
)

type Loggerr struct {
	INFO  *log.Logger
	WARN  *log.Logger
	ERROR *log.Logger
	DEBUG *log.Logger
	TRACE *log.Logger
}

func Logger(basepath, path string) *Loggerr {
	l := &Loggerr{}

	fullpath := basepath + "/" + path
	var file, err = os.OpenFile(fullpath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	multiWriter := io.MultiWriter(file, os.Stdout)

	l.INFO = log.New(multiWriter, "[INFO]  ", log.Lshortfile|log.LstdFlags)
	l.WARN = log.New(multiWriter, "[WARN]  ", log.Lshortfile|log.LstdFlags)
	l.ERROR = log.New(multiWriter, "[ERROR]  ", log.Lshortfile|log.LstdFlags)
	l.DEBUG = log.New(multiWriter, "[DEBUG]  ", log.Lshortfile|log.LstdFlags)
	l.TRACE = log.New(multiWriter, "[TRACE]  ", log.Lshortfile|log.LstdFlags)

	return l
}
