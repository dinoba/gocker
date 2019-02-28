package log

import (
	"log"
	"os"
	"sync"
	"time"
)

type logger struct {
	filename string
	*log.Logger
}

var logg *logger
var once sync.Once

//GetInstance retuns pointer to exactly one
func GetInstance() *logger {
	once.Do(func() {
		//TODO make this configurable
		logg = createLogger("dockergo.log")
		logg.SetFlags(0)
	})
	return logg
}

func createLogger(fname string) *logger {
	file, err := os.OpenFile(fname, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		//TODO add some desc
		panic(err)
	}
	return &logger{
		filename: fname,
		Logger:   log.New(file, "Consolidated logs ", log.Lshortfile),
	}
}

//WithPrefix - wrapper to add timestamp
func WithPrefix(l *logger, msg string, info string) {
	l.SetPrefix(time.Now().Format("2006-01-02 15:04:05 "))
	l.Print(info + " " + msg)
}
