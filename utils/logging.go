package utils

import (
	"io"
	"log"
	"os"
)

func LoggingSettings(logFile string) {
	// lgofileにread/create/appendの追記を許可 -> permissionは666
	logfile, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	//
	if err != nil {
		log.Fatalf("file=logFile err=%s", err.Error())
	}

	// MultiWriterでStdoutとlogfileに書き込み
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile) // 日付・時間・種類を記述
	log.SetOutput(multiLogFile)
}
