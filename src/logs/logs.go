package logs

import (
	"log"
)

func Sucess(m string) {
	log.Println(m)
}

func Error(err error) {
	log.Println(err)
}
