package core

import (
	"log"
)

func Logf(f string, args ...interface{}) {
	log.Printf(f, args...)

}

func Debug(f string, args ...interface{}) {
	log.Printf(f, args...)

}
