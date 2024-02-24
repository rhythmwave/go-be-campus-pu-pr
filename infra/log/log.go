package log

import (
	"fmt"
	"os"
	"time"

	"github.com/sccicitb/pupr-backend/constants"
	log "github.com/sirupsen/logrus"
)

const (
	gitignore = ".gitignore"
)

// PrintTimestamp function to print timestamp for every log
func PrintTimestamp() {
	formatter := &log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: constants.DateFromStd,
	}

	log.SetFormatter(formatter)
}

// PrintOutputToFile function to write log output to log file
// Returns *os.File
func PrintOutputToFile() *os.File {
	now := time.Now().Format(constants.DateFormatStd)
	filename := fmt.Sprintf("%s%s", constants.LogDir, now)

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	log.SetOutput(f)
	log.SetLevel(log.WarnLevel)

	return f
}

// // DeleteOldLogFile function to delete old log file
// // Default it is to delete log file older than 14 days
// func DeleteOldLogFile() {
// 	fileInfo, err := ioutil.ReadDir(constants.LogDir)
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}
// 	now := time.Now()
// 	cutoff := 14 * 24 * time.Hour

// 	daily := 24 * time.Hour
// 	jt := utils.NewJobTicker(daily)
// 	for {
// 		<-jt.T.C
// 		for _, info := range fileInfo {
// 			if diff := now.Sub(info.ModTime()); diff > cutoff && info.Name() != gitignore {
// 				filename := fmt.Sprintf("%s%s", constants.LogDir, info.Name())

// 				err = os.Remove(filename)
// 				if err != nil {
// 					log.Error(err)
// 				}
// 			}
// 		}
// 		jt.UpdateJobTicker(daily)
// 	}
// }
