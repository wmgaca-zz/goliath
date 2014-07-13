package main

import (
	"log"
	"os"
)

const disapprovalLook string = "ಠ_ಠ"

func ExitWithErr(message string) {
	log.Println(disapprovalLook, message)
	os.Exit(-1)
}
