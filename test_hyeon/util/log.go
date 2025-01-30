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

// 함수나 코드 블록의 진입을 로그로 기록
func Enterlog(s string) { log.Println("[INF]=====================<entering>: ", s) }

// 함수나 코드 블록의 종료를 로그로 기록
func Leavelog(s string) { log.Println("[INF]=====================<leaving>: ", s) }

/*
< 오류를 로그로 기록 >
s string: 함수나 코드 블록의 이름
c string: 오류 코드
r string: 오류 메시지
e interface{}: 오류 객체
*/
func Errlog(s, c, r string, e interface{}) {
	Enterlog(s)
	log.Println("[ERR]: Code:", c, ", Msg:", r)
	log.Println("[ERR]:", e)
	Leavelog(s)
}

// 현재 실행 중인 프로그램의 디렉토리 반환
func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

/*
< 지정된 경로에 디렉토리 생성 >
path string: 생성할 디렉토리의 경로
*/
func CreateDir(path string) bool {
	// 지정된 경로가 존재하지 않으면 디렉토리 생성
	if _, err := os.Stat(path); err != nil {
		err := os.MkdirAll(path, 0711)
		if err != nil {
			Errlog("err", "code", "reason", err)
			return false
		}
	}
	return true
}

// 현재 실행 중인 프로그램의 파일 이름 반환
func BaseName() string {
	return filepath.Base(os.Args[0])
}

// 로그를 저장할 디렉토리를 생성하고 그 경로 반환
func CreateLogDir() string {
	logPath := fmt.Sprintf("%s/log", GetCurrentDirectory())
	if CreateDir(logPath) {
		return logPath
	}
	return ""
}

// 로그 파일을 관리하는 구조체
type myWriter struct {
	createdDate string
	file        *os.File
}

/*
< 로그 데이터를 파일에 쓴다 >
- t *myWriter: 메서드 리시버로, myWriter 구조체의 인스턴스를 가리킨다
- p []byte: 로그 데이터가 담긴 바이트 슬라이스
- 쓰여진 바이트 수 n과 오류 err 반환
*/
func (t *myWriter) Write(p []byte) (n int, err error) {
	// 로그 데이터의 날짜 추출
	tt := string(p[5:10])
	// 날짜가 변경되었으면 로그 파일 회전
	if t.createdDate != tt {
		if err := t.RotateFile(time.Now()); err != nil {
			log.Printf("[ERRS] %s\n", err.Error())
		}
	}
	return t.file.Write(p)
}

/*
< 새로운 로그 파일을 생성하거나 기존 로그 파일 회전 >
- t *myWriter: 메서드 리시버로, myWriter 구조체의 인스턴스를 가리킨다
- now time.Time: 현재 시간을 나타내는 time.Time 타입의 인자
- 오류가 발생하면 error 반환
*/
func (t *myWriter) RotateFile(now time.Time) error {
	// 현재 날짜를 createdDate 필드에 저장
	t.createdDate = fmt.Sprintf("%02d/%02d", now.Month(), now.Day())

	// 로그 디렉토리를 생성하고 그 경로를 저장
	logDir := CreateLogDir()

	if len(logDir) != 0 {
		// baseName_YYYYMMDD.log
		// /path/to/file/<prefix>YYYYMMDD<suffix>
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

/*
< 로그 설정 초기화 >
- 출력 & 로그
- log.Ldate | log.Lmicroseconds
- ./log/baseName_YYYYMMDD.log
- 로그를 커스텀 로그 파일과 표준 출력에 동시에 출력하도록 설정
*/
func LogSetup() {
	// 로그 형식 설정
	log.SetFlags(log.Ldate | log.Lshortfile | log.Ltime)

	// 로그 출력 설정
	log.SetOutput(io.MultiWriter(&myWriter{} /*os.Stderr,*/, os.Stdout))
}
