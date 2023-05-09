package logger

import (
	"fmt"
	"log"
)

func Fatalf(str string, err error) {
	log.Fatalf("FATAL: "+str, err)
}

func Errorf(str string, err error) error {
	return fmt.Errorf("ERROR: "+str, err)
}
