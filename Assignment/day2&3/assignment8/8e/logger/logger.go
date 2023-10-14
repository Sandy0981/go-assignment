package logger

import (
	"log"
	"os"
)

type LogStore struct {
	log *log.Logger
}

func New() *LogStore {
	return &LogStore{
		log: log.New(os.Stdout, "Sandeep", log.Lshortfile),
	}
}
