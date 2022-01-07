package util

import (
	"log"
)

var (
	INFO *log.Logger
	WARN *log.Logger
	ERR  *log.Logger
)

// func init() {
// 	logfile, err := os.OpenFile(".log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
// 	if err != nil {
// 		panic(err)
// 	}
// 	INFO = log.New(logfile, "INFO: ", log.LstdFlags|log.Llongfile)
// 	WARN = log.New(logfile, "WARN: ", log.LstdFlags|log.Llongfile)
// 	ERR = log.New(logfile, "ERR: ", log.LstdFlags|log.Llongfile)
// }
