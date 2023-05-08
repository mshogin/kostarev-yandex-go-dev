package logger

import (
	"fmt"
	"log"
)

func Fatalf(str string, err ...any) {
	log.Fatalf("ERROR: "+str, err)
}

func Errorf(str string, err ...any) error {
	return fmt.Errorf(str, err)
}
