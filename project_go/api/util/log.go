package util

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func Enterlog(s string) { log.Println("[INF]=====================<entering>: ", s) }
func Leavelog(s string) { log.Println("[INF]=====================<leaving>: ", s) }

func Errlog(s, c, r string, e interface{}) {
	Enterlog(s)
	log.Println("[ERR]: Code:", c, ", Msg:", r)
	log.Println("[ERR]:", e)
	Leavelog(s)
}
func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}
func CreateDir(path string) bool {
	// check
	if _, err := os.Stat(path); err != nil {
		err := os.MkdirAll(path, 0711)
		if err != nil {
			Errlog("err", "code", "reason", err)
			return false
		}
	}
	return true
}
func BaseName() string {
	return filepath.Base(os.Args[0])
}
func CreateLogDir() string {
	logPath := fmt.Sprintf("%s/log", GetCurrentDirectory())
	if CreateDir(logPath) {
		return logPath
	}
	return ""
}

type myWriter struct {
	createdDate string
	file        *os.File
}

func (t *myWriter) Write(p []byte) (n int, err error) {
	tt := string(p[5:10])
	if t.createdDate != tt {
		if err := t.RotateFile(time.Now()); err != nil {
			log.Printf("[ERRS] %s\n", err.Error())
		}
	}
	return t.file.Write(p)
}
func (t *myWriter) RotateFile(now time.Time) error {
	t.createdDate = fmt.Sprintf("%02d/%02d", now.Month(), now.Day())
	logDir := CreateLogDir()
	if len(logDir) != 0 {
		//baseName_YYYYMMDD.log
		///path/to/file/<prefix>YYYYMMDD<suffix>
		filePath := fmt.Sprintf("%s/%s%s%s", logDir, BaseName()+"_", now.Format("20060102"), ".log")
		file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			Errlog("err", "code", "reason", err)
			return err
		}
		if t.file != nil {
			t.file.Close()
		}
		t.file = file
	}
	return nil
}

// log.Ldate | log.Lmicroseconds
// ./log/baseName_YYYYMMDD.log
// 출력 & 로그
func LogSetup() {
	log.SetFlags(log.Ldate | log.Lshortfile | log.Ltime)
	log.SetOutput(io.MultiWriter(&myWriter{} /*os.Stderr,*/, os.Stdout))
}
