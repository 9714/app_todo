package utils

import (
	"io"
	"log"
	"os"
)

func LoggingSettings(logFile string) {
	//OpenFileでlogFileを開き、読み書き・作成・追加
	logfile, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	//logの書き込み先をStdoutとlogfileに指定
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	//logのフォーマットを指定/日付・時間・フォーマット
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	//logの出力先を設定
	log.SetOutput(multiLogFile)
}
