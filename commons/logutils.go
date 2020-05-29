package commons

import (
	"log"
)

//the content of this utility function  can be easily commented,
//to avoid printing on standard out in case of errors
func InnerPrintln(v ...interface{}) {
	log.Println(v...)
}
