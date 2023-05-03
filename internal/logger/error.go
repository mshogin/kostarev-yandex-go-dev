package logger

import (
	"log"
)

func Error(str string, err ...any) {
	log.Fatalf("INFO: "+str, err)
}
