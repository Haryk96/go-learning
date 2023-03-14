package logger

import (
	"log"
	"os"
)

var Logger *log.Logger

func init() {
	Logger = log.New(os.Stdout, "JWT-Server ", log.LstdFlags)
}
