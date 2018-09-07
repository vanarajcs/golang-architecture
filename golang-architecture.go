package main

import (
	"golang-architecture/conf"
	"golang-architecture/logs"
	"golang-architecture/util"
	"os"
	"strings"

	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

// Starting point
func main() {
	// Get Process Id
	pid := os.Getpid()
	// Util object
	utilObj := util.Util{}
	utilObj.WritePidFile(pid)

	// Access Logspvc-a3602db0-9719-11e8-b812-a2ac76cc05ae
	accessLog := logs.AccessLog{}
	accessLog.WriteAccessLog()

	// Create Iris Object
	app := iris.New()

	// Middleware for authentication
	app.Use(AuthenticationMiddleWare)

	// Home Url
	app.Get("/", func(ctx context.Context) {
		// Response
		ctx.JSON(context.Map{"result": "Ok", "status": iris.StatusOK})
	})

	// Handle internal server error
	app.OnErrorCode(iris.StatusInternalServerError, func(ctx context.Context) {
		errMessage := ctx.Values().Get("error")
		ctx.Writef("%s", errMessage)
	})

	// Handle unauthorized server error
	app.OnErrorCode(iris.StatusUnauthorized, func(ctx context.Context) {
		errMessage := ctx.Values().Get("error")
		ctx.Writef("%s", errMessage)
	})

	// Run the app
	app.Run(
		iris.Addr(":6300"),
		iris.WithoutBanner,
		iris.WithoutStartupLog,
		iris.WithoutVersionChecker,
	)
}

// Before running the page.
func AuthenticationMiddleWare(ctx context.Context) {
	//Allow access to all
	ctx.Header("Access-Control-Allow-Origin", "*")
	// Write Acccess Log
	accessLog := logs.AccessLog{}
	accessLog = accessLog.GetAccessLog(ctx)
	accessLog.WriteAccessLog()
	// check the api key
	if ctx.GetHeader("Api-Key") == "" ||
		ctx.GetHeader("Api-Key") != strings.TrimSpace(conf.ApiKey) {
		// Return Error Message
		ctx.JSON(context.Map{"result": "Ok", "status": "unauthorized"})
		return
	}
	// Redirect to next page
	ctx.Next()
}
