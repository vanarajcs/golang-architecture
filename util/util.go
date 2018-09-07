package util

import (
	"golang-architecture/conf"
	"os"
	"strconv"
)

// Util Object
type Util struct{}

// Create Logs Directory
func (util Util) CreateDirectory(path string) {
	// Create all directories
	err := os.MkdirAll(path, 0777)

	if err != nil {
	}
}

// Replace contents to file
func (util Util) ReplaceFileContents(path string, content string) {
	// Open pid file
	f, err := os.OpenFile(
		path,
		os.O_RDWR|os.O_CREATE|os.O_RDWR,
		0666,
	)

	if err != nil {
		util.CreateDirectory(path)
		// Call Back
		util.ReplaceFileContents(path, content)
	}

	// Write the pid in to the file
	f.WriteString(content)

	// Close the file after ending the function
	f.Close()
}

// Replace contents to file
func (util Util) WritePidFile(pid int) {
	// Open pid file
	f, err := os.OpenFile(
		conf.PidPath+"/"+conf.Pid,
		os.O_RDWR|os.O_CREATE|os.O_RDWR,
		0666,
	)

	if err != nil {
		util.CreateDirectory(conf.PidPath)
		// Call Back
		util.WritePidFile(pid)
	}
	// Write the pid in to the file
	f.WriteString(strconv.Itoa(pid))

	// Close the file after ending the function
	f.Close()
}
