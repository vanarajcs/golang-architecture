package logs

import (
	"golang-architecture/conf"
	"golang-architecture/util"
	"log"
	"os"

	"github.com/kataras/iris/context"
)

// Access Log
type AccessLog struct {
	Method    string `json:"method"`
	ClientIp  string `json:"clientIp"`
	UserAgent string `json:"userAgent"`
}

// Write access log
func (accessLog AccessLog) WriteAccessLog() {
	// Open Log File
	f, err := os.OpenFile(
		conf.LogPath+"/"+conf.AccessLog,
		os.O_RDWR|os.O_CREATE|os.O_APPEND,
		0666,
	)

	// Check if Error
	if err != nil {
		// Create Logs Directory
		utilObject := util.Util{}
		utilObject.CreateDirectory(conf.LogPath)
		// Call Back
		accessLog.WriteAccessLog()
	}

	// Set Log
	log.SetOutput(f)

	if accessLog.Method == "" {
		log.Println("Agent started successfully...!")
	} else {
		log.Println(accessLog.Method + " " + accessLog.ClientIp + " " + accessLog.UserAgent)
	}
	// Close the file after ending the program
	f.Close()
}

// Get Access Log from context
func (accessLog AccessLog) GetAccessLog(ctx context.Context) AccessLog {
	// Access Log entity to write access log
	accesslog := AccessLog{}
	accesslog.Method = ctx.Method() + ctx.GetHeader("apiKey")
	accesslog.ClientIp = ctx.RemoteAddr() + ctx.Path()
	accesslog.UserAgent = ctx.GetHeader("User-Agent")
	return accesslog
}
