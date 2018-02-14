package logging

import (
	"io"
	"log"
	"os"
)

var (
	// Trace :
	Trace *log.Logger
	// Info :
	Info *log.Logger
	// Warning :
	Warning *log.Logger
	// Error :
	Error *log.Logger
)

//Init : Initialise the logging module
func Init(traceHandle io.Writer) {

	file, err := os.OpenFile("/var/log/coupons/coupons.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	if err != nil {
		log.Fatalln("Failed to open Coupon log file :", err)
	}

	Trace = log.New(traceHandle,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(file,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(file,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(file,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}
