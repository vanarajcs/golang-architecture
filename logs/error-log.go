package logs

import (
	"golang-architecture/conf"
	"log"
	"os"
)

// Error log Object
type ErrorLog struct {
	ErrorMessage string
}

// Write error log
func (errorLog ErrorLog) WriteErrorLog() {
	// Open Log File
	f, _ := os.OpenFile(
		conf.LogPath+"/"+conf.ErrorLog,
		os.O_RDWR|os.O_CREATE|os.O_APPEND,
		0666,
	)
	// Append the error to file
	log.SetOutput(f)

	log.Println(errorLog.ErrorMessage)
	// Close the file after ending the program
	f.Close()
}
